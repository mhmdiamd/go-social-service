# Use the official Golang image as the base image
FROM golang:1.22-alpine

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download all dependencies
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go application
RUN go build -o go-social-service ./cmd/api/main.go

# Expose the application port
EXPOSE 4000

# Copy the entrypoint script to the container
COPY entrypoint.sh /entrypoint.sh

# Ensure the entrypoint script is executable
RUN chmod +x /entrypoint.sh

# Set the entrypoint script as the container's entry point
ENTRYPOINT ["/entrypoint.sh"]
