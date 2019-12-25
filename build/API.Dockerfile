FROM golang:1.13.3

EXPOSE 9000

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go mod download && \
    go build -o /usr/local/bin/api ./cmd/api

