FROM golang:1.15 as build

ENV GONOSUMDB "github.com/cshep4"
ENV GOPRIVATE "github.com/cshep4"

WORKDIR /build

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags='-N -l' -ldflags "-s -w" -o /leagueservice ./cmd/league

FROM alpine as release

COPY --from=build /leagueservice /leagueservice

EXPOSE 3009 8089
ENTRYPOINT ["/leagueservice"]
