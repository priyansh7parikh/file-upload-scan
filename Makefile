# =========================
# Project metadata
# =========================
APP_NAME := file-upload-sca
CMD_DIR := cmd
MAIN_FILE := $(CMD_DIR)/main.go
SWAGGER_ENTRY := cmd/main.go

GO := go
SWAG := swag

# =========================
# Default target
# =========================
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make deps        - Download dependencies"
	@echo "  make swagger     - Generate Swagger docs"
	@echo "  make run         - Run the server"
	@echo "  make test        - Run all tests"
	@echo "  make test-cover  - Run tests with coverage"
	@echo "  make fmt         - Format Go code"
	@echo "  make vet         - Run go vet"
	@echo "  make clean       - Clean generated files"

# =========================
# Dependencies
# =========================
.PHONY: deps
deps:
	$(GO) mod tidy
	$(GO) mod download

# =========================
# Swagger
# =========================
.PHONY: swagger
swagger:
	swag init -g cmd/main.go -o ./docs

# =========================
# Run server
# =========================
.PHONY: run
run:
	$(GO) run $(MAIN_FILE)

# =========================
# Tests
# =========================
.PHONY: test
test:
	$(GO) test ./...

.PHONY: test-cover
test-cover:
	$(GO) test ./... -coverprofile=coverage.out
	$(GO) tool cover -func=coverage.out

# =========================
# Code quality
# =========================
.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: vet
vet:
	$(GO) vet ./...

# =========================
# Cleanup
# =========================
.PHONY: clean
clean:
	rm -rf docs coverage.out
