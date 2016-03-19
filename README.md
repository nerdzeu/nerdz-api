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
    "DbSSLMode"    : "disable",             // optional, i.e. "" -> fallback to disable
    "NERDZPath"  : "~/nerdz.eu/",           // required
    "NERDZHost"  : "www.nerdz.eu",          // required
    "EnableLog"  : false,		            // optional, default false
    "Host"       : "api.nerdz.eu",          // required
    "Scheme"     : "https",                 // required, in production must be https (mandatory for OAuth2)
    "Port"       : 8080                     // API port, optional -> fallback to 7536
}
```

# Tests

For back-end tests, see `nerdz/README.md`.

For front-end tests, you have to wait ;)

# Contributing

If you want to contribute, you should be at least a [NERDZ](http://www.nerdz.eu/) user.

Developers can go to `doc/CONTRIBUTING.md` to see the developer's guide to NERDZ-API.

# License
Copyright (C) 2016 Paolo Galeone; nessuno@nerdz.eu

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
