#!/usr/bin/env bash

#### BEGIN CONFIGURATION ####

# Enable/Disable GORM logging: default "1" enabled
ENABLE_LOG="0"
# You can change the go compiler to use. Default: "gc"
COMPILER="gc"
# Set the of of your NERDZ-API configuration file
CONF_FILE="$HOME/nerdz_env/confSample.json"
# Set the path of your nerdz-test-db repo's clone.
TEST_DB_PATH="$HOME/nerdz_env/nerdz-test-db/"

#### END CONFIGURATION ####

echo 'Creating new test database....'
echo 'Exising role (eg. postgres): '
read ROLE
echo 'Db name (eg. test_db): '
read DB_NAME
echo 'Password: '
read DB_PASS

LOCAL_PATH=$(pwd)

cd "$TEST_DB_PATH"
./initdb.sh "$ROLE" "$DB_NAME" "$DB_PASS"
cd "$LOCAL_PATH"

echo 'Test database created'; echo
echo 'Begin tests...'; echo

CONF_FILE=$CONF_FILE ENABLE_LOG=$ENABLE_LOG go test -compiler $COMPILER 
