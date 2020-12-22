module github.com/cshep4/premier-predictor-microservices/src/common

go 1.15

require (
	cloud.google.com/go v0.74.0
	contrib.go.opencensus.io/exporter/stackdriver v0.13.4
	github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen v0.0.1
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/palantir/witchcraft-go-logging v1.9.0
	github.com/rs/cors v1.7.0
	github.com/stretchr/testify v1.6.1
	go.mongodb.org/mongo-driver v1.4.4
	go.opencensus.io v0.22.5
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/grpc v1.34.0
	google.golang.org/protobuf v1.25.0
)
