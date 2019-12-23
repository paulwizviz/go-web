#!/bin/bash

COMMAND="$1"

export IMAGE_TAG="$2"

export COMPOSE_PROJECT_NAME=react-skeleton
export DEV_IMAGE_NAME=paulwizviz/react-ui-dev
export PROD_IMAGE_NAME=paulwizviz/react-ui

cli_message="$0 (build | start | stop | dev | clean) image_tag"

function checkImageTag() {
    if [ -z ${IMAGE_TAG} ]; then
        echo "Set image tag"
        echo $cli_message
        exit 1
    fi
}

function build() {
    docker-compose -f ./docker-composer/builder.yaml build development
}

function ops() {
    local command="$1"
    [ "$command" == "start" ] && ( docker-compose -f ./docker-composer/dev.yaml up )
    [ "$command" == "stop" ] && ( docker-compose -f ./docker-composer/dev.yaml down )
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
        echo $cli_message
        echo "image_tag   string values"
        ;;
esac