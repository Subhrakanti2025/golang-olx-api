.PHONY: build run migrate-build migrate-run-up migrate-run-down

build:
	@go build -o bin/api ./cmd/api

run: build
	@./bin/api

migrate-build:
	@go build -o bin/migrate ./cmd/migrate

migrate-run-up: migrate-build
	@./bin/migrate up

migrate-run-down: migrate-build
	@./bin/migrate down

migrate-force-0:
	@go run ./cmd/migrate force 0