FROM golang:1.15 as build

ENV GONOSUMDB "github.com/cshep4"
ENV GOPRIVATE "github.com/cshep4"

ARG GO_GIT_URL="git@github.com:"
RUN git config --system url.${GO_GIT_URL}.insteadOf "https://github.com/"

#ENV PROJECT github.com/cshep4/premier-predictor-microservices/src/leagueservice

WORKDIR /build

#WORKDIR $PROJECT

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -mod mod -gcflags='-N -l' -ldflags "-s -w" -o /leagueservice ./cmd/league

FROM alpine as release

COPY --from=build /leagueservice /leagueservice

EXPOSE 3009 8089
ENTRYPOINT ["/leagueservice"]
