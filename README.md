# Highload Users Service

Микросервис на Go для управления пользователями под высокой нагрузкой.

## Статус

Функциональность реализуется поэтапно.
- Шаг 2: подготовлена инфраструктура проекта.
- Шаг 3: приложение собирается и отвечает на GET /healthz.
- Шаг 4: настройки сервера загружаются из .env, добавлен unit-тест сборки адреса.
- Шаг 5: добавлен структурированный логгер logrus с уровнем и форматом из конфига.
- Шаг 6: добавлены таймауты сервера и graceful shutdown по SIGINT и SIGTERM.
- Шаг 7: добавлена доменная модель User с валидацией имени и email.

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

make run
curl http://localhost:8080/healthz

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
    internal/logger/     обёртка над logrus

| Элемент из задания | Где в проекте |
|---|---|
| main.go | cmd/api/main.go |
| handlers/ | internal/handler/ |
| services/ | internal/service/ |
| models/ | internal/domain/ |
| utils/ | internal/middleware/ |
| metrics/ | internal/metrics/ |

## API

Будет описано после реализации обработчиков.
