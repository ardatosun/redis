# Use official golang image as the base image
FROM golang:1.22.4 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o redis-server .

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory to /app in the container
WORKDIR /app

# Copy the compiled Go executable from the builder stage
COPY --from=builder /app/redis-server .

# Expose port 6379 on which the Redis server will run
EXPOSE 6379

# Command to run the Redis server
CMD ["./redis-server"]
