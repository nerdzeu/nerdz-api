package routes

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"strconv"
)

func Start(port int16, enableLog bool) {
	var e *echo.Echo = echo.New()

	if enableLog {
		e.Use(mw.Logger())
	}

	e.Get("/users/:id/posts", UserPosts)
	e.Run(":" + strconv.Itoa(int(port)))
}
