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

export IMAGE_NAME=paulwizviz/go-node-builder
export IMAGE_TAG=current

function package() {
    docker build -f ./build/node/Dockerfile -t ${IMAGE_NAME}:${IMAGE_TAG} .
}

function container() {
    id=$(docker create $CONTAINER_NAME ${IMAGE_NAME}:${IMAGE_TAG})
    CONTAINER_ID="${id:0:12}"
}

function node_modules() {
    if [ -d ./build/node_modules ]; then
        rm -rf ./build/node_modules
    fi
    docker cp $CONTAINER_ID:/opt/node_modules ./build
}
    
package
container
node_modules