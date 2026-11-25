# Цели, которые не являются файлами: Make выполняет их всегда.
.PHONY: help build run serve test fmt vet lint tidy docker-up docker-down

APP_NAME := highload-users-service
CMD_PATH := ./cmd/api
COMPOSE_FILE := deploy/docker-compose.yml

# gci установлен в GOPATH/bin, которого может не быть в PATH.
GCI := $(shell go env GOPATH)/bin/gci

# Показывает список доступных команд.
help:
	@echo "Доступные команды:"
	@echo "  make build       - собрать бинарник в bin/"
	@echo "  make run         - запустить сервис из исходников"
	@echo "  make serve       - собрать и запустить бинарник"
	@echo "  make test        - запустить тесты с race detector"
	@echo "  make fmt         - отформатировать код через gofumpt"
	@echo "  make vet         - проверить код через go vet"
	@echo "  make lint        - форматирование, импорты и golangci-lint"
	@echo "  make tidy        - обновить go.mod и go.sum"
	@echo "  make docker-up   - поднять стек в Docker"
	@echo "  make docker-down - остановить стек"

# Собирает статический бинарник в bin/. Каталог создаётся, если его нет.
build:
	mkdir -p bin
	CGO_ENABLED=0 go build -trimpath -o bin/$(APP_NAME) $(CMD_PATH)

# Запускает сервис из исходников, без отдельного бинарника.
run:
	go run $(CMD_PATH)

# Собирает и запускает готовый бинарник.
serve: build
	./bin/$(APP_NAME)

# Запускает тесты с детектором гонок; -count=1 отключает кэш результатов.
test:
	go test ./... -race -count=1

# Быстрое форматирование: gofumpt строже gofmt, -w пишет изменения в файлы.
fmt:
	gofumpt -w .

# Проверяет подозрительные конструкции, которые компилятор пропускает.
vet:
	go vet ./...

# Полная проверка перед коммитом: tidy, форматирование, импорты и автофиксы линтеров.
lint: tidy fmt
	$(GCI) write . --skip-generated -s standard -s default
	golangci-lint run --fix

# Удаляет неиспользуемые зависимости и добавляет недостающие.
tidy:
	go mod tidy

# Поднимает сервис и MinIO; файл compose появится на шаге контейнеризации.
docker-up:
	docker compose -f $(COMPOSE_FILE) up -d --build

# Останавливает контейнеры и сеть проекта.
docker-down:
	docker compose -f $(COMPOSE_FILE) down