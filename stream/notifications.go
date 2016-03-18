package stream

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	"github.com/labstack/echo/engine/standard"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"golang.org/x/net/websocket"
)

func StreamNotification() echo.HandlerFunc {
	return standard.WrapHandler(websocket.Handler(func(ws *websocket.Conn) {
	}))
}
