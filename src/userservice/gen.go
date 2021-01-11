package userservice

//go:generate mockgen -destination=internal/mocks/service/mock_service.gen.go -package=service_mock github.com/cshep4/premier-predictor-microservices/src/userservice/internal/service Service
//go:generate mockgen -destination=internal/mocks/store/mock_store.gen.go -package=store_mock github.com/cshep4/premier-predictor-microservices/src/userservice/internal/store/mongo Store
