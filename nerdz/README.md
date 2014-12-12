NERDZ-API - Back-end
====================

In this folder you can find the NERDZ back-end implementation.

# Back-end tests

Tests are based on [nerdz-test-db](https://github.com/nerdzeu/nerdz-test-db). If you want to run rests you must correctly setup this environment.

```sh
cd ~/nerdz_env/
git clone https://github.com/nerdzeu/nerdz-test-db.git
```

You don't need to do anything else in that folder.

Come back here and properly setup your JSON configuration file in order to use a new test-db.

Mine looks like:
```json
{
    "Username" : "test_db",
    "Password" : "db_test",
    "DbName"   : "test_db",
    "Host"     : "127.0.0.1",
    "Port"     : 0,
    "SSLMode"  : "disable",
    "NERDZPath": "/home/paolo/nerdz_env/nerdz.eu/",
	"EnableLog": true
}
```

After that, configure the nvironment variables into `test_all.sh`.


# Run the tests

To run all the test, you need a working database. If you wanto to automatically create a new database, use `./test_all.sh`.

If your nerdz-test-db is just ready thus you don't need to create a new one, you can lunch tests in these two ways:

```sh
CONF_FILE="/path/to/conf_file/conf_file_name" go test
```

If you want to see which queries are executed run tests with `EnableLog` parameter set to `true` 
in the configuration file and using the verbose mode for the test tool:


```sh
CONF_FILE="/path/to/conf_file/conf_file_name" go test -v |less
```
