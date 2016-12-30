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
	"fmt"
	"github.com/galeone/igor"
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/rest"
	"github.com/nerdzeu/nerdz-api/utils"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// authorization is the authorization middleware for users.
// It checks the access_token in the Authorization header or the access_token query parameter
// On success sets "me" = *User (current logged user) and "accessData" = current access data
// into the context. Sets even the scopes variable, the sorted slice of scopes in accessData
func authorization() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			var accessToken string
			auth := c.Request().Header.Get("Authorization")
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

			// fetch current logged user and store it into the context
			me, err := nerdz.NewUser(accessData.UserData.(uint64))
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			c.Set("me", me)

			// store the Access Data into the context
			c.Set("accessData", accessData)
			scopes := strings.Split(accessData.Scope, " ")
			sort.Strings(scopes)

			// store the sorted Scopes using the full format
			// eg: if accepted scope is profile:read,write
			// save 2 entries: profile:read and profile:write
			// each saved scope is always in the format <name>:<read|,write>
			var fullScopes []string
			for _, s := range scopes {
				//parts[0] = <scope>, parts[1] = <rw>
				parts := strings.Split(s, ":")
				rw := strings.Split(parts[1], ",")
				for _, perm := range rw {
					fullScopes = append(fullScopes, parts[0]+":"+perm)
				}
			}
			c.Set("scopes", fullScopes)

			// let next handler handle the context
			return next(c)
		})
	}
}

// setPostlist is the middleware that sets "postlistOptions" = *nerdz.PostlistOptions into the current Context
// handle GET parameters:
// following: if setted, requires posts from following users
// followers: if setted, requires posts from followers users
// lang: if setted to a supported language (nerdz.Configuration.Languages), requires
//       posts in that language
// older: if setted to an existing hpid, requires posts older than the "older" value
// olderType: if setted can be only "user" or "project". Represents a reference to the older hpid type
//		used when fetching from a view, where hpid can be from posts or groups_posts
// newer: if setted to an existing hpid, requires posts newer than the "newer" value7
// newerType: if setted can be only "user" or "project". Represents a reference to the newer hpid type
//		used when fetching from a view, where hpid can be from posts or groups_posts
// n: if setted, define the number of posts to retrieve. Follows the nerdz.atMostPost rules
func setPostlist() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			var following, followers bool
			if c.QueryParam("following") != "" {
				following = true
			}

			if c.QueryParam("followers") != "" {
				followers = true
			}

			for _, t := range []string{"olderType", "newerType"} {
				tValue := c.QueryParam(t)
				if tValue != "" {
					if tValue != "user" && tValue != "project" {
						message := fmt.Sprintf(`Unsupported %s %s. Only "user" or "project" are allowed`, t, tValue)
						return c.JSON(http.StatusBadRequest, &rest.Response{
							HumanMessage: message,
							Message:      message,
							Status:       http.StatusBadRequest,
							Success:      false,
						})
					}
				}
			}

			var olderModel, newerModel igor.DBModel
			if c.QueryParam("olderType") == "user" {
				olderModel = nerdz.UserPost{}
			} else {
				olderModel = nerdz.ProjectPost{}
			}

			if c.QueryParam("newerType") == "user" {
				newerModel = nerdz.UserPost{}
			} else {
				newerModel = nerdz.ProjectPost{}
			}

			var language string
			lang := c.QueryParam("lang")
			if lang == "" {
				language = ""
			} else {
				if !utils.InSlice(lang, nerdz.Configuration.Languages) {
					message := "Not supported language: " + lang
					return c.JSON(http.StatusBadRequest, &rest.Response{
						HumanMessage: message,
						Message:      message,
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
				Following:  following,
				Followers:  followers,
				Language:   language,
				N:          nerdz.AtMostPosts(n),
				Older:      older,
				OlderModel: olderModel,
				Newer:      newer,
				NewerModel: newerModel,
			})

			return next(c)
		})
	}
}

// setCommentList is the middleware that sets the "commentlistOptions" = *nerdz.CommentlistOptions into the current Context
// handle GET parameters:
// older: if setted to an existing hpid, requires posts older than the "older" value
// newer: if setted to an existing hpid, requires posts newer than the "newer" value
// n: if setted, define the number of comments to retrieve. Follows the nerdz.atMostComments rules
func setCommentList() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
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

			return next(c)
		})
	}
}

// setPmsOptions is the middleware that sets the "pmsOptions" = *nerdz.PmsOptions into the current context
// handle GET parameters:
func setPmsOptions() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			old := c.QueryParam("older")
			new := c.QueryParam("newer")

			older, _ := strconv.ParseUint(old, 10, 64)
			newer, _ := strconv.ParseUint(new, 10, 64)

			n, _ := strconv.ParseUint(c.QueryParam("n"), 10, 8)

			c.Set("pmsOptions", &nerdz.PmsOptions{
				N:     nerdz.AtMostComments(n),
				Older: older,
				Newer: newer,
			})
			return next(c)
		})
	}
}
