#!/bin/bash

FOLDER="$( cd $(dirname "${BASH_SOURCE[0]}"); pwd )"
cd $FOLDER/..

go run cmd/main.go
