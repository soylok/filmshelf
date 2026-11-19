include .env
export

export PROJECT_ROOT=$(shell pwd)

run-app:
	@go run $(PROJECT_ROOT)/cmd/filmshelf/main.go

env-up:
	@docker compose up -d env-postgres

env-down:
	@docker compose down

env-clean:
	@printf 'Остановить контейнеры и очистить данные бд? (Y/n): '; \
	if read -r answer && [ "$$answer" = "Y" ]; then \
		docker compose down --volumes; \
	else \
		printf 'Удаление отменено.\n'; \
	fi

migrate: NAME ?= default
migrate:
	@docker compose run --rm migration \
		create \
		-ext sql \
		-dir /migrations \
		-seq $(NAME)

migrate-up:
	@make migrate-do DO=up

migrate-down:
	@make migrate-do DO=down

migrate-do:
	@docker compose run --rm migration \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@env-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		$(DO)

logs-clean:
	@rm -rf $(PROJECT_ROOT)/logs
