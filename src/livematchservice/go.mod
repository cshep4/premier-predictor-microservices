module github.com/cshep4/premier-predictor-microservices/src/livematchservice

go 1.13

require (
	github.com/ahl5esoft/golang-underscore v2.0.0+incompatible
	github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen v0.0.0-00010101000000-000000000000
	github.com/cshep4/premier-predictor-microservices/src/common v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.4.0
	github.com/golang/protobuf v1.3.3
	github.com/gorilla/mux v1.7.3
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0
	github.com/xdg/stringprep v1.0.1-0.20180714160509-73f8eece6fdc // indirect
	go.mongodb.org/mongo-driver v1.3.0
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	google.golang.org/grpc v1.27.1
)

replace github.com/cshep4/premier-predictor-microservices/src/common => ../common

replace github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen => ../../proto-gen/model/gen
