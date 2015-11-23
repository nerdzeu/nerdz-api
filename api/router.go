package api

import (
	"strconv"

	"github.com/RangelReale/osin"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/nerdzeu/nerdz-api/nerdz"
)

// Start starts the API server on specified port.
// enableLog set to true enable echo middleware logger
func Start(port int16, enableLog bool) {
	e := echo.New()
	if enableLog {
		e.Use(mw.Logger())
	}

	// Configuring the Authorization server oauth
	authConfig := osin.NewServerConfig()
	authConfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
	authConfig.AllowedAccessTypes = osin.AllowedAccessType{
		osin.AUTHORIZATION_CODE,
		osin.REFRESH_TOKEN,
		osin.PASSWORD,
		osin.CLIENT_CREDENTIALS,
		//osin.ASSERTION: TODO ?
	}
	authConfig.AllowGetAccessRequest = true
	authConfig.AllowClientSecretInParams = true
	var authStorage nerdz.OAuth2Storage
	oauth = osin.NewServer(authConfig, &authStorage)

	// OAuth2 routes
	e.Get("/oauth/authorize", OAuth2Authorize)
	e.Get("/oauth/token", OAuth2Token)
	e.Get("/oauth/info", OAuth2Info)
	e.Get("/oauth/app", OAuth2App)
	e.Get("/oauth/appauth/code", OAuth2AppAuthCode)
	e.Get("/oauth/appauth/token", OAuth2AppAuthToken)
	e.Get("/oauth/appauth/password", OAuth2AppAuthPassword)
	e.Get("/oauth/appauth/client_credentials", OAuth2AppAuthClientCredentials)
	e.Get("/oauth/appauth/refresh", OAuth2AppAuthRefresh)
	e.Get("/oauth/appauth/info", OAuth2AppAuthInfo)

	// Content routes
	e.Get("/users/:id/posts", UserPosts)
	e.Get("/users/:id/friends", UserFriends)
	e.Get("/users/:id", UserInfo)
	e.Run(":" + strconv.Itoa(int(port)))
}
