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

package me

import (
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/rest/user"
)

// Posts handles the request and returns the required posts written by the specified user
func Posts() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Posts()(c)
	}
}

// Post handles the request and returns the single post required
func Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Post()(c)
	}
}

// NewPost handles the request and creates a new post
func NewPost() echo.HandlerFunc {

	// swagger:route POST /me/posts user post NewUserPost
	//
	// Creates a new post on the specified user board
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_messages:write
	//
	//	Responses:
	//		default: apiResponse
	return func(c echo.Context) error {
		return user.NewPost()(c)
	}
}

// EditPost handles the request and edits the post
func EditPost() echo.HandlerFunc {

	// swagger:route PUT /me/posts/{pid} user post EditUserPost
	//
	// Update the speficied post on the specified user board
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_messages:write
	//
	//	Responses:
	//		default: apiResponse
	return func(c echo.Context) error {
		return user.EditPost()(c)
	}
}

// DeletePostComment handles the request and deletes the comment
func DeletePostComment() echo.HandlerFunc {

	// swagger:route DELETE /me/posts/{pid}/comments/{cid} user post DeleteUserPostComment
	//
	// Delete the specified comment on the speficied user post
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_comments:write
	//
	//	Responses:
	//		default: apiResponse
	return func(c echo.Context) error {
		return user.DeletePostComment()(c)
	}
}

// DeletePost handles the request and deletes the post
func DeletePost() echo.HandlerFunc {

	// swagger:route DELETE /me/posts/{pid}/comments/{cid} user post DeleteUserPost
	//
	// Delete the specified comment on the speficied user post
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_comments:write
	//
	//	Responses:
	//		default: apiResponse
	return func(c echo.Context) error {
		return user.DeletePost()(c)
	}
}

// EditPostComment handles the request and edits the post comment
func EditPostComment() echo.HandlerFunc {

	// swagger:route PUT /posts/{pid}/comments/{cid} user post comment EditUserPost
	//
	// Update the speficied post on the specified user board
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_comments:write
	//
	//	Responses:
	//		default: apiResponse
	return func(c echo.Context) error {
		return user.EditPostComment()(c)
	}
}

// PostComments handles the request and returns the specified list of comments
func PostComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.PostComments()(c)
	}
}

// PostComment handles the request and returns the single comment required
func PostComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.PostComment()(c)
	}
}

// NewPostComment handles the request and creates a new post
func NewPostComment() echo.HandlerFunc {

	// swagger:route POST /me/posts/{pid}/comments user post NewUserPostComment
	//
	// Creates a new post on the specified user board
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_comments:write
	//
	//	Responses:
	//		default: apiResponse
	return func(c echo.Context) error {
		return user.NewPostComment()(c)
	}
}

// Info handles the request and returns all the basic information for the specified user
func Info() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Info()(c)
	}
}

// Friends handles the request and returns the user friends
func Friends() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Friends()(c)
	}
}

// Followers handles the request and returns the user followers
func Followers() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Followers()(c)
	}
}

// Following handles the request and returns the user following
func Following() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Following()(c)
	}
}

// Whitelist handles the request and returns the user whitelist
func Whitelist() echo.HandlerFunc {

	// swagger:route GET /me/whitelist user info whitelist getUserWhitelist
	//
	// Shows the whitelist informations for the current user
	//
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		return user.Whitelist()(c)
	}
}

// Whitelisting handles the request and returns the user whitelist
func Whitelisting() echo.HandlerFunc {

	// swagger:route GET /me/whitelisting user info whitelisting getUserWhitelisting
	//
	// Shows the whitelisting informations for the current user
	//
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		return user.Whitelisting()(c)
	}
}

// Blacklist handles the request and returns the user blacklist
func Blacklist() echo.HandlerFunc {

	// swagger:route GET /me/blacklist user info blacklist getUserBlacklist
	//
	// Shows the blacklist informations for the current user
	//
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		return user.Blacklist()(c)
	}
}

// Blacklisting handles the request and returns the user blacklist
func Blacklisting() echo.HandlerFunc {

	// swagger:route GET /me/blacklisting user info blacklisting getUserBlacklisting
	//
	// Shows the blacklisting informations for the current user
	//
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		return user.Blacklisting()(c)
	}
}

// Home handles the request and returns the user home
func Home() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Home()(c)
	}
}

// Conversations handles the request and returns the user private conversations
func Conversations() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Conversations()(c)
	}
}

// Conversation handles the request and returns the user private conversation with the other use
func Conversation() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Conversation()(c)
	}
}

// DeleteConversation handles the request and deletes the conversation
func DeleteConversation() echo.HandlerFunc {

	// swagger:route DELETE /me/pms/{other} user pms DeleteUserPms
	//
	// Delete the conversation beteen the current user and other
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: pms:write
	//
	//	Responses:
	//		default: apiResponse
	return func(c echo.Context) error {
		return user.DeleteConversation()(c)
	}
}

// Pm  handles the request and returns the specified Private Message
func Pm() echo.HandlerFunc {

	// swagger:route GET /me/pms/{other}/{pmid} user pms GetUserPm
	//
	// Update the speficied post on the specified user board
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: pms:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		return user.Pm()(c)
	}
}

// NewPm handles the request and creates a new pm
func NewPm() echo.HandlerFunc {

	// swagger:route POST /me/pms/{other} user pm NewUserPm
	//
	// Creates a new pm with from me to other user
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: pms:write
	//
	//	Responses:
	//		default: apiResponse
	return func(c echo.Context) error {
		return user.NewPm()(c)
	}
}

// EditPm handles the request and edits the pm
func EditPm() echo.HandlerFunc {

	// swagger:route PUT /me/pms/{other}/{pmid} user pm EditUserPm
	//
	// Update the speficied pm in the conversation with the other user
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: pms:write
	//
	//	Responses:
	//		default: apiResponse
	return func(c echo.Context) error {
		return user.EditPm()(c)
	}
}

// DeletePm handles the request and deletes the pm
func DeletePm() echo.HandlerFunc {

	// swagger:route DELETE /me/pms/{other}/{pmid} user pm DeleteUserPm
	//
	// Delete the speficied pm in the conversation with the other user
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: pms:write
	//
	//	Responses:
	//		default: apiResponse
	return func(c echo.Context) error {
		return user.DeletePm()(c)
	}
}

// PostVotes handles the request and returns the post votes
func PostVotes() echo.HandlerFunc {

	// swagger:route GET /me/posts/{pid}/votes user post votes GetUserPostVotes
	//
	// List the votes of the post
	//
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_messages:read
	//
	//	Responses:
	//		default: apiResponse
	return func(c echo.Context) error {
		return user.PostVotes()(c)
	}
}

// NewPostVote handles the request and creates a new vote for the post
func NewPostVote() echo.HandlerFunc {

	// swagger:route POST /me/posts/{pid}/votes user post vote NewUserPostVote
	//
	// Adds a new vote on the current post
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_messages:write
	//
	//	Responses:
	//		default: apiResponse
	return func(c echo.Context) error {
		return user.NewPostVote()(c)
	}
}

// PostCommentVotes handles the request and returns the comment votes
func PostCommentVotes() echo.HandlerFunc {

	// swagger:route GET /me/posts/{pid}/comments/{cid}/votes user post comments votes GetUserPostCommentsVotes
	//
	// List the votes on the comment
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_comments:read
	//
	//	Responses:
	//		default: apiResponse
	return func(c echo.Context) error {
		return user.PostCommentVotes()(c)
	}
}

// NewPostCommentVote handles the request and creates a new vote on the user comment post
func NewPostCommentVote() echo.HandlerFunc {

	// swagger:route POST /me/posts/{pid}/comments/{cid}votes user post comment vote NewUserPostCommentVote
	//
	// Adds a new vote on the current user post comment
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_comments:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		return user.NewPostCommentVote()(c)
	}
}

// PostBookmarks handles the request and returns the post bookmarks
func PostBookmarks() echo.HandlerFunc {

	// swagger:route GET /me/posts/{pid}/bookmarks user post bookmarks GetUserPostBookmarks
	//
	// List the bookmarks of the post
	//
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_messages:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		return user.PostBookmarks()(c)
	}
}

// NewPostBookmark handles the request and creates a new bookmark for the post
func NewPostBookmark() echo.HandlerFunc {

	// swagger:route POST /me/posts/{pid}/bookmarks user post vote NewUserPostBookmark
	//
	// Adds a new bookmark on the current post
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_messages:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		return user.NewPostBookmark()(c)
	}
}

// DeletePostBookmark handles the request and deletes the bookmark to the post
func DeletePostBookmark() echo.HandlerFunc {

	// swagger:route DELETE /me/posts/{pid}/bookmarks user post vote DeleteUserPostBookmark
	//
	// Deletes the bookmark on the current post
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_messages:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		return user.DeletePostBookmark()(c)
	}
}
