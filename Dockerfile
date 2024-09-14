# Build stage
FROM golang:1.22.1-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o qrserver main.go

# Final stage
FROM alpine:3.20

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/qrserver .

# Expose port 8080
EXPOSE 8080

# Run the executable
CMD ["./qrserver"]
