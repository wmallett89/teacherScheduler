# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY src /app/src
COPY stub /app/stub

WORKDIR /app/src/cmd

# Build the Go application
RUN go build -o scheduler .

# Stage 2: Run the Go binary in a lightweight container
FROM alpine:latest

# Set a working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/src/cmd/scheduler .
COPY --from=builder /app/src/cmd/config.yaml .

# Expose the port the server listens on
EXPOSE 443

# Run the Go binary
CMD ["./scheduler"]