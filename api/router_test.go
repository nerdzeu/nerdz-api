package api_test

import (
	"testing"

	"github.com/nerdzeu/nerdz-api/api"
	"github.com/nerdzeu/nerdz-api/nerdz"
)

func TestServer(t *testing.T) {
	api.Start(nerdz.Configuration.Port, nerdz.Configuration.EnableLog)
}
