# Highload Users Service

Микросервис на Go для управления пользователями под высокой нагрузкой.

## Статус

Функциональность реализуется поэтапно.

## Что будет реализовано

- [ ] CRUD API: GET /api/users, GET /api/users/{id}, POST /api/users,
  PUT /api/users/{id}, DELETE /api/users/{id}
- [ ] Асинхронный аудит операций через goroutines и канал
- [ ] Метрики Prometheus: RPS, latency, ошибки, endpoint /metrics
- [ ] Rate limiting 1000 запросов в секунду
- [ ] Интеграция с MinIO: загрузка и выдача аватара пользователя
- [ ] Dockerfile и docker-compose
- [ ] Нагрузочное тестирование wrk

## Стек

- Go 1.22+
- gorilla/mux
- minio-go
- golang.org/x/time/rate
- prometheus/client_golang
- logrus
- Docker, Docker Compose

## Запуск

Будет добавлен на шаге контейнеризации.

## Структура

Структура дополняется по шагам. Планируемая основа:

    cmd/api/             точка входа
    internal/app/        сборка приложения и запуск сервера
    internal/config/     конфигурация
    internal/domain/     модель User и валидация
    internal/repository/ хранение пользователей
    internal/service/    бизнес-логика
    internal/handler/    HTTP-обработчики
    internal/middleware/ логирование, recover, метрики, rate limit
    internal/metrics/    Prometheus
    internal/audit/      асинхронный аудит
    internal/storage/    клиент MinIO

## API

Будет описано после реализации обработчиков.
