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
	"errors"
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
			errstr := "Unable to fetch post list for the specified user"
			c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "other.Postlist error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
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
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
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
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
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
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		errstr := "Success"
		c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: errstr,
			Message:      errstr,
			Status:       http.StatusOK,
			Success:      true,
		})
		return nil
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
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
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
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
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
			errstr := "Unable to fetch comment list for the specified post"
			c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "UserPost.Comments(options) error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
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
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
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
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
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
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
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
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
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
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		errstr := "Success"
		c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: errstr,
			Message:      errstr,
			Status:       http.StatusOK,
			Success:      true,
		})
		return nil
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
		return rest.SelectFields(rest.GetUserInfo(other), c)
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
		return rest.SelectFields(rest.GetUsersInfo(friends), c)
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
		return rest.SelectFields(rest.GetUsersInfo(followers), c)
	}
}

// UserFollowing handles the request and returns the user following
func UserFollowing() echo.HandlerFunc {

	// swagger:route GET /users/{id}/following/users user info following GetUserFollowing
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
		following := c.Get("other").(*nerdz.User).UserFollowing()
		return rest.SelectFields(rest.GetUsersInfo(following), c)
	}
}

// ProjectFollowing handles the request and returns the user following
func ProjectFollowing() echo.HandlerFunc {

	// swagger:route GET /users/{id}/following/projects project info following GetProjectFollowing
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
		following := c.Get("other").(*nerdz.User).ProjectFollowing()
		return rest.SelectFields(rest.GetProjectsInfo(following), c)
	}
}

// Whitelist handles the request and returns the user whitelist
func Whitelist() echo.HandlerFunc {

	// swagger:route GET /users/{id}/whitelist user whitelist GetWhitelist
	//
	// Show the whitelist of the specified user
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
		whitelist := c.Get("other").(*nerdz.User).Whitelist()
		return rest.SelectFields(rest.GetUsersInfo(whitelist), c)
	}
}

// Whitelisting handles the request and returns the user whitelisting
func Whitelisting() echo.HandlerFunc {

	// swagger:route GET /users/{id}/whitelisting user whitelisting GetWhitelisting
	//
	// Show the user that placed the specified user in their whitelist
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
		whitelisting := c.Get("other").(*nerdz.User).Whitelisting()
		return rest.SelectFields(rest.GetUsersInfo(whitelisting), c)
	}
}

// Blacklist handles the request and returns the user blacklist
func Blacklist() echo.HandlerFunc {

	// swagger:route GET /users/{id}/blacklist user blacklist GetWhitelist
	//
	// Show the blacklist of the specified user
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
		blacklist := c.Get("other").(*nerdz.User).Blacklist()
		return rest.SelectFields(rest.GetUsersInfo(blacklist), c)
	}
}

// Blacklisting handles the request and returns the user blacklisting
func Blacklisting() echo.HandlerFunc {

	// swagger:route GET /users/{id}/blacklisting user blacklisting GetWhitelisting
	//
	// Show the user that placed the specified user in their blacklist
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
		blacklisting := c.Get("other").(*nerdz.User).Blacklisting()
		return rest.SelectFields(rest.GetUsersInfo(blacklisting), c)
	}
}

// PostVotes handles the request and returns the post votes
func PostVotes() echo.HandlerFunc {

	// swagger:route GET /users/{id}/posts/{pid}/votes user post votes GetUserPostVotes
	//
	// List the votes of the post
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
			errstr := "Unable to fetch votes for the specified post"
			c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "UserPost.Votes() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
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
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.UserPost)
		vote, err := me.Vote(post, body.Vote)
		if err != nil {
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
			errstr := "Unable to fetch votes for the specified post"
			c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "UserPostComment.Votes() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
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

	// swagger:route POST /users/{id}/posts/{pid}/comments/{cid}/votes user post comment vote NewUserPostCommentVote
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
			c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		// Send it
		me := c.Get("me").(*nerdz.User)
		comment := c.Get("comment").(*nerdz.UserPostComment)
		vote, err := me.Vote(comment, body.Vote)
		if err != nil {
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
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields(vote.(*nerdz.UserPostCommentVote).GetTO(me), c)
	}
}

// PostBookmarks handles the request and returns the post bookmarks
func PostBookmarks() echo.HandlerFunc {

	// swagger:route GET /users/{id}/posts/{pid}/bookmarks user post bookmarks GetUserPostBookmarks
	//
	// List the bookmarks of the post
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
		bookmarks := c.Get("post").(*nerdz.UserPost).Bookmarks()
		if bookmarks == nil {
			errstr := "Unable to fetch bookmarks for the specified post"
			c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "UserPost.Bookmarks() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		var bookmarksTO []*nerdz.UserPostBookmarkTO
		me := c.Get("me").(*nerdz.User)
		for _, v := range *bookmarks {
			// bookmarks contains Boook elements
			// we need to convert back to a UserPostBookmark in order to get a correct UserPostBookmarkTO
			if userPostBookmark := v.(*nerdz.UserPostBookmark); userPostBookmark != nil {
				bookmarksTO = append(bookmarksTO, userPostBookmark.GetTO(me))
			}
		}
		return rest.SelectFields(bookmarksTO, c)
	}
}

// NewPostBookmark handles the request and creates a new bookmark for the post
func NewPostBookmark() echo.HandlerFunc {

	// swagger:route POST /users/{id}/posts/{pid}/bookmarks user post vote NewUserPostBookmark
	//
	// Adds a new bookmark on the current post
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

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.UserPost)
		bookmark, err := me.Bookmark(post)
		if err != nil {
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
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields(bookmark.(*nerdz.UserPostBookmark).GetTO(me), c)
	}
}

// DeletePostBookmark handles the request and deletes the bookmark to the post
func DeletePostBookmark() echo.HandlerFunc {

	// swagger:route DELETE /users/{id}/posts/{pid}/bookmarks user post vote DeleteUserPostBookmark
	//
	// Deletes the bookmark on the current post
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

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.UserPost)
		err := me.Unbookmark(post)
		if err != nil {
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

		errstr := "Success"
		c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: errstr,
			Message:      errstr,
			Status:       http.StatusOK,
			Success:      true,
		})
		return nil
	}
}

// PostLurks handles the request and returns the post lurks
func PostLurks() echo.HandlerFunc {

	// swagger:route GET /users/{id}/posts/{pid}/lurks user post bookmarks GetUserPostLurks
	//
	// List the lurks of the post
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
		lurks := c.Get("post").(*nerdz.UserPost).Lurks()
		if lurks == nil {
			errstr := "Unable to fetch lurks for the specified post"
			c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "UserPost.Lurks() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		var lurksTO []*nerdz.UserPostLurkTO
		me := c.Get("me").(*nerdz.User)
		for _, v := range *lurks {
			// lurks contains Lurk elements
			// we need to convert back to a UserPostLurk in order to get a correct UserPostLurkTO
			if userPostLurk := v.(*nerdz.UserPostLurk); userPostLurk != nil {
				lurksTO = append(lurksTO, userPostLurk.GetTO(me))
			}
		}
		return rest.SelectFields(lurksTO, c)
	}
}

// NewPostLurk handles the request and creates a new lurk for the post
func NewPostLurk() echo.HandlerFunc {

	// swagger:route POST /users/{id}/posts/{pid}/lurks user post vote NewUserPostLurk
	//
	// Adds a new lurk on the current post
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

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.UserPost)
		lurk, err := me.Lurk(post)
		if err != nil {
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
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields(lurk.(*nerdz.UserPostLurk).GetTO(me), c)
	}
}

// DeletePostLurk handles the request and deletes the lurk to the post
func DeletePostLurk() echo.HandlerFunc {

	// swagger:route DELETE /users/{id}/posts/{pid}/lurks user post vote DeleteUserPostLurk
	//
	// Deletes the lurk on the current post
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

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.UserPost)
		err := me.Unlurk(post)
		if err != nil {
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

		errstr := "Success"
		c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: errstr,
			Message:      errstr,
			Status:       http.StatusOK,
			Success:      true,
		})
		return nil
	}
}

// PostLock handles the request and and a lock to the post
func PostLock() echo.HandlerFunc {

	// swagger:route GET /users/{id}/posts/{pid}/lurks user post lurks GetUserPostLock
	//
	// List the locks of the post
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
		locks := c.Get("post").(*nerdz.UserPost).Locks()
		if locks == nil {
			errstr := "Unable to fetch locks for the specified post"
			c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "UserPost.Lock() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
			return errors.New(errstr)
		}

		var locksTO []*nerdz.UserPostLockTO
		me := c.Get("me").(*nerdz.User)
		for _, v := range *locks {
			// locks contains Lock elements
			// we need to convert back to a UserPostLock in order to get a correct UserPostLockTO
			if userPostLock := v.(*nerdz.UserPostLock); userPostLock != nil {
				locksTO = append(locksTO, userPostLock.GetTO(me))
			}
		}
		return rest.SelectFields(locksTO, c)
	}
}

// NewPostLock handles the request and creates a new lock for the post
func NewPostLock() echo.HandlerFunc {

	// swagger:route POST /users/{id}/posts/{pid}/locks user post vote NewUserPostLock
	//
	// Adds a new lock on the current post
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

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.UserPost)
		lock, err := me.Lock(post)
		if err != nil {
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
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields((*lock)[0].(*nerdz.UserPostLock).GetTO(me), c)
	}
}

// DeletePostLock handles the request and deletes the lock to the post
func DeletePostLock() echo.HandlerFunc {

	// swagger:route DELETE /users/{id}/posts/{pid}/locks user post vote DeleteUserPostLock
	//
	// Deletes the lock on the current post
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

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.UserPost)
		err := me.Unlock(post)
		if err != nil {
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

		errstr := "Success"
		c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: errstr,
			Message:      errstr,
			Status:       http.StatusOK,
			Success:      true,
		})
		return nil
	}
}

// NewPostUserLock handles the request and creates a new lock for the notification
// caused by the target user
func NewPostUserLock() echo.HandlerFunc {

	// swagger:route POST /users/{id}/posts/{pid}/locks/{target} user post vote NewUserNewPostUserLock
	//
	// Locks the notification from the target user to the current logged user, on the specified post
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

		var target *nerdz.User
		var err error
		if target, err = rest.User("target", c); err != nil {
			return err
		}
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.UserPost)
		var lock *[]nerdz.Lock
		lock, err = me.Lock(post, target)
		if err != nil {
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
		// Extract the TO from the new lock and return selected fields.
		return rest.SelectFields((*lock)[0].(*nerdz.UserPostUserLock).GetTO(me), c)
	}
}

// DeletePostUserLock handles the request and deletes the lock for the notification of the target user
// on the specified post
func DeletePostUserLock() echo.HandlerFunc {

	// swagger:route DELETE /users/{id}/posts/{pid}/locks/{target} user post vote DeleteUserPostUserLock
	//
	// Deletes the lock for the notification of the target user on the specified post
	//
	// Consumes:
	// - application/json
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
		var target *nerdz.User
		var err error
		if target, err = rest.User("target", c); err != nil {
			return err
		}

		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.UserPost)
		err = me.Unlock(post, target)
		if err != nil {
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

		errstr := "Success"
		c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: errstr,
			Message:      errstr,
			Status:       http.StatusOK,
			Success:      true,
		})
		return nil
	}
}
