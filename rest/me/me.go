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

// Posts handles the request and returns the required posts written by the specified user
func Posts() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Posts()(c)
	}
}

// Post handles the request and returns the single post required
func Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Post()(c)
	}
}

// PostComments handles the request and returns the specified list of comments
func PostComments() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.PostComments()(c)
	}
}

// PostComment handles the request and returns the single comment required
func PostComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.PostComment()(c)
	}
}

// Info handles the request and returns all the basic information for the specified user
func Info() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Info()(c)
	}
}

// Friends handles the request and returns the user friends
func Friends() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Friends()(c)
	}
}

// Followers handles the request and returns the user followers
func Followers() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Followers()(c)
	}
}

// Following handles the request and returns the user following
func Following() echo.HandlerFunc {
	return func(c echo.Context) error {
		return user.Following()(c)
	}
}
