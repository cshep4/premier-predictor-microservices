FROM golang:1.14-alpine as build

ENV PROJECT github.com/cshep4/premier-predictor-microservices/src/userservice
WORKDIR /go/src/$PROJECT

COPY . .

RUN go build -mod vendor -gcflags='-N -l' -ldflags "-s -w" -o /userservice ./cmd/user

ENTRYPOINT ["/userservice"]

FROM gcr.io/distroless/base-debian10

COPY --from=build /userservice /userservice

EXPOSE 3005 8085
ENTRYPOINT ["/userservice"]
