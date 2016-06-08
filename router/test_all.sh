#!/usr/bin/env bash

#### BEGIN CONFIGURATION ####
# Set the of of your NERDZ-API configuration file
CONF_FILE="$HOME/projects/nerdz/confSample.json"
# Set the path of your nerdz-test-db repo's clone.
TEST_DB_PATH="$HOME/projects/nerdz/nerdz-test-db/"
# Set the username of an exising postgres role, usually postgres
ROLE=postgres
#### END CONFIGURATION ####

LOCAL_PATH=$( cd $(dirname $0) ; pwd -P )
DB_NAME=$(cat $CONF_FILE | jq ".DbName" | tr -d '"')
DB_PASS=$(cat $CONF_FILE | jq ".DbPassword" | tr -d '"')

cd "$TEST_DB_PATH"

./initdb.sh "$ROLE" "$DB_NAME" "$DB_PASS"
cd "$LOCAL_PATH"

echo 'Test database created'; echo
echo 'Begin tests...'; echo

CONF_FILE="$CONF_FILE" go test

echo 'Restoring test db...'
cd "$TEST_DB_PATH"
./initdb.sh "$ROLE" "$DB_NAME" "$DB_PASS"
