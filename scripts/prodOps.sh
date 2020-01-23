#!/bin/bash

COMMAND="$1"

export IMAGE_TAG="$2"

export COMPOSE_PROJECT_NAME=react-skeleton
export IMAGE_NAME=paulwizviz/go-react

function checkImageTag() {
    if [ -z ${IMAGE_TAG} ]; then
        echo "echo $0 $COMMAND version_number"
        exit 1
    fi
}

function build() {
    docker build -f ./build/Dockerfile -t ${IMAGE_NAME}:${IMAGE_TAG} .
}

function clean() {
    docker-compose -f ./deployment/compose/prod.yaml down
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

checkImageTag
case $COMMAND in
    "build")
        build
        ;;
    "start")
        docker-compose -f ./deployment/compose/prod.yaml up
        ;;
    "stop")
        docker-compose -f ./deployment/compose/prod.yaml down
        ;;
    "clean")
        clean
        ;;
    *)
        echo "$0 (build | start | stop | clean) version_number"
        echo "image_tag   string values"
        ;;
esac