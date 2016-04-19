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
	"github.com/nerdzeu/nerdz-api/rest/project"
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
	o.GET("/authorize", oauth2.Authorize())
	o.Post("/authorize", oauth2.Authorize())
	o.GET("/token", oauth2.Token())
	o.GET("/info", oauth2.Info())

	/**************************************************************************
	* ROUTE /users/:id
	* Authorization required
	***************************************************************************/
	usersG := basePath.Group("/users") // users Group
	usersG.Use(authorization())
	usersG.Use(user.SetOther())
	usersG.GET("/:id", user.Info())
	usersG.GET("/:id/friends", user.Friends())
	usersG.GET("/:id/followers", user.Followers())
	usersG.GET("/:id/following", user.Following())
	// uses setPostlist middleware
	usersG.GET("/:id/posts", user.Posts(), setPostlist())
	// requests below uses the user.SetPost() middleware to refers to the requested post
	usersG.GET("/:id/posts/:pid", user.Post(), user.SetPost())
	// uses setCommentList middleware
	usersG.GET("/:id/posts/:pid/comments", user.PostComments(), user.SetPost(), setCommentList())
	usersG.GET("/:id/posts/:pid/comments/:cid", user.PostComment(), user.SetPost())

	/**************************************************************************
	* ROUTE /me
	* Authorization required
	***************************************************************************/
	meG := basePath.Group("/me")
	meG.Use(authorization())
	meG.Use(me.SetOther())
	meG.GET("", me.Info())
	meG.GET("/friends", me.Friends())
	meG.GET("/followers", me.Followers())
	meG.GET("/following", me.Following())
	meG.GET("/whitelist", me.Whitelist())
	meG.GET("/whitelisting", me.Whitelisting())
	meG.GET("/blacklist", me.Blacklist())
	meG.GET("/blacklisting", me.Blacklisting())
	meG.GET("/home", me.Home(), setPostlist())
	meG.GET("/pms", me.Conversations())
	// uses setPmsOptions middleware
	meG.GET("/pms/:other", me.Conversation(), setPmsOptions())
	meG.GET("/pms/:other/:pmid", me.Pm())

	// uses setPostlist middleware
	meG.GET("/posts", me.Posts(), setPostlist())
	// requests below uses the user.SetPost() middleware to refers to the requested post
	meG.GET("/posts/:pid", me.Post(), me.SetPost())
	// uses setCommentList middleware
	meG.GET("/posts/:pid/comments", me.PostComments(), me.SetPost(), setCommentList())
	meG.GET("/posts/:pid/comments/:cid", me.PostComment(), me.SetPost())

	/**************************************************************************
	* ROUTE /projects/:id
	* Authorization required
	***************************************************************************/
	projectG := basePath.Group("/projects") // users Group
	projectG.Use(authorization())
	projectG.Use(project.SetProject())
	projectG.GET("/:id", project.Info())
	projectG.GET("/:id/members", project.Members())
	projectG.GET("/:id/followers", project.Followers())
	// uses setPostlist middleware
	projectG.GET("/:id/posts", project.Posts(), setPostlist())
	// requests below uses the project.SetPost() middleware to refers to the requested post
	projectG.GET("/:id/posts/:pid", project.Post(), project.SetPost())
	// uses setCommentList middleware
	projectG.GET("/:id/posts/:pid/comments", project.PostComments(), project.SetPost(), setCommentList())
	projectG.GET("/:id/posts/:pid/comments/:cid", project.PostComment(), project.SetPost())

	/**************************************************************************
	* Stream API
	* ROUTE /stream/me
	* Authorization required
	***************************************************************************/
	s := basePath.Group("/stream/me")
	s.Use(authorization())
	// notification for current logged in user
	s.GET("/notifications", stream.Notifications())
	// TODO
	// /stream/users group
	//streamUsers := s.Group("/users/:id")
	// live update of current open profile
	//streamUsers.GET("/", stream.UserPosts())
	// life update of comments on current post
	//streamUsers.GET("/:pid/comments", stream.UserPostComments())

	return e
}
