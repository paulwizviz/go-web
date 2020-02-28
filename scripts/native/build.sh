#!/bin/bash

# Copyright 2020 Paul Sitoh
# 
# Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This is under development

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