# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.12-alpine as build

# Project path
ENV PACKAGENAME 'gatewayservice'
ENV PACKAGE github.com/cshep4/premier-predictor-microservices/src/${PACKAGENAME}

# Environment Variables
ENV GOPATH /go
ENV GOPROJECT ${GOPATH}/src/${PACKAGE}

# Add go files to the project
RUN mkdir -p ${GOPROJECT}
ADD . ${GOPROJECT}

# GO Project directory
WORKDIR ${GOPROJECT}/cmd/gatewayservice

# Build application with custom ldflags
RUN apk update && \
    apk --no-cache add ca-certificates && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build -mod vendor -ldflags "-s -w" -o /app

ENTRYPOINT ["/app"]

# Build Alpine Image
FROM scratch

# User bin directory
WORKDIR /usr/bin/

# Copy certificates
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy executable to alpine container
COPY --from=build /app .

ENTRYPOINT ["./app"]