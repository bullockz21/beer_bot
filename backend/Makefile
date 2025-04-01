# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
include .env
export $(shell sed 's/=.*//' .env)
endif

# Project variables
BINARY_NAME := beer_bot
MIGRATIONS_DIR := ./migrations
GOOSE_DRIVER := postgres
DOCKER_COMPOSE_FILE := docker-compose.yml

# Go build flags
LDFLAGS := -w -s

.PHONY: all build run test clean migrate-up migrate-down docker-up docker-down lint format deps help

all: build

## Build application binary
build:
	@echo "Building binary..."
	@go build -ldflags="$(LDFLAGS)" -o $(BINARY_NAME) ./cmd

## Build and run application
run: build
	@echo "Starting application..."
	@./$(BINARY_NAME)

## Run tests
test:
	@echo "Running tests..."
	@go test -v -race ./...

## Clean build artifacts
clean:
	@echo "Cleaning..."
	@go clean
	@rm -f $(BINARY_NAME)

## Apply database migrations
migrate-up:
	@echo "Applying migrations..."
	@goose -dir $(MIGRATIONS_DIR) $(GOOSE_DRIVER) "$(DB_DSN)" up

## Rollback last migration
migrate-down:
	@echo "Rolling back migrations..."
	@goose -dir $(MIGRATIONS_DIR) $(GOOSE_DRIVER) "$(DB_DSN)" down

## Start Docker containers
docker-up:
	@echo "Starting Docker services..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

## Stop Docker containers
docker-down:
	@echo "Stopping Docker services..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down

## Install development dependencies
deps:
	@echo "Installing tools..."
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@go install github.com/volatiletech/sqlboiler/v4@latest
	@go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

## Run linters
lint:
	@echo "Running linters..."
	@golangci-lint run ./...

## Format source code
format:
	@echo "Formatting code..."
	@gofmt -s -w .

## Show help
help:
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help