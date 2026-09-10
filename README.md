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

Create the local environment file first:

```bash
cp .env.example .env
```

Update the passwords in `.env`, then start the infrastructure, apply migrations,
and run the server:

```bash
make dev
```

Health check:

```bash
curl http://localhost:8080/ping
```

`make infra-down` stops the containers without deleting their volumes.

## Database Migrations

Database schema changes are managed by `golang-migrate`. The commonly used
commands are:

```bash
make migrate-up
make migrate-down
make migrate-version
make migrate-create name=add_message_metadata
```

Do not edit an applied migration. Create a new migration for every later schema
change. The application uses `MYSQL_DSN`, while the migrate CLI uses
`MIGRATE_DATABASE_URL`; both are defined in the ignored local `.env` file.

If a MySQL volume was created before migrations were introduced, its old tables
remain present. Back up important local data before deciding whether to reuse or
recreate that volume.

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
migrations          versioned database migrations
docs                architecture notes
```
