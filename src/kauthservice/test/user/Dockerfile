FROM golang:1.14-alpine

WORKDIR /build

COPY proto-gen/model/gen github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen

COPY src/kauthservice/test/user github.com/cshep4/premier-predictor-microservices/src/kauthservice/test/user

WORKDIR github.com/cshep4/premier-predictor-microservices/src/kauthservice/test/user

RUN go build -mod=mod -o /app .

ENTRYPOINT ["/app"]