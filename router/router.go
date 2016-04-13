/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

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
*/

package router

import (
	"github.com/RangelReale/osin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/oauth2"
	"github.com/nerdzeu/nerdz-api/rest/me"
	"github.com/nerdzeu/nerdz-api/rest/user"
	"github.com/nerdzeu/nerdz-api/stream"
	"strconv"
)

// VERSION is the API version and base path of requests: /v<VERSION>/
const VERSION = 1

// Init configures the router and returns the *echo.Echo struct
// enableLog set to true enable echo middleware logger
func Init(enableLog bool) *echo.Echo {
	e := echo.New()
	if enableLog {
		e.Use(middleware.Logger())
	}
	e.Pre(middleware.RemoveTrailingSlash())

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

	basePath := e.Group("/v" + strconv.Itoa(VERSION))

	/**************************************************************************
	* ROUTE /oauth2
	* Authorization not required.
	***************************************************************************/
	o := basePath.Group("/oauth2")
	o.Get("/authorize", oauth2.Authorize())
	o.Post("/authorize", oauth2.Authorize())
	o.Get("/token", oauth2.Token())
	o.Get("/info", oauth2.Info())
	/**************************************************************************
	* ROUTE /users/:id
	* Authorization required
	***************************************************************************/
	usersG := basePath.Group("/users") // users Group
	usersG.Use(authorization())
	usersG.Use(user.SetOther())
	usersG.Get("/:id", user.Info())
	usersG.Get("/:id/friends", user.Friends())
	usersG.Get("/:id/followers", user.Followers())
	usersG.Get("/:id/following", user.Following())
	// uses setPostlist middleware
	usersG.Get("/:id/posts", user.Posts(), setPostlist())
	// requests below uses the user.SetPost() middleware to refers to the requested post
	usersG.Get("/:id/posts/:pid", user.Post(), user.SetPost())
	// uses setCommentList middleware
	usersG.Get("/:id/posts/:pid/comments", user.PostComments(), user.SetPost(), setCommentList())
	usersG.Get("/:id/posts/:pid/comments/:cid", user.PostComment(), user.SetPost())

	/**************************************************************************
	* ROUTE /me
	* Authorization required
	***************************************************************************/
	meG := basePath.Group("/me")
	meG.Use(authorization())
	meG.Use(me.SetOther())
	meG.Get("", me.Info())
	meG.Get("/friends", me.Friends())
	meG.Get("/followers", me.Followers())
	meG.Get("/following", me.Following())
	meG.Get("/whitelist", me.Whitelist())
	meG.Get("/whitelisting", me.Whitelisting())
	meG.Get("/blacklist", me.Blacklist())
	meG.Get("/blacklisting", me.Blacklisting())
	meG.Get("/home", me.Home(), setPostlist())
	meG.Get("/pms", me.Conversations())
	// uses setPmsOptions middleware
	meG.Get("/pms/:other", me.Conversation(), setPmsOptions())
	meG.Get("/pms/:other/:pmid", me.Pm())

	// uses setPostlist middleware
	meG.Get("/posts", me.Posts(), setPostlist())
	// requests below uses the user.SetPost() middleware to refers to the requested post
	meG.Get("/posts/:pid", me.Post(), me.SetPost())
	// uses setCommentList middleware
	meG.Get("/posts/:pid/comments", me.PostComments(), me.SetPost(), setCommentList())
	meG.Get("/posts/:pid/comments/:cid", me.PostComment(), me.SetPost())

	/**************************************************************************
	* Stream API
	* ROUTE /stream/me
	* Authorization required
	***************************************************************************/
	s := basePath.Group("/stream/me")
	s.Use(authorization())
	// notification for current logged in user
	s.Get("/notifications", stream.Notifications())
	// TODO
	// /stream/users group
	//streamUsers := s.Group("/users/:id")
	// live update of current open profile
	//streamUsers.Get("/", stream.UserPosts())
	// life update of comments on current post
	//streamUsers.Get("/:pid/comments", stream.UserPostComments())

	return e
}
