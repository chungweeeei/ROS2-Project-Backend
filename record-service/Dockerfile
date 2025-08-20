# Multi-stage Dockerfile for building and running the authenticate service
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . /app

RUN go mod download
RUN CGO_ENABLED=0  go build -ldflags="-s -w" -o recordApp /app/cmd/api
RUN chmod +x /app/recordApp

# build a tiny docker image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/recordApp /app

CMD ["/app/recordApp"]