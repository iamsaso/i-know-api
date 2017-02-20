#!/bin/bash
set -e

export DOCKER_HOST_IP=$(ip route | awk '/default/ { print $3 }')

exec "$@"
