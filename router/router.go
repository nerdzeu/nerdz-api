package router

import (
	"github.com/RangelReale/osin"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/oauth2"
	"github.com/nerdzeu/nerdz-api/rest"
)

// Init configures the router and returns the *echo.Echo struct
// enableLog set to true enable echo middleware logger
func Init(enableLog bool) *echo.Echo {
	e := echo.New()
	if enableLog {
		e.Use(mw.Logger())
	}

	// Create the Authorization server for OAuth2
	authConfig := osin.NewServerConfig()
	// Configure the Authorization server
	authConfig.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN}
	authConfig.AllowedAccessTypes = osin.AllowedAccessType{
		osin.AUTHORIZATION_CODE,
		osin.REFRESH_TOKEN,
		osin.PASSWORD,
		osin.CLIENT_CREDENTIALS,
	}
	authConfig.AllowGetAccessRequest = true
	authConfig.AllowClientSecretInParams = true

	// Create the storage for osin (where to save oauth infos)
	var authStorage nerdz.OAuth2Storage
	authorizationServer := osin.NewServer(authConfig, &authStorage)

	// Initialize oauth2 server implementation
	oauth2.Init(authorizationServer)

	// OAuth2 routes
	e.Get("/oauth2/authorize", oauth2.Authorize())
	e.Post("/oauth2/authorize", oauth2.Authorize())
	e.Get("/oauth2/token", oauth2.Token())
	e.Get("/oauth2/info", oauth2.Info())
	e.Get("/oauth2/app", oauth2.App())
	e.Get("/oauth2/appauth/code", oauth2.AppAuthCode())
	e.Get("/oauth2/appauth/token", oauth2.AppAuthToken())
	e.Get("/oauth2/appauth/password", oauth2.AppAuthPassword())
	e.Get("/oauth2/appauth/client_credentials", oauth2.AppAuthClientCredentials())
	e.Get("/oauth2/appauth/refresh", oauth2.AppAuthRefresh())
	e.Get("/oauth2/appauth/info", oauth2.AppAuthInfo())

	// Content routes
	e.Get("/users/:id/posts", rest.UserPosts())
	e.Get("/users/:id/friends", rest.UserFriends())
	e.Get("/users/:id", rest.UserInfo())

	return e
}
