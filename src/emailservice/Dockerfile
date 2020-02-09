FROM golang:1.13-alpine as builder

ENV PROJECT github.com/cshep4/premier-predictor-microservices/src/emailservice
WORKDIR /go/src/$PROJECT

COPY . .

RUN apk add --no-cache ca-certificates && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -mod vendor -gcflags='-N -l' -o /emailservice .

FROM alpine as release

COPY --from=builder /emailservice /emailservice

EXPOSE 3011
ENTRYPOINT ["/emailservice"]