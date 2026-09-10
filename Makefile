-include .env
export

MIGRATIONS_DIR := migrations

.PHONY: help check-env infra-up infra-down infra-logs migrate-up migrate-down migrate-version migrate-create run dev test fmt tidy

help:
	@echo "Available targets:"
	@echo "  infra-up             Start infrastructure and wait until it is ready"
	@echo "  infra-down           Stop infrastructure without deleting data"
	@echo "  infra-logs           Follow infrastructure logs"
	@echo "  migrate-up           Apply all pending migrations"
	@echo "  migrate-down         Roll back one migration"
	@echo "  migrate-version      Show the current migration version"
	@echo "  migrate-create name= Create a sequential SQL migration"
	@echo "  run                  Run the Go server"
	@echo "  dev                  Start infrastructure, migrate, and run the server"
	@echo "  test                 Run tests"
	@echo "  fmt                  Format Go source files"
	@echo "  tidy                 Tidy Go modules"

check-env:
	@test -f .env || (echo "missing .env; copy .env.example to .env" && exit 1)
	@test -n "$(MYSQL_ROOT_PASSWORD)" || (echo "MYSQL_ROOT_PASSWORD is required in .env" && exit 1)
	@test -n "$(MYSQL_DATABASE)" || (echo "MYSQL_DATABASE is required in .env" && exit 1)
	@test -n "$(MYSQL_USER)" || (echo "MYSQL_USER is required in .env" && exit 1)
	@test -n "$(MYSQL_PASSWORD)" || (echo "MYSQL_PASSWORD is required in .env" && exit 1)
	@test -n "$(MYSQL_DSN)" || (echo "MYSQL_DSN is required in .env" && exit 1)
	@test -n "$(MIGRATE_DATABASE_URL)" || (echo "MIGRATE_DATABASE_URL is required in .env" && exit 1)

infra-up: check-env
	docker compose up -d --wait

infra-down:
	docker compose down

infra-logs:
	docker compose logs -f

migrate-up: check-env
	@echo "Applying pending database migrations..."
	@migrate -path $(MIGRATIONS_DIR) -database "$(MIGRATE_DATABASE_URL)" up

migrate-down: check-env
	@echo "Rolling back one database migration..."
	@migrate -path $(MIGRATIONS_DIR) -database "$(MIGRATE_DATABASE_URL)" down 1

migrate-version: check-env
	@migrate -path $(MIGRATIONS_DIR) -database "$(MIGRATE_DATABASE_URL)" version

migrate-create:
	@test -n "$(name)" || (echo "usage: make migrate-create name=add_something" && exit 1)
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq "$(name)"

run: check-env
	go run ./cmd/server

dev: infra-up
	@$(MAKE) migrate-up
	@$(MAKE) run

test:
	go test ./...

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

tidy:
	go mod tidy
