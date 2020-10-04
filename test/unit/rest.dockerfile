ARG GO_VER

FROM golang:${GO_VER}

WORKDIR /opt

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go mod download &&\
    go test ./...
