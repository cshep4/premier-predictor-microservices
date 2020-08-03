package http

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/cshep4/premier-predictor-microservices/src/common/log"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	"github.com/gorilla/mux"
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
	router.HandleFunc("/fixtures/predicted/{id}", h.getFixturesWithPredictions).
		Methods(http.MethodGet)
	router.HandleFunc("/predictions/{id}", h.getPredictorData).
		Methods(http.MethodGet)
	router.HandleFunc("/predictions", h.updatePredictions).
		Methods(http.MethodPost)
	router.HandleFunc("/predictions/summary/{id}", h.getUsersPastPredictions).
		Methods(http.MethodGet)
	router.HandleFunc("/predictions/{userId}/{matchId}", h.getPrediction).
		Methods(http.MethodGet)
}

func (h *router) getFixturesWithPredictions(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	fixturePredictions, err := h.service.GetFixturesWithPredictions(r.Context(), id)

	h.sendResponse(r.Context(), fixturePredictions, err, w)
}

func (h *router) getPredictorData(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	predictorData, err := h.service.GetPredictorData(r.Context(), id)

	h.sendResponse(r.Context(), predictorData, err, w)
}

func (h *router) getUsersPastPredictions(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	predictions, err := h.service.GetUsersPastPredictions(r.Context(), id)

	h.sendResponse(r.Context(), predictions, err, w)
}

func (h *router) updatePredictions(w http.ResponseWriter, r *http.Request) {
	var predictions []common.Prediction

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(r.Context(), "cannot_read_request", log.ErrorParam(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, &predictions); err != nil {
		log.Error(r.Context(), "cannot_decode_request", log.ErrorParam(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.service.UpdatePredictions(r.Context(), predictions); err != nil {
		log.Error(r.Context(), "error_updating_predictions", log.ErrorParam(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *router) getPrediction(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	matchId := mux.Vars(r)["matchId"]

	predictions, err := h.service.GetPrediction(r.Context(), userId, matchId)

	h.sendResponse(r.Context(), predictions, err, w)
}

func (h *router) sendResponse(ctx context.Context, data interface{}, err error, w http.ResponseWriter) {
	switch {
	case err == nil:
		if data != nil {
			json.NewEncoder(w).Encode(data)
			return
		}
		w.WriteHeader(http.StatusOK)
	case errors.Is(err, model.ErrPredictionNotFound):
		log.Debug(ctx, "prediction_not_found", log.ErrorParam(err))
		w.WriteHeader(http.StatusNotFound)
	default:
		log.Error(ctx, "server_error", log.ErrorParam(err))
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ServerError{
			Message: errors.Unwrap(err).Error(),
		})
	}
}

func (h *router) IsUnauthenticatedEndpoint(path string) bool {
	return false
}

func (h *router) GetRequestAudience(r *http.Request) string {
	return ""
}
