FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum go.work go.work.sum* ./

COPY core/go.mod core/go.sum* ./core/

COPY services/rate-service/go.mod services/rate-service/go.sum* ./services/rate-service/

RUN go mod download

COPY core ./core

COPY gen/go/rates/v1 ./gen/go/rates/v1

COPY services/rate-service ./services/rate-service

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app ./services/rate-service/cmd/api/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /bin/app .

COPY services/rate-service/migrations ./migrations

RUN chmod +x ./app
CMD ["./app"]