#!/bin/bash

set -e

FOLDER="$( cd "$(dirname "${BASH_SOURCE[0]}")" && pwd )"

docker cp kafka:/etc/kafka/secrets/ca.crt $FOLDER/local-dev-ca.crt
