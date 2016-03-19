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

package router_test

import (
	"encoding/json"
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/galeone/igor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/router"
	"net/http"
	"strings"
	"testing"
	"time"
)

var e *echo.Echo
var oauth *nerdz.OAuth2Storage

var app1, app2 osin.Client

func createOAuth2Client(app_name, secret, redirect_uri string, app_owner uint64) *nerdz.OAuth2Client {
	create := &osin.DefaultClient{
		Secret:      secret,
		RedirectUri: redirect_uri,
		UserData:    app_owner,
	}
	var client *nerdz.OAuth2Client
	var err error
	if client, err = oauth.CreateClient(create, app_name); err != nil {
		panic(fmt.Sprintf("Unable to create application %s: %s\n", app_name, err.Error()))
	}

	return client
}

func deleteOAuth2Client(client_id uint64) {
	if e := oauth.RemoveClient(client_id); e != nil {
		panic(e.Error())
	}
}

func init() {
	// Initialize every route to test
	e = router.Init(nerdz.Configuration.EnableLog)

	// Create two test app for OAuth 2 request
	//app1 = createOAuth2Client("app 1", "secret 1", "http://localhost/", uint64(1))
	//app2 = createOAuth2Client("app 2", "secret 2", "http://localhost/", uint64(1))
	// Get two test app
	app1, _ = oauth.GetClient("1")
	app2, _ = oauth.GetClient("2")
}

// Test REST requests

// Test GET on Group /users/

func TestGETOnGroupUsers(t *testing.T) {
	t.Log("GET /users/1")
	req := test.NewRequest(echo.GET, "/users/1", nil)
	res := test.NewResponseRecorder()
	e.ServeHTTP(req, res)

	if res.Status() != http.StatusUnauthorized {
		t.Fatalf("Error in GET request: should't be authorized to GET /users/1 but got status code: %d", res.Status())
	}

	// Authorize
	// extract stored access_token because osin can't work with engine/test.Request
	// thus I manually generated and stored an access token for app1
	// this is done here only, in a real world application the user follow the OAuth flows and get the access token
	var at nerdz.OAuth2AccessData
	nerdz.Db().First(&at, uint64(1))
	req = test.NewRequest(echo.GET, "/users/1", nil)
	req.Header().Set("Authorization", "Bearer "+at.AccessToken)
	res = test.NewResponseRecorder()
	e.ServeHTTP(req, res)

	// This request should fail, because access token is expired
	// A real Application will handle this, requesting a new access token or (better) a refresh token
	if !strings.Contains(res.Body.String(), "expired") {
		t.Fatalf("The access token used in Authorization Bearer should be expired, but it's not")
	}

	// since we got db access, we update the created_at field and make the request again
	at.CreatedAt = time.Now()
	if err := nerdz.Db().Updates(&at); err != nil {
		t.Fatal(err.Error())
	}

	req = test.NewRequest(echo.GET, "/users/1", nil)
	req.Header().Set("Authorization", "Bearer "+at.AccessToken)
	res = test.NewResponseRecorder()
	e.ServeHTTP(req, res)

	dec := json.NewDecoder(res.Body)

	var mapData igor.JSON

	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %+v", err)
	}

	t.Log("GET /users/1/friends")
	req = test.NewRequest(echo.GET, "/users/1/friends", nil)
	req.Header().Set("Authorization", "Bearer "+at.AccessToken)
	res = test.NewResponseRecorder()
	e.ServeHTTP(req, res)

	if res.Status() != http.StatusOK {
		t.Fatalf("Error in GET request: status code=%d", res.Status())
	}

	dec = json.NewDecoder(res.Body)

	friendsData := make(map[string]interface{})

	if err := dec.Decode(&friendsData); err != nil {
		t.Fatalf("Unable to decode received data: %+v", err)
	}

	// User 1 has 3 friends
	if lenData := len(friendsData["data"].(map[string]interface{})); lenData != 3 {
		t.Errorf("Incorrect retrived friends. User(1) has 3 friends, got %d", lenData)
	}

	t.Log("GET /users/1/posts")
	req = test.NewRequest(echo.GET, "/users/1/posts", nil)
	req.Header().Set("Authorization", "Bearer "+at.AccessToken)
	res = test.NewResponseRecorder()
	e.ServeHTTP(req, res)

	if res.Status() != http.StatusOK {
		t.Fatalf("Error in GET request: status code=%d", res.Status())
	}

	dec = json.NewDecoder(res.Body)

	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %+v", err)
	}

	if len(mapData["data"].(map[string]interface{})) != 20 {
		t.Errorf("Expected 20 posts, but got: %d\n", len(mapData["data"].(map[string]interface{})))
	}

	t.Logf("GET /users/1/posts?n=10")
	req = test.NewRequest(echo.GET, "/users/1/posts?n=10", nil)
	req.Header().Set("Authorization", "Bearer "+at.AccessToken)
	res = test.NewResponseRecorder()
	e.ServeHTTP(req, res)

	if res.Status() != http.StatusOK {
		t.Fatalf("Error in GET request: status code=%d", res.Status())
	}

	dec = json.NewDecoder(res.Body)

	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %+v", err)
	}

	if lenData := len(mapData["data"].(map[string]interface{})); lenData != 10 {
		t.Fatalf("Unable to retrieve correctly posts: lenData=%d != 10", lenData)
	}
}
