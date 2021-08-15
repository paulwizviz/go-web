#!/bin/bash

# Copyright 2020 [go-web] Authors
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

export APP_IMAGE_NAME=paulwizviz/goreact:current
export APP_TEST_CONTAINER_NAME=test_container

function build() {
    docker-compose -f ./build/package/goreact/builder.yaml build container
}

function clean() {
    docker rm -f ${APP_TEST_CONTAINER_NAME}
    docker rmi -f ${APP_IMAGE_NAME}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

function run(){
    docker-compose -f ./build/package/goreact/builder.yaml up -d container
}

function stop(){
    docker rm -f $(docker ps -aqf "name=${APP_TEST_CONTAINER_NAME}")
}

message="$0 build | clean | run "

if [ "$#" -ne 1 ]; then
    echo $message
    exit 1
fi

case $COMMAND in
    "build")
        build
        ;;
    "clean")
        clean
        ;;
    "run")
        run
        ;;
    "stop")
        stop ${id}
        ;;
    *)
        echo $message
        ;;
esac