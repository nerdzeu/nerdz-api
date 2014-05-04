#!/bin/bash

# Sample environment variables, change them according to your environment
ENABLE_LOG="1"
CONF_FILE="$HOME/nerdz_env/confSample.json"

for i in $(find . -name '*_test.go'); do
    echo $i
    echo 'Using gc:'
    CONF_FILE=$CONF_FILE ENABLE_LOG=$ENABLE_LOG go test -compiler gc $i $@
    if hash gccgo 2>/dev/null; then
      echo 'Using gccgo:'
      CONF_FILE=$CONF_FILE ENABLE_LOG=$ENABLE_LOG go test -compiler gccgo $i $@
    fi
done
