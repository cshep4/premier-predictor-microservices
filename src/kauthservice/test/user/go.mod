module github.com/cshep4/premier-predictor-microservices/src/kauthservice/test/user

require (
	github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.3.3
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	google.golang.org/grpc v1.27.1
)

go 1.14

replace github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen => ../../../../proto-gen/model/gen
