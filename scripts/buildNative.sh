#!/bin/bash

COMMAND="$1"

export IMAGE_TAG="$2"

export COMPOSE_PROJECT_NAME=go-react-prod
export IMAGE_NAME=paulwizviz/go-react

function checkImageTag() {
    if [ -z "$IMAGE_TAG" ]; then
        echo "$0 $COMMAND version_number"
        exit 2
    fi
}

function build() {
    docker build -f ./build/Dockerfile.native -t ${IMAGE_NAME}:${IMAGE_TAG} .
}

function container() {
    id=$(docker create $CONTAINER_NAME ${IMAGE_NAME}:${IMAGE_TAG})
    CONTAINER_ID="${id:0:12}"
}

function package() {
    if [ -d ./build/package/ ]; then
        rm -rf ./build/package/
    fi
    mkdir -p ./build/package/linux
    docker cp $CONTAINER_ID:/opt/build/package/linux/ ./build/package/
    mkdir -p ./build/package/macOS
    docker cp $CONTAINER_ID:/opt/build/package/macOS/ ./build/package/
    mkdir -p ./build/package/windows
    docker cp $CONTAINER_ID:/opt/build/package/windows/ ./build/package/
}

function cleanDocker() {
    docker-compose -f ./deployment/compose/prod.yaml down
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

case $COMMAND in
    "package")
        checkImageTag
        build
        container
        package
        ;;
    "clean-package")
        if [ -d ./build/package ]; then
            rm -rf ./build/package
        fi
        ;;
    "clean-docker")
        checkImageTag
        cleanDocker
        ;;
    *)
        echo "$0 [package | clean-package | clean-docker] version_number"
        ;;
esac