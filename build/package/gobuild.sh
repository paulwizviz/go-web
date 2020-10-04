#!/bin/bash

go get github.com/GeertJohan/go.rice/rice
./build/go-rice.sh
go mod download
env GOOS=linux GOARCH=amd64 go build -o ./build/package/linux/${APP_NAME} ./cmd/${APP_NAME}
env GOOS=darwin GOARCH=amd64 go build -o ./build/package/macOS/${APP_NAME} ./cmd/${APP_NAME}
env GOOS=windows GOARCH=amd64 go build -o ./build/package/windows/${APP_NAME}.exe ./cmd/${APP_NAME}