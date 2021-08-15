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

export REST_TEST_IMAGE=paulwizviz/rest_test_image:current

function restUnit(){
    docker build -f ./test/unit/rest.dockerfile --build-arg GO_VER=1.13.3 -t ${REST_TEST_IMAGE} .
    docker rmi -f ${REST_TEST_IMAGE}
}

message="Usage: $0 rest"

if [ "$#" -ne 1 ]; then
    echo $message
    exit 1
fi

case $COMMAND in
    "unit")
        restUnit
        ;;
    *)
        echo $message
        ;;
esac