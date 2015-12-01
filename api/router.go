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

	// Configuring the Authorization server oauth2
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
	OAuth = osin.NewServer(authConfig, &authStorage)

	// OAuth2 routes
	e.Get("/oauth2/authorize", OAuth2Authorize)
	e.Post("/oauth2/authorize", OAuth2Authorize)
	e.Get("/oauth2/token", OAuth2Token)
	e.Get("/oauth2/info", OAuth2Info)
	e.Get("/oauth2/app", OAuth2App)
	e.Get("/oauth2/appauth/code", OAuth2AppAuthCode)
	e.Get("/oauth2/appauth/token", OAuth2AppAuthToken)
	e.Get("/oauth2/appauth/password", OAuth2AppAuthPassword)
	e.Get("/oauth2/appauth/client_credentials", OAuth2AppAuthClientCredentials)
	e.Get("/oauth2/appauth/refresh", OAuth2AppAuthRefresh)
	e.Get("/oauth2/appauth/info", OAuth2AppAuthInfo)

	// Content routes
	e.Get("/users/:id/posts", UserPosts)
	e.Get("/users/:id/friends", UserFriends)
	e.Get("/users/:id", UserInfo)
	e.Run(":" + strconv.Itoa(int(port)))
}
