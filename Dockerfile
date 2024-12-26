# Build stage
FROM golang:1.23-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Set working directory
WORKDIR /app

# Copy go mod files
COPY src/go.mod src/go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY src/ ./

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o epusdt .

# Final stage
FROM alpine:3.14

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/epusdt .

# Copy required files
COPY src/static ./static

# Create required directories
RUN mkdir -p runtime/logs

# Expose port (adjust if needed based on your .env configuration)
EXPOSE 8000

# Run binary with the specified command
CMD ["./epusdt", "http", "start"]
