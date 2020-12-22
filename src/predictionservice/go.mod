module github.com/cshep4/premier-predictor-microservices/src/predictionservice

require (
	github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen v0.0.1
	github.com/cshep4/premier-predictor-microservices/src/common v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.8.0
	github.com/kisielk/errcheck v1.2.0 // indirect
	github.com/stretchr/testify v1.6.1
	go.mongodb.org/mongo-driver v1.4.4
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	google.golang.org/grpc v1.34.0
)

go 1.13

replace github.com/cshep4/premier-predictor-microservices/src/common => ../common

replace github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen => ../../proto-gen/model/gen
