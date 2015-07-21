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

	var posts_n uint64

	n := c.Request().FormValue("n")
	if n == "" {
		posts_n = MAX_POSTS
	} else {
		if posts_n, e = strconv.ParseUint(n, 10, 8); e != nil {
			posts_n = MIN_POSTS
		} else {
			if posts_n > MAX_POSTS {
				posts_n = MAX_POSTS
			}
		}
	}

	if posts := user.Postlist(&nerdz.PostlistOptions{N: uint8(posts_n)}); posts == nil {
		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "Unable to fetch post list for the specified user",
			Message:      "user.Postlist error",
			Status:       http.StatusBadRequest,
			Success:      false,
		})
	} else {
		return c.JSON(http.StatusOK, &Response{
			Data:         posts,
			HumanMessage: "Correctly fetched post list for the specified user",
			Message:      "user.Postlist ok",
			Status:       http.StatusOK,
			Success:      true,
		})
	}
}

func User(c *echo.Context) error {
	return echo.NewHTTPError(http.StatusOK)
}
