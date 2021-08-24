.PHONY: build, run, lint, test, generate

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
