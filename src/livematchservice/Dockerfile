FROM golang:1.15 as build

ENV GONOSUMDB "github.com/cshep4"
ENV GOPRIVATE "github.com/cshep4"

WORKDIR /build

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags='-N -l' -ldflags "-s -w" -o /livematchservice .

FROM alpine as release

COPY --from=build /livematchservice /livematchservice

EXPOSE 3008 8088
ENTRYPOINT ["/livematchservice"]