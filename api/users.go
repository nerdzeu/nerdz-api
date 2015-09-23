package api

import (
	"net/http"
	"strconv"

	"fmt"
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
)

//UserPosts handles the request and returns all the posts written
//by the specified user
func UserPosts(c *echo.Context) error {
	var id uint64
	var e error
	if id, e = strconv.ParseUint(c.Param("id"), 10, 64); e != nil {
		return c.JSON(http.StatusBadRequest, &nerdz.Response{
			HumanMessage: "Invalid user identifier specified",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	var user *nerdz.User
	if user, e = nerdz.NewUser(id); e != nil {
		return c.JSON(http.StatusBadRequest, &nerdz.Response{
			HumanMessage: "User does not exists",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	var options *nerdz.PostlistOptions
	if options, e = NewPostlistOptions(c); e != nil {
		return c.JSON(http.StatusBadRequest, &nerdz.Response{
			HumanMessage: e.Error(),
			Message:      "NewPostlistOptions error",
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	options.User = true
	posts := user.Postlist(options)

	if posts == nil {
		return c.JSON(http.StatusBadRequest, &nerdz.Response{
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
			fmt.Println(userPost.GetTO().(*nerdz.UserPostTO))
			postsAPI = append(postsAPI, userPost.GetTO().(*nerdz.UserPostTO))
		}
	}

	out, err := SelectFields(postsAPI, c)
	if err == nil {
		return c.JSON(http.StatusOK, &nerdz.Response{
			Data:         out,
			HumanMessage: "Correctly fetched post list for the specified user",
			Message:      "user.Postlist ok",
			Status:       http.StatusOK,
			Success:      true,
		})
	}

	return c.JSON(http.StatusBadRequest, &nerdz.Response{
		HumanMessage: "Error selecting required fields",
		Message:      err.Error(),
		Status:       http.StatusBadRequest,
		Success:      false,
	})
}

//UserInfo handles the request and returns all the basic information for the
//specified user
func UserInfo(c *echo.Context) error {
	var id uint64
	var e error
	if id, e = strconv.ParseUint(c.Param("id"), 10, 64); e != nil {
		return c.JSON(http.StatusBadRequest, &nerdz.Response{
			HumanMessage: "Invalid user identifier specified",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	var user *nerdz.User
	if user, e = nerdz.NewUser(id); e != nil {
		return c.JSON(http.StatusBadRequest, &nerdz.Response{
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

	out, err := SelectFields(info, c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &nerdz.Response{
			HumanMessage: "Unable to fetch information for the specified user",
			Message:      "user.Info unable to get fields",
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	return c.JSON(http.StatusOK, &nerdz.Response{
		HumanMessage: "Correctly retrieved user information",
		Data:         out,
		Message:      "User.Info ok",
		Status:       http.StatusOK,
		Success:      true,
	})

}

//UserFriends handles the request and returns the friend's of the specified user
func UserFriends(c *echo.Context) error {
	var id uint64
	var e error
	if id, e = strconv.ParseUint(c.Param("id"), 10, 64); e != nil {
		return c.JSON(http.StatusBadRequest, &nerdz.Response{
			HumanMessage: "Invalid user identifier specified",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	var user *nerdz.User
	if user, e = nerdz.NewUser(id); e != nil {
		return c.JSON(http.StatusBadRequest, &nerdz.Response{
			HumanMessage: "User does not exists",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	users := user.Friends()

	// Ops. No friends found
	if len(*users) == 0 {
		return c.JSON(http.StatusBadRequest, &nerdz.Response{
			HumanMessage: "Unable to retrieve friends for the specified user",
			Message:      "User.Friends empty friends data",
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	var friendsInfo []*UserInformations

	for _, u := range *users {
		friendsInfo = append(friendsInfo, &UserInformations{
			Info:     u.Info().GetTO().(*nerdz.InfoTO),
			Contacts: u.ContactInfo().GetTO().(*nerdz.ContactInfoTO),
			Personal: u.PersonalInfo().GetTO().(*nerdz.PersonalInfoTO),
		})
	}

	out, err := SelectFields(friendsInfo, c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &nerdz.Response{
			HumanMessage: "Unable to retrieve friends for the specified user",
			Message:      "User.Friends select fields",
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	return c.JSON(http.StatusOK, &nerdz.Response{
		HumanMessage: "Correctly retrieved friends",
		Data:         out,
		Message:      "User.Friends ok",
		Status:       http.StatusOK,
		Success:      true,
	})

}
