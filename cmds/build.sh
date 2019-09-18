#!/usr/bin/env bash
# generate swagger document
swag init -d /opt/dev/code/ws-go/src/candidate_service/resources/ -g /opt/dev/code/ws-go/src/candidate_service/main/main.go

# create docker image
docker build --tag candidate_service_v1 .