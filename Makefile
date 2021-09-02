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
	GOBIN=$(LOCAL_BIN) go get -u github.com/pressly/goose/v3/cmd/goose

migration: # make migration name=create_user_table
	$(LOCAL_BIN)/goose -dir=db/migrations create $(name) sql
	$(LOCAL_BIN)/goose -dir=db/migrations fix

migrate:
	$(LOCAL_BIN)/goose -dir=db/migrations postgres "user=account password=secret port=54321  dbname=account sslmode=disable" up

migrate-down:
	$(LOCAL_BIN)/goose -dir=db/migrations postgres "user=account password=secret port=54321  dbname=account sslmode=disable" down

up:
	docker-compose up -d

down:
	docker-compose down
