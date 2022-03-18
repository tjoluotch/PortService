#!/bin/bash

IMG="portservice"

run(){
echo "building docker container..."
docker run \
--restart unless-stopped \
-p 8000:8000/tcp \
--network microservice \
--name portservice \
rpc:0.0.1
}

build(){
echo "building docker image for portservice..."
docker build \
--rm \
-t ${IMG}:0.0.1 .
echo "docker image build for ${IMG} completed..."
}

if [[ $1 == "build" ]]
then
  build
elif [[ $1 == "run" ]]
then
    run
fi