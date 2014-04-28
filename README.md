NERDZ API
=========

Yay, my  [pull request](https://github.com/jinzhu/gorm/pull/85) has been merged! Thus,

```sh
go get github.com/jinzhu/gorm/
```

Configuration file
=================
Because JSON standard prohibits comments, your must remove the comments if you are willing to use the sample configFile.json below (comments are there as an explanation).

```JSON
{
    "Username" : "nerdz",     // required
    "Password" : "pass",      // required if any, otherwise ""
    "DbName"   : "nerdz",     // required
    "Host"     : "127.0.0.1", // opional, i.e. "" -> fallback to localhost
    "Port"     : 0,           // optional, i.e. 0 -> fallback to 5432
    "SSLMode"  : "disable"    // optional, i.e. "" -> fallback to disable
}
```

Tests
=====

Tests are based on [nerdz-test-db](https://github.com/nerdzeu/nerdz-test-db). If you want to run rests you must correctly setup this environment.
Thus, first of all you have to:
```sh
cd ~
git clone https://github.com/nerdzeu/nerdz-test-db.git
cd nerdz-test-db
./initdb.sh
```
This will display the script's usage. Add the two requried parameters to setup the test-db and run the script.

Once the database is up and running, you have to properly setup your JSON configuration file in order to use this database.

Mine looks like:
```json
{
    "Username" : "test_db",
    "Password" : "db_test",
    "DbName"   : "test_db",
    "Host"     : "127.0.0.1",
    "Port"     : 0,
    "SSLMode"  : "disable"
}
```

Since is required a working database to test the API, you need to specify the configuration file path to the go test command, for example:

```sh
CONF_FILE="$HOME/confSample.json" go test orm/tests/user_test.go
```

If you want to see which queries are executed run tests with `ENABLE_LOG` environment value not empty

```sh
ENABLE_LOG="1" CONF_FILE="$HOME/confSample.json" go test orm/tests/project_test.go -v |less
```

You can run all the tests after having configured the enviroinment variables in `testAll.sh`, after you can simply run the tests running the script
```sh
./testAll.sh
```
By default the `testAll.sh` script doesn't print anything as output except the status of the test (ok or fail, with a long output in this second case).

If you wan't to enable the verbose mode, you can launch the script with the -v flag
```sh
./testAll.sh -v
```
In general the flag passed to `testAll.sh` are passed to `go test`
