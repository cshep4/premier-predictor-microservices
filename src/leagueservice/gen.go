package leagueservice

// Internal
//go:generate mockgen -destination=internal/mock/service/http/mock_service.gen.go -package=service_mock github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/handler/http Service
//go:generate mockgen -destination=internal/mock/service/grpc/mock_service.gen.go -package=service_mock github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/handler/grpc Service
//go:generate mockgen -destination=internal/mock/store/mock_store.gen.go -package=store_mock github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/service LeagueStore,UserStore
//go:generate mockgen -destination=internal/mock/time/mock_time.gen.go -package=time_mock github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/service Timer

// External
//go:generate mockgen -destination=internal/mock/auth/mock_auth_client.gen.go -package=auth_mock github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen AuthServiceClient
