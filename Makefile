LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user_api
	protoc --proto_path api/user_api \
	--go_out=pkg/user_api --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_api --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_api/user_api.proto

build:
	GOOS=linux GOARCH=amd64 go build -o auth cmd/grpc_server/main.go

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t ${REGISTRY}/auth:v0.0.1 .
	aws ecr get-login-password --region eu-north-1 | docker login --username AWS --password-stdin ${REGISTRY}
	docker push ${REGISTRY}
