FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app ./services/rate-service/cmd/api/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /bin/app .
COPY services/rate-service/migrations ./migrations

RUN chmod +x ./app
CMD ["./app"]