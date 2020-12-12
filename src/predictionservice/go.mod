module github.com/cshep4/premier-predictor-microservices/src/predictionservice

require (
	github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen v0.0.0-00010101000000-000000000000
	github.com/cshep4/premier-predictor-microservices/src/common v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.1
	github.com/gorilla/mux v1.7.3
	github.com/stretchr/testify v1.4.0
	go.mongodb.org/mongo-driver v1.3.0
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/tools v0.0.0-20200214225126-5916a50871fb // indirect
	google.golang.org/grpc v1.27.1
)

go 1.13

replace github.com/cshep4/premier-predictor-microservices/src/common => ../common

replace github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen => ../../proto-gen/model/gen
