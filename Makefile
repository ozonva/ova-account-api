.PHONY: build, run, lint

default: build

build:
	go build -o ./bin/service ./cmd/ova-account-api

run:
	go run ./cmd/ova-account-api

lint:
	$(info ******************** running lint tools ********************)
	golangci-lint run -v
