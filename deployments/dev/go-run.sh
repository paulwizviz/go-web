#!/bin/bash

go mod download && \
go build -o /usr/local/bin/${REST_SERVER_NAME} ./cmd/${REST_SERVER_NAME}

/usr/local/bin/${REST_SERVER_NAME} start noui -p 9000