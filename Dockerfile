# Multi-stage build for Go backend
FROM golang:1.21-alpine AS builder

# Install git for private modules (if needed)
RUN apk add --no-cache git ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=1 for SQLite support
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Production stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests and sqlite
RUN apk --no-cache add ca-certificates sqlite

# Create appuser for security
RUN addgroup -g 1001 appgroup && \
    adduser -u 1001 -G appgroup -s /bin/sh -D appuser

WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/main .

# Copy config files if needed
COPY --from=builder /app/configs ./configs

# Create directory for SQLite database
RUN mkdir -p /app/data && chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port 8083
EXPOSE 8083

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8083/health || exit 1

# Command to run the application
CMD ["./main"] 