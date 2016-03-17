package main

import (
	"github.com/labstack/echo/engine/standard"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/router"
	"strconv"
)

func main() {
	// Configure the router
	r := router.Init(nerdz.Configuration.EnableLog)
	// Start the router using standard (net/http) engine
	r.Run(standard.New(":" + strconv.Itoa(int(nerdz.Configuration.Port))))
}
