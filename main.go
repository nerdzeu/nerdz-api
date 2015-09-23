package main

import (
	"github.com/nerdzeu/nerdz-api/api"
	"github.com/nerdzeu/nerdz-api/nerdz"
)

//"github.com/RangelReale/osin"
//"github.com/nerdzeu/nerdz-api/nerdz"

//main starts the server on the specified port
func main() {
	api.Start(nerdz.Configuration.Port, nerdz.Configuration.EnableLog)

	/*
		    ######################## DRAFT FOR OAUTH ##############################
		    osinConfig := osin.NewServerConfig()
			osinConfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
			osinConfig.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN, osin.PASSWORD, osin.CLIENT_CREDENTIALS, osin.ASSERTION}
			osinConfig.AllowGetAccessRequest = true
			osinServer := osin.NewServer(osinConfig, &nerdz.OAuth2Storage{})
		    #######################################################################
	*/

}
