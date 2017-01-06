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
	"errors"
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/rest"
	"github.com/nerdzeu/nerdz-api/rest/user"
	"net/http"
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

// UserFollowing handles the request and returns the user following
func UserFollowing() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.UserFollowing()(c)
	}
}

// ProjectFollowing handles the request and returns the project following
func ProjectFollowing() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.ProjectFollowing()(c)
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

// Whitelisting handles the request and returns the user whitelistings
func Whitelisting() echo.HandlerFunc {

	// swagger:route GET /me/whitelisting user info whitelisting getUserWhitelisted
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

// Blacklisting handles the request and returns the user blacklistings
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
		if !rest.IsGranted("pms:read", c) {
			return rest.InvalidScopeResponse("pms:read", c)
		}

		me := c.Get("me").(*nerdz.User)
		conversations, e := me.Conversations()
		if e != nil {
			errstr := "Unable to fetch conversations for the specified user"
			c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "other.Conversations error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		var conversationsTO []*nerdz.ConversationTO
		for _, c := range *conversations {
			conversationsTO = append(conversationsTO, c.GetTO(me))
		}
		return rest.SelectFields(conversationsTO, c)
	}
}

// Conversation handles the request and returns the user private conversation with the other use
func Conversation() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("pms:read", c) {
			return rest.InvalidScopeResponse("pms:read", c)
		}

		var other *nerdz.User
		var err error
		if other, err = rest.User("other", c); err != nil {
			return err
		}

		// fetch conversation between me and other
		var conversation *[]nerdz.Pm
		options := c.Get("pmsOptions").(*nerdz.PmsOptions)
		me := c.Get("me").(*nerdz.User)
		conversation, err = me.Pms(other.ID(), *options)

		if err != nil {
			errstr := "Unable to fetch conversation with the specified user"
			c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "me.Conversation error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}
		var conversationTO []*nerdz.PmTO
		for _, pm := range *conversation {
			conversationTO = append(conversationTO, pm.GetTO(me))
		}
		return rest.SelectFields(conversationTO, c)
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
		if !rest.IsGranted("pms:write", c) {
			return rest.InvalidScopeResponse("pms:write", c)
		}

		var other *nerdz.User
		var err error
		if other, err = rest.User("other", c); err != nil {
			return err
		}

		me := c.Get("me").(*nerdz.User)
		if err = me.DeleteConversation(other.ID()); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		message := "Success"
		c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: message,
			Message:      message,
			Status:       http.StatusOK,
			Success:      true,
		})
		return nil
	}
}

// Pm handles the request and returns the specified Private Message
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
		if !rest.IsGranted("pms:read", c) {
			return rest.InvalidScopeResponse("pms:read", c)
		}
		pm := c.Get("pm").(*nerdz.Pm)
		me := c.Get("me").(*nerdz.User)
		return rest.SelectFields(pm.GetTO(me), c)
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
		if !rest.IsGranted("pms:write", c) {
			return rest.InvalidScopeResponse("pms:write", c)
		}

		// Read a rest.NewMessage from the body request.
		message := rest.NewMessage{}
		if err := c.Bind(&message); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		var other *nerdz.User
		var err error
		if other, err = rest.User("other", c); err != nil {
			return err
		}

		// Create a nerdz.Pm from the message
		// and current context.
		pm := nerdz.Pm{}
		pm.Message = message.Message
		pm.Lang = message.Lang
		pm.To = other.ID()

		// Send it
		me := c.Get("me").(*nerdz.User)
		if err = me.Add(&pm); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}
		// Extract the TO from the new pm and return
		// selected fields.
		return rest.SelectFields(pm.GetTO(me), c)
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
		if !rest.IsGranted("pms:write", c) {
			return rest.InvalidScopeResponse("pms:write", c)
		}

		// Read a rest.NewMessage from the body request.
		message := rest.NewMessage{}
		if err := c.Bind(&message); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		// Update fields
		pm := c.Get("pm").(*nerdz.Pm)
		pm.Message = message.Message
		if message.Lang != "" {
			pm.Lang = message.Lang
		}

		// Edit
		me := c.Get("me").(*nerdz.User)
		if err := me.Edit(pm); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		// Extract the TO from the pm and return selected fields.
		return rest.SelectFields(pm.GetTO(me), c)
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
		if !rest.IsGranted("pms:write", c) {
			return rest.InvalidScopeResponse("pms:write", c)
		}
		pm := c.Get("pm").(*nerdz.Pm)
		me := c.Get("me").(*nerdz.User)
		if err := me.Delete(pm); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		message := "Success"
		c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: message,
			Message:      message,
			Status:       http.StatusOK,
			Success:      true,
		})
		return nil
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

// PostLurks handles the request and returns the post lurks
func PostLurks() echo.HandlerFunc {

	// swagger:route GET /me/posts/{pid}/lurks user post bookmarks GetUserPostLurks
	//
	// List the lurks of the post
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
		return user.PostLurks()(c)
	}
}

// NewPostLurk handles the request and creates a new lurk for the post
func NewPostLurk() echo.HandlerFunc {

	// swagger:route POST /me/posts/{pid}/lurks user post vote NewUserPostLurk
	//
	// Adds a new lurk on the current post
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
		return user.NewPostLurk()(c)
	}
}

// DeletePostLurk handles the request and deletes the lurk to the post
func DeletePostLurk() echo.HandlerFunc {

	// swagger:route DELETE /me/posts/{pid}/lurks user post vote DeleteUserPostLurk
	//
	// Deletes the lurk on the current post
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
		return user.DeletePostLurk()(c)
	}
}

// PostLock handles the request and and a lock to the post
func PostLock() echo.HandlerFunc {

	// swagger:route GET /me/posts/{pid}/lurks user post lurks GetUserPostLock
	//
	// List the locks of the post
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
		return user.PostLock()(c)
	}
}

// NewPostLock handles the request and creates a new lock for the post
func NewPostLock() echo.HandlerFunc {

	// swagger:route POST /me/posts/{pid}/locks user post vote NewUserPostLock
	//
	// Adds a new lock on the current post
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
		return user.NewPostLock()(c)
	}
}

// DeletePostLock handles the request and deletes the lock to the post
func DeletePostLock() echo.HandlerFunc {

	// swagger:route DELETE /me/posts/{pid}/locks user post vote DeleteUserPostLock
	//
	// Deletes the lock on the current post
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
		return user.DeletePostLock()(c)
	}
}

// NewPostUserLock handles the request and creates a new lock for the notification
// caused by the target user
func NewPostUserLock() echo.HandlerFunc {

	// swagger:route POST /me/posts/{pid}/locks/{target} user post vote NewUserNewPostUserLock
	//
	// Locks the notification from the target user to the current logged user, on the specified post
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
		return user.NewPostUserLock()(c)
	}
}

// DeletePostUserLock handles the request and deletes the lock for the notification of the target user
// on the specified post
func DeletePostUserLock() echo.HandlerFunc {

	// swagger:route DELETE /me/posts/{pid}/locks/{target} user post vote DeleteUserPostUserLock
	//
	// Deletes the lock  for the notification of the target user on the specified post
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
		return user.DeletePostUserLock()(c)
	}
}

// NewUserFollowing handles the request and creates and adds target to the following list of the current user
func NewUserFollowing() echo.HandlerFunc {

	// swagger:route POST /me/following/users/{target} userfollowing NewUserFollowing
	//
	// Adds target to the following list of the current user
	//
	//  Produces:
	//  - application/json
	//
	//  Security:
	//      oauth: following:write
	//
	//  Responses:
	//      default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("following:write", c) {
			return rest.InvalidScopeResponse("following:write", c)
		}

		var target *nerdz.User
		var err error
		if target, err = rest.User("target", c); err != nil {
			return err
		}
		me := c.Get("me").(*nerdz.User)
		if err = me.Follow(target); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}
		// Return selected field from the followed User
		return rest.SelectFields(target.GetTO(me), c)
	}
}

// DeleteUserFollowing handles the request and deletes the target user from the current user following list
func DeleteUserFollowing() echo.HandlerFunc {

	// swagger:route DELETE /me/following/users/{target} user following DeleteUserFollowing
	//
	// Deletes target user from the current user following list
	//
	// Consumes:
	// - application/json
	//
	//  Security:
	//      oauth: following:write
	//
	//  Responses:
	//      default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("following:write", c) {
			return rest.InvalidScopeResponse("following:write", c)
		}
		var target *nerdz.User
		var err error
		if target, err = rest.User("target", c); err != nil {
			return err
		}

		me := c.Get("me").(*nerdz.User)
		if err = me.Unfollow(target); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		message := "Success"
		c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: message,
			Message:      message,
			Status:       http.StatusOK,
			Success:      true,
		})
		return nil
	}
}

// NewProjectFollowing handles the request and creates and adds target to the following list of the current user
func NewProjectFollowing() echo.HandlerFunc {

	// swagger:route POST /me/following/projects/{target} project following NewProjectFollowing
	//
	// Adds target project to the following list of the current user
	//
	//  Produces:
	//  - application/json
	//
	//  Security:
	//      oauth: following:write
	//
	//  Responses:
	//      default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("following:write", c) {
			return rest.InvalidScopeResponse("following:write", c)
		}

		var target *nerdz.Project
		var err error
		if target, err = rest.Project("target", c); err != nil {
			return err
		}
		me := c.Get("me").(*nerdz.User)
		if err = me.Follow(target); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}
		// Return selected field from the followed Project
		return rest.SelectFields(target.GetTO(me), c)
	}
}

// DeleteProjectFollowing handles the request and deletes the target project from the current list following list
func DeleteProjectFollowing() echo.HandlerFunc {

	// swagger:route DELETE /me/following/users/{target} user DeleteProjectFollowing
	//
	// Deletes target project from the current user following list
	//
	// Consumes:
	// - application/json
	//
	//  Security:
	//      oauth: following:write
	//
	//  Responses:
	//      default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("following:write", c) {
			return rest.InvalidScopeResponse("following:write", c)
		}
		var target *nerdz.Project
		var err error
		if target, err = rest.Project("target", c); err != nil {
			return err
		}

		me := c.Get("me").(*nerdz.User)
		if err = me.Unfollow(target); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		message := "Success"
		c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: message,
			Message:      message,
			Status:       http.StatusOK,
			Success:      true,
		})
		return nil
	}
}

// NewWhitelisted handles the request and creates and adds target user to current user whitelist
func NewWhitelisted() echo.HandlerFunc {

	// swagger:route POST /me/whitelist/{target} user whitelist NewWhitelisted
	//
	// Adds target to the whitelist of the current user
	//
	//  Produces:
	//  - application/json
	//
	//  Security:
	//      oauth: profile:write
	//
	//  Responses:
	//      default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile:write", c) {
			return rest.InvalidScopeResponse("profile:write", c)
		}

		var target *nerdz.User
		var err error
		if target, err = rest.User("target", c); err != nil {
			return err
		}
		me := c.Get("me").(*nerdz.User)
		if err = me.WhitelistUser(target); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}
		// Return selected field from the followed User
		return rest.SelectFields(target.GetTO(me), c)
	}
}

// DeleteWhitelisted handles the request and deletes the target user from the current user whitelist
func DeleteWhitelisted() echo.HandlerFunc {

	// swagger:route DELETE /me/following/users/{target} user whitelist DeleteWhitelisted
	//
	// Deletes target user from the current user whitelist
	//
	// Consumes:
	// - application/json
	//
	//  Security:
	//      oauth: profile:write
	//
	//  Responses:
	//      default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile:write", c) {
			return rest.InvalidScopeResponse("profile:write", c)
		}
		var target *nerdz.User
		var err error
		if target, err = rest.User("target", c); err != nil {
			return err
		}

		me := c.Get("me").(*nerdz.User)
		if err = me.UnwhitelistUser(target); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		message := "Success"
		c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: message,
			Message:      message,
			Status:       http.StatusOK,
			Success:      true,
		})
		return nil
	}
}

// NewBlacklisted handles the request and creates and adds target user to current user whitelist
func NewBlacklisted() echo.HandlerFunc {

	// swagger:route POST /me/whitelist/{target} user whitelist NewBlacklisted
	//
	// Adds target to the whitelist of the current user
	//
	//  Produces:
	//  - application/json
	//
	//  Consumes:
	//  - application/json
	//
	//  Security:
	//      oauth: profile:write
	//
	//  Responses:
	//      default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile:write", c) {
			return rest.InvalidScopeResponse("profile:write", c)
		}

		var target *nerdz.User
		var err error
		if target, err = rest.User("target", c); err != nil {
			return err
		}

		// Read a rest.NewMessage from the body request.
		message := rest.NewMessage{}
		if err := c.Bind(&message); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		me := c.Get("me").(*nerdz.User)
		if err = me.BlacklistUser(target, message.Message); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}
		// Return selected field from the followed User
		return rest.SelectFields(target.GetTO(me), c)
	}
}

// DeleteBlacklisted handles the request and deletes the target user from the current user whitelist
func DeleteBlacklisted() echo.HandlerFunc {

	// swagger:route DELETE /me/following/users/{target} user blacklist DeleteBlacklisted
	//
	// Deletes target user from the current user blacklist
	//
	// Consumes:
	// - application/json
	//
	//  Security:
	//      oauth: profile:write
	//
	//  Responses:
	//      default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile:write", c) {
			return rest.InvalidScopeResponse("profile:write", c)
		}
		var target *nerdz.User
		var err error
		if target, err = rest.User("target", c); err != nil {
			return err
		}

		me := c.Get("me").(*nerdz.User)
		if err = me.UnblacklistUser(target); err != nil {
			errstr := err.Error()
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		message := "Success"
		c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: message,
			Message:      message,
			Status:       http.StatusOK,
			Success:      true,
		})
		return nil
	}
}
