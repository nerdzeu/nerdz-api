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
	"errors"
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/rest"
	"net/http"
	"strconv"
)

// SetProject is the middleware that checks if the current logged user can see the required project
// and if the required project exists. On success sets the "project" = *Project variable in the context
func SetProject() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			var project *nerdz.Project
			var err error
			if project, err = rest.Project("id", c); err != nil {
				return err
			}
			// store the project Project into the context
			c.Set("project", project)
			// pass context to the next handler
			return next(c)
		})
	}
}

// SetPost is the middleware that checks if the required post, on the project board, exists.
// If it exists, set the "post" = *ProjectPost in the current context
func SetPost() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			var e error
			var pid uint64

			if pid, e = strconv.ParseUint(c.Param("pid"), 10, 64); e != nil {
				c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "Invalid post identifier specified",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
				return e
			}

			projectID := c.Get("project").(*nerdz.Project).ID()
			var post *nerdz.ProjectPost

			if post, e = nerdz.NewProjectPostWhere(&nerdz.ProjectPost{nerdz.Post{To: projectID, Pid: pid}}); e != nil {
				c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "Required post does not exists",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
				return e
			}

			c.Set("post", post)
			return next(c)
		})
	}
}

// SetComment is the middleware that check if the required comment, on the project board, exists.
// If it exists, set the "comment" =  *ProjectPostComment in the current context
func SetComment() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			var cid uint64
			var e error
			if cid, e = strconv.ParseUint(c.Param("cid"), 10, 64); e != nil {
				c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "Invalid comment identifier specified",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
				return e
			}

			var comment *nerdz.ProjectPostComment
			if comment, e = nerdz.NewProjectPostComment(cid); e != nil {
				c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "Invalid comment identifier specified",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
				return e
			}

			post := c.Get("post").(*nerdz.ProjectPost)
			if comment.Hpid != post.Hpid {
				errstr := "Mismatch between comment ID and post ID. Comment not related to the post"
				c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: errstr,
					Message:      errstr,
					Status:       http.StatusBadRequest,
					Success:      false,
				})
				return errors.New(errstr)
			}
			c.Set("comment", comment)
			return next(c)
		})
	}
}
