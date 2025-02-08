# Multi-stage build for optimized production image
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod .
# Comment it out for future use if needed
# COPY go.sum .

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./main.go

# Final stage
FROM alpine:latest

# Install certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

# Copy web assets
COPY cmd/ web/

# Expose port
EXPOSE 8080

# Run the application
CMD ["./server"]
