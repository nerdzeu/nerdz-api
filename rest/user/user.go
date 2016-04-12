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

package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/rest"
)

// Posts handles the request and returns the required posts written by the specified user
func Posts() echo.HandlerFunc {

	// swagger:route GET /users/{id}/posts user posts getUseosts
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
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile_messages:read", c) {
			return rest.InvalidScopeResponse("profile_messages:read", c)
		}

		other := c.Get("other").(*nerdz.User)
		options := c.Get("postlistOptions").(*nerdz.PostlistOptions)
		options.Model = nerdz.UserPost{}
		posts := other.Postlist(*options)

		if posts == nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Unable to fetch post list for the specified user",
				Message:      "other.Postlist error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		me := c.Get("me").(*nerdz.User)
		var postsAPI []*nerdz.PostTO
		for _, p := range *posts {
			// posts contains ExistingPost elements
			// we need to convert back to a UserPost in order to get a correct PostTO
			if userPost := p.(*nerdz.UserPost); userPost != nil {
				postsAPI = append(postsAPI, userPost.GetTO(me))
			}
		}

		return rest.SelectFields(postsAPI, c)
	}
}

// Post handles the request and returns the single post required
func Post() echo.HandlerFunc {

	// swagger:route GET /users/{id}/posts/{pid} user post getUserPost
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
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile_messages:read", c) {
			return rest.InvalidScopeResponse("profile_messages:read", c)
		}
		me := c.Get("me").(*nerdz.User)
		postTO := c.Get("post").(*nerdz.UserPost).GetTO(me)
		return rest.SelectFields(postTO, c)
	}
}

// PostComments handles the request and returns the specified list of comments
func PostComments() echo.HandlerFunc {

	// swagger:route GET /users/{id}/posts/{pid}/comments user post comments getUserPostComments
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
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile_comments:read", c) {
			return rest.InvalidScopeResponse("profile_comments:read", c)
		}
		comments := c.Get("post").(*nerdz.UserPost).Comments(*(c.Get("commentlistOptions").(*nerdz.CommentlistOptions)))
		if comments == nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Unable to fetch comment list for the specified post",
				Message:      "UserPost.Comments(options) error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var commentsAPI []*nerdz.UserPostCommentTO
		me := c.Get("me").(*nerdz.User)
		for _, p := range *comments {
			// comments contains ExistingPost elements
			// we need to convert back to a UserPostComment in order to get a correct UserPostCommentTO
			if userPostComment := p.(*nerdz.UserPostComment); userPostComment != nil {
				commentsAPI = append(commentsAPI, userPostComment.GetTO(me))
			}
		}
		return rest.SelectFields(commentsAPI, c)
	}
}

// PostComment handles the request and returns the single comment required
func PostComment() echo.HandlerFunc {

	// swagger:route GET /users/{id}/posts/{pid}/comments/{cid} user post comment getUserPostComment
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
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile_comments:read", c) {
			return rest.InvalidScopeResponse("profile_comments:read", c)
		}
		var cid uint64
		var e error
		if cid, e = strconv.ParseUint(c.Param("cid"), 10, 64); e != nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Invalid comment identifier specified",
				Message:      e.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var comment *nerdz.UserPostComment
		if comment, e = nerdz.NewUserPostComment(cid); e != nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Invalid comment identifier specified",
				Message:      e.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		post := c.Get("post").(*nerdz.UserPost)
		if comment.Hpid != post.Hpid {
			message := "Mismatch between comment ID and post ID. Comment not related to the post"
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: message,
				Message:      message,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		me := c.Get("me").(*nerdz.User)
		return rest.SelectFields(comment.GetTO(me), c)
	}
}

// Info handles the request and returns all the basic informations for the specified user
func Info() echo.HandlerFunc {

	// swagger:route GET /users/{id} user info getUserInfo
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
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile:read", c) {
			return rest.InvalidScopeResponse("profile:read", c)
		}
		other := c.Get("other").(*nerdz.User)
		return rest.SelectFields(getInfo(other), c)
	}
}

// Friends handles the request and returns the user friends
func Friends() echo.HandlerFunc {

	// swagger:route GET /users/{id}/friends user info friends getUserFriends
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
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile:read", c) {
			return rest.InvalidScopeResponse("profile:read", c)
		}
		friends := c.Get("other").(*nerdz.User).Friends()
		return rest.SelectFields(getUsersInfo(friends), c)
	}
}

// Followers handles the request and returns the user followers
func Followers() echo.HandlerFunc {

	// swagger:route GET /users/{id}/followers user info followers getUserFollowers
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
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile:read", c) {
			return rest.InvalidScopeResponse("profile:read", c)
		}
		followers := c.Get("other").(*nerdz.User).Followers()
		return rest.SelectFields(getUsersInfo(followers), c)
	}
}

// Following handles the request and returns the user following
func Following() echo.HandlerFunc {

	// swagger:route GET /users/{id}/following user info following getUserFollowing
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
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile:read", c) {
			return rest.InvalidScopeResponse("profile:read", c)
		}
		following := c.Get("other").(*nerdz.User).Following()
		return rest.SelectFields(getUsersInfo(following), c)
	}
}

// Whitelist handles the request and returns the user whitelist
func Whitelist() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("profile:read", c) {
			return rest.InvalidScopeResponse("profile:read", c)
		}
		whitelist := c.Get("other").(*nerdz.User).Whitelist()
		return rest.SelectFields(getUsersInfo(whitelist), c)
	}
}

// Whitelisting handles the request and returns the user whitelisting
func Whitelisting() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("profile:read", c) {
			return rest.InvalidScopeResponse("profile:read", c)
		}
		whitelisting := c.Get("other").(*nerdz.User).Whitelisting()
		return rest.SelectFields(getUsersInfo(whitelisting), c)
	}
}

// Blacklist handles the request and returns the user blacklist
func Blacklist() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("profile:read", c) {
			return rest.InvalidScopeResponse("profile:read", c)
		}
		blacklist := c.Get("other").(*nerdz.User).Blacklist()
		return rest.SelectFields(getUsersInfo(blacklist), c)
	}
}

// Blacklisting handles the request and returns the user blacklisting
func Blacklisting() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("profile:read", c) {
			return rest.InvalidScopeResponse("profile:read", c)
		}
		blacklisting := c.Get("other").(*nerdz.User).Blacklisting()
		return rest.SelectFields(getUsersInfo(blacklisting), c)
	}
}

// Home handles the request and returns the user home
func Home() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("messages:read", c) {
			return rest.InvalidScopeResponse("messages:read", c)
		}

		other := c.Get("other").(*nerdz.User)
		options := c.Get("postlistOptions").(*nerdz.PostlistOptions)
		posts := other.Home(*options)

		if posts == nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Unable to fetch home page for the specified user",
				Message:      "other.Home error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		me := c.Get("me").(*nerdz.User)
		var postsAPI []*nerdz.PostTO
		for _, p := range *posts {
			postsAPI = append(postsAPI, p.GetTO(me))
		}

		return rest.SelectFields(postsAPI, c)
	}
}

// Conversations handles the request and returns the user private conversations
func Conversations() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("pms:read", c) {
			return rest.InvalidScopeResponse("pms:read", c)
		}

		other := c.Get("other").(*nerdz.User)
		conversations, e := other.Conversations()

		if e != nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Unable to fetch conversations for the specified user",
				Message:      "other.Conversations error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var conversationsTO []*nerdz.ConversationTO
		me := c.Get("me").(*nerdz.User)
		for _, c := range *conversations {
			conversationsTO = append(conversationsTO, c.GetTO(me))
		}

		return rest.SelectFields(conversationsTO, c)
	}
}

// Conversation handles the request and returns the private conversation with the other user
func Conversation() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("pms:read", c) {
			return rest.InvalidScopeResponse("pms:read", c)
		}

		// other is the owner of the pm list
		other := c.Get("other").(*nerdz.User)
		// otherID is the ID of the second actor in the conversation
		var otherID uint64
		var e error
		if otherID, e = strconv.ParseUint(c.Param("other"), 10, 64); e != nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Invalid user identifier specified",
				Message:      e.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		// fetch conversation between c.Get("other") = "me" in the /me context
		// and the "other" user
		options := c.Get("pmsOptions").(*nerdz.PmsOptions)

		var conversation *[]nerdz.Pm
		conversation, e = other.Pms(otherID, *options)

		if e != nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Unable to fetch conversation with the specified user",
				Message:      "other.Conversation error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var conversationTO []*nerdz.PmTO
		me := c.Get("me").(*nerdz.User)
		for _, pm := range *conversation {
			conversationTO = append(conversationTO, pm.GetTO(me))
		}

		return rest.SelectFields(conversationTO, c)
	}
}

// Pm handles the request and returns the specified Private Message
func Pm() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("pms:read", c) {
			return rest.InvalidScopeResponse("pms:read", c)
		}
		// other is the owner of the pm list
		other := c.Get("other").(*nerdz.User)
		// otherID is the ID of the second actor in the conversation
		var otherID, pmID uint64
		var e error
		if otherID, e = strconv.ParseUint(c.Param("other"), 10, 64); e != nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Invalid user identifier specified",
				Message:      e.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		if pmID, e = strconv.ParseUint(c.Param("pmid"), 10, 64); e != nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Invalid PM identifier specified",
				Message:      e.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var pm *nerdz.Pm
		if pm, e = nerdz.NewPm(pmID); e != nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: e.Error(),
				Message:      e.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		me := c.Get("me").(*nerdz.User)
		if (pm.From == otherID && pm.To == other.ID()) || (pm.From == other.ID() && pm.To == otherID) {
			return rest.SelectFields(pm.GetTO(me), c)
		}

		message := "You're not autorized to see the requested PM"
		return c.JSON(http.StatusUnauthorized, &rest.Response{
			HumanMessage: message,
			Message:      message,
			Status:       http.StatusUnauthorized,
			Success:      false,
		})
	}
}
