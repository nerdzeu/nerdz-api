package api_test

import (
	"github.com/nerdzeu/nerdz-api/api"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"testing"
)

func TestServer(t *testing.T) {
	api.Start(nerdz.Configuration.Port, nerdz.Configuration.EnableLog)
}
