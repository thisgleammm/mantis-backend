# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app
RUN apk add --no-cache gcc musl-dev
# Install goose
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd
RUN go build -o migrate ./cmd/migrate

# Final stage
FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .
# Copy migrations
COPY --from=builder /app/internal/adapters/postgresql/migrations ./internal/adapters/postgresql/migrations

EXPOSE 8080
CMD ["./main"]
