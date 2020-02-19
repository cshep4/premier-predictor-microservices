package main

//go:generate mockgen -destination=internal/mocks/prediction/mock_prediction.gen.go -package=prediction_mocks github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/service Predictor
//go:generate mockgen -destination=internal/mocks/live/mock_store.gen.go -package=live_mocks github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/service Store
//go:generate mockgen -destination=internal/mocks/live/mock_service.gen.go -package=live_mocks github.com/cshep4/premier-predictor-microservices/src/livematchservice/internal/handler Servicer
//go:generate mockgen -destination=internal/mocks/prediction/mock_prediction_service_client.gen.go -package=prediction_mocks github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen PredictionServiceClient
