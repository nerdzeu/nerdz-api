NERDZ API
=========

# Configuration file

Because JSON standard prohibits comments, your must remove the comments if you are willing to use the sample configFile.json below (comments are there as an explanation).

```JSON
{
    "DbUsername" : "nerdz",                 // required
    "DbPassword" : "pass",                  // required if any, otherwise ""
    "DbName"     : "nerdz",                 // required
    "DbHost"     : "127.0.0.1",             // optional, i.e. "" -> fallback to localhost
    "DbPort"     : 0,                       // optional, i.e. 0 -> fallback to 5432
    "SSLMode"    : "disable",               // optional, i.e. "" -> fallback to disable
    "NERDZPath"  : "~/nerdz.eu/",           // required
    "NERDZUrl"   : "https://www.nerdz.eu/", // required
    "EnableLog"  : false,		            // optional, default false
    "Port"       : 8080                     // API port, optional -> fallback to 7536
}
```

# Tests

For back-end tests, see `nerdz/README.md`.

For front-end tests, you have to wait ;)

# Contributing

If you want to contribute, you should be at least a [NERDZ](http://www.nerdz.eu/) user.

Developers can go to `doc/CONTRIBUTING.md` to see the developer's guide to NERDZ-API.
