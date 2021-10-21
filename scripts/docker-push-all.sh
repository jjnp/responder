#!/usr/bin/env bash

image=jjnp/responder

if [[ $1 ]]; then
	version="$1"
else
	version="latest"
fi

basetag="${image}:${version}"


# # push em all
docker push ${basetag}-amd64 &
docker push ${basetag}-arm32v7 &
docker push ${basetag}-aarch64 &

wait

export DOCKER_CLI_EXPERIMENTAL=enabled

# create the manifest
docker manifest create ${basetag} \
	${basetag}-amd64 \
	${basetag}-arm32v7 \
	${basetag}-aarch64

# explicit annotations
docker manifest annotate ${basetag} ${basetag}-arm32v7 --os "linux" --arch "arm" --variant "v7"

# ship it
docker manifest push --purge ${basetag}
