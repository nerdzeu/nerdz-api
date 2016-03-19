package main

import (
	"github.com/labstack/echo/engine/standard"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/router"
	"github.com/rs/cors"
	"strconv"
)

func main() {
	// Initialize routes
	r := router.Init(nerdz.Configuration.EnableLog)
	// Enalble CORS globally
	r.Use(standard.WrapMiddleware(cors.New(cors.Options{}).Handler))
	// Start the router using standard (net/http) engine
	r.Run(standard.New(":" + strconv.Itoa(int(nerdz.Configuration.Port))))
}
