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

package router

import (
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/rest"
	"github.com/nerdzeu/nerdz-api/utils"
	"net/http"
	"strconv"
	"strings"
)

// authorization is the authorization middleware for users.
// It checks the access_token in the Authorization header or the access_token query parameter
// On success sets "me" = *User (current logged user) and "accessData" = current access data
// into the context
func authorization() echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) error {
			var accessToken string
			auth := c.Request().Header().Get("Authorization")
			if auth == "" {
				// Check if there's the parameter access_token in the URL
				// this makes the bearer authentication with websockets compatible with OAuth2
				accessToken = c.QueryParam("access_token")
				if accessToken == "" {
					return c.String(http.StatusUnauthorized, "access_token required")
				}
			} else {
				if !strings.HasPrefix(auth, "Bearer ") {
					return echo.ErrUnauthorized
				}
				ss := strings.Split(auth, " ")
				if len(ss) != 2 {
					return echo.ErrUnauthorized
				}
				accessToken = ss[1]
			}

			accessData, err := (&nerdz.OAuth2Storage{}).LoadAccess(accessToken)
			if err != nil {
				return c.String(http.StatusUnauthorized, err.Error())
			}

			// store the Access Data into the context
			c.Set("accessData", accessData)

			// fetch current logged user and store it into the context
			me, err := nerdz.NewUser(accessData.UserData.(uint64))
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			c.Set("me", me)

			// let next handler handle the context
			return next.Handle(c)
		})
	}
}

// users is the middleware that checks if the current logged user can see the required profile
// and if the required profile exist.
// On success sets the "other" = *User variable in the context
func users() echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
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
				return c.JSON(http.StatusUnauthorized, &rest.Response{
					HumanMessage: "You can't see the required profile",
					Message:      "You can't see the required profile",
					Status:       http.StatusUnauthorized,
					Success:      false,
				})
			}

			// store the other User into the context
			c.Set("other", other)
			// pass context to the next handler
			return next.Handle(c)
		})
	}
}

// userPost is the middleware that checks if the required post, on the user board, exists.
// If it exists, set the "post" = *UserPost in the current context
func userPost() echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
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

			otherID := c.Get("other").(*nerdz.User).Counter
			var post nerdz.UserPost

			if e = nerdz.Db().Model(nerdz.UserPost{}).Where(&nerdz.UserPost{nerdz.Post{To: otherID, Pid: pid}}).Scan(&post); e != nil {
				return c.JSON(http.StatusBadRequest, &rest.Response{
					HumanMessage: "Required post does not exists",
					Message:      e.Error(),
					Status:       http.StatusBadRequest,
					Success:      false,
				})
			}

			c.Set("post", &post)
			return next.Handle(c)
		})
	}
}

// postlist is the middleware that sets "postlistOptions" = *nerdz.PostlistOptions into the current Context
// handle GET parameters:
// following: if setted, requires posts from following users
// followers: if setted, requires posts from followers users
// lang: if setted to a supported language (nerdz.Configuration.Languages), requires
//       posts in that language
// older: if setted to an existing hpid, requires posts older than the "older" value
// newer: if setted to an existing hpid, requires posts newer than the "newer" value
// n: if setted, define the number of posts to retriete. Follows the nerdz.atMostPost rules
func postlist() echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) error {
			var following, followers bool
			if c.QueryParam("following") != "" {
				following = true
			}

			if c.QueryParam("followers") != "" {
				followers = true
			}

			var language string
			lang := c.QueryParam("lang")
			if lang == "" {
				language = ""
			} else {
				if !utils.InSlice(lang, nerdz.Configuration.Languages) {
					return c.JSON(http.StatusBadRequest, &rest.Response{
						HumanMessage: "Not supported language: " + lang,
						Message:      "Not supported language: " + lang,
						Status:       http.StatusBadRequest,
						Success:      false,
					})
				}
				language = lang
			}

			old := c.QueryParam("older")
			new := c.QueryParam("newer")

			older, _ := strconv.ParseUint(old, 10, 64)
			newer, _ := strconv.ParseUint(new, 10, 64)

			n, _ := strconv.ParseUint(c.QueryParam("n"), 10, 8)

			c.Set("postlistOptions", &nerdz.PostlistOptions{
				Following: following,
				Followers: followers,
				Language:  language,
				N:         nerdz.AtMostPosts(n),
				Older:     older,
				Newer:     newer,
			})

			return next.Handle(c)
		})
	}
}

// commentlist is the middleware that sets the "commentlistOptions" = *nerdz.CommentlistOptions into the current ContextX
// handle GET parameters:
// older: if setted to an existing hpid, requires posts older than the "older" value
// newer: if setted to an existing hpid, requires posts newer than the "newer" value
// n: if setted, define the number of posts to retriete. Follows the nerdz.atMostComments rules
func commentlist() echo.MiddlewareFunc {
	return func(next echo.Handler) echo.Handler {
		return echo.HandlerFunc(func(c echo.Context) error {
			old := c.QueryParam("older")
			new := c.QueryParam("newer")

			older, _ := strconv.ParseUint(old, 10, 64)
			newer, _ := strconv.ParseUint(new, 10, 64)

			n, _ := strconv.ParseUint(c.QueryParam("n"), 10, 8)

			c.Set("commentlistOptions", &nerdz.CommentlistOptions{
				N:     nerdz.AtMostComments(n),
				Older: older,
				Newer: newer,
			})

			return next.Handle(c)
		})
	}
}
