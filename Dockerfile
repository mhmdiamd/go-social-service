# FROM golang:1.22 AS builder
#
# # Set working directory for the build stage
# WORKDIR /go/src/app
#
# # Copy your Go source code (replace with your actual directory)
# COPY . .
#
# # Install dependencies (Replace with your actual commands)
# RUN go mod download
#
# # Build the go binary (Replace with your actual Build commands)
# RUN go build -o main ./cmd/api
#
# EXPOSE 4000
FROM golang:1.22

WORKDIR /go/src/app

COPY . .

ENV APP_ENV=staging

RUN go mod download

EXPOSE 4000

CMD ["go", "run", "./cmd/api/main.go"]
