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

package oauth2

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/rest"
	"github.com/openshift/osin"
)

var oauth *osin.Server

// Init initialize OAuth2 Authorization server. This is the first function to call before using OAuth2
func Init(server *osin.Server) {
	oauth = server
}

// Authorize is the action of GET /oauth2/authorize and POST /oauth2/authorize when authentication is required
func Authorize() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := oauth.NewResponse()
		defer resp.Close()

		if ar := oauth.HandleAuthorizeRequest(resp, c.Request()); ar != nil {
			if c.QueryParam("authorized") == "" || c.QueryParam("authorized_code") == "" {
				return c.Redirect(http.StatusFound, fmt.Sprintf("%s/oauth2/authorize.php?client_id=%s&response_type=%s&redirect_uri=%s&scope=%s&state=%s",
					nerdz.Configuration.NERDZURL().String(),
					url.QueryEscape(c.QueryParam("client_id")),
					url.QueryEscape(c.QueryParam("response_type")),
					url.QueryEscape(c.QueryParam("redirect_uri")),
					url.QueryEscape(c.QueryParam("scope")),
					url.QueryEscape(c.QueryParam("state"))))
			} else {
				var e error
				var userID uint64
				if userID, e = strconv.ParseUint(c.QueryParam("authorized"), 10, 64); e != nil {
					return c.JSON(http.StatusInternalServerError, &rest.Response{
						HumanMessage: "Invalid authorized (user id) value",
						Message:      e.Error(),
						Status:       http.StatusInternalServerError,
						Success:      false,
					})
				}

				var user *nerdz.User
				if user, e = nerdz.NewUser(userID); e != nil {
					return c.JSON(http.StatusInternalServerError, &rest.Response{
						HumanMessage: "Problem retrieving specified user",
						Message:      e.Error(),
						Status:       http.StatusInternalServerError,
						Success:      false,
					})
				}
				sha1_sum := fmt.Sprintf("%x", sha1.Sum([]byte(user.Username+user.Password+user.Email)))
				if sha1_sum != c.QueryParam("authorized_code") {
					message := "invalid authorization code"
					return c.JSON(http.StatusInternalServerError, &rest.Response{
						HumanMessage: message,
						Message:      message,
						Status:       http.StatusInternalServerError,
						Success:      false,
					})
				}

				ar.UserData = user.Counter
				ar.Authorized = true
				oauth.FinishAuthorizeRequest(resp, c.Request(), ar)
			}
		}

		if resp.IsError && resp.InternalError != nil {
			return c.JSON(http.StatusInternalServerError, &rest.Response{
				HumanMessage: "Internal Server error",
				Message:      resp.InternalError.Error(),
				Status:       http.StatusInternalServerError,
				Success:      false,
			})
		}

		return osin.OutputJSON(resp, c.Response(), c.Request())
	}
}

// Token is the action of Get /oauth2/token
func Token() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := oauth.NewResponse()
		defer resp.Close()

		if ar := oauth.HandleAccessRequest(resp, c.Request()); ar != nil {
			switch ar.Type {
			case osin.AUTHORIZATION_CODE:
				ar.Authorized = true
			case osin.REFRESH_TOKEN:
				ar.Authorized = true
			case osin.PASSWORD:
				if _, err := nerdz.Login(ar.Username, ar.Password); err == nil {
					ar.Authorized = true
				}
			case osin.CLIENT_CREDENTIALS:
				ar.Authorized = true
			}
			oauth.FinishAccessRequest(resp, c.Request(), ar)
		}

		if resp.IsError && resp.InternalError != nil {
			return c.JSON(http.StatusInternalServerError, &rest.Response{
				HumanMessage: "Internal Server error",
				Message:      resp.InternalError.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		return osin.OutputJSON(resp, c.Response(), c.Request())
	}
}

// Info is the action of GET /oauth2/info
func Info() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := oauth.NewResponse()
		defer resp.Close()

		if ir := oauth.HandleInfoRequest(resp, c.Request()); ir != nil {
			oauth.FinishInfoRequest(resp, c.Request(), ir)
		}

		return osin.OutputJSON(resp, c.Response(), c.Request())
	}
}
