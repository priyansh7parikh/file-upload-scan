APP_NAME := file-upload-scan
MAIN_FILE := cmd/main.go
SWAGGER_ENTRY := cmd/main.go

GO := go
SWAG := swag

.PHONY: help
help:
	@echo "make run          - Run app (no swagger)"
	@echo "make run-dev      - Run app with swagger"
	@echo "make swagger      - Generate swagger docs"
	@echo "make test         - Run tests"
	@echo "make clean        - Cleanup generated files"

.PHONY: run
run:
	$(GO) run $(MAIN_FILE)

.PHONY: run-dev
run-dev:
	$(GO) run -tags=dev $(MAIN_FILE)

.PHONY: swagger
swagger:
	$(SWAG) init -g $(SWAGGER_ENTRY) -o docs

.PHONY: test
test:
	$(GO) test ./...

.PHONY: clean
clean:
	rm -rf docs coverage.out
