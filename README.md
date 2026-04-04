# USDT Rate Service

[![Go Version](https://img.shields.io/badge/go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org)
[![gRPC](https://img.shields.io/badge/gRPC-red?style=flat&logo=grpc)](https://grpc.io)
[![Docker](https://img.shields.io/badge/docker-blue?style=flat&logo=docker)](https://www.docker.com)
[![PostgreSQL](https://img.shields.io/badge/postgresql-336791?style=flat&logo=postgresql)](https://www.postgresql.org)

A high-performance gRPC service built with **Go 1.25+** that fetches, processes,
and persists USDT rates from the Grinex exchange.

## 🛠 Tech Stack

- **Language:** Go 1.25+
- **Communication:** gRPC & Protocol Buffers (managed via **buf**)
- **Database:** PostgreSQL with **goose** migrations
- **Observability:** OpenTelemetry (Traces) with **Jaeger** & **zap** logging
- **Tools:** Docker & Docker Compose, Resty (HTTP client), Cleanenv
  (configuration)

## 🚀 Quick Start

### 1. Clone the repository

```bash
git clone https://github.com/WhilsoM/test-go-senior.git
cd test-go-senior
```

## Setup Environment Variables (IMPORTANT)

The project requires .env files to be present in two locations. You must copy
them from the provided examples:

Global / Docker Compose Config:

```bash
cp .env.example .env
```

Service Internal Config:

```bash
cp services/rate-service/.env.example services/rate-service/.env
```

Note: In services/rate-service/.env, ensure DATABASE_URL uses the hostname db
(e.g., postgres://user:pass@db:5432/rate_service) for the Docker environment.

## Launch

Build and start all containers (Service, PostgreSQL, Jaeger):

```bash
docker-compose up -d --build
```

or via Makefile

```bash
make docker-build
```

or via flags

```bash
docker compose up -d db jaeger
```

and then

```bash
cd services/rate-service && go run cmd/api/main.go \                                                            4s
  --db-url="postgres://postgres:password@localhost:5432/rate_service?sslmode=disable" \
  --grpc-port=":50051" \
  --exchange-url="https://grinex.io/api/v1/spot/depth?symbol=usdta7a5"
```

## 🏗 Makefile Commands

Command Description

- make build Downloads dependencies and builds the Linux binary.
- make test Runs all unit tests recursively with verbose output.
- make docker-build Rebuilds and restarts the application via Docker Compose.
- make run Runs the service locally (requires a running Database).
- make lint Runs golangci-lint to ensure code quality.

## 🧪 Testing the API

Once the service is running, you can interact with the gRPC methods using Evans:

```bash
evans --path ./proto --proto rates/v1/rates.proto --port 50051 repl
```

Inside Evans:

```bash
call GetRates
```

## 🔍 Observability

Jaeger UI: Access http://localhost:16686 to visualize request traces.

Database Check: To verify saved rates directly in PostgreSQL:

```bash
docker exec -it rate_db psql -U postgres -d rate_service -c "SELECT * FROM rates;"
```

Logs: Structured JSON logging is handled by zap.
