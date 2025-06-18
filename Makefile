.PHONY: help dev prod down clean logs build test

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

dev: ## Start development environment with hot reload
	docker compose -f docker-compose.dev.yaml -f docker-compose.override.yaml up --build

dev-detached: ## Start development environment in background
	docker compose -f docker-compose.dev.yaml -f docker-compose.override.yaml up -d --build

prod: ## Start production environment
	docker compose -f docker-compose.prod.yaml up -d --build

down: ## Stop all services
	docker compose -f docker-compose.dev.yaml -f docker-compose.override.yaml down
	docker compose -f docker-compose.prod.yaml down

clean: ## Remove all containers, volumes, and images
	docker compose -f docker-compose.dev.yaml -f docker-compose.override.yaml down -v --rmi all
	docker compose -f docker-compose.prod.yaml down -v --rmi all

logs: ## Show logs for development environment
	docker compose -f docker-compose.dev.yaml -f docker-compose.override.yaml logs -f

logs-prod: ## Show logs for production environment
	docker compose -f docker-compose.prod.yaml logs -f

build: ## Build Go application
	go build -o bin/main cmd/main.go

test: ## Run tests
	go test ./...

db-dev: ## Start only databases for development
	docker compose -f docker-compose.dev.yaml up -d pgsql redis mongodb

db-prod: ## Start only databases for production
	docker compose -f docker-compose.prod.yaml up -d pgsql redis mongodb

# Development helpers
dev-app: ## Run only the app in development mode
	docker compose -f docker-compose.dev.yaml -f docker-compose.override.yaml up app

dev-tools: ## Start development tools (adminer, redis-commander, mongo-express)
	docker compose -f docker-compose.dev.yaml -f docker-compose.override.yaml up -d adminer redis-commander mongo-express