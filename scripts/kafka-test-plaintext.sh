#!/bin/bash

set -euo pipefail

FOLDER="$( cd $(dirname "${BASH_SOURCE[0]}"); pwd )"
cd $FOLDER/..

docker compose -f "./docker/docker-compose.yml" up -d
INIT_CONTAINER_ID=$(docker compose -f "./docker/docker-compose.yml" ps -q kafka-init)
docker wait "$INIT_CONTAINER_ID" > /dev/null

if command -v jq >/dev/null 2>&1; then
  go run cmd/kafka-plaintext/main.go | jq -R "fromjson? | . "
else
  go run cmd/kafka-plaintext/main.go
fi
