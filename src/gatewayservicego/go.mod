module github.com/cshep4/premier-predictor-microservices/src/gatewayservice

go 1.12

require (
	github.com/99designs/gqlgen v0.9.3
	github.com/agnivade/levenshtein v1.0.2 // indirect
	github.com/cshep4/premier-predictor-microservices v0.0.0-20190810105329-1655ffce5253
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/gorilla/handlers v1.4.2 // indirect
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/websocket v1.4.1
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/stretchr/testify v1.3.0
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859 // indirect
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys v0.0.0-20190606122018-79a91cf218c4 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/grpc v1.22.1
)

replace github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen => ../../proto-gen/model/gen

replace github.com/cshep4/premier-predictor-microservices/src/common => ../common
