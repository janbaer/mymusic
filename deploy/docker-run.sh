#!/bin/sh

DOCKER_COMPOSE=/usr/local/bin/docker-compose

${DOCKER_COMPOSE} down

PWD=$(pwd) ${DOCKER_COMPOSE} up -d
