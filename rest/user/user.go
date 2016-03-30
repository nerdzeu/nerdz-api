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
	return func(c echo.Context) error {
		other := c.Get("other").(*nerdz.User)
		options := c.Get("postlistOptions").(*nerdz.PostlistOptions)
		options.User = true
		posts := other.Postlist(*options)

		if posts == nil {
			return c.JSON(http.StatusBadRequest, &rest.Response{
				HumanMessage: "Unable to fetch post list for the specified user",
				Message:      "other.Postlist error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var postsAPI []*nerdz.UserPostTO
		for _, p := range *posts {
			// posts contains ExistingPost elements
			// we need to convert back to a UserPost in order to get a correct UserPostTO
			if userPost := p.(*nerdz.UserPost); userPost != nil {
				postsAPI = append(postsAPI, userPost.GetTO().(*nerdz.UserPostTO))
			}
		}

		return rest.SelectFields(postsAPI, c)
	}
}

// Post handles the request and returns the single post required
func Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		postTO := c.Get("post").(*nerdz.UserPost).GetTO().(*nerdz.UserPostTO)
		return rest.SelectFields(postTO, c)
	}
}

// PostComments handles the request and returns the specified list of comments
func PostComments() echo.HandlerFunc {
	return func(c echo.Context) error {
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
		for _, p := range *comments {
			// comments contains ExistingPost elements
			// we need to convert back to a UserPostComment in order to get a correct UserPostCommentTO
			if userPostComment := p.(*nerdz.UserPostComment); userPostComment != nil {
				commentsAPI = append(commentsAPI, userPostComment.GetTO().(*nerdz.UserPostCommentTO))
			}
		}
		return rest.SelectFields(commentsAPI, c)
	}
}

// PostComment handles the request and returns the single comment required
func PostComment() echo.HandlerFunc {
	return func(c echo.Context) error {
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

		return rest.SelectFields(comment.GetTO().(*nerdz.UserPostCommentTO), c)
	}
}

// Info handles the request and returns all the basic information for the specified user
func Info() echo.HandlerFunc {
	return func(c echo.Context) error {
		other := c.Get("other").(*nerdz.User)
		return rest.SelectFields(getInfo(other), c)
	}
}

// Friends handles the request and returns the user friends
func Friends() echo.HandlerFunc {
	return func(c echo.Context) error {
		friends := c.Get("other").(*nerdz.User).Friends()

		var usersInfo []*Informations
		for _, u := range friends {
			usersInfo = append(usersInfo, getInfo(u))
		}
		return rest.SelectFields(usersInfo, c)
	}
}

// Followers handles the request and returns the user followers
func Followers() echo.HandlerFunc {
	return func(c echo.Context) error {
		friends := c.Get("other").(*nerdz.User).Followers()

		var usersInfo []*Informations
		for _, u := range friends {
			usersInfo = append(usersInfo, getInfo(u))
		}
		return rest.SelectFields(usersInfo, c)
	}
}

// Following handles the request and returns the user following
func Following() echo.HandlerFunc {
	return func(c echo.Context) error {
		friends := c.Get("other").(*nerdz.User).Following()

		var usersInfo []*Informations
		for _, u := range friends {
			usersInfo = append(usersInfo, getInfo(u))
		}
		return rest.SelectFields(usersInfo, c)
	}
}
