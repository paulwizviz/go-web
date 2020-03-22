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

export IMAGE_TAG="$2"

# Modify image name to suit your project
export IMAGE_NAME=paulwizviz/go-react-container

function checkImageTag() {
    if [ -z "$IMAGE_TAG" ]; then
        echo "$0 $COMMAND image_tag"
        exit 2
    fi
}

function package() {
    docker build -f ./build/package/production/container/Dockerfile -t ${IMAGE_NAME}:${IMAGE_TAG} .
}

function clean() {
    docker rmi -f ${IMAGE_NAME}:${IMAGE_TAG}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

case $COMMAND in
    "package")
        checkImageTag
        package
        ;;
    "clean")
        checkImageTag
        clean
        ;;
    *)
        echo "$0 [package | clean ] version_number"
        ;;
esac