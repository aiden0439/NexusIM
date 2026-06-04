# im-server

Go IM backend skeleton for future WebSocket single chat, Kafka message pipeline, Redis sequence allocation, MySQL persistence, ACK, and offline message compensation.

## Tech Stack

- Go 1.21+
- Gin
- GORM
- MySQL 8
- Redis 7
- Kafka + Zookeeper
- Viper
- zap
- Docker Compose

## Quick Start

```bash
docker compose up -d
go mod tidy
go run ./cmd/server
```

Health check:

```bash
curl http://localhost:8080/ping
```

## Project Layout

```text
cmd/server          application entry
config              YAML configuration
internal/handler    HTTP and WebSocket handlers
internal/service    business service layer
internal/repo       persistence repositories
internal/model      database models
internal/ws         WebSocket manager and client skeleton
internal/router     Gin router wiring
pkg                 shared infrastructure packages
scripts             SQL initialization
docs                architecture notes
```
