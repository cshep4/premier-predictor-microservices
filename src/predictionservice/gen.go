package main

//go:generate mockgen -destination=internal/mocks/fixture/mock_fixture.gen.go -package=fixture_mocks github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/service FixtureService
//go:generate mockgen -destination=internal/mocks/prediction/mock_store.gen.go -package=prediction_mocks github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/service Store
//go:generate mockgen -destination=internal/mocks/prediction/mock_service.gen.go -package=prediction_mocks github.com/cshep4/premier-predictor-microservices/src/predictionservice/internal/handler Servicer
//go:generate mockgen -destination=internal/mocks/fixture/mock_fixture_service_client.gen.go -package=fixture_mocks github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen FixtureServiceClient
