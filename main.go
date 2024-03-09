/*
Copyright (C) 2016-2020 Paolo Galeone <nessuno@nerdz.eu>

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

package main

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/router"
	"github.com/rs/cors"
)

func main() {
	// Initialize routes
	r := router.Init(nerdz.Configuration.EnableLog)
	// Enable CORS globally
	r.Use(echo.WrapMiddleware(cors.New(cors.Options{}).Handler))
	// Recover from panics
	r.Use(middleware.Recover())
	// Start the router
	if err := r.Start(":" + strconv.Itoa(int(nerdz.Configuration.Port))); err != nil {
		log.Fatal(err)
	}
}
