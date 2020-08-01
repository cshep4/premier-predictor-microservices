FROM golang:1.14

WORKDIR /build

COPY proto-gen/model/gen github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen

COPY src/kauthservice/test/integration github.com/cshep4/premier-predictor-microservices/src/kauthservice/test/integration

WORKDIR github.com/cshep4/premier-predictor-microservices/src/kauthservice/test/integration

ENTRYPOINT ["go", "test", "-v", "-mod", "mod", "-race", "-tags", "integration", "."]