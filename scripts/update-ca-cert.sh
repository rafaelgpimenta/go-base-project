#!/bin/bash

set -e

FOLDER="$( cd "$(dirname "${BASH_SOURCE[0]}")" && pwd )"

sudo cp $FOLDER/local-dev-ca.crt /usr/local/share/ca-certificates/local-dev-ca.crt
sudo update-ca-certificates
