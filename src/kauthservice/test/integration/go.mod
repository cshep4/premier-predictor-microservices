module github.com/cshep4/premier-predictor-microservices/src/kauthservice/test/integration

require (
	github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.27.1
)

go 1.14

replace github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen => ../../../../proto-gen/model/gen
