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
			var id uint64
			var e error
			if id, e = strconv.ParseUint(c.Param("id"), 10, 64); e != nil {
				return c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "Invalid project identifier specified",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
			}

			var project *nerdz.Project
			if project, e = nerdz.NewProject(id); e != nil {
				return c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "Project does not exists",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
			}

			me := c.Get("me").(*nerdz.User)
			if !me.CanSee(project) {
				message := "You can't see the required project"
				return c.JSON(http.StatusUnauthorized, &rest.Response{
					HumanMessage: message,
					Message:      message,
					Status:       http.StatusUnauthorized,
					Success:      false,
				})
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
				return c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "Invalid post identifier specified",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
			}

			projectID := c.Get("project").(*nerdz.Project).ID()
			var post *nerdz.ProjectPost

			if post, e = nerdz.NewProjectPostWhere(&nerdz.ProjectPost{nerdz.Post{To: projectID, Pid: pid}}); e != nil {
				return c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "Required post does not exists",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
			}

			c.Set("post", post)
			return next(c)
		})
	}
}
