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
