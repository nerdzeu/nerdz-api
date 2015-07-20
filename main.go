package main

import (
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/nerdzeu/nerdz-api/nerdz"
)

func main() {
	osinConfig := osin.NewServerConfig()
	osinConfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeTypes{osin.CODE, osin.TOKEN}
	osinConfig.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN, osin.PASSWORD, osin.CLIENT_CREDENTIALS, osin.ASSERTION}
	osinConfig.AllowGetAccessRequest = true
	osinServer := osin.NewServer(osinConfig, &nerdz.OAuth2Storage{})
}
