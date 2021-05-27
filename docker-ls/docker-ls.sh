#!/usr/bin/env bash

docker-ls repositories --basic-auth -r https://gcr.io/jetstack-paul -l1
docker-ls tags --basic-auth -t taglist -r https://gcr.io jetstack-paul/hello -l1 -c config.yaml
