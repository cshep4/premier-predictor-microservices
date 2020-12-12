FROM golang:1.14-alpine as build

ENV PROJECT github.com/cshep4/premier-predictor-microservices/src/userservice

WORKDIR /build

COPY proto-gen/model/gen github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen

COPY src/common github.com/cshep4/premier-predictor-microservices/src/common

COPY src/userservice $PROJECT

WORKDIR $PROJECT

RUN go build -mod mod -gcflags='-N -l' -ldflags "-s -w" -o /userservice ./cmd/user

FROM alpine as release

COPY --from=build /userservice /userservice

RUN ls

EXPOSE 3005 8085
ENTRYPOINT ["/userservice"]
