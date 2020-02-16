package http

import (
	"encoding/json"
	"errors"
	common "github.com/cshep4/premier-predictor-microservices/src/common/model"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/handler"
	"github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/model"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type router struct {
	service handler.Servicer
}

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

	fixturePredictions, err := h.service.GetFixturesWithPredictions(id)

	h.sendResponse(fixturePredictions, err, w)
}

func (h *router) getPredictorData(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	predictorData, err := h.service.GetPredictorData(id)

	h.sendResponse(predictorData, err, w)
}

func (h *router) getUsersPastPredictions(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	predictions, err := h.service.GetUsersPastPredictions(id)

	h.sendResponse(predictions, err, w)
}

func (h *router) updatePredictions(w http.ResponseWriter, r *http.Request) {
	var predictions []common.Prediction

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("cannot read request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, &predictions); err != nil {
		log.Println("cannot decode request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.service.UpdatePredictions(predictions); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *router) getPrediction(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["userId"]
	matchId := mux.Vars(r)["matchId"]

	predictions, err := h.service.GetPrediction(userId, matchId)

	h.sendResponse(predictions, err, w)
}

func (h *router) sendResponse(data interface{}, err error, w http.ResponseWriter) {
	if err == model.ErrPredictionNotFound {
		log.Println("prediction not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
