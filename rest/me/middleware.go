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

package me

import (
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/rest/user"
)

// SetOther is the middleware that sets the context variable "other" to "me"
// therfore we can use the package user methods in the me package
func SetOther() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			c.Set("other", c.Get("me"))
			return next(c)
		})
	}
}

// SetPost is the middleware that wraps user.SetPost middlware, thus:
// checks if the required post, on the user board, exists. If it exists,
// set the "post" = *UserPost in the current context
func SetPost() echo.MiddlewareFunc {
	return user.SetPost()
}
