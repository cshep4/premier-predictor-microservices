package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/userservice/internal/model"

	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	"github.com/gorilla/mux"
)

type (
	serverError struct {
		Message string `json:"message"`
	}

	router struct {
		service handler.Service
	}

	// InvalidParameterError is returned when a required parameter passed to New is invalid.
	InvalidParameterError struct {
		Parameter string
	}
)

func (i InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter %s", i.Parameter)
}

func New(service handler.Service) (*router, error) {
	if service == nil {
		return nil, InvalidParameterError{Parameter: "service"}
	}

	return &router{
		service: service,
	}, nil
}

func (h *router) Route(router *mux.Router) {
	router.HandleFunc("/users/{id}", h.getUser).
		Methods(http.MethodGet)
	router.HandleFunc("/users", h.updateUserInfo).
		Methods(http.MethodPut)
	router.HandleFunc("/users/password", h.updatePassword).
		Methods(http.MethodPut)
	router.HandleFunc("/users/score/{id}", h.getUserScore).
		Methods(http.MethodGet)
	//TODO - add tests
	router.HandleFunc("/legacy", h.storeLegacyUser).
		Methods(http.MethodPost)
}

func (h *router) getUser(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		h.errorResponse(r.Context(), http.StatusBadRequest, model.InvalidParameterError{
			Parameter: "id",
		}.Error(), w)
		return
	}

	user, err := h.service.GetUserById(r.Context(), id)
	switch {
	case err == nil:
		if err := json.NewEncoder(w).Encode(user); err != nil {
			log.Error(r.Context(), "encode_response_error", log.ErrorParam(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case errors.Is(err, model.ErrUserNotFound):
		h.errorResponse(r.Context(), http.StatusNotFound, "user not found", w)
	case isInvalidParameterErr(err):
		h.errorResponse(r.Context(), http.StatusBadRequest, err.Error(), w)
	default:
		log.Error(r.Context(), "error_getting_user", log.ErrorParam(err))
		h.errorResponse(r.Context(), http.StatusInternalServerError, "could not get user", w)
	}
}

func (h *router) updateUserInfo(w http.ResponseWriter, r *http.Request) {
	var userInfo model.UserInfo
	if err := json.NewDecoder(r.Body).Decode(&userInfo); err != nil {
		h.errorResponse(r.Context(), http.StatusBadRequest, "invalid request body", w)
		return
	}

	err := h.service.UpdateUserInfo(r.Context(), userInfo)
	switch {
	case err == nil:
		// 200 OK
	case errors.Is(err, model.ErrUserNotFound):
		h.errorResponse(r.Context(), http.StatusNotFound, "user not found", w)
	case isInvalidParameterErr(err):
		h.errorResponse(r.Context(), http.StatusBadRequest, err.Error(), w)
	default:
		log.Error(r.Context(), "error_updating_user_info", log.ErrorParam(err))
		h.errorResponse(r.Context(), http.StatusInternalServerError, "could not update user info", w)
	}
}

func (h *router) updatePassword(w http.ResponseWriter, r *http.Request) {
	var updatePassword model.UpdatePassword
	if err := json.NewDecoder(r.Body).Decode(&updatePassword); err != nil {
		h.errorResponse(r.Context(), http.StatusBadRequest, "invalid request body", w)
		return
	}

	err := h.service.UpdateUserPassword(r.Context(), updatePassword)
	switch {
	case err == nil:
		// 200 OK
	case errors.Is(err, model.ErrUserNotFound):
		h.errorResponse(r.Context(), http.StatusNotFound, "user not found", w)
	case isInvalidParameterErr(err):
		h.errorResponse(r.Context(), http.StatusBadRequest, err.Error(), w)
	default:
		log.Error(r.Context(), "error_updating_password", log.ErrorParam(err))
		h.errorResponse(r.Context(), http.StatusInternalServerError, "could not update password", w)
	}
}

func (h *router) getUserScore(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		h.errorResponse(r.Context(), http.StatusBadRequest, model.InvalidParameterError{
			Parameter: "id",
		}.Error(), w)
		return
	}

	score, err := h.service.GetUserScore(r.Context(), id)
	switch {
	case err == nil:
		if err := json.NewEncoder(w).Encode(model.UserScore{Score: score}); err != nil {
			log.Error(r.Context(), "encode_response_error", log.ErrorParam(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case errors.Is(err, model.ErrUserNotFound):
		h.errorResponse(r.Context(), http.StatusNotFound, "user not found", w)
	case isInvalidParameterErr(err):
		h.errorResponse(r.Context(), http.StatusBadRequest, err.Error(), w)
	default:
		log.Error(r.Context(), "error_getting_user_score", log.ErrorParam(err))
		h.errorResponse(r.Context(), http.StatusInternalServerError, "could not get user score", w)
	}
}

func (h *router) storeLegacyUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.errorResponse(r.Context(), http.StatusBadRequest, "invalid request body", w)
		return
	}

	_, err := h.service.StoreUser(r.Context(), user)
	switch {
	case err == nil:
		// 200 OK
	case errors.Is(err, model.ErrUserNotFound):
		h.errorResponse(r.Context(), http.StatusNotFound, "user not found", w)
	case isInvalidParameterErr(err):
		h.errorResponse(r.Context(), http.StatusBadRequest, err.Error(), w)
	default:
		log.Error(r.Context(), "error_storing_legacy_user", log.ErrorParam(err))
		h.errorResponse(r.Context(), http.StatusInternalServerError, "could not store legacy user", w)
	}
}

func (h *router) errorResponse(ctx context.Context, status int, message string, w http.ResponseWriter) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(serverError{Message: message}); err != nil {
		log.Error(ctx, "encode_response_error", log.ErrorParam(err))
		return
	}
}

func isInvalidParameterErr(err error) bool {
	_, ok := err.(model.InvalidParameterError)
	return ok
}

func (h *router) IsUnauthenticatedEndpoint(path string) bool {
	return path == "/legacy"
}

func (h *router) GetRequestAudience(r *http.Request) string {
	p, err := mux.CurrentRoute(r).GetPathTemplate()
	if err != nil {
		return ""
	}

	switch p {
	case "/users/{id}":
		fallthrough
	case "/users/score/{id}":
		return mux.Vars(r)["id"]
	case "/users":
		fallthrough
	case "/users/password":
		var req struct {
			Id string `json:"id"`
		}

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return ""
		}

		defer func() {
			r.Body.Close()
			r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}()

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return ""
		}
		return req.Id
	}

	return ""
}
