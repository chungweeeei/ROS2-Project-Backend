# Multi-stage Dockerfile for building and running the authentication service
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . /app

RUN go mod download
RUN CGO_ENABLED=0 go build -o authApp /app/api
RUN chmod +x /app/authApp

# build a tiny docker image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/authApp /app/authApp

CMD ["bash"]
