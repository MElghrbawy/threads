# Load environment variables from .env file (cross-platform support)
include .env

# Go build command
GO_BUILD := go build -v

# Go run command
GO_RUN := go run

# Migrate command
MIGRATE_CMD := $(GO_RUN) cmd/migrate/main.go

export MIGRATION_DIR= internal/db/migrations

.PHONY: build
build:
	$(GO_BUILD) -o bin/server cmd/server/main.go
	$(GO_BUILD) -o bin/migrate cmd/migrate/main.go

.PHONY: run
run:
	$(GO_RUN) cmd/server/main.go

.PHONY: migrate-up
migrate-up:
	$(MIGRATE_CMD) up

.PHONY: migrate-down
migrate-down:
	$(MIGRATE_CMD) down 1

.PHONY: migrate-create
migrate-create:
	$(MIGRATE_CMD) create $(name) sql

.PHONY: migrate-status
migrate-status:
	$(MIGRATE_CMD) status

.PHONY: migrate-reset
migrate-reset:
	$(MIGRATE_CMD) reset

.PHONY: sqlc
sqlc:
	@echo "Generating sqlc..."
	@sqlc generate


.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build              - Build the server and migrate binaries"
	@echo "  make run                - Run the server"
	@echo "  make migrate-up         - Apply all available migrations"
	@echo "  make migrate-down       - Rollback the last applied migration"
	@echo "  make migrate-create     - Create a new migration file"
	@echo "  make migrate-status     - Show current migration status"
	@echo "  make migrate-reset      - Rollback all migrations"
	@echo "  make help               - Show this help message"
	@echo " make sqlc               - Generate sqlc"