PROJECT_NAME = premier-predictor-microservices

_gendart:
	@mkdir -p model/gen/ship/dart
	@protoc -I=model/proto --plugin=/Users/chrisshepherd/.pub-cache/bin/protoc-gen-dart --dart_out=grpc:model/gen/ship/dart model/proto/*.proto

_gengo:
	@mkdir -p model/gen
	@protoc -I=model/proto --go_out=plugins=grpc:model/gen model/proto/*.proto

_copyjava:
	@cp -R model/proto/auth.proto model/proto/notification.proto model/proto/request.proto ../src/notificationservice/proto
	@cp -R model/proto/auth.proto model/proto/user.proto  ../src/kauthservice/src/main/proto

_copyts:
	@cp -R model/proto/auth.proto model/proto/request.proto model/proto/fixture.proto ../src/fixtureservice/proto

gen: _gengo _copyjava _copyts

build: get gen
	@env CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags '-w -extldflags "-static"' -o build/${PROJECT_NAME}_linux_amd64 .
	@env GOARCH=amd64 go build -ldflags '-w -extldflags "-static"' -o build/${PROJECT_NAME}_macosx_amd64 .

get:
	@go get -u github.com/golang/dep/cmd/dep
	@dep ensure

install: get gen
	@cp config_template.json config.json
