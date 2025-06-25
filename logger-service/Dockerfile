# Multi-stage Dockerfile for building and running the authentication service
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . /app

RUN go mod download
RUN CGO_ENABLED=0 go build -o logApp /app/cmd/api
RUN chmod +x /app/logApp

# build a tiny docker image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/logApp /app

CMD ["/app/logApp"]
