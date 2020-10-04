#!/bin/bash

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