module github.com/cshep4/premier-predictor-microservices/src/userservice

go 1.15

require (
	github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen v0.0.1
	github.com/cshep4/premier-predictor-microservices/src/common v0.0.0-00010101000000-000000000000
	github.com/golang/mock v1.4.4
	github.com/gorilla/mux v1.8.0
	github.com/kisielk/errcheck v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	github.com/tidwall/pretty v1.0.1 // indirect
	github.com/xdg/stringprep v1.0.1-0.20180714160509-73f8eece6fdc // indirect
	go.mongodb.org/mongo-driver v1.4.4
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/grpc v1.34.0
)

replace github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen => ../../proto-gen/model/gen

replace github.com/cshep4/premier-predictor-microservices/src/common => ../common
