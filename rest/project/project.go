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
	"strconv"

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
	//		oauth: profile_messages:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile_messages:read", c) {
			return rest.InvalidScopeResponse("profile_messages:read", c)
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
	//		oauth: profile_messages:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile_messages:read", c) {
			return rest.InvalidScopeResponse("profile_messages:read", c)
		}
		me := c.Get("me").(*nerdz.User)
		postTO := c.Get("post").(*nerdz.ProjectPost).GetTO(me)
		return rest.SelectFields(postTO, c)
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
	//		oauth: profile_comments:read
	//
	//	Responses:
	//		default: apiResponse

	return func(c echo.Context) error {
		if !rest.IsGranted("profile_comments:read", c) {
			return rest.InvalidScopeResponse("profile_comments:read", c)
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

	// swagger:route GET /projects/{id}/posts/{pid}/comments/{cid} project post comment getProjectPostComment
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

		var comment *nerdz.ProjectPostComment
		if comment, e = nerdz.NewProjectPostComment(cid); e != nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Invalid comment identifier specified",
				Message:      e.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		post := c.Get("post").(*nerdz.ProjectPost)
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
