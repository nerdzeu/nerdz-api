package routes

import (
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"net/http"
	"strconv"
)

func UserPosts(c *echo.Context) error {
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
	if options, e = NewPostlistOptions(c.Request()); e != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: e.Error(),
			Message:      "NewPostlistOptions error",
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	}

	if posts := user.Postlist(options); posts == nil {
		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "Unable to fetch post list for the specified user",
			Message:      "user.Postlist error",
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	} else {
		if out, err := SelectFields(posts, c.Request()); err == nil {
			return c.JSON(http.StatusOK, &Response{
				Data:         out,
				HumanMessage: "Correctly fetched post list for the specified user",
				Message:      "user.Postlist ok",
				Status:       http.StatusOK,
				Success:      true,
			})
		} else {
			return c.JSON(http.StatusBadRequest, &Response{
				HumanMessage: "Error selecting required fields",
				Message:      err.Error(),
				Status:       http.StatusBadRequest,
				Success:      false,
			})
		}
	}
}

func User(c *echo.Context) error {
	return echo.NewHTTPError(http.StatusOK)
}
