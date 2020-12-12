package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/model"

	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	"github.com/gorilla/mux"
)

type handler struct {
	service Service
}

func New(service Service) (*handler, error) {
	if service == nil {
		return nil, InvalidParameterError{Parameter: "service"}
	}

	return &handler{
		service: service,
	}, nil
}

func (h *handler) Route(router *mux.Router) {
	router.HandleFunc("/", h.addLeague).
		Methods(http.MethodPost)
	router.HandleFunc("/join", h.joinLeague).
		Methods(http.MethodPut)
	router.HandleFunc("/leave", h.leaveLeague).
		Methods(http.MethodPut)
	router.HandleFunc("/rename", h.renameLeague).
		Methods(http.MethodPut)
	router.HandleFunc("/standings/{id}", h.getLeagueTable).
		Methods(http.MethodGet)
	router.HandleFunc("/standings", h.getOverallTable).
		Methods(http.MethodGet)
	router.HandleFunc("/{id}", h.getUsersLeagueList).
		Methods(http.MethodGet)

	router.HandleFunc("/rebuild", h.rebuild).
		Methods(http.MethodPost)
}

func (h *handler) getUsersLeagueList(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		h.errorResponse(r.Context(), http.StatusBadRequest, model.InvalidParameterError{
			Parameter: "id",
		}.Error(), w)
		return
	}

	leagues, err := h.service.GetUsersLeagueList(r.Context(), id)
	switch {
	case err == nil:
		if err := json.NewEncoder(w).Encode(leagues); err != nil {
			log.Error(r.Context(), "encode_response_error", log.ErrorParam(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case errors.Is(err, model.ErrLeagueNotFound):
		h.errorResponse(r.Context(), http.StatusNotFound, "league not found", w)
	case isInvalidParameterErr(err):
		h.errorResponse(r.Context(), http.StatusBadRequest, err.Error(), w)
	default:
		log.Error(r.Context(), "error_getting_leagues", log.ErrorParam(err))
		h.errorResponse(r.Context(), http.StatusInternalServerError, "could not get leagues", w)
	}
}

func (h *handler) addLeague(w http.ResponseWriter, r *http.Request) {
	var req addLeagueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.errorResponse(r.Context(), http.StatusBadRequest, "invalid request body", w)
		return
	}

	league, err := h.service.AddUserLeague(r.Context(), req.Id, req.Name)
	switch {
	case err == nil:
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(league); err != nil {
			log.Error(r.Context(), "encode_response_error", log.ErrorParam(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case isInvalidParameterErr(err):
		h.errorResponse(r.Context(), http.StatusBadRequest, err.Error(), w)
	default:
		log.Error(r.Context(), "error_adding_league", log.ErrorParam(err))
		h.errorResponse(r.Context(), http.StatusInternalServerError, "could not add league", w)
	}
}

func (h *handler) joinLeague(w http.ResponseWriter, r *http.Request) {
	var req leagueRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.errorResponse(r.Context(), http.StatusBadRequest, "invalid request body", w)
		return
	}

	leagueOverview, err := h.service.JoinUserLeague(r.Context(), req.Id, req.Pin)
	switch {
	case err == nil:
		if err := json.NewEncoder(w).Encode(leagueOverview); err != nil {
			log.Error(r.Context(), "encode_response_error", log.ErrorParam(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case errors.Is(err, model.ErrLeagueNotFound):
		h.errorResponse(r.Context(), http.StatusNotFound, "league not found", w)
	case isInvalidParameterErr(err):
		h.errorResponse(r.Context(), http.StatusBadRequest, err.Error(), w)
	default:
		log.Error(r.Context(), "error_joining_league", log.ErrorParam(err))
		h.errorResponse(r.Context(), http.StatusInternalServerError, "could not join league", w)
	}
}

func (h *handler) leaveLeague(w http.ResponseWriter, r *http.Request) {
	var req leagueRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.errorResponse(r.Context(), http.StatusBadRequest, "invalid request body", w)
		return
	}

	err = h.service.LeaveUserLeague(r.Context(), req.Id, req.Pin)
	switch {
	case err == nil:
		// 200 OK
	case errors.Is(err, model.ErrLeagueNotFound):
		h.errorResponse(r.Context(), http.StatusNotFound, "league not found", w)
	case isInvalidParameterErr(err):
		h.errorResponse(r.Context(), http.StatusBadRequest, err.Error(), w)
	default:
		log.Error(r.Context(), "error_leaving_league", log.ErrorParam(err))
		h.errorResponse(r.Context(), http.StatusInternalServerError, "could not leave league", w)
	}
}

func (h *handler) renameLeague(w http.ResponseWriter, r *http.Request) {
	var req renameRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.errorResponse(r.Context(), http.StatusBadRequest, "invalid request body", w)
		return
	}

	err = h.service.RenameUserLeague(r.Context(), req.Pin, req.Name)
	switch {
	case err == nil:
		// 200 OK
	case errors.Is(err, model.ErrLeagueNotFound):
		h.errorResponse(r.Context(), http.StatusNotFound, "league not found", w)
	case isInvalidParameterErr(err):
		h.errorResponse(r.Context(), http.StatusBadRequest, err.Error(), w)
	default:
		log.Error(r.Context(), "error_renaming_league", log.ErrorParam(err))
		h.errorResponse(r.Context(), http.StatusInternalServerError, "could not rename league", w)
	}
}

func (h *handler) getLeagueTable(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		h.errorResponse(r.Context(), http.StatusBadRequest, model.InvalidParameterError{
			Parameter: "id",
		}.Error(), w)
		return
	}

	pin, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		h.errorResponse(r.Context(), http.StatusBadRequest, model.InvalidParameterError{
			Parameter: "pin",
		}.Error(), w)
		return
	}

	league, err := h.service.GetLeagueTable(r.Context(), pin)
	switch {
	case err == nil:
		if err := json.NewEncoder(w).Encode(league); err != nil {
			log.Error(r.Context(), "encode_response_error", log.ErrorParam(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case errors.Is(err, model.ErrLeagueNotFound):
		h.errorResponse(r.Context(), http.StatusNotFound, "league not found", w)
	case isInvalidParameterErr(err):
		h.errorResponse(r.Context(), http.StatusBadRequest, err.Error(), w)
	default:
		log.Error(r.Context(), "error_getting_league_table", log.ErrorParam(err))
		h.errorResponse(r.Context(), http.StatusInternalServerError, "could not get league table", w)
	}
}

func (h *handler) getOverallTable(w http.ResponseWriter, r *http.Request) {
	league, err := h.service.GetOverallLeagueTable(r.Context())
	switch {
	case err == nil:
		if err := json.NewEncoder(w).Encode(league); err != nil {
			log.Error(r.Context(), "encode_response_error", log.ErrorParam(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		log.Error(r.Context(), "error_getting_overall_table", log.ErrorParam(err))
		h.errorResponse(r.Context(), http.StatusInternalServerError, "could not get overall table", w)
	}
}

func (h *handler) rebuild(w http.ResponseWriter, r *http.Request) {
	if err := h.service.RebuildOverallLeagueTable(r.Context()); err != nil {
		log.Error(r.Context(), "error_rebuilding_overall_table", log.ErrorParam(err))
		h.errorResponse(r.Context(), http.StatusInternalServerError, "could not rebuild overall league table", w)
	}
}

func (h *handler) errorResponse(ctx context.Context, status int, message string, w http.ResponseWriter) {
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

func (h *handler) IsUnauthenticatedEndpoint(string) bool {
	return false
}

func (h *handler) GetRequestAudience(r *http.Request) string {
	p, err := mux.CurrentRoute(r).GetPathTemplate()
	if err != nil {
		return ""
	}

	switch p {
	case "/rebuild":
		return "user-updater"
	case "/{id}":
		return mux.Vars(r)["id"]
	case "/leave":
		fallthrough
	case "/join":
		fallthrough
	case "/":
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
