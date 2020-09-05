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

export REACT_IMAGE_NAME=paulwizviz/goweb-dev-react
export REST_IMAGE_NAME=paulwizviz/goweb-dev-rest
export IMAGE_TAG=dev
export NODE_IMAGE_TAG=13.10.1

COMMAND="$1"

function package() {
    docker build -f ./build/package/dev/react.dockerfile -t ${REACT_IMAGE_NAME}:${IMAGE_TAG} .
    docker build -f ./build/package/dev/rest.dockerfile -t ${REST_IMAGE_NAME}:${IMAGE_TAG} .
}

function run() {
    docker-compose -f ./deployments/dev/docker-compose.yaml up -d
}

function stop(){
    docker-compose -f ./deployments/dev/docker-compose.yaml down
}

function clean(){
    docker-compose -f ./deployments/dev/docker-compose.yaml down
    docker rmi -f ${REACT_IMAGE_NAME}:${IMAGE_TAG}
    docker rmi -f ${REST_IMAGE_NAME}:${IMAGE_TAG}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

message="Usage: $0 [package | run | stop | clean ]"

if [ "$#" -ne 1 ]; then
    echo $message
    exit 1
fi

case $COMMAND in
    "clean")
        clean
        ;;
    "package")
        package
        ;;
    "run")
        run
        ;;
    "stop")
        stop
        ;;
    *)
        echo $message
        ;;
esac