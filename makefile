.PHONY: default run build docs clean
# Variables
APP_NAME=server
MIGRATE_NAME=migration

# Tasks
default: run

fmt:
	go fmt ./...
vet: fmt
	go vet ./...
docs:
	@swag init -g cmd/main.go
run: docs
	@go run migrate/migrate.go
	@go run cmd/main.go
build: vet
	@go build -o $(MIGRATE_NAME) migrate/migrate.go
	@go build -o $(APP_NAME) cmd/main.go
clean:
	@rm -f $(APP_NAME)
	@rm -f $(MIGRATE_NAME)
	@rm -rf ./docs
