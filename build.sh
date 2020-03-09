#!/usr/bin/env bash

set -e -x

IMAGE_NAME="quotanizer"

docker build -t ${IMAGE_NAME} .
