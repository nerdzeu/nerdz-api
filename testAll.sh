#!/bin/bash

# Sample environment variables, change them according to your environment
ENABLE_LOG="1"
CONF_FILE="$HOME/nerdz_env/confSample.json"

for i in $(find . -name '*_test.go'); do
    echo $i
    CONF_FILE=$CONF_FILE ENABLE_LOG=$ENABLE_LOG go test $i $@
done
