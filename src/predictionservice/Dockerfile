FROM golang:1.13-alpine as builder

ENV PROJECT github.com/cshep4/premier-predictor-microservices/src/predictionservice

WORKDIR /build

COPY proto-gen/model/gen github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen

COPY src/common github.com/cshep4/premier-predictor-microservices/src/common

COPY src/predictionservice $PROJECT

WORKDIR $PROJECT

RUN apk add --no-cache ca-certificates && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -mod vendor -gcflags='-N -l' -o /predictionservice .

FROM alpine as release

COPY --from=builder /predictionservice /predictionservice
EXPOSE 3007 8087
ENTRYPOINT ["/predictionservice"]