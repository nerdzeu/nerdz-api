/*
Copyright (C) 2016-2020 Paolo Galeone <nessuno@nerdz.eu>

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
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/oauth2"
	"github.com/nerdzeu/nerdz-api/rest/me"
	"github.com/nerdzeu/nerdz-api/rest/project"
	"github.com/nerdzeu/nerdz-api/rest/user"
	"github.com/nerdzeu/nerdz-api/stream"
	"github.com/openshift/osin"
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
	o.POST("/authorize", oauth2.Authorize())
	o.GET("/token", oauth2.Token())
	o.POST("/token", oauth2.Token())
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
	usersG.GET("/:id/whitelist", user.Whitelist())
	usersG.GET("/:id/whitelisting", user.Whitelisting())
	usersG.GET("/:id/blacklist", user.Blacklist())
	usersG.GET("/:id/blacklisting", user.Blacklisting())
	usersG.GET("/:id/following/users", user.UserFollowing())
	usersG.GET("/:id/following/projects", user.ProjectFollowing())
	// uses setPostlist middleware
	usersG.GET("/:id/posts", user.Posts(), setPostlist())
	usersG.POST("/:id/posts", user.NewPost())
	// requests below uses the user.SetPost() middleware to refer to the requested post
	usersG.GET("/:id/posts/:pid", user.Post(), user.SetPost())
	usersG.PUT("/:id/posts/:pid", user.EditPost(), user.SetPost())
	usersG.DELETE("/:id/posts/:pid", user.DeletePost(), user.SetPost())
	// Votes
	usersG.GET("/:id/posts/:pid/votes", user.PostVotes(), user.SetPost())
	// Vote can be used to add/edit/delete the vote, just changing the vote value
	usersG.POST("/:id/posts/:pid/votes", user.NewPostVote(), user.SetPost())
	// Bookmark
	usersG.GET("/:id/posts/:pid/bookmarks", user.PostBookmarks(), user.SetPost())
	usersG.POST("/:id/posts/:pid/bookmarks", user.NewPostBookmark(), user.SetPost())
	usersG.DELETE("/:id/posts/:pid/bookmarks", user.DeletePostBookmark(), user.SetPost())
	// Lurk
	usersG.GET("/:id/posts/:pid/lurks", user.PostLurks(), user.SetPost())
	usersG.POST("/:id/posts/:pid/lurks", user.NewPostLurk(), user.SetPost())
	usersG.DELETE("/:id/posts/:pid/lurks", user.DeletePostLurk(), user.SetPost())
	// Lock
	usersG.GET("/:id/posts/:pid/locks", user.PostLock(), user.SetPost())
	usersG.POST("/:id/posts/:pid/locks", user.NewPostLock(), user.SetPost())
	usersG.DELETE("/:id/posts/:pid/locks", user.DeletePostLock(), user.SetPost())
	usersG.POST("/:id/posts/:pid/locks/:target", user.NewPostUserLock(), user.SetPost())
	usersG.DELETE("/:id/posts/:pid/locks/:target", user.DeletePostUserLock(), user.SetPost())
	// uses setCommentList middleware
	usersG.GET("/:id/posts/:pid/comments", user.PostComments(), user.SetPost(), setCommentList())
	usersG.POST("/:id/posts/:pid/comments", user.NewPostComment(), user.SetPost())
	// requests below uses user.SetComment middleware
	usersG.GET("/:id/posts/:pid/comments/:cid", user.PostComment(), user.SetPost(), user.SetComment())
	usersG.PUT("/:id/posts/:pid/comments/:cid", user.EditPostComment(), user.SetPost(), user.SetComment())
	usersG.DELETE("/:id/posts/:pid/comments/:cid", user.DeletePostComment(), user.SetPost(), user.SetComment())
	// Votes
	usersG.GET("/:id/posts/:pid/comments/:cid/votes", user.PostCommentVotes(), user.SetPost(), user.SetComment())
	usersG.POST("/:id/posts/:pid/comments/:cid/votes", user.NewPostCommentVote(), user.SetPost(), user.SetComment())

	/**************************************************************************
	* ROUTE /me
	* Authorization required
	***************************************************************************/
	meG := basePath.Group("/me")
	meG.Use(authorization())
	meG.Use(me.SetOther())
	// Read only
	meG.GET("", me.Info())
	meG.GET("/friends", me.Friends())
	meG.GET("/followers", me.Followers())
	// Read & write
	meG.GET("/following/users", me.UserFollowing())
	meG.POST("/following/users/:target", me.NewUserFollowing())
	meG.DELETE("/following/users/:target", me.DeleteUserFollowing())
	meG.GET("/following/projects", me.ProjectFollowing())
	meG.POST("/following/projects/:target", me.NewProjectFollowing())
	meG.DELETE("/following/projects/:target", me.DeleteProjectFollowing())
	meG.GET("/whitelist", me.Whitelist())
	meG.POST("/whitelist/:target", me.NewWhitelisted())
	meG.DELETE("/whitelist/:target", me.DeleteWhitelisted())
	// Read only
	meG.GET("/whitelisting", me.Whitelisting())
	// Read & write
	meG.GET("/blacklist", me.Blacklist())
	meG.POST("/blacklist/:target", me.NewBlacklisted())
	meG.DELETE("/blacklist/:target", me.DeleteBlacklisted())
	// Read only
	meG.GET("/blacklisting", me.Blacklisting())
	meG.GET("/home", me.Home(), setPostlist())
	meG.GET("/pms", me.Conversations())
	// uses setPmsOptions middleware
	meG.GET("/pms/:other", me.Conversation(), setPmsOptions())
	meG.POST("/pms/:other", me.NewPm())
	meG.DELETE("/pms/:other", me.DeleteConversation())
	// requests below uses the user.SetPm() middleware to refer to the requested pm
	meG.GET("/pms/:other/:pmid", me.Pm(), me.SetPm())
	// Disabled. At the moment pms' edit is not supported
	//meG.PUT("/pms/:other/:pmid", me.EditPm(), me.SetPm())
	meG.DELETE("/pms/:other/:pmid", me.DeletePm(), me.SetPm())

	// uses setPostlist middleware
	meG.GET("/posts", me.Posts(), setPostlist())
	meG.POST("/posts", me.NewPost())
	// requests below uses the user.SetPost() middleware to refer to the requested post
	meG.GET("/posts/:pid", me.Post(), me.SetPost())
	meG.PUT("/posts/:pid", me.EditPost(), me.SetPost())
	meG.DELETE("/posts/:pid", me.DeletePost(), me.SetPost())
	// Votes
	meG.GET("/posts/:pid/votes", me.PostVotes(), me.SetPost())
	// Vote can be used to add/edit/delete the vote, just changing the vote value
	meG.POST("/posts/:pid/votes", me.NewPostVote(), me.SetPost())
	// Bookmark
	meG.GET("/posts/:pid/bookmarks", me.PostBookmarks(), me.SetPost())
	meG.POST("/posts/:pid/bookmarks", me.NewPostBookmark(), me.SetPost())
	meG.DELETE("/posts/:pid/bookmarks", me.DeletePostBookmark(), me.SetPost())
	// Lurk
	meG.GET("/posts/:pid/lurks", me.PostLurks(), me.SetPost())
	meG.POST("/posts/:pid/lurks", me.NewPostLurk(), me.SetPost())
	meG.DELETE("/posts/:pid/lurks", me.DeletePostLurk(), me.SetPost())
	// Lock
	meG.GET("/posts/:pid/locks", me.PostLock(), me.SetPost())
	meG.POST("/posts/:pid/locks", me.NewPostLock(), me.SetPost())
	meG.DELETE("/posts/:pid/locks", me.DeletePostLock(), me.SetPost())
	meG.POST("/posts/:pid/locks/:target", me.NewPostUserLock(), me.SetPost())
	meG.DELETE("/posts/:pid/locks/:target", me.DeletePostUserLock(), me.SetPost())
	// uses setCommentList middleware
	meG.GET("/posts/:pid/comments", me.PostComments(), me.SetPost(), setCommentList())
	meG.POST("/posts/:pid/comments", me.NewPostComment(), me.SetPost())
	// requests below uses me.SetComment middleware
	meG.GET("/posts/:pid/comments/:cid", me.PostComment(), me.SetPost(), me.SetComment())
	meG.PUT("/posts/:pid/comments/:cid", me.EditPostComment(), me.SetPost(), me.SetComment())
	meG.DELETE("/posts/:pid/comments/:cid", me.DeletePostComment(), me.SetPost(), me.SetComment())
	// Votes
	meG.GET("/posts/:pid/comments/:cid/votes", me.PostCommentVotes(), me.SetPost(), me.SetComment())
	meG.POST("/posts/:pid/comments/:cid/votes", me.NewPostCommentVote(), me.SetPost(), me.SetComment())

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
	projectG.POST("/:id/posts", project.NewPost())
	// requests below uses the project.SetPost() middleware to refer to the requested post
	projectG.GET("/:id/posts/:pid", project.Post(), project.SetPost())
	projectG.PUT("/:id/posts/:pid", project.EditPost(), project.SetPost())
	projectG.DELETE("/:id/posts/:pid", project.DeletePost(), project.SetPost())
	// Votes
	projectG.GET("/:id/posts/:pid/votes", project.PostVotes(), project.SetPost())
	// Vote can be used to add/edit/delete the vote, just changing the vote value
	projectG.POST("/:id/posts/:pid/votes", project.NewPostVote(), project.SetPost())
	// Bookmark
	projectG.GET("/:id/posts/:pid/bookmarks", project.PostBookmarks(), project.SetPost())
	projectG.POST("/:id/posts/:pid/bookmarks", project.NewPostBookmark(), project.SetPost())
	projectG.DELETE("/:id/posts/:pid/bookmarks", project.DeletePostBookmark(), project.SetPost())
	// Lurk
	projectG.GET("/:id/posts/:pid/lurks", project.PostLurks(), project.SetPost())
	projectG.POST("/:id/posts/:pid/lurks", project.NewPostLurk(), project.SetPost())
	projectG.DELETE("/:id/posts/:pid/lurks", project.DeletePostLurk(), project.SetPost())
	// Lock
	projectG.GET("/:id/posts/:pid/locks", project.PostLock(), project.SetPost())
	projectG.POST("/:id/posts/:pid/locks", project.NewPostLock(), project.SetPost())
	projectG.DELETE("/:id/posts/:pid/locks", project.DeletePostLock(), project.SetPost())
	projectG.POST("/:id/posts/:pid/locks/:target", project.NewPostUserLock(), project.SetPost())
	projectG.DELETE("/:id/posts/:pid/locks/:target", project.DeletePostUserLock(), project.SetPost())
	// uses setCommentList middleware
	projectG.GET("/:id/posts/:pid/comments", project.PostComments(), project.SetPost(), setCommentList())
	projectG.POST("/:id/posts/:pid/comments", project.NewPostComment(), project.SetPost())
	// requests below uses project.SetComment middleware
	projectG.GET("/:id/posts/:pid/comments/:cid", project.PostComment(), project.SetPost(), project.SetComment())
	projectG.PUT("/:id/posts/:pid/comments/:cid", project.EditPostComment(), project.SetPost(), project.SetComment())
	projectG.DELETE("/:id/posts/:pid/comments/:cid", project.DeletePostComment(), project.SetPost(), project.SetComment())
	// Votes
	projectG.GET("/:id/posts/:pid/comments/:cid/votes", project.PostCommentVotes(), project.SetPost(), project.SetComment())
	projectG.POST("/:id/posts/:pid/comments/:cid/votes", project.NewPostCommentVote(), project.SetPost(), project.SetComment())

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
	// live update of comments on current post
	//streamUsers.GET("/:pid/comments", stream.UserPostComments())

	return e
}
