package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/mocks/service"
	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	pin        = int64(12345)
	userId     = "🆔"
	leagueName = "🏆🏆🏆🏆🏆🏆"
)

func TestHttpHandler_getLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := service_mock.NewMockService(ctrl)

	handler, err := New(service)
	require.NoError(t, err)

	t.Run("it should return bad request if an invalid parameter error", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/"+userId,
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": userId,
		})

		service.EXPECT().GetUsersLeagueList(req.Context(), userId).Return(nil, model.InvalidParameterError{Parameter: "param"})

		handler.getUsersLeagueList(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "invalid parameter: param", responseBody.Message)
	})

	t.Run("it should return internal server error if an error occurred", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/"+userId,
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": userId,
		})

		service.EXPECT().GetUsersLeagueList(req.Context(), userId).Return(nil, errors.New("some error"))

		handler.getUsersLeagueList(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "could not get leagues", responseBody.Message)
	})

	t.Run("it should return ok with league overview in body", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/"+userId,
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": userId,
		})

		leagueOverview := &model.StandingsOverview{
			OverallLeagueOverview: model.OverallLeagueOverview{
				UserCount: 50,
			},
		}

		service.EXPECT().GetUsersLeagueList(req.Context(), userId).Return(leagueOverview, nil)

		handler.getUsersLeagueList(rr, req)

		var responseBody model.StandingsOverview
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, leagueOverview, &responseBody)
	})
}

func TestHttpHandler_addLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := service_mock.NewMockService(ctrl)

	handler, err := New(service)
	require.NoError(t, err)

	t.Run("it should return bad request if the request body is invalid", func(t *testing.T) {
		b, err := json.Marshal("invalid body")
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		handler.addLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "invalid request body", responseBody.Message)
	})

	t.Run("it should return bad request if an invalid parameter error", func(t *testing.T) {
		b, err := json.Marshal(addLeagueRequest{
			Id:   userId,
			Name: leagueName,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().AddUserLeague(req.Context(), userId, leagueName).Return(nil, model.InvalidParameterError{Parameter: "param"})

		handler.addLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "invalid parameter: param", responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		b, err := json.Marshal(addLeagueRequest{
			Id:   userId,
			Name: leagueName,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().AddUserLeague(req.Context(), userId, leagueName).Return(nil, errors.New("some error"))

		handler.addLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "could not add league", responseBody.Message)
	})

	t.Run("it should return created with league in body", func(t *testing.T) {
		b, err := json.Marshal(addLeagueRequest{
			Id:   userId,
			Name: leagueName,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPost,
			"/",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		league := &model.League{
			Name:  leagueName,
			Users: []string{userId},
			Pin:   pin,
		}

		service.EXPECT().AddUserLeague(req.Context(), userId, leagueName).Return(league, nil)

		handler.addLeague(rr, req)

		var responseBody model.League
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, league, &responseBody)
	})
}

func TestHttpHandler_joinLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := service_mock.NewMockService(ctrl)

	handler, err := New(service)
	require.NoError(t, err)

	t.Run("it should return bad request if the request body is invalid", func(t *testing.T) {
		b, err := json.Marshal("invalid body")
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/join",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		handler.joinLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "invalid request body", responseBody.Message)
	})

	t.Run("it should return bad request if an invalid parameter error", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:  userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/join",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().JoinUserLeague(req.Context(), userId, pin).Return(nil, model.InvalidParameterError{Parameter: "param"})

		handler.joinLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "invalid parameter: param", responseBody.Message)
	})

	t.Run("it should return not found if league not found", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:  userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/join",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().JoinUserLeague(req.Context(), userId, pin).Return(nil, model.ErrLeagueNotFound)

		handler.joinLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, model.ErrLeagueNotFound.Error(), responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:  userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/join",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().JoinUserLeague(req.Context(), userId, pin).Return(nil, errors.New("some error"))

		handler.joinLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "could not join league", responseBody.Message)
	})

	t.Run("it should return ok with league overview in body", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:  userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/join",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		leagueOverview := &model.LeagueOverview{
			Pin:        pin,
			LeagueName: leagueName,
			Rank:       1,
		}

		service.EXPECT().JoinUserLeague(req.Context(), userId, pin).Return(leagueOverview, nil)

		handler.joinLeague(rr, req)

		var responseBody model.LeagueOverview
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, leagueOverview, &responseBody)
	})
}

func TestHttpHandler_leaveLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := service_mock.NewMockService(ctrl)

	handler, err := New(service)
	require.NoError(t, err)

	t.Run("it should return bad request if the request body is invalid", func(t *testing.T) {
		b, err := json.Marshal("invalid body")
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/leave",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		handler.leaveLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "invalid request body", responseBody.Message)
	})

	t.Run("it should return bad request if an invalid parameter error", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:  userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/leave",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().LeaveUserLeague(req.Context(), userId, pin).Return(model.InvalidParameterError{Parameter: "param"})

		handler.leaveLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "invalid parameter: param", responseBody.Message)
	})

	t.Run("it should return not found if league not found", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:  userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/leave",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().LeaveUserLeague(req.Context(), userId, pin).Return(model.ErrLeagueNotFound)

		handler.leaveLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, model.ErrLeagueNotFound.Error(), responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:  userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/leave",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().LeaveUserLeague(req.Context(), userId, pin).Return(errors.New("some error"))

		handler.leaveLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "could not leave league", responseBody.Message)
	})

	t.Run("it should return ok", func(t *testing.T) {
		b, err := json.Marshal(leagueRequest{
			Id:  userId,
			Pin: pin,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/leave",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().LeaveUserLeague(req.Context(), userId, pin).Return(nil)

		handler.leaveLeague(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestHttpHandler_renameLeague(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := service_mock.NewMockService(ctrl)

	handler, err := New(service)
	require.NoError(t, err)

	t.Run("it should return bad request if the request body is invalid", func(t *testing.T) {
		b, err := json.Marshal("invalid body")
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/rename",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		handler.renameLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "invalid request body", responseBody.Message)
	})

	t.Run("it should return bad request if an invalid parameter error", func(t *testing.T) {
		b, err := json.Marshal(renameRequest{
			Pin:  pin,
			Name: leagueName,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/rename",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().RenameUserLeague(req.Context(), pin, leagueName).Return(model.InvalidParameterError{Parameter: "param"})

		handler.renameLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "invalid parameter: param", responseBody.Message)
	})

	t.Run("it should return not found if league not found", func(t *testing.T) {
		b, err := json.Marshal(renameRequest{
			Pin:  pin,
			Name: leagueName,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/rename",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().RenameUserLeague(req.Context(), pin, leagueName).Return(model.ErrLeagueNotFound)

		handler.renameLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, model.ErrLeagueNotFound.Error(), responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		b, err := json.Marshal(renameRequest{
			Pin:  pin,
			Name: leagueName,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/rename",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().RenameUserLeague(req.Context(), pin, leagueName).Return(errors.New("some error"))

		handler.renameLeague(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "could not rename league", responseBody.Message)
	})

	t.Run("it should return ok", func(t *testing.T) {
		b, err := json.Marshal(renameRequest{
			Pin:  pin,
			Name: leagueName,
		})
		require.NoError(t, err)

		req := httptest.NewRequest(
			http.MethodPut,
			"/rename",
			bytes.NewReader(b),
		)
		rr := httptest.NewRecorder()

		service.EXPECT().RenameUserLeague(req.Context(), pin, leagueName).Return(nil)

		handler.renameLeague(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestHttpHandler_getLeagueTable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := service_mock.NewMockService(ctrl)

	handler, err := New(service)
	require.NoError(t, err)

	t.Run("it should return bad request if pin is invalid", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/standings/invalidPin",
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": "invalid pin",
		})

		handler.getLeagueTable(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "invalid parameter: pin", responseBody.Message)
	})

	t.Run("it should return bad request if an invalid parameter error", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/standings/"+strconv.FormatInt(pin, 10),
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": strconv.FormatInt(pin, 10),
		})

		service.EXPECT().GetLeagueTable(req.Context(), pin).Return(nil, model.InvalidParameterError{Parameter: "param"})

		handler.getLeagueTable(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "invalid parameter: param", responseBody.Message)
	})

	t.Run("it should return not found if league not found", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/standings/"+strconv.FormatInt(pin, 10),
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": strconv.FormatInt(pin, 10),
		})

		service.EXPECT().GetLeagueTable(req.Context(), pin).Return(nil, model.ErrLeagueNotFound)

		handler.getLeagueTable(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, model.ErrLeagueNotFound.Error(), responseBody.Message)
	})

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/standings/"+strconv.FormatInt(pin, 10),
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": strconv.FormatInt(pin, 10),
		})

		service.EXPECT().GetLeagueTable(req.Context(), pin).Return(nil, errors.New("some error"))

		handler.getLeagueTable(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "could not get league table", responseBody.Message)
	})

	t.Run("it should return ok", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/standings/"+strconv.FormatInt(pin, 10),
			nil,
		)
		rr := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": strconv.FormatInt(pin, 10),
		})

		leagueTable := []model.LeagueUser{
			{
				Id:    userId,
				Score: 1,
			},
		}

		service.EXPECT().GetLeagueTable(req.Context(), pin).Return(leagueTable, nil)

		handler.getLeagueTable(rr, req)

		var responseBody []model.LeagueUser
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, leagueTable, responseBody)
	})
}

func TestHttpHandler_getOverallTable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := service_mock.NewMockService(ctrl)

	handler, err := New(service)
	require.NoError(t, err)

	t.Run("it should return internal server error if another error occurred", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/standings",
			nil,
		)
		rr := httptest.NewRecorder()

		service.EXPECT().GetOverallLeagueTable(req.Context()).Return(nil, errors.New("some error"))

		handler.getOverallTable(rr, req)

		var responseBody serverError
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "could not get overall table", responseBody.Message)
	})

	t.Run("it should return ok", func(t *testing.T) {
		req := httptest.NewRequest(
			http.MethodGet,
			"/standings",
			nil,
		)
		rr := httptest.NewRecorder()

		leagueTable := []model.LeagueUser{
			{
				Id:    userId,
				Score: 1,
			},
		}

		service.EXPECT().GetOverallLeagueTable(req.Context()).Return(leagueTable, nil)

		handler.getOverallTable(rr, req)

		var responseBody []model.LeagueUser
		err = json.NewDecoder(rr.Result().Body).Decode(&responseBody)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, leagueTable, responseBody)
	})
}
