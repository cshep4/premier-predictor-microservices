package leagueservice

//go:generate mockgen -destination=internal/mocks/service/mock_http_service.gen.go -package=service_mock github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/handler/http Service
//go:generate mockgen -destination=internal/mocks/service/mock_gprc_service.gen.go -package=service_mock github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/handler/grpc Service
//go:generate mockgen -destination=internal/mocks/store/mock_store.gen.go -package=store_mock github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/service Store
//go:generate mockgen -destination=internal/mocks/user/mock_user_service.gen.go -package=user_mock github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/service UserService
//go:generate mockgen -destination=internal/mocks/table/mock_table.gen.go -package=table_mock github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/service LeagueTable
//go:generate mockgen -destination=internal/mocks/time/mock_time.gen.go -package=time_mock github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/service Timer
//go:generate mockgen -destination=internal/mocks/token/mock_generator.gen.go -package=token_mock github.com/cshep4/premier-predictor-microservices/src/leagueservice/internal/user TokenGenerator

//go:generate mockgen -destination=internal/mocks/user/mock_user_client.gen.go -package=user_mock github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen UserServiceClient
//go:generate mockgen -destination=internal/mocks/auth/mock_auth_client.gen.go -package=auth_mock github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen AuthServiceClient
