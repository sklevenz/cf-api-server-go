SHELL := /usr/bin/env bash

APP_NAME := cf-api-server
BIN_DIR := bin
LOG_DIR := logs
DOCKER_IMAGE := $(APP_NAME):latest
GEN_OUTPUT=./internal/gen/api_gen.go

.PHONY: default all build generate test verbose run clean copy-openapi

default: help

all: clean generate test build

help:
	@echo "Available targets:"
	@echo "  build         Build the binary"
	@echo "  generate      Run code generation"
	@echo "  test          Run tests"
	@echo "  verbose       Run tests verbose"
	@echo "  run           Build and run the server (logs to ./$(LOG_DIR))"
	@echo "  clean         Remove binaries and logs"
	@echo "  copy-openapi  Copy openapi spec"


build: prepare-dirs
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/server

generate:
	@mkdir -p ./internal/gen
	@command -v oapi-codegen >/dev/null || (echo "Fehler: oapi-codegen fehlt. Installiere mit: go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest" && exit 1)
	@echo "Generiere API-Code..."
	@oapi-codegen --config=./cfg/oapi-codegen-config.yaml ./spec/openapi.yaml > $(GEN_OUTPUT)
	@echo "✅ API-Code erfolgreich generiert: $(GEN_OUTPUT)"

test:
	go test ./... 

verbose:
	go test ./... -v

run: prepare-dirs
	go run ./cmd/server/main.go

prepare-dirs:
	mkdir -p $(BIN_DIR) $(LOG_DIR)

clean:
	rm -rf $(BIN_DIR) $(LOG_DIR)

copy-openapi:
	@SRC="../cf-api-spec/gen/openapi.yaml"; \
	DST="./spec"; \
	if [ ! -f "$$SRC" ]; then \
		echo "OpenAPI spec not found at '$$SRC'"; \
		exit 1; \
	fi; \
	echo "📄 Copying OpenAPI spec from '$$SRC' to '$$DST'..."; \
	cp "$$SRC" "$$DST"; \
	echo "OpenAPI spec copied successfully."; \
	echo "To generate a new server, run: make generate"