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

package rest

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
)

// UserPosts handles the request and returns all the posts written by the specified user
func UserPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id uint64
		var e error
		if id, e = strconv.ParseUint(c.Param("id"), 10, 64); e != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				HumanMessage: "Invalid user identifier specified",
				Message:      e.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var user *nerdz.User
		if user, e = nerdz.NewUser(id); e != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				HumanMessage: "User does not exists",
				Message:      e.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var options *nerdz.PostlistOptions
		if options, e = newPostlistOptions(c); e != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				HumanMessage: e.Error(),
				Message:      "newPostlistOptions error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		options.User = true
		posts := user.Postlist(options)

		if posts == nil {
			return c.JSON(http.StatusBadRequest, &Response{
				HumanMessage: "Unable to fetch post list for the specified user",
				Message:      "user.Postlist error",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var postsAPI []*nerdz.UserPostTO

		for _, p := range *posts {
			// posts contains ExistingPost elements
			// we need to convert back to a UserPost in order to
			// get a correct UserPostTO
			if userPost := p.(*nerdz.UserPost); userPost != nil {
				postsAPI = append(postsAPI, userPost.GetTO().(*nerdz.UserPostTO))
			}
		}

		out, err := selectFields(postsAPI, c)
		if err == nil {
			return c.JSON(http.StatusOK, &Response{
				Data:         out,
				HumanMessage: "Correctly fetched post list for the specified user",
				Message:      "user.Postlist ok",
				Status:       http.StatusOK,
				Success:      true,
			})
		}

		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "Error selecting required fields",
			Message:      err.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}
}

//UserInfo handles the request and returns all the basic information for the specified user
func UserInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id uint64
		var e error

		if id, e = strconv.ParseUint(c.Param("id"), 10, 64); e != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				HumanMessage: "Invalid user identifier specified",
				Message:      e.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var user *nerdz.User
		if user, e = nerdz.NewUser(id); e != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				HumanMessage: "User does not exists",
				Message:      e.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var info UserInformations
		info.Info = user.Info().GetTO().(*nerdz.InfoTO)
		info.Contacts = user.ContactInfo().GetTO().(*nerdz.ContactInfoTO)
		info.Personal = user.PersonalInfo().GetTO().(*nerdz.PersonalInfoTO)

		out, err := selectFields(info, c)

		if err != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				HumanMessage: "Unable to fetch information for the specified user",
				Message:      "user.Info unable to get fields",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		return c.JSON(http.StatusOK, &Response{
			HumanMessage: "Correctly retrieved user information",
			Data:         out,
			Message:      "User.Info ok",
			Status:       http.StatusOK,
			Success:      true,
		})
	}
}

//UserFriends handles the request and returns the friend's of the specified user
func UserFriends() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id uint64
		var e error
		if id, e = strconv.ParseUint(c.Param("id"), 10, 64); e != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				HumanMessage: "Invalid user identifier specified",
				Message:      e.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var user *nerdz.User
		if user, e = nerdz.NewUser(id); e != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				HumanMessage: "User does not exists",
				Message:      e.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		friends := user.Friends()

		// Ops. No friends found
		if len(friends) == 0 {
			return c.JSON(http.StatusBadRequest, &Response{
				HumanMessage: "Unable to retrieve friends for the specified user",
				Message:      "User.Friends empty friends data",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		var friendsInfo []*UserInformations

		for _, u := range friends {
			friendsInfo = append(friendsInfo, &UserInformations{
				Info:     u.Info().GetTO().(*nerdz.InfoTO),
				Contacts: u.ContactInfo().GetTO().(*nerdz.ContactInfoTO),
				Personal: u.PersonalInfo().GetTO().(*nerdz.PersonalInfoTO),
			})
		}

		out, err := selectFields(friendsInfo, c)

		if err != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				HumanMessage: "Unable to retrieve friends for the specified user",
				Message:      "User.Friends select fields",
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}

		return c.JSON(http.StatusOK, &Response{
			HumanMessage: "Correctly retrieved friends",
			Data:         out,
			Message:      "User.Friends ok",
			Status:       http.StatusOK,
			Success:      true,
		})
	}
}
