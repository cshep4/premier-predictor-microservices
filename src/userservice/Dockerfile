FROM golang:1.15 as build

ENV GONOSUMDB "github.com/cshep4"
ENV GOPRIVATE "github.com/cshep4"

WORKDIR /build

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -gcflags='-N -l' -ldflags "-s -w" -o /userservice ./cmd/user

FROM alpine as release

COPY --from=build /userservice /userservice

EXPOSE 3005 8085
ENTRYPOINT ["/userservice"]
