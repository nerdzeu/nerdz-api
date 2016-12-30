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

package stream

import (
	"github.com/RangelReale/osin"
	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
)

// Notifications is the route for the stream of notifications for the current user
func Notifications() echo.HandlerFunc {
	return func(c echo.Context) error {
		accessData := c.Get("accessData").(*osin.AccessData)
		if accessData == nil {
			return c.String(http.StatusInternalServerError, "Invalid authorization")
		}

		websocket.Handler(func(ws *websocket.Conn) {
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
		}).ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
