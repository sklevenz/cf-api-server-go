# cf-api-server-go

[![Test Status](https://github.com/sklevenz/cf-api-server-go/actions/workflows/test.yaml/badge.svg)](https://github.com/sklevenz/cf-api-server-go/actions)
[![Build Status](https://github.com/sklevenz/cf-api-server-go/actions/workflows/build.yaml/badge.svg)](https://github.com/sklevenz/cf-api-server-go/actions)

mplementation of the [cf-api-spec](https://github.com/sklevenz/cf-api-spec) in Go. This project provides a Go-based solution adhering to the Cloud Foundry API specifications, enabling seamless integration and usage for Cloud Foundry-related development.

# Project Directory Structure

This project is structured to facilitate maintainability, scalability, and clarity. Below is an overview of the directory structure:

```
cf-api-server/
├── script/              # Utility scripts for run, build, test, code generation ...
├── cmd/
│   └── server/
│       └── main.go       # Entry point of the application
├── internal/
│   ├── gen/              # Package for generated code by oapi-codegen
│   │   └── cfapi_gen.go  # Generated file by ./script/generate.sh script
│   ├── handlers/         # Handlers for HTTP routes
│   │   ├── health.go     # Health check handler
│   │   ├── api.go        # API-specific handlers
│   │   └── ...           # Other handlers
│   ├── services/         # Business logic or service layer
│   │   ├── api_service.go
│   │   └── ...
│   ├── server/           # HTTP server implementation
│   ├── middleware/       # Middleware for request processing
│   │   ├── logging.go
│   │   ├── auth.go
│   │   └── ...
│   └── config/           # Configuration management
│       ├── config.go
│       └── ...
├── pkg/                  # Shared utilities and libraries
│   ├── logger/           # Logging utilities
│   │   ├── logger.go
│   │   └── ...
│   ├── response/         # Utilities for HTTP responses
│   │   ├── json_response.go
│   │   └── ...
│   └── ...
├── spec/                 # OpenAPI scecification 
│   └── openapi.yaml      # The concrete spec file used for generator 
├── test/                 # Tests and test utilities
│   ├── integration/      # Integration tests
│   └── ...
├── go.mod                # Go module file
├── go.sum                # Go module checksum file
└── README.md             # Documentation for the project
```

## Directory Breakdown

### `cmd/server/`
The `cmd` directory hosts the main entry point for the application. The `server/main.go` file initializes and starts the HTTP server.

### `internal/`
This directory contains core application logic that is not intended to be exposed outside of this project.

- **`gen/`**: We are using oapi-codegen to generate code from the ./spec/openapi.yaml file.
- **`handlers/`**: Defines HTTP route handlers that process incoming requests and return responses.
- **`services/`**: Implements the business logic and acts as an intermediary between handlers and the data layer.
- **`server/`**: Implements the HTTP server code to offload the main entry point.
- **`middleware/`**: Contains reusable middleware functions for processing HTTP requests, such as logging or authentication.
- **`config/`**: Manages application configurations, such as environment variables and settings.

### `pkg/`
The `pkg` directory includes reusable utilities and libraries that can be used across the project. 

- **`logger/`**: Utilities for structured logging.
- **`response/`**: Helper functions for building consistent HTTP responses.

### `test/`
Organized directory for tests:
- **`integration/`**: Integration tests to verify end-to-end functionality.
- **`unit/`**: Unit tests for individual components.

### `go.mod` and `go.sum`
These files define and manage the project's dependencies using Go modules.

### `README.md`
The project documentation, including setup instructions, usage details, and additional information.

---
