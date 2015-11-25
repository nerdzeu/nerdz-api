package api_test

import (
	"testing"

	"github.com/RangelReale/osin"
	"github.com/nerdzeu/nerdz-api/api"
	"github.com/nerdzeu/nerdz-api/nerdz"
)

func TestServer(t *testing.T) {
	var storage *nerdz.OAuth2Storage
	var err error

	create := &osin.DefaultClient{
		Secret:      "secret 1",
		RedirectUri: "http://localhost/",
		UserData:    uint64(1),
	}

	if _, err = storage.CreateClient(create, "App 1"); err != nil {
		t.Errorf("Unable to create application client1: %s\n", err.Error())
	}

	create2 := &osin.DefaultClient{
		Secret:      "secret 2",
		RedirectUri: "http://localhost/",
		UserData:    uint64(2),
	}

	if _, err = storage.CreateClient(create2, "App 2"); err != nil {
		t.Errorf("Unable to create application client2: %s\n", err.Error())
	}

	api.Start(nerdz.Configuration.Port, nerdz.Configuration.EnableLog)
}
