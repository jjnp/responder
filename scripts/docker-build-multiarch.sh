#!/usr/bin/env bash

image=jjnp/responder

if [[ $1 ]]; then
	version="$1"
else
	version="latest"
fi

basetag="${image}:${version}"

# change into project root
BASE="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PROJECT_ROOT=$(realpath "${BASE}/../")
cd $PROJECT_ROOT

# build all the images
docker build -t ${basetag}-amd64 -f build/package/responder/Dockerfile.amd64 .
docker build -t ${basetag}-arm32v7 -f build/package/responder/Dockerfile.arm32v7 .
docker build -t ${basetag}-aarch64 -f build/package/responder/Dockerfile.aarch64 .

exit 0

