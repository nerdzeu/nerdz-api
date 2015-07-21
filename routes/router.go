package routes

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"strconv"
)

var e *echo.Echo = echo.New()

func Start(port int16) {
	e.Use(mw.Logger())
	e.Get("/users/:id/posts", UserPosts)
	e.Run(":" + strconv.Itoa(int(port)))
}
