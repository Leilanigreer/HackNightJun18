# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app ./app
COPY modus.json ./modus.json
EXPOSE 8080
CMD ["./app"] 