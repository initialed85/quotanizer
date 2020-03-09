#!/usr/bin/env bash

set -e -x

IMAGE_NAME="quotanizer"
CONTAINER_NAME="${IMAGE_NAME}-test"

docker run --rm -t --name ${CONTAINER_NAME} --entrypoint /srv/_test.sh ${IMAGE_NAME}
