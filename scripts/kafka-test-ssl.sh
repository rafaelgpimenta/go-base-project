#!/bin/bash

set -euo pipefail

FOLDER="$( cd $(dirname "${BASH_SOURCE[0]}"); pwd )"
cd $FOLDER/..

docker compose -f "./docker/docker-compose.yml" up -d
INIT_CONTAINER_ID=$(docker compose -f "./docker/docker-compose.yml" ps -q kafka-init)
docker wait "$INIT_CONTAINER_ID" > /dev/null

"${FOLDER}/cp-ca.sh"

if command -v jq >/dev/null 2>&1; then
  SSL_CERT_FILE="$FOLDER/local-dev-ca.crt" go run cmd/kafka-ssl/main.go | jq -R "fromjson? | . "
else
  SSL_CERT_FILE="$FOLDER/local-dev-ca.crt" go run cmd/kafka-ssl/main.go
fi
