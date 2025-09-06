#!/usr/bin/bash

RASP_HOST="192.168.1.24"

docker context create rasp --docker host=tcp://${RASP_HOST}:2375
docker --context rasp compose -f docker/rasp/docker-compose.yml build
docker --context rasp compose -f docker/rasp/docker-compose.yml up -d