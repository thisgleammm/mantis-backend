# Build stage
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build applications
RUN go build -o main ./cmd
RUN go build -o migrate ./cmd/migrate

# Final stage
FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache ca-certificates

# Copy binaries from builder
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .

# Copy migrations (maintained structure)
COPY --from=builder /app/internal/adapters/postgresql/migrations ./internal/adapters/postgresql/migrations

EXPOSE 8080

CMD ["./main"]
