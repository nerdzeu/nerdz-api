NERDZ API
=========

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
