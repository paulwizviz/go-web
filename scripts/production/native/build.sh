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

COMMAND="$1"

# Feel free to modify these variables accordingly
IMAGE_TAG=current
IMAGE_NAME=paulwizviz/go-react-native

function build() {
    docker build -f ./build/production/native/Dockerfile -t ${IMAGE_NAME}:${IMAGE_TAG} .
}

function container() {
    id=$(docker create ${IMAGE_NAME}:${IMAGE_TAG})
    CONTAINER_ID="${id:0:12}"
}

function cleanPackage() {
    if [ -d ./build/package ]; then
        rm -rf ./build/package
    fi
}

function package() {
    cleanPackage
    mkdir -p ./build/package/linux
    docker cp $CONTAINER_ID:/opt/build/package/linux/ ./build/package/
    mkdir -p ./build/package/macOS
    docker cp $CONTAINER_ID:/opt/build/package/macOS/ ./build/package/
    mkdir -p ./build/package/windows
    docker cp $CONTAINER_ID:/opt/build/package/windows/ ./build/package/
    docker rm -f $CONTAINER_ID
}

function cleanDocker() {
    docker rmi -f ${IMAGE_NAME}:${IMAGE_TAG}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

case $COMMAND in
    "package")
        build
        container
        package
        ;;
    "clean")
        cleanPackage
        cleanDocker
        ;;
    *)
        echo "$0 [package | clean]"
        ;;
esac