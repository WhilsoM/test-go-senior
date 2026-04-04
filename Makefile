.PHONY: build test docker-build run lint

build:
	go mod download
	CGO_ENABLED=0 GOOS=linux go build -o app ./services/rate-service/cmd/api/main.go

test:
	go test -v ./services/rate-service/internal/service/

docker-build:
	docker compose up --build -d

run:
	go run ./services/rate-service/cmd/api/main.go

lint:
	golangci-lint run ./...