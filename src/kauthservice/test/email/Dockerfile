FROM golang:1.14-alpine

WORKDIR /build

COPY . .

RUN go build -o /app ./test/email

ENTRYPOINT ["/app"]
