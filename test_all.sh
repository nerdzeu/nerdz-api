#!/usr/bin/env bash

# Sample environment variables, change them according to your environment
ENABLE_LOG="1"
CONF_FILE="$HOME/nerdz_env/confSample.json"
TEST_DB_PATH="$HOME/nerdz_env/nerdz-test-db/"

echo 'Creating new test database....'
echo 'Exising role (eg. postgres): '
read ROLE
echo 'Db name (eg. test_db): '
read DB_NAME

LOCAL_PATH=$(pwd)

cd $TEST_DB_PATH
./initdb.sh $ROLE $DB_NAME
cd $LOCAL_PATH
echo $LOCAL_PATH

echo 'Test database created'; echo

#if hash gccgo 2>/dev/null; then
#    COMPILER=gccgo
#else
    COMPILER=gc
#fi

echo "Using $COMPILER"; echo

for i in $(find . -name '*_test.go'); do
    echo $i
    CONF_FILE=$CONF_FILE ENABLE_LOG=$ENABLE_LOG go test -compiler $COMPILER $i $@
done
