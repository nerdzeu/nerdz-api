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
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/rest"
	"net/http"
	"strconv"
)

// SetOther is the middleware that checks if the current logged user can see the required profile
// and if the required profile exists. On success sets the "other" = *User variable in the context
func SetOther() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			var id uint64
			var e error
			if id, e = strconv.ParseUint(c.Param("id"), 10, 64); e != nil {
				return c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "Invalid user identifier specified",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
			}

			var other *nerdz.User
			if other, e = nerdz.NewUser(id); e != nil {
				return c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "User does not exists",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
			}

			me := c.Get("me").(*nerdz.User)
			if !me.CanSee(other) {
				message := "You can't see the required profile"
				return c.JSON(http.StatusUnauthorized, &rest.Response{
					HumanMessage: message,
					Message:      message,
					Status:       http.StatusUnauthorized,
					Success:      false,
				})
			}

			// store the other User into the context
			c.Set("other", other)
			// pass context to the next handler
			return next(c)
		})
	}
}

// SetPost is the middleware that checks if the required post, on the user board, exists.
// If it exists, set the "post" = *UserPost in the current context
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

			otherID := c.Get("other").(*nerdz.User).ID()
			var post *nerdz.UserPost

			if post, e = nerdz.NewUserPostWhere(&nerdz.UserPost{nerdz.Post{To: otherID, Pid: pid}}); e != nil {
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

// SetComment is the middleware that check if the required comment, on the user board, exists.
// If it exists, set the "comment" =  *UserPostComment in the current context
func SetComment() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

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
			c.Set("comment", comment)
			return next(c)
		})
	}
}

// SetPm is the middleware that check if the required pm exists.
// If it exists, set the "pm" = *Pm in the current context
func SetPm() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

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

			if (pm.From == otherID && pm.To == other.ID()) || (pm.From == other.ID() && pm.To == otherID) {
				c.Set("pm", pm)
				return next(c)
			}

			message := "You're not autorized to see the requested PM"
			return c.JSON(http.StatusUnauthorized, &rest.Response{
				HumanMessage: message,
				Message:      message,
				Status:       http.StatusUnauthorized,
				Success:      false,
			})
		})
	}
}
