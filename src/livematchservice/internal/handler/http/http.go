package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/metadata"
)

type (
	ServerError struct {
		Message string `json:"message"`
	}

	router struct {
		service handler.Servicer
	}
)

func New(service handler.Servicer) (*router, error) {
	if service == nil {
		return nil, errors.New("service_is_nil")
	}

	return &router{
		service: service,
	}, nil
}

func (h *router) Route(router *mux.Router) {
	router.HandleFunc("/upcoming", h.getUpcomingMatches).
		Methods(http.MethodGet)
	router.HandleFunc("/match/{matchId}/user/{userId}", h.getMatchSummary).
		Methods(http.MethodGet)
}

func (h *router) getUpcomingMatches(w http.ResponseWriter, r *http.Request) {
	upcomingMatches, err := h.service.GetUpcomingMatches()

	h.sendResponse(r.Context(), upcomingMatches, err, w)
}

func (h *router) getMatchSummary(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	ctx := metadata.NewIncomingContext(r.Context(), metadata.MD{"token": []string{token}})

	matchId := mux.Vars(r)["matchId"]
	userId := mux.Vars(r)["userId"]

	req := model.PredictionRequest{
		UserId:  userId,
		MatchId: matchId,
	}

	league, err := h.service.GetMatchSummary(ctx, req)
	h.sendResponse(r.Context(), league, err, w)
}

func (h *router) sendResponse(ctx context.Context, data interface{}, err error, w http.ResponseWriter) {
	switch {
	case err == nil:
		if data != nil {
			json.NewEncoder(w).Encode(data)
			return
		}
		w.WriteHeader(http.StatusOK)
	case errors.Is(err, model.ErrMatchNotFound):
		log.Debug(ctx, "match_not_found", log.ErrorParam(err))
		w.WriteHeader(http.StatusNotFound)
	case errors.Is(err, common.ErrInvalidRequestData):
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ServerError{
			Message: errors.Unwrap(err).Error(),
		})
	default:
		log.Error(ctx, "server_error", log.ErrorParam(err))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ServerError{
			Message: errors.Unwrap(err).Error(),
		})
	}
}
