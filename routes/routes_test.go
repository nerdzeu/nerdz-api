package routes_test

import (
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/routes"
	"testing"
)

func TestServer(t *testing.T) {
	routes.Start(nerdz.Configuration.Port, nerdz.Configuration.EnableLog)
}
