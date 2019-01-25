#!/bin/bash

if [ -z ${SHA+x}]; then SHA=$(git rev-parse HEAD); fi

docker build -t andrewlm/echo1:latest-dev -f ./echoservers/echo1/Dockerfile.dev ./echoservers/echo1
docker build -t andrewlm/echo1:$SHA -f ./echoservers/echo1/Dockerfile.dev ./echoservers/echo1

docker push andrewlm/echo1:latest-dev
docker push andrewlm/echo1:$SHA
