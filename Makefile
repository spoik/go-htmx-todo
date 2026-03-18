ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: up
up: ## Start the services with hot-reloading
	docker compose up --build -d

.PHONY: down
down: ## Stop and remove the services
	docker compose down

.PHONY: logs
logs: ## Follow the logs from the services
	docker compose logs -f

.PHONY: shell
shell: ## Open a shell inside the app container
	docker compose exec app sh

.PHONY: db-shell
db-shell: ## Open a psql shell in the database container
	docker compose exec db psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

.PHONY: generate
generate: ## Run templ and sqlc generation inside the app container
	docker compose exec app templ generate
	docker compose exec app sqlc generate

.PHONY: migrate-up
migrate-up: ## Run database migrations inside the app container
	docker compose exec app migrate -path migrations -database "$(DATABASE_URL)" up

.PHONY: migrate-down
migrate-down: ## Rollback database migrations inside the app container
	docker compose exec app migrate -path migrations -database "$(DATABASE_URL)" down

.PHONY: test
test: ## Run tests inside the app container
	docker compose exec app go test ./...
