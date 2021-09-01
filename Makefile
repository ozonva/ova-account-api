.PHONY: build, run, lint, test, generate, generate-grpc, bin-deps, up, down

LOCAL_BIN:=$(CURDIR)/bin

default: build

build:
	go build -o ./bin/service ./cmd/ova-account-api

run:
	go run ./cmd/ova-account-api

lint:
	$(info ******************** running lint tools ********************)
	golangci-lint run -v

test:
	go test ./...

generate:
	go generate ./...

generate-grpc:
	GOBIN=$(LOCAL_BIN) protoc \
		--proto_path=api/ \
        --go_out=pkg/ova-account-api --go_opt=paths=source_relative \
        --go-grpc_out=pkg/ova-account-api --go-grpc_opt=paths=source_relative \
        api/*.proto

bin-deps:
	GOBIN=$(LOCAL_BIN) go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/proto
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

up:
	docker-compose up -d

down:
	docker-compose down
