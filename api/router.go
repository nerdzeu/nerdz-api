package api

import (
	"strconv"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

// Start starts the API server on specified port.
// enableLog set to true enable echo middleware logger
func Start(port int16, enableLog bool) {
	e := echo.New()
	if enableLog {
		e.Use(mw.Logger())
	}

	e.Get("/users/:id/posts", UserPosts)
	e.Get("/users/:id/friends", UserFriends)
	e.Get("/users/:id", UserInfo)
	e.Run(":" + strconv.Itoa(int(port)))
}
