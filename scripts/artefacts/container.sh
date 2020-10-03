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

export APP_IMAGE_NAME=paulwizviz/goreact:current

function build() {
    docker-compose -f ./build/package/artefacts/builder.yaml build
}

function clean() {
    docker rmi -f ${APP_IMAGE_NAME}
    docker rmi -f $(docker images --filter "dangling=true" -q)
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
        clean
        ;;
    *)
        echo $message
        ;;
esac