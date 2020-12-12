FROM golang:1.13-alpine as builder

ENV PROJECT github.com/cshep4/premier-predictor-microservices/src/chatservice
WORKDIR /go/src/$PROJECT

COPY . .

RUN apk add --no-cache ca-certificates && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -mod vendor -gcflags='-N -l' -o /chatservice .

FROM alpine as release

COPY --from=builder /chatservice /chatservice

EXPOSE 3003
ENTRYPOINT ["/chatservice"]