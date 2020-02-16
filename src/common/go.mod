module github.com/cshep4/premier-predictor-microservices/src/common

go 1.14

replace github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen => ../../proto-gen/model/gen

require (
	cloud.google.com/go v0.53.0
	contrib.go.opencensus.io/exporter/stackdriver v0.13.0
	github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.4.0
	github.com/golang/protobuf v1.3.3
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/rs/cors v1.7.0
	github.com/stretchr/testify v1.4.0
	go.mongodb.org/mongo-driver v1.3.0
	go.opencensus.io v0.22.3
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	google.golang.org/api v0.17.0 // indirect
	google.golang.org/grpc v1.27.1
)