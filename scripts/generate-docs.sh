#!/bin/bash

FOLDER="$( cd $(dirname "${BASH_SOURCE[0]}"); pwd )"
cd $FOLDER/..

# npm --loglevel verbose install -g @asyncapi/cli
asyncapi generate fromTemplate ./docs/asyncapi/spec.yaml @asyncapi/html-template@3.0.0 --use-new-generator -o ./docs/asyncapi

# npm install -g @redocly/cli@latest
redocly build-docs ./docs/openapi/spec.yaml -o ./docs/openapi/index.html
