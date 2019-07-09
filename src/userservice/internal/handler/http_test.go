package handler

import (
	"bytes"
	"encoding/json"
	authmocks "github.com/cshep4/premier-predictor-microservices/src/common/auth/mocks"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"
	usermocks "github.com/cshep4/premier-predictor-microservices/src/userservice/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	userId = "userId"
)

func TestHttpHandler_getUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := usermocks.NewMockService(ctrl)
	auth := authmocks.NewMockAuthenticator(ctrl)

	handler, err  := NewHttpHandler(service, auth)
	require.NoError(t, err)

	t.Run("it should return not found if no user exists", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/users/" + userId,
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": userId,
		})

		service.EXPECT().GetUser(userId).Return(nil, model.ErrUserNotFound)

		handler.getUser(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, model.ErrUserNotFound.Error(), responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/users/" + userId,
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": userId,
		})

		service.EXPECT().GetUser(userId).Return(nil, errors.New("some error"))

		handler.getUser(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "some error", responseBody.Message)
	})

	t.Run("it should return ok with user in body", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/users/" + userId,
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": userId,
		})

		user := &model.User{
			Id: userId,
		}

		service.EXPECT().GetUser(userId).Return(user, nil)

		handler.getUser(rr, req)

		var responseBody model.User
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, user, &responseBody)
	})
}

func TestHttpHandler_updateUserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := usermocks.NewMockService(ctrl)
	auth := authmocks.NewMockAuthenticator(ctrl)

	handler, err  := NewHttpHandler(service, auth)
	require.NoError(t, err)

	userInfo := model.UserInfo{
		Id: userId,
	}

	b, err := json.Marshal(userInfo)
	require.NoError(t, err)

	t.Run("it should return bad request if request body is not valid", func(t *testing.T) {
		b, err := json.Marshal("invalid body")
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/users",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		handler.updateUserInfo(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, invalidRequestBody, responseBody.Message)
	})

	t.Run("it should return not found if no user exists", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodPut,
			"/users",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().UpdateUserInfo(userInfo).Return(model.ErrUserNotFound)

		handler.updateUserInfo(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, model.ErrUserNotFound.Error(), responseBody.Message)
	})

	t.Run("it should return bad request if input is invalid", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodPut,
			"/users",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		const errorMessage = "error message"
		e := errors.Wrap(common.ErrInvalidRequestData, errorMessage)

		service.EXPECT().UpdateUserInfo(userInfo).Return(e)

		handler.updateUserInfo(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, errorMessage, responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodPut,
			"/users",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		e := errors.New("some error")

		service.EXPECT().UpdateUserInfo(userInfo).Return(e)

		handler.updateUserInfo(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, e.Error(), responseBody.Message)
	})

	t.Run("it should update user info and return ok", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodPut,
			"/users",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().UpdateUserInfo(userInfo).Return(nil)

		handler.updateUserInfo(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestHttpHandler_updatePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := usermocks.NewMockService(ctrl)
	auth := authmocks.NewMockAuthenticator(ctrl)

	handler, err  := NewHttpHandler(service, auth)
	require.NoError(t, err)

	updatePassword := model.UpdatePassword{
		Id: userId,
	}

	b, err := json.Marshal(updatePassword)
	require.NoError(t, err)

	t.Run("it should return bad request if request body is not valid", func(t *testing.T) {
		b, err := json.Marshal("invalid body")
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/users/password",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		handler.updatePassword(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, invalidRequestBody, responseBody.Message)
	})

	t.Run("it should return not found if no user exists", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodPut,
			"/users/password",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().UpdatePassword(updatePassword).Return(model.ErrUserNotFound)

		handler.updatePassword(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, model.ErrUserNotFound.Error(), responseBody.Message)
	})

	t.Run("it should return bad request if input is invalid", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodPut,
			"/users/password",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		const errorMessage = "error message"
		e := errors.Wrap(common.ErrInvalidRequestData, errorMessage)

		service.EXPECT().UpdatePassword(updatePassword).Return(e)

		handler.updatePassword(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, errorMessage, responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodPut,
			"/users/password",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		e := errors.New("some error")

		service.EXPECT().UpdatePassword(updatePassword).Return(e)

		handler.updatePassword(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, e.Error(), responseBody.Message)
	})

	t.Run("it should update password and return ok", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodPut,
			"/users/password",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().UpdatePassword(updatePassword).Return(nil)

		handler.updatePassword(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestHttpHandler_getUserScore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := usermocks.NewMockService(ctrl)
	auth := authmocks.NewMockAuthenticator(ctrl)

	handler, err  := NewHttpHandler(service, auth)
	require.NoError(t, err)

	t.Run("it should return not found if no user exists", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/users/score/" + userId,
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": userId,
		})

		service.EXPECT().GetUserScore(userId).Return(0, model.ErrUserNotFound)

		handler.getUserScore(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, model.ErrUserNotFound.Error(), responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/users/score/" + userId,
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": userId,
		})

		service.EXPECT().GetUserScore(userId).Return(0, errors.New("some error"))

		handler.getUserScore(rr, req)

		var responseBody ServerError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "some error", responseBody.Message)
	})

	t.Run("it should return ok with user in body", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/users/score/" + userId,
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": userId,
		})

		const score = 1234
		service.EXPECT().GetUserScore(userId).Return(score, nil)

		handler.getUserScore(rr, req)

		var responseBody model.UserScore
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, score, responseBody.Score)
	})
}