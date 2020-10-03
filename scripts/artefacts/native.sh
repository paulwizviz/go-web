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

. ./scripts/common.sh

COMMAND="$1"

export NATIVE_BUILD_IMAGE=paulwizviz/native_build_image:current
export NATIVE_BUILD_CONTAINER=native_build_container

function cleanBuildContainer(){
    docker rm -f ${NATIVE_BUILD_CONTAINER}
}

function cleanNative() {
    if [ -d ./build/native ]; then
        rm -rf ./build/native
    fi
}

function cleanImages() {
    docker rmi -f ${NATIVE_BUILD_IMAGE}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

function build(){
    cleanNative
    cleanBuildContainer
    cleanImages
    docker-compose -f ./build/package/artefacts/builder.yaml build
    docker-compose -f ./build/package/artefacts/builder.yaml up
}

message="$0 build | clean "

if [ "$#" -ne 1 ]; then
    echo $message
    exit 1
fi

case $COMMAND in
    "build")
        build
        ;;
    "clean")
        cleanBuildContainer
        cleanNative
        cleanImages
        ;;
    *)
        echo $message
        ;;
esac