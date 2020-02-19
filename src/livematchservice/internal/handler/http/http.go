package http

import (
	"encoding/json"
	"log"
	"net/http"

	m "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/common/util"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/model"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
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

	h.sendResponse(upcomingMatches, err, w)
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
	h.sendResponse(league, err, w)
}

func (h *router) sendResponse(data interface{}, err error, w http.ResponseWriter) {
	switch {
	case err == nil:
		w.WriteHeader(http.StatusOK)
		if data != nil {
			_ = json.NewEncoder(w).Encode(data)
			return
		}
		return

	case err == model.ErrMatchNotFound:
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(ServerError{
			Message: err.Error(),
		})

	case errors.Cause(err) == m.ErrInvalidRequestData:
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(ServerError{
			Message: util.GetErrorMessage(err),
		})

	default:
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(ServerError{
			Message: err.Error(),
		})
	}

	log.Println(err.Error())
}
