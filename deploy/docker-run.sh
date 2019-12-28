#!/bin/sh

DOCKER_COMPOSE=/usr/local/bin/docker-compose

${DOCKER_COMPOSE} down

./create-certificate.sh

PWD=$(pwd) ${DOCKER_COMPOSE} up -d
