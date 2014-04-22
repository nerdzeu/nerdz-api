NERDZ API
=========

Yay, my  [pull request](https://github.com/jinzhu/gorm/pull/85) has been merged! Thus,

```sh
go get github.com/jinzhu/gorm/
```

Configuration file
=================
Go can't handle properly comments in JSON files, so your configFile.json must be without comment. The one below is commented only to explain the fileds

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

Since is required a working database to test the API, you need to specify the configuration file path to the test command, for example:

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


TODO
====
Tests works only with my local copy of the nerdz database.

After completing the develop of the API, I'll make a test database avaiable to everyone (with false and testing values) in a new repository, I'll add this repository as submodule. In that way you can build your own test database and running all the tests.
