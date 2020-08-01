// +build integration

package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	"time"

	model "github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestHTTP_POST_Login(t *testing.T) {
	t.Run("will return error if email is empty", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":    "",
			"password": "",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("login"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "email is empty", res.Message)
	})

	t.Run("will return error if password is empty", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":    validEmail,
			"password": "",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("login"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "password is empty", res.Message)
	})

	t.Run("will return error if error getting user", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":    getUserErrorEmail,
			"password": password,
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("login"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not login", res.Message)
	})

	t.Run("will return error if user not found", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":    notFoundErrorEmail,
			"password": password,
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("login"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not login", res.Message)
	})

	t.Run("will return error if password does not match", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":    validEmail,
			"password": "some different password",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("login"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not login", res.Message)
	})

	t.Run("will return user's ID and token if the credentials are correct", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":    validEmail,
			"password": password,
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("login"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res loginResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, validEmail, res.Id)

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token:    res.Token,
				Audience: validEmail,
				Role:     model.Role_ROLE_USER,
			})
		require.NoError(t, err)
	})
}

func TestHTTP_POST_Register(t *testing.T) {
	t.Run("will return first name if email is empty", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"firstName":       "",
			"surname":         "",
			"email":           "",
			"password":        "",
			"confirmation":    "",
			"predictedWinner": "",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("sign-up"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "first name is empty", res.Message)
	})

	t.Run("will return last name if email is empty", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"firstName":       "first name",
			"surname":         "",
			"email":           "",
			"password":        "",
			"confirmation":    "",
			"predictedWinner": "",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("sign-up"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "surname is empty", res.Message)
	})

	t.Run("will return error if email is empty", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"firstName":       "first name",
			"surname":         "last name",
			"email":           "",
			"password":        "",
			"confirmation":    "",
			"predictedWinner": "",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("sign-up"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "email is empty", res.Message)
	})

	t.Run("will return error if password is empty", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"firstName":       "first name",
			"surname":         "last name",
			"email":           "email",
			"password":        "",
			"confirmation":    "",
			"predictedWinner": "",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("sign-up"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "password is empty", res.Message)
	})

	t.Run("will return error if confirmation is empty", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"firstName":       "first name",
			"surname":         "last name",
			"email":           "email",
			"password":        "password",
			"confirmation":    "",
			"predictedWinner": "",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("sign-up"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "confirmation is empty", res.Message)
	})

	t.Run("will return error if predicted winner is empty", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"firstName":       "first name",
			"surname":         "last name",
			"email":           "email",
			"password":        "password",
			"confirmation":    "confirmation",
			"predictedWinner": "",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("sign-up"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "predicted winner is empty", res.Message)
	})

	t.Run("will return error if email is invalid", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"firstName":       "first name",
			"surname":         "last name",
			"email":           "email",
			"password":        "password",
			"confirmation":    "confirmation",
			"predictedWinner": "predicted winner",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("sign-up"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "email address is invalid", res.Message)
	})

	t.Run("will return error if password is invalid", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"firstName":       "first name",
			"surname":         "last name",
			"email":           validEmail,
			"password":        "password",
			"confirmation":    "confirmation",
			"predictedWinner": "predicted winner",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("sign-up"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "password is invalid", res.Message)
	})

	t.Run("will return error if password and confirmation does not match", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"firstName":       "first name",
			"surname":         "last name",
			"email":           validEmail,
			"password":        password,
			"confirmation":    "confirmation",
			"predictedWinner": "predicted winner",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("sign-up"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "password and confirmation do not match", res.Message)
	})

	t.Run("will return error if user already exists with email", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"firstName":       "first name",
			"surname":         "last name",
			"email":           validEmail,
			"password":        password,
			"confirmation":    password,
			"predictedWinner": "predicted winner",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("sign-up"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusConflict, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "email already exists", res.Message)
	})

	t.Run("will return error if error getting user by email", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"firstName":       "first name",
			"surname":         "last name",
			"email":           getUserErrorEmail,
			"password":        password,
			"confirmation":    password,
			"predictedWinner": "predicted winner",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("sign-up"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not register", res.Message)
	})

	t.Run("will return error if error creating user", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"firstName":       "first name",
			"surname":         "last name",
			"email":           failCreateErrorEmail,
			"password":        password,
			"confirmation":    password,
			"predictedWinner": "predicted winner",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("sign-up"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not register", res.Message)
	})

	t.Run("will return user's ID and token if user created successfully", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"firstName":       "first name",
			"surname":         "last name",
			"email":           notFoundErrorEmail,
			"password":        password,
			"confirmation":    password,
			"predictedWinner": "predicted winner",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("sign-up"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res registerResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, createdUserId, res.Id)

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, grpcHostname, grpc.WithInsecure(), grpc.WithBlock())
		require.NoError(t, err)
		defer conn.Close()

		_, err = model.NewAuthServiceClient(conn).
			Validate(ctx, &model.ValidateRequest{
				Token:    res.Token,
				Audience: createdUserId,
				Role:     model.Role_ROLE_USER,
			})
		require.NoError(t, err)
	})
}

func TestHTTP_POST_InitiatePasswordReset(t *testing.T) {
	t.Run("will return error if email is empty", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email": "",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("initiate-password-reset"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "email is empty", res.Message)
	})

	t.Run("will return error if email is invalid", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email": "invalid email",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("initiate-password-reset"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "email address is invalid", res.Message)
	})

	t.Run("will return error if error getting user", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email": notFoundErrorEmail,
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("initiate-password-reset"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not initiate password reset", res.Message)
	})

	t.Run("will return error if error updating signature user", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email": failUpdateSignatureErrorEmail,
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("initiate-password-reset"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not initiate password reset", res.Message)
	})

	t.Run("will return error if error sending email", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email": failEmailErrorEmail,
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("initiate-password-reset"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not initiate password reset", res.Message)
	})

	t.Run("will update user's signature and send password reset email", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email": validEmail,
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("initiate-password-reset"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestHTTP_POST_ResetPassword(t *testing.T) {
	t.Run("will return error if email is empty", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":        "",
			"signature":    "",
			"password":     "",
			"confirmation": "",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("reset-password"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "email is empty", res.Message)
	})

	t.Run("will return error if signature is empty", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":        validEmail,
			"signature":    "",
			"password":     "",
			"confirmation": "",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("reset-password"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "signature is empty", res.Message)
	})

	t.Run("will return error if password is empty", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":        validEmail,
			"signature":    "signature",
			"password":     "",
			"confirmation": "",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("reset-password"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "password is empty", res.Message)
	})

	t.Run("will return error if confirmation is empty", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":        validEmail,
			"signature":    "signature",
			"password":     "password",
			"confirmation": "",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("reset-password"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "confirmation is empty", res.Message)
	})

	t.Run("will return error if password is invalid", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":        validEmail,
			"signature":    "signature",
			"password":     "password",
			"confirmation": "confirmation",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("reset-password"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "password is invalid", res.Message)
	})

	t.Run("will return error if password and confirmation does not match", func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":        validEmail,
			"signature":    "signature",
			"password":     password,
			"confirmation": "confirmation",
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("reset-password"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "password and confirmation do not match", res.Message)
	})

	t.Run("will return error if signature not valid", func(t *testing.T) {
		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
			SignedString([]byte("different secret"))
		require.NoError(t, err)

		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":        validEmail,
			"signature":    signature,
			"password":     password,
			"confirmation": password,
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("reset-password"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not reset password", res.Message)
	})

	t.Run("will return error if error getting user", func(t *testing.T) {
		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":        getUserErrorEmail,
			"signature":    signature,
			"password":     password,
			"confirmation": password,
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("reset-password"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not reset password", res.Message)
	})

	t.Run("will return error if signature does not match", func(t *testing.T) {
		claims := &jwt.StandardClaims{
			Audience: "will create different signature",
		}
		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":        validEmail,
			"signature":    signature,
			"password":     password,
			"confirmation": password,
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("reset-password"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not reset password", res.Message)
	})

	t.Run("will return error if error updating password", func(t *testing.T) {
		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":        failUpdatePasswordErrorEmail,
			"signature":    signature,
			"password":     password,
			"confirmation": password,
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("reset-password"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not reset password", res.Message)
	})

	t.Run("will successfully update user's password", func(t *testing.T) {
		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(map[string]interface{}{
			"email":        validEmail,
			"signature":    signature,
			"password":     password,
			"confirmation": password,
		}))

		req, err := http.NewRequest(http.MethodPost, buildUrl("reset-password"), bytes.NewReader(b.Bytes()))
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestHTTP_GET_ResetPassword(t *testing.T) {
	t.Run("will return error if email is empty", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, buildUrl("reset-password"), nil)
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "email is empty", res.Message)
	})

	t.Run("will return error if signature is empty", func(t *testing.T) {
		u, err := url.Parse(buildUrl("reset-password"))
		require.NoError(t, err)

		q := u.Query()
		q.Add("email", validEmail)
		u.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "signature is empty", res.Message)
	})

	t.Run("will return error if password is empty", func(t *testing.T) {
		u, err := url.Parse(buildUrl("reset-password"))
		require.NoError(t, err)

		q := u.Query()
		q.Add("email", validEmail)
		q.Add("signature", "signature")
		u.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "password is empty", res.Message)
	})

	t.Run("will return error if confirmation is empty", func(t *testing.T) {
		u, err := url.Parse(buildUrl("reset-password"))
		require.NoError(t, err)

		q := u.Query()
		q.Add("email", validEmail)
		q.Add("signature", "signature")
		q.Add("password", "password")
		u.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "confirmation is empty", res.Message)
	})

	t.Run("will return error if password is invalid", func(t *testing.T) {
		u, err := url.Parse(buildUrl("reset-password"))
		require.NoError(t, err)

		q := u.Query()
		q.Add("email", validEmail)
		q.Add("signature", "signature")
		q.Add("password", "password")
		q.Add("confirmation", "confirmation")
		u.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "password is invalid", res.Message)
	})

	t.Run("will return error if password and confirmation does not match", func(t *testing.T) {
		u, err := url.Parse(buildUrl("reset-password"))
		require.NoError(t, err)

		q := u.Query()
		q.Add("email", validEmail)
		q.Add("signature", "signature")
		q.Add("password", password)
		q.Add("confirmation", "confirmation")
		u.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "password and confirmation do not match", res.Message)
	})

	t.Run("will return error if signature not valid", func(t *testing.T) {
		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
			SignedString([]byte("different secret"))
		require.NoError(t, err)

		u, err := url.Parse(buildUrl("reset-password"))
		require.NoError(t, err)

		q := u.Query()
		q.Add("email", validEmail)
		q.Add("signature", signature)
		q.Add("password", password)
		q.Add("confirmation", password)
		u.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not reset password", res.Message)
	})

	t.Run("will return error if error getting user", func(t *testing.T) {
		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		u, err := url.Parse(buildUrl("reset-password"))
		require.NoError(t, err)

		q := u.Query()
		q.Add("email", getUserErrorEmail)
		q.Add("signature", signature)
		q.Add("password", password)
		q.Add("confirmation", password)
		u.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not reset password", res.Message)
	})

	t.Run("will return error if signature does not match", func(t *testing.T) {
		claims := &jwt.StandardClaims{
			Audience: "will create different signature",
		}
		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, claims).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		u, err := url.Parse(buildUrl("reset-password"))
		require.NoError(t, err)

		q := u.Query()
		q.Add("email", validEmail)
		q.Add("signature", signature)
		q.Add("password", password)
		q.Add("confirmation", password)
		u.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not reset password", res.Message)
	})

	t.Run("will return error if error updating password", func(t *testing.T) {
		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		u, err := url.Parse(buildUrl("reset-password"))
		require.NoError(t, err)

		q := u.Query()
		q.Add("email", failUpdatePasswordErrorEmail)
		q.Add("signature", signature)
		q.Add("password", password)
		q.Add("confirmation", password)
		u.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		defer resp.Body.Close()
		buf, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)

		var res errorResponse
		err = json.Unmarshal(buf, &res)
		require.NoError(t, err)

		assert.Equal(t, "could not reset password", res.Message)
	})

	t.Run("will successfully update user's password", func(t *testing.T) {
		signature, err := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.StandardClaims{}).
			SignedString([]byte(jwtSecret))
		require.NoError(t, err)

		u, err := url.Parse(buildUrl("reset-password"))
		require.NoError(t, err)

		q := u.Query()
		q.Add("email", validEmail)
		q.Add("signature", signature)
		q.Add("password", password)
		q.Add("confirmation", password)
		u.RawQuery = q.Encode()

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		require.NoError(t, err)

		req.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
