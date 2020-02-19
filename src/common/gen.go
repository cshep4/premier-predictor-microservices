package common

//go:generate mockgen -destination=internal/mocks/auth/auth_client.gen.go -package=auth_mocks github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen AuthServiceClient
//go:generate mockgen -destination=internal/mocks/notification/notification_client.gen.go -package=notification_mocks github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen NotificationServiceClient
