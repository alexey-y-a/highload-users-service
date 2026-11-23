.PHONY: help build run test fmt lint tidy docker-up docker-down

APP_NAME := highload-users-service
CMD_PATH := ./cmd/api

help:
	@echo "Доступные команды:"
	@echo "  make build      - собрать бинарник в bin/"
	@echo "  make run        - запустить сервис локально"
	@echo "  make test       - запустить тесты с race detector"
	@echo "  make fmt        - отформатировать код"
	@echo "  make lint       - запустить golangci-lint"
	@echo "  make tidy       - обновить go.mod и go.sum"
	@echo "  make docker-up  - поднять стек в Docker"
	@echo "  make docker-down - остановить стек"

build:
	go build -o bin/$(APP_NAME) $(CMD_PATH)

run:
	go run $(CMD_PATH)

test:
	go test ./... -race -count=1

fmt:
	gofmt -w .

lint:
	golangci-lint run

tidy:
	go mod tidy

docker-up:
	docker compose up -d --build

docker-down:
	docker compose down