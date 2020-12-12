FROM golang:1.14-alpine as build

ENV PROJECT github.com/cshep4/premier-predictor-microservices/src/leagueservice

WORKDIR /build

COPY proto-gen/model/gen github.com/cshep4/premier-predictor-microservices/proto-gen/model/gen

COPY src/common github.com/cshep4/premier-predictor-microservices/src/common

COPY src/leagueservice $PROJECT

WORKDIR $PROJECT

RUN go build -mod mod -gcflags='-N -l' -ldflags "-s -w" -o /leagueservice ./cmd/league

FROM alpine as release

COPY --from=build /leagueservice /leagueservice

RUN ls

EXPOSE 3009 8089
ENTRYPOINT ["/leagueservice"]
