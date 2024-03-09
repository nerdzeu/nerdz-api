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

package project

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/rest"
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
	//		default: ProjectsIdPosts

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:read", c) {
			return rest.InvalidScopeResponse("project_messages:read", c)
		}

		project := c.Get("project").(*nerdz.Project)
		options := c.Get("postlistOptions").(*nerdz.PostlistOptions)
		options.Model = nerdz.ProjectPost{}
		posts := project.Postlist(*options)

		if posts == nil {
			errstr := "unable to fetch post list for the specified project"
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "project.Postlist error",
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
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
	//		default: ProjectsIdPostsPid

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
	//		default: ProjectsIdPostsPid

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

		errstr := "success"
		if err := c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: errstr,
			Message:      errstr,
			Status:       http.StatusOK,
			Success:      true,
		}); err != nil {
			log.Errorf("Error while writing response: %s", err.Error())
		}
		return nil
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
	//		default: ProjectsIdPostsPid

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
	//		default: ProjectsIdPostsPidComments

	return func(c echo.Context) error {
		if !rest.IsGranted("project_comments:read", c) {
			return rest.InvalidScopeResponse("project_comments:read", c)
		}
		comments := c.Get("post").(*nerdz.ProjectPost).Comments(*(c.Get("commentlistOptions").(*nerdz.CommentlistOptions)))
		if comments == nil {
			errstr := "unable to fetch comment list for the specified post"
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "ProjectPost.Comments(options) error",
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
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
	//		default: ProjectsIdPostsPidCommentsCid

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
	//		default: ProjectsIdPostsPidCommentsCid

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

	// swagger:route PUT /projects/{id}/posts/{pid}/comments/{cid} project post comment EditProjectPostComment
	//
	// Update the speficied comment on the specified project post
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
	//		default: ProjectsIdPostsPidCommentsCid

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

		errstr := "success"
		if err := c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: errstr,
			Message:      errstr,
			Status:       http.StatusOK,
			Success:      true,
		}); err != nil {
			log.Errorf("Error while writing response: %s", err.Error())
		}
		return nil
	}
}

// Info handles the request and returns all the basic information for the specified project
func Info() echo.HandlerFunc {

	// swagger:route GET /projects/{id} project info getProjectInfo
	//
	// Shows the basic information for the specified project
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
	//		default: ProjectsId

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
	// Shows the members information for the specified project
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
	//		default: ProjectsIdMembers

	return func(c echo.Context) error {
		if !rest.IsGranted("projects:read", c) {
			return rest.InvalidScopeResponse("projects:read", c)
		}
		members := c.Get("project").(*nerdz.Project).Members()
		return rest.SelectFields(rest.GetUsersInfo(members), c)
	}
}

// Followers handles the request and returns the project followers
func Followers() echo.HandlerFunc {

	// swagger:route GET /projects/{id}/followers project info followers getProjectFollowers
	//
	// Shows the followers information for the specified project
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
	//		default: ProjectsIdFollowers

	return func(c echo.Context) error {
		if !rest.IsGranted("projects:read", c) {
			return rest.InvalidScopeResponse("projects:read", c)
		}
		followers := c.Get("project").(*nerdz.Project).Followers()
		return rest.SelectFields(rest.GetUsersInfo(followers), c)
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
	//		default: ProjectsIdPostsPidVotes

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:read", c) {
			return rest.InvalidScopeResponse("project_messages:read", c)
		}
		votes := c.Get("post").(*nerdz.ProjectPost).Votes()
		if votes == nil {
			errstr := "unable to fetch votes for the specified post"
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "ProjectPost.Votes() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
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
	//		default: ProjectsIdPostsPidVotes

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:write", c) {
			return rest.InvalidScopeResponse("project_messages:write", c)
		}

		// Read a rest.NewVote from the body request.
		body := rest.NewVote{}
		if err := c.Bind(&body); err != nil {
			errstr := err.Error()
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.ProjectPost)
		vote, err := me.Vote(post, body.Vote)
		if err != nil {
			errstr := err.Error()
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
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
	//		default: ProjectsIdPostsPidCommentsCidVotes

	return func(c echo.Context) error {
		if !rest.IsGranted("project_comments:read", c) {
			return rest.InvalidScopeResponse("project_comments:read", c)
		}
		votes := c.Get("comment").(*nerdz.ProjectPostComment).Votes()
		if votes == nil {
			errstr := "unable to fetch votes for the specified post"
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "ProjectPostComment.Votes() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
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

	// swagger:route POST /projects/{id}/posts/{pid}/comments/{cid}/votes project post comment vote NewProjectPostCommentVote
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
	//		default: ProjectsIdPostsPidCommentsCidVotes

	return func(c echo.Context) error {
		if !rest.IsGranted("project_comments:write", c) {
			return rest.InvalidScopeResponse("project_comments:write", c)
		}

		// Read a rest.NewVote from the body request.
		body := rest.NewVote{}
		if err := c.Bind(&body); err != nil {
			errstr := err.Error()
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}

		// Send it
		me := c.Get("me").(*nerdz.User)
		comment := c.Get("comment").(*nerdz.ProjectPostComment)
		vote, err := me.Vote(comment, body.Vote)
		if err != nil {
			errstr := err.Error()
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields(vote.(*nerdz.ProjectPostCommentVote).GetTO(me), c)
	}
}

// PostBookmarks handles the request and returns the post bookmarks
func PostBookmarks() echo.HandlerFunc {

	// swagger:route GET /projects/{id}/posts/{pid}/bookmarks project post bookmarks GetProjectPostBookmarks
	//
	// List the bookmarks of the post
	//
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_messages:read
	//
	//	Responses:
	//		default: ProjectsIdPostsPidBookmarks

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:read", c) {
			return rest.InvalidScopeResponse("project_messages:read", c)
		}
		bookmarks := c.Get("post").(*nerdz.ProjectPost).Bookmarks()
		if bookmarks == nil {
			errstr := "unable to fetch bookmarks for the specified post"
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "ProjectPost.Bookmarks() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}

		var bookmarksTO []*nerdz.ProjectPostBookmarkTO
		me := c.Get("me").(*nerdz.User)
		for _, v := range *bookmarks {
			// bookmarks contains Boook elements
			// we need to convert back to a ProjectPostBookmark in order to get a correct ProjectPostBookmarkTO
			if userPostBookmark := v.(*nerdz.ProjectPostBookmark); userPostBookmark != nil {
				bookmarksTO = append(bookmarksTO, userPostBookmark.GetTO(me))
			}
		}
		return rest.SelectFields(bookmarksTO, c)
	}
}

// NewPostBookmark handles the request and creates a new bookmark for the post
func NewPostBookmark() echo.HandlerFunc {

	// swagger:route POST /projects/{id}/posts/{pid}/bookmarks project post vote NewProjectPostBookmark
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
	//		oauth: project_messages:write
	//
	//	Responses:
	//		default: ProjectsIdPostsPidBookmarks

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:write", c) {
			return rest.InvalidScopeResponse("project_messages:write", c)
		}

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.ProjectPost)
		bookmark, err := me.Bookmark(post)
		if err != nil {
			errstr := err.Error()
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields(bookmark.(*nerdz.ProjectPostBookmark).GetTO(me), c)
	}
}

// DeletePostBookmark handles the request and deletes the bookmark to the post
func DeletePostBookmark() echo.HandlerFunc {

	// swagger:route DELETE /projects/{id}/posts/{pid}/bookmarks project post vote DeleteProjectPostBookmark
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
	//		oauth: project_messages:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:write", c) {
			return rest.InvalidScopeResponse("project_messages:write", c)
		}

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.ProjectPost)
		err := me.Unbookmark(post)
		if err != nil {
			errstr := err.Error()
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}

		errstr := "success"
		return c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: errstr,
			Message:      errstr,
			Status:       http.StatusOK,
			Success:      true,
		})
	}
}

// PostLurks handles the request and returns the post lurks
func PostLurks() echo.HandlerFunc {

	// swagger:route GET /projects/{id}/posts/{pid}/lurks project post bookmarks GetProjectPostLurks
	//
	// List the lurks of the post
	//
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_messages:read
	//
	//	Responses:
	//		default: ProjectsIdPostsPidLurks

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:read", c) {
			return rest.InvalidScopeResponse("project_messages:read", c)
		}
		lurks := c.Get("post").(*nerdz.ProjectPost).Lurks()
		if lurks == nil {
			errstr := "unable to fetch lurks for the specified post"
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "ProjectPost.Lurks() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}

		var lurksTO []*nerdz.ProjectPostLurkTO
		me := c.Get("me").(*nerdz.User)
		for _, v := range *lurks {
			// lurks contains Lurk elements
			// we need to convert back to a ProjectPostLurk in order to get a correct ProjectPostLurkTO
			if userPostLurk := v.(*nerdz.ProjectPostLurk); userPostLurk != nil {
				lurksTO = append(lurksTO, userPostLurk.GetTO(me))
			}
		}
		return rest.SelectFields(lurksTO, c)
	}
}

// NewPostLurk handles the request and creates a new lurk for the post
func NewPostLurk() echo.HandlerFunc {

	// swagger:route POST /projects/{id}/posts/{pid}/lurks project post vote NewProjectPostLurk
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
	//		oauth: project_messages:write
	//
	//	Responses:
	//		default: ProjectsIdPostsPidLurks

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:write", c) {
			return rest.InvalidScopeResponse("project_messages:write", c)
		}

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.ProjectPost)
		lurk, err := me.Lurk(post)
		if err != nil {
			errstr := err.Error()
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields(lurk.(*nerdz.ProjectPostLurk).GetTO(me), c)
	}
}

// DeletePostLurk handles the request and deletes the lurk to the post
func DeletePostLurk() echo.HandlerFunc {

	// swagger:route DELETE /projects/{id}/posts/{pid}/lurks project post vote DeleteProjectPostLurk
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
	//		oauth: project_messages:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:write", c) {
			return rest.InvalidScopeResponse("project_messages:write", c)
		}

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.ProjectPost)
		err := me.Unlurk(post)
		if err != nil {
			errstr := err.Error()
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}

		errstr := "success"
		return c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: errstr,
			Message:      errstr,
			Status:       http.StatusOK,
			Success:      true,
		})
	}
}

// PostLock handles the request and and a lock to the post
func PostLock() echo.HandlerFunc {

	// swagger:route GET /projects/{id}/posts/{pid}/locks project post locks GetProjectPostLock
	//
	// List the locks of the post
	//
	//	Produces:
	//	- application/json
	//
	//	Security:
	//		oauth: project_messages:read
	//
	//	Responses:
	//		default: ProjectsIdPostsPidLocks

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:read", c) {
			return rest.InvalidScopeResponse("project_messages:read", c)
		}
		locks := c.Get("post").(*nerdz.ProjectPost).Locks()
		if locks == nil {
			errstr := "unable to fetch locks for the specified post"
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: errstr,
				Message:      "ProjectPost.Lock() error",
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}

		var locksTO []*nerdz.ProjectPostLockTO
		me := c.Get("me").(*nerdz.User)
		for _, v := range *locks {
			// locks contains Lock elements
			// we need to convert back to a ProjectPostLock in order to get a correct ProjectPostLockTO
			if userPostLock := v.(*nerdz.ProjectPostLock); userPostLock != nil {
				locksTO = append(locksTO, userPostLock.GetTO(me))
			}
		}
		return rest.SelectFields(locksTO, c)
	}
}

// NewPostLock handles the request and creates a new lock for the post
func NewPostLock() echo.HandlerFunc {

	// swagger:route POST /projects/{id}/posts/{pid}/locks project post vote NewProjectPostLock
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
	//		oauth: project_messages:write
	//
	//	Responses:
	//		default: ProjectsIdPostsPidLocks

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:write", c) {
			return rest.InvalidScopeResponse("project_messages:write", c)
		}

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.ProjectPost)
		lock, err := me.Lock(post)
		if err != nil {
			errstr := err.Error()
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}
		// Extract the TO from the new post and return
		// selected fields.
		return rest.SelectFields((*lock)[0].(*nerdz.ProjectPostLock).GetTO(me), c)
	}
}

// DeletePostLock handles the request and deletes the lock to the post
func DeletePostLock() echo.HandlerFunc {

	// swagger:route DELETE /projects/{id}/posts/{pid}/locks project post vote DeleteProjectPostLock
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
	//		oauth: project_messages:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:write", c) {
			return rest.InvalidScopeResponse("project_messages:write", c)
		}

		// Send it
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.ProjectPost)
		err := me.Unlock(post)
		if err != nil {
			errstr := err.Error()
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}

		errstr := "success"
		return c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: errstr,
			Message:      errstr,
			Status:       http.StatusOK,
			Success:      true,
		})
	}
}

// NewPostUserLock handles the request and creates a new lock for the notification
// caused by the target user
func NewPostUserLock() echo.HandlerFunc {

	// swagger:route POST /projects/{id}/posts/{pid}/locks/{target} project post vote NewUserNewPostProjectLock
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
	//		oauth: project_messages:write
	//
	//	Responses:
	//		default: ProjectsIdPostsPidLocks

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:write", c) {
			return rest.InvalidScopeResponse("project_messages:write", c)
		}

		var target *nerdz.User
		var err error
		if target, err = rest.User("target", c); err != nil {
			return err
		}
		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.ProjectPost)
		var lock *[]nerdz.Lock
		lock, err = me.Lock(post, target)
		if err != nil {
			errstr := err.Error()
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}
		// Extract the TO from the new lock and return selected fields.
		return rest.SelectFields((*lock)[0].(*nerdz.ProjectPostUserLock).GetTO(me), c)
	}
}

// DeletePostUserLock handles the request and deletes the lock for the notification of the target user
// on the specified post
func DeletePostUserLock() echo.HandlerFunc {

	// swagger:route DELETE /projects/{id}/posts/{pid}/locks/{target} project post vote DeleteProjectPostUserLock
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
	//		oauth: project_messages:write
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("project_messages:write", c) {
			return rest.InvalidScopeResponse("project_messages:write", c)
		}
		var target *nerdz.User
		var err error
		if target, err = rest.User("target", c); err != nil {
			return err
		}

		me := c.Get("me").(*nerdz.User)
		post := c.Get("post").(*nerdz.ProjectPost)
		err = me.Unlock(post, target)
		if err != nil {
			errstr := err.Error()
			if err := c.JSON(http.StatusBadRequest, &rest.Response{
				Data:         nil,
				HumanMessage: errstr,
				Message:      errstr,
				Status:       http.StatusBadRequest,
				Success:      false,
			}); err != nil {
				log.Errorf("Error while writing response: %s", err.Error())
			}
			return errors.New(errstr)
		}

		errstr := "success"
		return c.JSON(http.StatusOK, &rest.Response{
			Data:         nil,
			HumanMessage: errstr,
			Message:      errstr,
			Status:       http.StatusOK,
			Success:      true,
		})
	}
}
