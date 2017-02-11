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

	// swagger:route GET /me/posts me posts GetMePosts
	//
	// List posts on user board, filtered by some parameters.
	//
	// This will show the last posts on the user board by default.
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_messages:read
	//
	//	Responses:
	//		default: MePosts

	return func(c echo.Context) error {
		return user.Posts()(c)
	}
}

// Post handles the request and returns the single post required
func Post() echo.HandlerFunc {

	// swagger:route GET /me/posts/{pid} me post GetMePost
	//
	// Shows selected posts with id pid on specified user board
	//
	// This will show the last comments on the post by default.
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_messages:read
	//
	//	Responses:
	//		default: MePostsPid

	return func(c echo.Context) error {
		return user.Post()(c)
	}
}

// NewPost handles the request and creates a new post
func NewPost() echo.HandlerFunc {

	// swagger:route POST /me/posts me post NewMePost
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
	//		default: MePostsPid

	return func(c echo.Context) error {
		return user.NewPost()(c)
	}
}

// EditPost handles the request and edits the post
func EditPost() echo.HandlerFunc {

	// swagger:route PUT /me/posts/{pid} me post EditMePost
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
	//		default: MePostsPid

	return func(c echo.Context) error {
		return user.EditPost()(c)
	}
}

// DeletePostComment handles the request and deletes the comment
func DeletePostComment() echo.HandlerFunc {

	// swagger:route DELETE /me/posts/{pid}/comments/{cid} me post DeleteMePostComment
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

	// swagger:route DELETE /me/posts/{pid}/comments/{cid} me post DeleteMePost
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

	// swagger:route PUT /me/posts/{pid}/comments/{cid} me post comment EditMeComment
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
	//		default: MePostsPidCommentsCid

	return func(c echo.Context) error {
		return user.EditPostComment()(c)
	}
}

// PostComments handles the request and returns the specified list of comments
func PostComments() echo.HandlerFunc {

	// swagger:route GET /me/posts/{pid}/comments me post comments GetMePostComments
	//
	// List comments on specified post, filtered by some parameters.
	//
	// This will show the last posts on the user board by default.
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_comments:read
	//
	//	Responses:
	//		default: MePostsPidComments

	return func(c echo.Context) error {
		return user.PostComments()(c)
	}
}

// PostComment handles the request and returns the single comment required
func PostComment() echo.HandlerFunc {

	// swagger:route GET /me/posts/{pid}/comments/{cid} me post comment GetMePostComment
	//
	// Shows selected comment on specified post, filtered by some parameters.
	//
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: profile_comments:read
	//
	//	Responses:
	//		default: MePostsPidCommentsCid

	return func(c echo.Context) error {
		return user.PostComment()(c)
	}
}

// NewPostComment handles the request and creates a new post
func NewPostComment() echo.HandlerFunc {

	// swagger:route POST /me/posts/{pid}/comments me post NewMePostComment
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
	//		default: MePostsPidCommentsCid

	return func(c echo.Context) error {
		return user.NewPostComment()(c)
	}
}

// Info handles the request and returns all the basic information for the specified user
func Info() echo.HandlerFunc {

	// swagger:route GET /me me info GetMeInfo
	//
	// Shows the basic informations for the specified user
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
	//		default: Me

	return func(c echo.Context) error {
		return user.Info()(c)
	}
}

// Friends handles the request and returns the user friends
func Friends() echo.HandlerFunc {

	// swagger:route GET /me/friends me info friends GetMeFriends
	//
	// Shows the friends informations for the specified user
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
	//		default: MeFriends

	return func(c echo.Context) error {
		return user.Friends()(c)
	}
}

// Followers handles the request and returns the user followers
func Followers() echo.HandlerFunc {

	// swagger:route GET /me/followers me info followers GetMeFollowers
	//
	// Shows the followers informations for the specified user
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
	//		default: MeFollowers

	return func(c echo.Context) error {
		return user.Followers()(c)
	}
}

// UserFollowing handles the request and returns the user following
func UserFollowing() echo.HandlerFunc {

	// swagger:route GET /me/following/users me info following GetMeFollowing
	//
	// Shows the following informations for the specified user
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
	//		default: MeFollowingUsers

	return func(c echo.Context) error {
		return user.UserFollowing()(c)
	}
}

// ProjectFollowing handles the request and returns the project following
func ProjectFollowing() echo.HandlerFunc {

	// swagger:route GET /me/following/projects project info following GetMeProjectFollowing
	//
	// Shows the following informations for the specified user
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
	//		default: MeFollowingProjects

	return func(c echo.Context) error {
		return user.ProjectFollowing()(c)
	}
}

// Whitelist handles the request and returns the user whitelist
func Whitelist() echo.HandlerFunc {

	// swagger:route GET /me/whitelist me info whitelist getMeWhitelist
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
	//		default: MeWhitelist

	return func(c echo.Context) error {
		return user.Whitelist()(c)
	}
}

// Whitelisting handles the request and returns the user whitelistings
func Whitelisting() echo.HandlerFunc {

	// swagger:route GET /me/whitelisting me info whitelisting getMeWhitelisted
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
	//		default: MeWhitelisting

	return func(c echo.Context) error {
		return user.Whitelisting()(c)
	}
}

// Blacklist handles the request and returns the user blacklist
func Blacklist() echo.HandlerFunc {

	// swagger:route GET /me/blacklist me info blacklist getMeBlacklist
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
	//		default: MeBlacklist

	return func(c echo.Context) error {
		return user.Blacklist()(c)
	}
}

// Blacklisting handles the request and returns the user blacklistings
func Blacklisting() echo.HandlerFunc {

	// swagger:route GET /me/blacklisting me info blacklisting getMeBlacklisting
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
	//		default: MeBlacklisting

	return func(c echo.Context) error {
		return user.Blacklisting()(c)
	}
}

// Home handles the request and returns the user home
func Home() echo.HandlerFunc {

	// swagger:route GET /me/home me post home getMeHome
	//
	// Shows the homepage of the current user, mixing projects and users posts
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
	//		default: MeHome

	return func(c echo.Context) error {
		if !rest.IsGranted("messages:read", c) {
			return rest.InvalidScopeResponse("messages:read", c)
		}

		me := c.Get("me").(*nerdz.User)
		options := c.Get("postlistOptions").(*nerdz.PostlistOptions)
		posts := me.Home(*options)

		if posts == nil {
			errstr := "Unable to fetch home page for the specified user"
			c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "me.Home error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		var postsAPI []*nerdz.PostTO
		for _, p := range *posts {
			postsAPI = append(postsAPI, p.GetTO(me))
		}

		return rest.SelectFields(postsAPI, c)
	}
}

// Conversations handles the request and returns the user private conversations
func Conversations() echo.HandlerFunc {

	// swagger:route GET /me/pms me post pms getMePms
	//
	// Shows the list of the private conversation of the current user
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
	//		default: MePms

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
				Message:      "me.Conversations error",
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

// Conversation handles the request and returns the user private conversation with the other user
func Conversation() echo.HandlerFunc {

	// swagger:route GET /me/pms/{other} me post pms getMeConversation
	//
	// Returns the private conversation of the current user with the other user
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
	//		default: MePmsOther

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

	// swagger:route DELETE /me/pms/{other} me pms DeleteMePms
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

	// swagger:route GET /me/pms/{other}/{pmid} me pms GetMePm
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
	//		default: MePmsOtherPmid

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

	// swagger:route POST /me/pms/{other} me pm NewMePm
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
	//		default: MePmsOtherPmid

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

	// swagger:route PUT /me/pms/{other}/{pmid} me pm EditMePm
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
	//		default: MePmsOtherPmid

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

	// swagger:route DELETE /me/pms/{other}/{pmid} me pm DeleteMePm
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

	// swagger:route GET /me/posts/{pid}/votes me post votes GetMePostVotes
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
	//		default: MePostsPidVotes

	return func(c echo.Context) error {
		return user.PostVotes()(c)
	}
}

// NewPostVote handles the request and creates a new vote for the post
func NewPostVote() echo.HandlerFunc {

	// swagger:route POST /me/posts/{pid}/votes me post vote NewMePostVote
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
	//		default: MePostsPidVotes

	return func(c echo.Context) error {
		return user.NewPostVote()(c)
	}
}

// PostCommentVotes handles the request and returns the comment votes
func PostCommentVotes() echo.HandlerFunc {

	// swagger:route GET /me/posts/{pid}/comments/{cid}/votes me post comments votes GetMePostCommentsVotes
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
	//		default: MePostsPidCommentsCidVotes

	return func(c echo.Context) error {
		return user.PostCommentVotes()(c)
	}
}

// NewPostCommentVote handles the request and creates a new vote on the user comment post
func NewPostCommentVote() echo.HandlerFunc {

	// swagger:route POST /me/posts/{pid}/comments/{cid}/votes me post comment vote NewMePostCommentVote
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
	//		default: MePostsPidCommentsCidVotes

	return func(c echo.Context) error {
		return user.NewPostCommentVote()(c)
	}
}

// PostBookmarks handles the request and returns the post bookmarks
func PostBookmarks() echo.HandlerFunc {

	// swagger:route GET /me/posts/{pid}/bookmarks me post bookmarks GetMePostBookmarks
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
	//		default: MePostsPidBookmarks

	return func(c echo.Context) error {
		return user.PostBookmarks()(c)
	}
}

// NewPostBookmark handles the request and creates a new bookmark for the post
func NewPostBookmark() echo.HandlerFunc {

	// swagger:route POST /me/posts/{pid}/bookmarks me post vote NewMePostBookmark
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
	//		default: MePostsPidBookmarks

	return func(c echo.Context) error {
		return user.NewPostBookmark()(c)
	}
}

// DeletePostBookmark handles the request and deletes the bookmark to the post
func DeletePostBookmark() echo.HandlerFunc {

	// swagger:route DELETE /me/posts/{pid}/bookmarks me post vote DeleteMePostBookmark
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

	// swagger:route GET /me/posts/{pid}/lurks me post bookmarks GetMePostLurks
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
	//		default: MePostsPidLurks

	return func(c echo.Context) error {
		return user.PostLurks()(c)
	}
}

// NewPostLurk handles the request and creates a new lurk for the post
func NewPostLurk() echo.HandlerFunc {

	// swagger:route POST /me/posts/{pid}/lurks me post vote NewMePostLurk
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
	//		default: MePostsPidLurks

	return func(c echo.Context) error {
		return user.NewPostLurk()(c)
	}
}

// DeletePostLurk handles the request and deletes the lurk to the post
func DeletePostLurk() echo.HandlerFunc {

	// swagger:route DELETE /me/posts/{pid}/lurks me post vote DeleteMePostLurk
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

	// swagger:route GET /me/posts/{pid}/locks me post locks GetMePostLock
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
	//		default: MePostsPidLocks

	return func(c echo.Context) error {
		return user.PostLock()(c)
	}
}

// NewPostLock handles the request and creates a new lock for the post
func NewPostLock() echo.HandlerFunc {

	// swagger:route POST /me/posts/{pid}/locks me post vote NewMePostLock
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
	//		default: MePostsPidLocks

	return func(c echo.Context) error {
		return user.NewPostLock()(c)
	}
}

// DeletePostLock handles the request and deletes the lock to the post
func DeletePostLock() echo.HandlerFunc {

	// swagger:route DELETE /me/posts/{pid}/locks me post vote DeleteMePostLock
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

	// swagger:route POST /me/posts/{pid}/locks/{target} me post vote NewMeNewPostUserLock
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
	//		default: MePostsPidLocks

	return func(c echo.Context) error {
		return user.NewPostUserLock()(c)
	}
}

// DeletePostUserLock handles the request and deletes the lock for the notification of the target user
// on the specified post
func DeletePostUserLock() echo.HandlerFunc {

	// swagger:route DELETE /me/posts/{pid}/locks/{target} me post vote DeleteMePostUserLock
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

	// swagger:route POST /me/following/users/{target} me userfollowing NewMeFollowing
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
	//      default: MeFollowingUsers

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

	// swagger:route DELETE /me/following/users/{target} me following DeleteMeFollowing
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
	//      default: MeFollowingProjects

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

	// swagger:route DELETE /me/following/users/{target} me DeleteProjectFollowing
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

	// swagger:route POST /me/whitelist/{target} me whitelist NewWhitelisted
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
	//      default: MeWhitelist

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

	// swagger:route DELETE /me/following/users/{target} me whitelist DeleteWhitelisted
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

	// swagger:route POST /me/blacklist/{target} me blacklist NewBlacklisted
	//
	// Adds target to the blacklist of the current user
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
	//      default: MeBlacklist

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

	// swagger:route DELETE /me/following/users/{target} me blacklist DeleteBlacklisted
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
