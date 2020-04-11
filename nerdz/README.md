NERDZ-API - Back-end
====================

In this folder you can find the NERDZ back-end implementation.

# Back-end types architecture

In the implementation of the NERDZ's type hierarchy, a lot of data structure are generated in order to manipulate all the
information that the social network manages. All the data structure are filled by a database ORM that is used
to avoid to rely on the specific query dialect when creating each query. In order to work with it, a specific type has been defined in the file
`types.go` which contains specific details that grant to the ORM to manage all the database's logic.

All the information it's not simply generated and consumed inside the system. In making available an API, we need to decouple all the internal data structure
from the one that are returned to the user by the system. It's absolutely **NOT** correct to return all the data exactly has they are stored in the database. For this reason,
according to the state-of-the-art pattern of the [Transfer Object](http://www.oracle.com/technetwork/java/transferobject-139757.html), we have defined a main structure
and a specific *transfer object type* which can be generated from it.

In particular, each type defined in the file `types.go` implements a specific interface, called `Transferable`, which let to it to define how will be generated
the companion transfer object type defined in the file `api_types.go`. Each transfer object type associated to the main struct, doesn't have all the ORM details
and is completely seperated from the database's logic.

Working in this way, will be necessary, to get each transfer object specific type for a data structure, to work a little with type conversion and type switch.
In a first moment this could be really tedious, but this approach has a lot of benefits that are appreciable only in the long-run.

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
    "DbUsername" : "test_db",
    "DbPassword" : "test_db",
    "DbName"     : "test_db",
    "DbHost"     : "127.0.0.1",
    "DbPort"     : 0,
    "DbSSLMode"    : "disable",
    "NERDZPath"  : "/home/paolo/nerdz_env/nerdz.eu/",
    "NERDZHost"  : "local.nerdz.eu",
    "EnableLog"  : true,
    "Port"       : 9090,
    "Scheme"     : "http",
    "Host"       : "local.api.nerdz.eu"
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
