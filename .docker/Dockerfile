FROM golang:1.23.3-alpine AS builder

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

COPY . /app

WORKDIR /app

RUN go build -o /usr/local/bin/migrate -a ./cmd/migrate/main.go
RUN go build -o /usr/local/bin/server -a ./cmd/server/main.go

FROM alpine:latest

COPY --from=builder /usr/local/bin /usr/local/bin
