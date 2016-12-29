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

	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/rest"
)

// Posts handles the request and returns the required posts written by the specified user
func Posts() echo.HandlerFunc {

	// swagger:route GET /users/{id}/posts user posts GetUserPosts
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

	// swagger:route GET /users/{id}/posts/{pid} user post GetUserPost
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
		return rest.SelectFields(c.Get("post").(*nerdz.UserPost).GetTO(me), c)
	}
}

// NewPost handles the request and creates a new post
func NewPost() echo.HandlerFunc {

	// swagger:route POST /users/{id}/posts user post NewUserPost
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
		if !rest.IsGranted("profile_messages:write", c) {
			return rest.InvalidScopeResponse("profile_messages:write", c)
		}

		// Read a rest.NewMessage from the body request.
		message := rest.NewMessage{}
		if err := c.Bind(&message); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		// Create a nerdz.UserPost from the message
		// and current context.
		post := nerdz.UserPost{}
		post.Message = message.Message
		post.Lang = message.Lang
		post.To = c.Get("other").(*nerdz.User).ID()

		// Send it
		me := c.Get("me").(*nerdz.User)
		if err := me.Add(&post); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields(post.GetTO(me), c)
	}
}

// DeletePost handles the request and deletes the specified post
func DeletePost() echo.HandlerFunc {

	// swagger:route DELETE /users/{id}/posts/{pid} user post DeleteUserPost
	//
	// Delete the post on the specified user board
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
		if !rest.IsGranted("profile_messages:write", c) {
			return rest.InvalidScopeResponse("profile_messages:write", c)
		}

		post := c.Get("post").(*nerdz.UserPost)
		me := c.Get("me").(*nerdz.User)
		if err := me.Delete(post); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		message := "Success"
		return c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: message,
			Message:      message,
			Status:       http.StatusOK,
			Success:      true,
		})
	}
}

// EditPost handles the request and edits the post
func EditPost() echo.HandlerFunc {

	// swagger:route PUT /users/{id}/posts/{pid} user post EditUserPost
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
		if !rest.IsGranted("profile_messages:write", c) {
			return rest.InvalidScopeResponse("profile_messages:write", c)
		}

		// Read a rest.NewMessage from the body request.
		message := rest.NewMessage{}
		if err := c.Bind(&message); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}
		post := c.Get("post").(*nerdz.UserPost)

		// Update filds
		post.Message = message.Message
		if message.Lang != "" {
			post.Lang = message.Lang
		}

		// Edit
		me := c.Get("me").(*nerdz.User)
		if err := me.Edit(post); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		// Extract the TO from the post and return selected fields.
		return rest.SelectFields(post.GetTO(me), c)
	}
}

// PostComments handles the request and returns the specified list of comments
func PostComments() echo.HandlerFunc {

	// swagger:route GET /users/{id}/posts/{pid}/comments user post comments GetUserPostComments
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

	// swagger:route GET /users/{id}/posts/{pid}/comments/{cid} user post comment GetUserPostComment
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
		comment := c.Get("comment").(*nerdz.UserPostComment)
		me := c.Get("me").(*nerdz.User)
		return rest.SelectFields(comment.GetTO(me), c)
	}
}

// NewPostComment handles the request and creates a new post
func NewPostComment() echo.HandlerFunc {

	// swagger:route POST /users/{id}/posts/{pid}/comments user post comment NewUserPostComment
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
		if !rest.IsGranted("profile_comments:write", c) {
			return rest.InvalidScopeResponse("profile_comments:write", c)
		}

		// Read a rest.NewMessage from the body request.
		message := rest.NewMessage{}
		if err := c.Bind(&message); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		// Create a nerdz.UserPostComment from the message
		// and current context.
		comment := nerdz.UserPostComment{}
		comment.Message = message.Message
		comment.Lang = message.Lang
		comment.To = c.Get("other").(*nerdz.User).ID()
		comment.Hpid = c.Get("post").(*nerdz.UserPost).ID()

		// Send it
		me := c.Get("me").(*nerdz.User)
		if err := me.Add(&comment); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields(comment.GetTO(me), c)
	}
}

// EditPostComment handles the request and edits the post comment
func EditPostComment() echo.HandlerFunc {

	// swagger:route PUT /users/{id}/posts/{pid}/comments/{cid} user post comment EditUserPost
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
		if !rest.IsGranted("profile_comments:write", c) {
			return rest.InvalidScopeResponse("profile_comments:write", c)
		}

		// Read a rest.NewMessage from the body request.
		message := rest.NewMessage{}
		if err := c.Bind(&message); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}
		comment := c.Get("comment").(*nerdz.UserPostComment)

		// Update filds
		comment.Message = message.Message
		if comment.Lang != "" {
			comment.Lang = message.Lang
		}

		// Edit
		me := c.Get("me").(*nerdz.User)
		if err := me.Edit(comment); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		// Extract the TO from the comment and return selected fields.
		return rest.SelectFields(comment.GetTO(me), c)
	}
}

// DeletePostComment handles the request and deletes the specified
// comment on the speficied post
func DeletePostComment() echo.HandlerFunc {

	// swagger:route DELETE /users/{id}/posts/{pid}/comments/{cid} user post DeleteUserPostComment
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
		if !rest.IsGranted("profile_comments:write", c) {
			return rest.InvalidScopeResponse("profile_comments:write", c)
		}

		comment := c.Get("comment").(*nerdz.UserPostComment)
		me := c.Get("me").(*nerdz.User)
		if err := me.Delete(comment); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		message := "Success"
		return c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: message,
			Message:      message,
			Status:       http.StatusOK,
			Success:      true,
		})
	}
}

// Info handles the request and returns all the basic informations for the specified user
func Info() echo.HandlerFunc {

	// swagger:route GET /users/{id} user info GetUserInfo
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
		return rest.SelectFields(GetInfo(other), c)
	}
}

// Friends handles the request and returns the user friends
func Friends() echo.HandlerFunc {

	// swagger:route GET /users/{id}/friends user info friends GetUserFriends
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
		return rest.SelectFields(GetUsersInfo(friends), c)
	}
}

// Followers handles the request and returns the user followers
func Followers() echo.HandlerFunc {

	// swagger:route GET /users/{id}/followers user info followers GetUserFollowers
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
		return rest.SelectFields(GetUsersInfo(followers), c)
	}
}

// Following handles the request and returns the user following
func Following() echo.HandlerFunc {

	// swagger:route GET /users/{id}/following user info following GetUserFollowing
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
		return rest.SelectFields(GetUsersInfo(following), c)
	}
}

// Whitelist handles the request and returns the user whitelist
func Whitelist() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("profile:read", c) {
			return rest.InvalidScopeResponse("profile:read", c)
		}
		whitelist := c.Get("other").(*nerdz.User).Whitelist()
		return rest.SelectFields(GetUsersInfo(whitelist), c)
	}
}

// Whitelisting handles the request and returns the user whitelisting
func Whitelisting() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("profile:read", c) {
			return rest.InvalidScopeResponse("profile:read", c)
		}
		whitelisting := c.Get("other").(*nerdz.User).Whitelisting()
		return rest.SelectFields(GetUsersInfo(whitelisting), c)
	}
}

// Blacklist handles the request and returns the user blacklist
func Blacklist() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("profile:read", c) {
			return rest.InvalidScopeResponse("profile:read", c)
		}
		blacklist := c.Get("other").(*nerdz.User).Blacklist()
		return rest.SelectFields(GetUsersInfo(blacklist), c)
	}
}

// Blacklisting handles the request and returns the user blacklisting
func Blacklisting() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("profile:read", c) {
			return rest.InvalidScopeResponse("profile:read", c)
		}
		blacklisting := c.Get("other").(*nerdz.User).Blacklisting()
		return rest.SelectFields(GetUsersInfo(blacklisting), c)
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
		otherID := other.ID()
		// fetch conversation between c.Get("other") = "me" in the /me context
		// and the "other" user
		options := c.Get("pmsOptions").(*nerdz.PmsOptions)

		var conversation *[]nerdz.Pm
		var err error
		conversation, err = other.Pms(otherID, *options)

		if err != nil {
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

// DeleteConversation handles the request and delets the private conversation with the other user
func DeleteConversation() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rest.IsGranted("pms:write", c) {
			return rest.InvalidScopeResponse("pms:write", c)
		}

		// other is the owner of the pm list
		other := c.Get("other").(*nerdz.User)
		otherID := other.ID()
		if _, err := other.Pms(otherID, nerdz.PmsOptions{}); err != nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Unable to fetch conversation with the specified user",
				Message:      "other.Conversation error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		me := c.Get("me").(*nerdz.User)
		if err := me.DeleteConversation(otherID); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		message := "Success"
		return c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: message,
			Message:      message,
			Status:       http.StatusOK,
			Success:      true,
		})

	}
}

// Pm handles the request and returns the specified Private Message
func Pm() echo.HandlerFunc {

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
	return func(c echo.Context) error {
		if !rest.IsGranted("pms:write", c) {
			return rest.InvalidScopeResponse("pms:write", c)
		}

		// Read a rest.NewMessage from the body request.
		message := rest.NewMessage{}
		if err := c.Bind(&message); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		// Create a nerdz.Pm from the message
		// and current context.
		pm := nerdz.Pm{}
		pm.Message = message.Message
		pm.Lang = message.Lang
		pm.To = c.Get("other").(*nerdz.User).ID()

		// Send it
		me := c.Get("me").(*nerdz.User)
		if err := me.Add(&pm); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}
		// Extract the TO from the new pm and return
		// selected fields.
		return rest.SelectFields(pm.GetTO(me), c)
	}
}

// DeletePm handles the request and deletes the specified pm
func DeletePm() echo.HandlerFunc {

	return func(c echo.Context) error {
		if !rest.IsGranted("pms:write", c) {
			return rest.InvalidScopeResponse("pms:write", c)
		}
		pm := c.Get("pm").(*nerdz.Pm)
		me := c.Get("me").(*nerdz.User)
		if err := me.Delete(pm); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		message := "Success"
		return c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: message,
			Message:      message,
			Status:       http.StatusOK,
			Success:      true,
		})
	}
}

// EditPm handles the request and edits the pm
func EditPm() echo.HandlerFunc {

	return func(c echo.Context) error {
		if !rest.IsGranted("pms:write", c) {
			return rest.InvalidScopeResponse("pms:write", c)
		}

		// Read a rest.NewMessage from the body request.
		message := rest.NewMessage{}
		if err := c.Bind(&message); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		pm := c.Get("pm").(*nerdz.Pm)

		// Update fields
		pm.Message = message.Message
		if message.Lang != "" {
			pm.Lang = message.Lang
		}

		// Edit
		me := c.Get("me").(*nerdz.User)
		if err := me.Edit(pm); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		// Extract the TO from the pm and return selected fields.
		return rest.SelectFields(pm.GetTO(me), c)
	}
}

// PostVotes handles the request and returns the post votes
func PostVotes() echo.HandlerFunc {

	// swagger:route GET /users/{id}/posts/{pid}/votes user post votes GetUserPostVotes
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
		if !rest.IsGranted("profile_messages:read", c) {
			return rest.InvalidScopeResponse("profile_messages:read", c)
		}
		votes := c.Get("post").(*nerdz.UserPost).Votes()
		if votes == nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Unable to fetch votes for the specified post",
				Message:      "UserPost.Votes() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var votesTO []*nerdz.UserPostVoteTO
		me := c.Get("me").(*nerdz.User)
		for _, v := range *votes {
			// votes contains Vote elements
			// we need to convert back to a UserPostVote in order to get a correct UserPostVoteTO
			if userPostVote := v.(*nerdz.UserPostVote); userPostVote != nil {
				votesTO = append(votesTO, userPostVote.GetTO(me))
			}
		}
		return rest.SelectFields(votesTO, c)
	}
}

// NewPostVote handles the request and creates a new vote for the post
func NewPostVote() echo.HandlerFunc {

	// swagger:route POST /users/{id}/posts/{pid}/votes user post vote NewUserPostVote
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
		if !rest.IsGranted("profile_messages:write", c) {
			return rest.InvalidScopeResponse("profile_messages:write", c)
		}

		// Read a rest.NewVote from the body request.
		body := rest.NewVote{}
		if err := c.Bind(&body); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.UserPost)
		vote, err := me.Vote(post, body.Vote)
		if err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields(vote.(*nerdz.UserPostVote).GetTO(me), c)
	}
}

// PostCommentVotes handles the request and returns the comment votes
func PostCommentVotes() echo.HandlerFunc {

	// swagger:route GET /users/{id}/posts/{pid}/comments/{cid}/votes user post comments votes GetUserPostCommentsVotes
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
		if !rest.IsGranted("profile_comments:read", c) {
			return rest.InvalidScopeResponse("profile_comments:read", c)
		}
		votes := c.Get("comment").(*nerdz.UserPostComment).Votes()
		if votes == nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Unable to fetch votes for the specified post",
				Message:      "UserPostComment.Votes() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var votesTO []*nerdz.UserPostCommentVoteTO
		me := c.Get("me").(*nerdz.User)
		for _, v := range *votes {
			// votes contains Vote elements
			// we need to convert back to a UserPostCommentVote in order to get a correct UserPostCommentVoteTO
			if userPostCommentVote := v.(*nerdz.UserPostCommentVote); userPostCommentVote != nil {
				votesTO = append(votesTO, userPostCommentVote.GetTO(me))
			}
		}
		return rest.SelectFields(votesTO, c)
	}
}

// NewPostCommentVote handles the request and creates a new vote on the user comment post
func NewPostCommentVote() echo.HandlerFunc {

	// swagger:route POST /users/{id}/posts/{pid}/comments/{cid}votes user post comment vote NewUserPostCommentVote
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
		if !rest.IsGranted("profile_comments:write", c) {
			return rest.InvalidScopeResponse("profile_comments:write", c)
		}

		// Read a rest.NewVote from the body request.
		body := rest.NewVote{}
		if err := c.Bind(&body); err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		// Send it
		me := c.Get("me").(*nerdz.User)
		comment := c.Get("comment").(*nerdz.UserPostComment)
		vote, err := me.Vote(comment, body.Vote)
		if err != nil {
			errstr := err.Error()
			return c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields(vote.(*nerdz.UserPostCommentVote).GetTO(me), c)
	}
}
