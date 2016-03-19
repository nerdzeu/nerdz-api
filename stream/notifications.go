package stream

import (
	"bufio"
	"errors"
	"github.com/RangelReale/osin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"golang.org/x/net/websocket"
	"io"
	"net"
	"net/http"
	"strconv"
)

type responseAdapter struct {
	http.ResponseWriter
	writer io.Writer
}

func (r *responseAdapter) Write(b []byte) (n int, err error) {
	return r.writer.Write(b)
}

func (r *responseAdapter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hj, ok := r.ResponseWriter.(http.Hijacker); ok {
		return hj.Hijack()
	}
	return nil, nil, errors.New("I'm not a Hijacker")
}

// Notifications is the route for the stream of notifications for the current user
func Notifications() echo.HandlerFunc {
	// Can't use standard.WrapHandler(websocket.Handler(func(ws *websocket.Conn)
	// because we need the context to fetch the stored accessData variable
	return func(c echo.Context) error {
		accessData := c.Get("accessData").(*osin.AccessData)
		if accessData == nil {
			return c.String(http.StatusInternalServerError, "Invalid authorization")
		}

		wsHandler := websocket.Handler(func(ws *websocket.Conn) {
			// Listen from notification sent on DB channel u<ID>
			nerdz.Db().Listen("u"+strconv.Itoa(int(accessData.UserData.(uint64))), func(payload ...string) {
				if len(payload) == 1 {
					if websocket.Message.Send(ws, payload[0]) != nil {
						return
					}
				}
			})

			// try to read from client (we dont' expect a message) to prevent websocket closing
			for {
				var m string
				if websocket.Message.Receive(ws, &m) != nil {
					// If here, client closed the connection
					break
				}
			}
		})

		w := &responseAdapter{
			c.Response().(*standard.Response).ResponseWriter,
			c.Response(),
		}
		r := c.Request().(*standard.Request).Request

		wsHandler.ServeHTTP(w, r)
		return nil
	}
}
