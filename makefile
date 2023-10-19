.PHONY: default run build test clean

APP_NAME=chi-rest-api

default: run

run:
	@go run ./cmd/server/main.go

build:
	@go build -o $(APP_NAME) ./cmd/server/main.go

test:
	@go test ./...

clean:
	@rm -f $(APP_NAME)