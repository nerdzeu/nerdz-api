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

package project

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/rest"
	"github.com/nerdzeu/nerdz-api/rest/user"
)

// Posts handles the request and returns the required posts written by the specified project
func Posts() echo.HandlerFunc {

	// swagger:route GET /projects/{id}/posts project posts getProjectPosts
	//
	// List posts on project board, filtered by some parameters.
	//
	// This will show the last posts on the project board by default.
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_messages:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:read", c) {
			return rest.InvalidScopeResponse("project_messages:read", c)
		}

		project := c.Get("project").(*nerdz.Project)
		options := c.Get("postlistOptions").(*nerdz.PostlistOptions)
		options.Model = nerdz.ProjectPost{}
		posts := project.Postlist(*options)

		if posts == nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Unable to fetch post list for the specified project",
				Message:      "project.Postlist error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		me := c.Get("me").(*nerdz.User)
		var postsAPI []*nerdz.PostTO
		for _, p := range *posts {
			// posts contains ExistingPost elements
			// we need to convert back to a ProjectPost in order to get a correct PostTO
			if projectPost := p.(*nerdz.ProjectPost); projectPost != nil {
				postsAPI = append(postsAPI, projectPost.GetTO(me))
			}
		}

		return rest.SelectFields(postsAPI, c)
	}
}

// Post handles the request and returns the single post required
func Post() echo.HandlerFunc {

	// swagger:route GET /projects/{id}/posts/{pid} project post getProjectPost
	//
	// Shows selected posts with id pid on specified project board
	//
	// This will show the last comments on the post by default.
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_messages:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:read", c) {
			return rest.InvalidScopeResponse("project_messages:read", c)
		}
		me := c.Get("me").(*nerdz.User)
		postTO := c.Get("post").(*nerdz.ProjectPost).GetTO(me)
		return rest.SelectFields(postTO, c)
	}
}

// NewPost handles the request and creates a new post
func NewPost() echo.HandlerFunc {

	// swagger:route POST /projects/{id}/posts project post NewProjectPost
	//
	// Creates a new post on the specified project board
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_messages:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:write", c) {
			return rest.InvalidScopeResponse("project_messages:write", c)
		}

		// Read a rest.Message from the body request.
		message := rest.NewMessage{}
		if err := c.Bind(&message); err != nil {
			return err
		}

		// Create a nerdz.ProjectPost from the message
		// and current context.
		post := nerdz.ProjectPost{}
		post.Message = message.Message
		post.Lang = message.Lang
		post.To = c.Get("project").(*nerdz.Project).ID()

		// Send it
		me := c.Get("me").(*nerdz.User)
		if err := me.Add(&post); err != nil {
			return err
		}
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields(post.GetTO(me), c)
	}
}

// DeletePost handles the request and deletes the specified post
func DeletePost() echo.HandlerFunc {

	// swagger:route DELETE /projects/{id}/posts/{pid} project post DeleteProjectPost
	//
	// Delete the post on the specified project board
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_messages:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:write", c) {
			return rest.InvalidScopeResponse("project_messages:write", c)
		}

		post := c.Get("post").(*nerdz.ProjectPost)
		me := c.Get("me").(*nerdz.User)
		if err := me.Delete(post); err != nil {
			return err
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

	// swagger:route PUT /projects/{id}/posts/{pid} project post EditProjectPost
	//
	// Update the speficied post on the specified project board
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_messages:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:write", c) {
			return rest.InvalidScopeResponse("project_messages:write", c)
		}

		// Read a rest.Message from the body request.
		message := rest.NewMessage{}
		if err := c.Bind(&message); err != nil {
			return err
		}
		post := c.Get("post").(*nerdz.ProjectPost)

		// Update fields
		post.Message = message.Message
		if message.Lang != "" {
			post.Lang = message.Lang
		}

		// Edit
		me := c.Get("me").(*nerdz.User)
		if err := me.Edit(post); err != nil {
			return err
		}

		// Extract the TO from the post and return selected fields.
		return rest.SelectFields(post.GetTO(me), c)
	}
}

// PostComments handles the request and returns the specified list of comments
func PostComments() echo.HandlerFunc {

	// swagger:route GET /projects/{id}/posts/{pid}/comments project post comments getProjectPostComments
	//
	// List comments on specified post, filtered by some parameters.
	//
	// This will show the last posts on the project board by default.
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_comments:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_comments:read", c) {
			return rest.InvalidScopeResponse("project_comments:read", c)
		}
		comments := c.Get("post").(*nerdz.ProjectPost).Comments(*(c.Get("commentlistOptions").(*nerdz.CommentlistOptions)))
		if comments == nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Unable to fetch comment list for the specified post",
				Message:      "ProjectPost.Comments(options) error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var commentsAPI []*nerdz.ProjectPostCommentTO
		me := c.Get("me").(*nerdz.User)
		for _, p := range *comments {
			// comments contains ExistingPost elements
			// we need to convert back to a ProjectPostComment in order to get a correct ProjectPostCommentTO
			if projectPostComment := p.(*nerdz.ProjectPostComment); projectPostComment != nil {
				commentsAPI = append(commentsAPI, projectPostComment.GetTO(me))
			}
		}
		return rest.SelectFields(commentsAPI, c)
	}
}

// PostComment handles the request and returns the single comment required
func PostComment() echo.HandlerFunc {

	// swagger:route GET /projects/{id}/posts/{pid}/comments/{cid} project post comment GetProjectPostComment
	//
	// Shows selected comment on specified post, filtered by some parameters.
	//
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_comments:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_comments:read", c) {
			return rest.InvalidScopeResponse("project_comments:read", c)
		}
		comment := c.Get("comment").(*nerdz.ProjectPostComment)
		me := c.Get("me").(*nerdz.User)
		return rest.SelectFields(comment.GetTO(me), c)
	}
}

// NewPostComment handles the request and creates a new post
func NewPostComment() echo.HandlerFunc {

	// swagger:route POST /projects/{id}/posts/{pid}/comments project post comment NewProjectPostComment
	//
	// Creates a new post on the specified project board
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_comments:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_comments:write", c) {
			return rest.InvalidScopeResponse("project_comments:write", c)
		}

		// Read a rest.Message from the body request.
		message := rest.NewMessage{}
		if err := c.Bind(&message); err != nil {
			return err
		}

		// Create a nerdz.ProjectPostComment from the message
		// and current context.
		comment := nerdz.ProjectPostComment{}
		comment.Message = message.Message
		comment.Lang = message.Lang
		comment.To = c.Get("project").(*nerdz.Project).ID()
		comment.Hpid = c.Get("post").(*nerdz.ProjectPost).ID()

		// Send it
		me := c.Get("me").(*nerdz.User)
		if err := me.Add(&comment); err != nil {
			return err
		}
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields(comment.GetTO(me), c)
	}
}

// EditPostComment handles the request and edits the post comment
func EditPostComment() echo.HandlerFunc {

	// swagger:route PUT /projects/{id}/posts/{pid}/comments/{cid} project post comment EditProjectPost
	//
	// Update the speficied post on the specified project board
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_comments:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_comments:write", c) {
			return rest.InvalidScopeResponse("project_comments:write", c)
		}

		// Read a rest.Message from the body request.
		message := rest.NewMessage{}
		if err := c.Bind(&message); err != nil {
			return err
		}
		comment := c.Get("comment").(*nerdz.ProjectPostComment)

		// Update filds
		comment.Message = message.Message
		if comment.Lang != "" {
			comment.Lang = message.Lang
		}

		// Edit
		me := c.Get("me").(*nerdz.User)
		if err := me.Edit(comment); err != nil {
			return err
		}

		// Extract the TO from the comment and return selected fields.
		return rest.SelectFields(comment.GetTO(me), c)
	}
}

// DeletePostComment handles the request and deletes the specified
// comment on the speficied post
func DeletePostComment() echo.HandlerFunc {

	// swagger:route DELETE /projects/{id}/posts/{pid}/comments/{cid} project post DeleteProjectPostComment
	//
	// Delete the specified comment on the speficied project post
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_comments:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_comments:write", c) {
			return rest.InvalidScopeResponse("project_comments:write", c)
		}

		comment := c.Get("comment").(*nerdz.ProjectPostComment)
		me := c.Get("me").(*nerdz.User)
		if err := me.Delete(comment); err != nil {
			return err
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

// Info handles the request and returns all the basic informations for the specified project
func Info() echo.HandlerFunc {

	// swagger:route GET /projects/{id} project info getProjectInfo
	//
	// Shows the basic informations for the specified project
	//
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: projects:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("projects:read", c) {
			return rest.InvalidScopeResponse("projects:read", c)
		}
		project := c.Get("project").(*nerdz.Project)
		return rest.SelectFields(project.Info().GetTO(), c)
	}
}

// Members handles the request and returns the project members
func Members() echo.HandlerFunc {

	// swagger:route GET /projects/{id}/members project info members getProjectMembers
	//
	// Shows the members informations for the specified project
	//
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: projects:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("projects:read", c) {
			return rest.InvalidScopeResponse("projects:read", c)
		}
		members := c.Get("project").(*nerdz.Project).Members()
		return rest.SelectFields(user.GetUsersInfo(members), c)
	}
}

// Followers handles the request and returns the project followers
func Followers() echo.HandlerFunc {

	// swagger:route GET /projects/{id}/followers project info followers getProjectFollowers
	//
	// Shows the followers informations for the specified project
	//
	// You can personalize the request via query string parameters
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: projects:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("projects:read", c) {
			return rest.InvalidScopeResponse("projects:read", c)
		}
		followers := c.Get("project").(*nerdz.Project).Followers()
		return rest.SelectFields(user.GetUsersInfo(followers), c)
	}
}

// PostVotes handles the request and returns the post votes
func PostVotes() echo.HandlerFunc {

	// swagger:route GET /projects/{id}/posts/{pid}/votes project post votes GetProjectPostVotes
	//
	// List the votes of the post
	//
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_messages:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:read", c) {
			return rest.InvalidScopeResponse("project_messages:read", c)
		}
		votes := c.Get("post").(*nerdz.ProjectPost).Votes()
		if votes == nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Unable to fetch votes for the specified post",
				Message:      "ProjectPost.Votes() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var votesTO []*nerdz.ProjectPostVoteTO
		me := c.Get("me").(*nerdz.User)
		for _, v := range *votes {
			// votes contains Vote elements
			// we need to convert back to a ProjectPostVote in order to get a correct ProjectPostVoteTO
			if userPostVote := v.(*nerdz.ProjectPostVote); userPostVote != nil {
				votesTO = append(votesTO, userPostVote.GetTO(me))
			}
		}
		return rest.SelectFields(votesTO, c)
	}
}

// NewPostVote handles the request and creates a new vote for the post
func NewPostVote() echo.HandlerFunc {

	// swagger:route POST /projects/{id}/posts/{pid}/votes project post vote NewProjectPostVote
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
	//		oauth: project_messages:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:write", c) {
			return rest.InvalidScopeResponse("project_messages:write", c)
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
		post := c.Get("post").(*nerdz.ProjectPost)
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
		return rest.SelectFields(vote.(*nerdz.ProjectPostVote).GetTO(me), c)
	}
}

// PostCommentVotes handles the request and returns the comment votes
func PostCommentVotes() echo.HandlerFunc {

	// swagger:route GET /projects/{id}/posts/{pid}/comments/{cid}/votes project post comments votes GetProjectPostCommentsVotes
	//
	// List the votes on the comment
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_comments:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_comments:read", c) {
			return rest.InvalidScopeResponse("project_comments:read", c)
		}
		votes := c.Get("comment").(*nerdz.ProjectPostComment).Votes()
		if votes == nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Unable to fetch votes for the specified post",
				Message:      "ProjectPostComment.Votes() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var votesTO []*nerdz.ProjectPostCommentVoteTO
		me := c.Get("me").(*nerdz.User)
		for _, v := range *votes {
			// votes contains Vote elements
			// we need to convert back to a ProjectPostCommentVote in order to get a correct ProjectPostCommentVoteTO
			if userPostCommentVote := v.(*nerdz.ProjectPostCommentVote); userPostCommentVote != nil {
				votesTO = append(votesTO, userPostCommentVote.GetTO(me))
			}
		}
		return rest.SelectFields(votesTO, c)
	}
}

// NewPostCommentVote handles the request and creates a new vote on the user comment post
func NewPostCommentVote() echo.HandlerFunc {

	// swagger:route POST /projects/{id}/posts/{pid}/comments/{cid}votes project post comment vote NewProjectPostCommentVote
	//
	// Adds a new vote on the current project post comment
	//
	// Consumes:
	// - application/json
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_comments:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_comments:write", c) {
			return rest.InvalidScopeResponse("project_comments:write", c)
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
		comment := c.Get("comment").(*nerdz.ProjectPostComment)
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
		return rest.SelectFields(vote.(*nerdz.ProjectPostCommentVote).GetTO(me), c)
	}
}
