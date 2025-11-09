FROM golang:1.24-alpine AS builder

RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/shortener /app/cmd/url-shortener

EXPOSE 8080

USER root
CMD ["./bin/shortener"]