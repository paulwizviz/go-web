#!/bin/bash

COMMAND="$1"

export IMAGE_TAG="$2"

export COMPOSE_PROJECT_NAME=react-skeleton
export IMAGE_NAME=paulwizviz/go-react

function checkImageTag() {
    if [ -z ${IMAGE_TAG} ]; then
        echo "Set image tag"
        echo "echo $0 $COMMAND image_tag"
        exit 1
    fi
}

function build() {
    docker build -f ./build/Dockerfile -t ${IMAGE_NAME}:${IMAGE_TAG} .
}

function clean() {
    docker-compose -f ./deployment/compose/dev.yaml down
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

checkImageTag
case $COMMAND in
    "build")
        build
        ;;
    "start")
        docker-compose -f ./deployment/compose/dev.yaml up
        ;;
    "stop")
        docker-compose -f ./deployment/compose/dev.yaml down
        ;;
    "clean")
        clean
        ;;
    *)
        echo "$0 (build | start | stop | clean) image_tag"
        echo "image_tag   string values"
        ;;
esac