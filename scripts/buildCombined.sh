#!/bin/bash

COMMAND="$1"

export IMAGE_TAG="$2"

export COMPOSE_PROJECT_NAME=go-react-prod
export IMAGE_NAME=paulwizviz/go-react

function checkImageTag() {
    if [ -z "$IMAGE_TAG" ]; then
        echo "$0 $COMMAND image_tag"
        exit 2
    fi
}

function build() {
    docker build -f ./build/combined/Dockerfile -t ${IMAGE_NAME}:${IMAGE_TAG} .
}

function cleanDocker() {
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

case $COMMAND in
    "build")
        checkImageTag
        build
        ;;
    "clean")
        checkImageTag
        cleanDocker
        ;;
    *)
        echo "$0 [build | clean] version_number"
        ;;
esac