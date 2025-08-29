#!/bin/bash

set -euo pipefail

FOLDER="$( cd $(dirname "${BASH_SOURCE[0]}"); pwd )"
cd $FOLDER/..

if command -v jq >/dev/null 2>&1; then
  go run cmd/main.go | jq -R "fromjson? | . "
else
  go run cmd/main.go
fi
