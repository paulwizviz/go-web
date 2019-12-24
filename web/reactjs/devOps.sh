#!/bin/bash

COMMAND="$1"

export IMAGE_TAG="$2"

export COMPOSE_PROJECT_NAME=react-skeleton
export DEV_IMAGE_NAME=paulwizviz/react-ui-dev
export PROD_IMAGE_NAME=paulwizviz/react-ui

function checkImageTag() {
    if [ -z ${IMAGE_TAG} ]; then
        echo "Set image tag"
        echo "echo $0 $COMMAND image_tag"
        exit 1
    fi
}

function build() {
    docker-compose -f ./docker-compose.builder.yaml build development
}

function ops() {
    local command="$1"
    [ "$command" == "start" ] && ( docker-compose -f ./docker-compose.dev.yaml up )
    [ "$command" == "stop" ] && ( docker-compose -f ./docker-compose.dev.yaml down )
}

function clean() {
    docker rmi -f $(docker images --filter "dangling=true" -q)
    docker rmi -f ${DEV_IMAGE_NAME}:${IMAGE_TAG}
    docker rmi -f ${PROD_IMAGE_NAME}:${IMAGE_TAG}
}

checkImageTag
case $COMMAND in
   "build")
        build
        ;;
    "start")
        ops start
        ;;
    "stop")
        ops stop
        ;;
    "clean")
        clean
        ;;
    *)
        echo "$0 (build | start | stop | dev | clean) image_tag"
        echo "image_tag   string values"
        ;;
esac