NERDZ API
=========

This repository contains the API of nerdz (incredible). Swagger-generated documentation available at: https://api.nerdz.eu/docs

## Create a client

After having registered an application on NERDZ, you can use the OAuth2 authorization flow to get the correct tokens. The following is a minimal example in Python that shows how to interact with the NERDZ API.

```python
import logging
from oauth2_client.credentials_manager import (CredentialManager,
                                               ServiceInformation)

logging.basicConfig(level=logging.DEBUG)
_logger = logging.getLogger("client")

scopes = ["pms:read", "pms:write"]
service_information = ServiceInformation(
    "https://api.nerdz.eu/v1/oauth2/authorize",
    "https://api.nerdz.eu/v1/oauth2/token",
    "1",  # client_id
    "$2a$07$F4HMU60OX0Tc5bsufMOY7OZBXjItcd7VzmN2r89Uwezf0Fasdasd",  # client_secret
    scopes,
)

manager = CredentialManager(service_information)
redirect_uri = "http://localhost:8080/oauth/code"

# Builds the authorization url and starts the local server according to the redirect_uri parameter
url = manager.init_authorize_code_process(redirect_uri, "state_test")
_logger.info("Open this url in your browser\n%s", url)
code = manager.wait_and_terminate_authorize_code_process()
_logger.debug("Code got = %s", code)
manager.init_with_authorize_code(redirect_uri, code)
_logger.debug("Access got = %s", manager._access_token)
```

# Contributing

If you want to contribute, you should be at least a [NERDZ](http://www.nerdz.eu/) user.

Developers can go to `doc/CONTRIBUTING.md` to see the developer's guide to NERDZ-API.

# License

Copyright (C) 2016-2020 Paolo Galeone; nessuno@nerdz.eu

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
