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
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/router"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func GETRequest(path, accessToken string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(echo.GET, path, nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+accessToken)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	return res
}

func DELETERequest(path, accessToken string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(echo.DELETE, path, nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+accessToken)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	return res
}

func POSTRequest(path, accessToken, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(echo.POST, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+accessToken)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	return res
}

func PUTRequest(path, accessToken, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(echo.PUT, path, strings.NewReader(body))
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+accessToken)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	res := httptest.NewRecorder()
	e.ServeHTTP(res, req)
	return res
}

func setUP() nerdz.OAuth2AccessData {
	var at nerdz.OAuth2AccessData
	nerdz.Db().First(&at, uint64(1))

	// since we got db access, we update the created_at field and make the request again
	at.CreatedAt = time.Now()
	if err := nerdz.Db().Updates(&at); err != nil {
		panic(err.Error())
	}
	return at
}

func cleanUP() {
	var at nerdz.OAuth2AccessData
	nerdz.Db().First(&at, uint64(1))

	// since we got db access, we update the created_at field and make the request again
	at.CreatedAt = time.Now()
	// Make the access token expire again to make next tests
	at.CreatedAt = time.Date(2010, 1, 1, 1, 1, 1, 1, time.UTC)
	if err := nerdz.Db().Updates(&at); err != nil {
		panic(err.Error())
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

func TestGETOnProjects(t *testing.T) {
	at := setUP()
	endpoint := "/v1/projects/1"
	var mapData igor.JSON
	//mapData := make(map[string]interface{})

	// Project 1 has 0 followers
	res := GETRequest(endpoint+"/followers", at.AccessToken)
	if res.Code != http.StatusOK {
		t.Fatalf("Error in GET request: status code=%d", res.Code)
	}
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %+v", err)
	}
	// User 1 has 5 followers
	if lenData := len(mapData["data"].([]interface{})); lenData != 0 {
		t.Fatalf("Incorrect retrived followers. Project(1) has 0 followers, but got %d", lenData)
	}

	res = GETRequest(endpoint+"/posts", at.AccessToken)

	if res.Code != http.StatusOK {
		t.Fatalf("Error in GET request: status code=%d", res.Code)
	}

	dec = json.NewDecoder(res.Body)

	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %+v", err)
	}

	if len(mapData["data"].([]interface{})) != 4 {
		t.Fatalf("Expected 4 posts, but got: %d\n", len(mapData["data"].([]interface{})))
	}

	res = GETRequest(endpoint+"/posts?n=2", at.AccessToken)

	if res.Code != http.StatusOK {
		t.Fatalf("Error in GET request: status code=%d", res.Code)
	}

	dec = json.NewDecoder(res.Body)

	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %+v", err)
	}

	if lenData := len(mapData["data"].([]interface{})); lenData != 2 {
		t.Fatalf("Unable to retrieve correctly posts: lenData=%d != 2", lenData)
	}

	res = GETRequest(endpoint+"/posts/2", at.AccessToken)
	if res.Code != http.StatusOK {
		t.Fatalf("Error in GET request: status code=%d, body: %s", res.Code, res.Body)
	}

	dec = json.NewDecoder(res.Body)

	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %+v", err)
	}

	if !strings.Contains(mapData["data"].(map[string]interface{})["message"].(string), "PROGETTO") {
		t.Fatalf("expected the PROGETTO:2 post, but got: %v", mapData["data"])
	}

	// PROGETTO.2 has 2 comments
	res = GETRequest(endpoint+"/posts/2/comments", at.AccessToken)

	if res.Code != http.StatusOK {
		t.Fatalf("Error in GET request: status code=%d", res.Code)
	}

	dec = json.NewDecoder(res.Body)

	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %+v", err)
	}

	if lenData := len(mapData["data"].([]interface{})); lenData != 2 {
		t.Fatalf("Incorrect number of comments in GET "+endpoint+"/posts/2/comments. Expected 2 got %d", lenData)
	}

	res = GETRequest(endpoint+"/posts/2/comments?n=1&fields=message", at.AccessToken)

	if res.Code != http.StatusOK {
		t.Fatalf("Error in GET request: status code=%d", res.Code)
	}

	dec = json.NewDecoder(res.Body)

	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %+v", err)
	}

	if lenData := len(mapData["data"].([]interface{})); lenData != 1 {
		t.Fatalf("Incorrect number of comments in GET "+endpoint+"/posts/20/comments?n=1&fields=message. Expected 1 got %d", lenData)
	}

	// test single comment based on comment id (hcid), extract hcid and message only
	res = GETRequest(endpoint+"/posts/2/comments/2?fields=message,hcid", at.AccessToken)
	dec = json.NewDecoder(res.Body)
	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %+v", err)
	}
	if lenData := len(mapData["data"].(map[string]interface{})); lenData != 2 {
		t.Fatalf("Incorrect number of comments in GET "+endpoint+"/posts/2/comments/2?fields=message,hcid. Expected 1 got %d", lenData)
	}

	data := mapData["data"].(map[string]interface{})
	if !strings.Contains(data["message"].(string), "Defollow") {
		t.Fatalf("Expected a message that contains Defollow, but got %s\n", data["message"].(string))
	}

	cleanUP()

}

// Test GET on Group /users and /me /projects
// we expect the same responses if :id = current logged user
func TestGETAndOnGroupUsers(t *testing.T) {
	endpoints := []string{"/v1/users/1", "/v1/me"}
	for _, endpoint := range endpoints {
		// Authorize
		// extract stored access_token
		// thus I manually generated and stored an access token for app1
		// this is done here only, in a real world application the user follow the OAuth flows and get the access token
		at := setUP()
		res := GETRequest(endpoint, at.AccessToken)

		// since we got db access, we update the created_at field and make the request again
		at.CreatedAt = time.Now()
		if err := nerdz.Db().Updates(&at); err != nil {
			t.Fatal(err.Error())
		}

		res = GETRequest(endpoint, at.AccessToken)

		dec := json.NewDecoder(res.Body)

		var mapData igor.JSON

		if err := dec.Decode(&mapData); err != nil {
			t.Fatalf("Unable to decode received data: %+v", err)
		}

		//mapData := make(map[string]interface{})

		res = GETRequest(endpoint+"/friends", at.AccessToken)

		if res.Code != http.StatusOK {
			t.Fatalf("Error in GET request: status code=%d", res.Code)
		}

		dec = json.NewDecoder(res.Body)

		if err := dec.Decode(&mapData); err != nil {
			t.Fatalf("Unable to decode received data: %+v", err)
		}

		// User 1 has 3 friends
		if lenData := len(mapData["data"].([]interface{})); lenData != 3 {
			t.Fatalf("Incorrect retrived friends. User(1) has 3 friends, got %d", lenData)
		}

		res = GETRequest(endpoint+"/following", at.AccessToken)
		if res.Code != http.StatusOK {
			t.Fatalf("Error in GET request: status code=%d", res.Code)
		}
		dec = json.NewDecoder(res.Body)
		if err := dec.Decode(&mapData); err != nil {
			t.Fatalf("Unable to decode received data: %+v", err)
		}
		// User 1 has 4 followers
		if lenData := len(mapData["data"].([]interface{})); lenData != 4 {
			t.Fatalf("Incorrect retrived following. User(1) has 5 followers, got %d", lenData)
		}

		// User 1 has 5 followers and 4 following
		res = GETRequest(endpoint+"/followers", at.AccessToken)
		if res.Code != http.StatusOK {
			t.Fatalf("Error in GET request: status code=%d", res.Code)
		}
		dec = json.NewDecoder(res.Body)
		if err := dec.Decode(&mapData); err != nil {
			t.Fatalf("Unable to decode received data: %+v", err)
		}
		// User 1 has 5 followers
		if lenData := len(mapData["data"].([]interface{})); lenData != 5 {
			t.Fatalf("Incorrect retrived followers. User(1) has 5 followers, got %d", lenData)
		}

		res = GETRequest(endpoint+"/posts", at.AccessToken)

		if res.Code != http.StatusOK {
			t.Fatalf("Error in GET request: status code=%d", res.Code)
		}

		dec = json.NewDecoder(res.Body)

		if err := dec.Decode(&mapData); err != nil {
			t.Fatalf("Unable to decode received data: %+v", err)
		}

		if len(mapData["data"].([]interface{})) != 20 {
			t.Fatalf("Expected 20 posts, but got: %d\n", len(mapData["data"].([]interface{})))
		}

		res = GETRequest(endpoint+"/posts?n=10", at.AccessToken)

		if res.Code != http.StatusOK {
			t.Fatalf("Error in GET request: status code=%d", res.Code)
		}

		dec = json.NewDecoder(res.Body)

		if err := dec.Decode(&mapData); err != nil {
			t.Fatalf("Unable to decode received data: %+v", err)
		}

		if lenData := len(mapData["data"].([]interface{})); lenData != 10 {
			t.Fatalf("Unable to retrieve correctly posts: lenData=%d != 10", lenData)
		}

		res = GETRequest(endpoint+"/posts/6", at.AccessToken)
		if res.Code != http.StatusOK {
			t.Fatalf("Error in GET request: status code=%d, body: %s", res.Code, res.Body)
		}

		dec = json.NewDecoder(res.Body)

		if err := dec.Decode(&mapData); err != nil {
			t.Fatalf("Unable to decode received data: %+v", err)
		}

		if !strings.Contains(mapData["data"].(map[string]interface{})["message"].(string), "PROGETTO") {
			t.Fatalf("expected the admin.6 post, but got: %v", mapData["data"])
		}

		// admin.20 has 3 comments
		res = GETRequest(endpoint+"/posts/20/comments", at.AccessToken)

		if res.Code != http.StatusOK {
			t.Fatalf("Error in GET request: status code=%d", res.Code)
		}

		dec = json.NewDecoder(res.Body)

		if err := dec.Decode(&mapData); err != nil {
			t.Fatalf("Unable to decode received data: %+v", err)
		}

		if lenData := len(mapData["data"].([]interface{})); lenData != 3 {
			t.Fatalf("Incorrect number of comments in GET "+endpoint+"/posts/20/comments. Expected 3 got %d", lenData)
		}

		res = GETRequest(endpoint+"/posts/20/comments?n=1&fields=message", at.AccessToken)

		if res.Code != http.StatusOK {
			t.Fatalf("Error in GET request: status code=%d", res.Code)
		}

		dec = json.NewDecoder(res.Body)

		if err := dec.Decode(&mapData); err != nil {
			t.Fatalf("Unable to decode received data: %+v", err)
		}

		if lenData := len(mapData["data"].([]interface{})); lenData != 1 {
			t.Fatalf("Incorrect number of comments in GET "+endpoint+"/posts/20/comments?n=1&fields=message. Expected 1 got %d", lenData)
		}

		// test single comment based on comment id (hcid), extract hcid and message only
		res = GETRequest(endpoint+"/posts/20/comments/224?fields=message,hcid", at.AccessToken)
		dec = json.NewDecoder(res.Body)
		if err := dec.Decode(&mapData); err != nil {
			t.Fatalf("Unable to decode received data: %+v", err)
		}
		if lenData := len(mapData["data"].(map[string]interface{})); lenData != 2 {
			t.Fatalf("Incorrect number of comments in GET "+endpoint+"/posts/20/comments/224?fields=message,hcid. Expected 1 got %d", lenData)
		}

		data := mapData["data"].(map[string]interface{})
		if !strings.Contains(data["message"].(string), "VEDERE GENTE") {
			t.Fatalf("Expected a message that contains VEDERE GENTE, but got %s\n", data["message"].(string))
		}
		cleanUP()
	}
}

func TestMeOnlyRoute(t *testing.T) {
	var mapData igor.JSON
	at := setUP()

	res := GETRequest("/v1/me/pms/4/11", at.AccessToken)
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %+v", err)
	}

	data := mapData["data"].(map[string]interface{})
	if !strings.Contains(data["message"].(string), "GABEN UNLEASHED") {
		t.Fatalf("Expected a message that contains GABEN UNLEASHED but got %s\n", data["message"].(string))
	}

	// Update: not supported from pms
	/*
		res = PUTRequest("/v1/me/pms/4/11", at.AccessToken, `{"message": "GABBANA", "lang": "it"}`)
		dec = json.NewDecoder(res.Body)
		if err := dec.Decode(&mapData); err != nil {
			t.Fatalf("Unable to decode received data: %+v", err)
		}

		data = mapData["data"].(map[string]interface{})
		if !strings.Contains(data["message"].(string), "GABBANA") {
			t.Fatalf("Expected a message that contains GABBANA but got %s\n", data["message"].(string))
		}

		if data["lang"].(string) != "it" {
			t.Fatalf("Expected a message in italian, but got: %s\n", data["lang"].(string))
		}
	*/

	// Delete
	res = DELETERequest("/v1/me/pms/4/11", at.AccessToken)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected a successfull DELETE, but got status: %d", res.Code)
	}

	// Delete the whole conversation
	res = DELETERequest("/v1/me/pms/4/", at.AccessToken)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected a successfull DELETE of conversation, but got status: %d", res.Code)
	}
	cleanUP()
}

func postCommentActions(t *testing.T, endpoint string) {
	var mapData igor.JSON
	at := setUP()

	// Post tests BEGIN
	res := POSTRequest(endpoint, at.AccessToken, `{"message": "POST TEST YEAH"}`)

	if res.Code == http.StatusUnauthorized {
		t.Fatalf("Error in POST request: should be authorized to POST "+endpoint+" but got status code: %d", res.Code)
	}

	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %v", err)
	}

	data := mapData["data"].(map[string]interface{})
	if !strings.Contains(data["message"].(string), "POST TEST YEAH") {
		t.Fatalf("Expected a message that contains POST TEST YEAH but got %s\n", data["message"].(string))
	}

	// extraxt the post id
	pid := strconv.Itoa(int(data["pid"].(float64)))

	// edit the message
	editEndpoint := endpoint + "/" + pid
	postEndpoint := editEndpoint
	res = PUTRequest(editEndpoint, at.AccessToken, `{"message": "post evviva evviva", "lang": "it"}`)
	if res.Code == http.StatusUnauthorized {
		t.Fatalf("Error in PUT request: should be authorized to PUT "+editEndpoint+" but got status code: %d", res.Code)
	}

	dec = json.NewDecoder(res.Body)
	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %v", err)
	}

	data = mapData["data"].(map[string]interface{})
	if !strings.Contains(data["message"].(string), "evviva") {
		t.Fatalf("Expected a message that contains evviva but got %s\n", data["message"].(string))
	}

	if data["lang"].(string) != "it" {
		t.Fatalf("Language should be updated to 'it', but found %s instead", data["lang"].(string))
	}

	// Upvote the post
	res = POSTRequest(editEndpoint+"/votes", at.AccessToken, `{"vote": 1}`)
	dec = json.NewDecoder(res.Body)
	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %v, %s, %s", err, editEndpoint+"/votes", res.Body.String())
	}
	data = mapData["data"].(map[string]interface{})
	if res.Code != http.StatusOK {
		t.Fatalf("Expected OK to upvote, but got status: %d: %s, %v", res.Code, editEndpoint+"/votes", res.Body.String())
	}

	// Delete the vote
	res = POSTRequest(editEndpoint+"/votes", at.AccessToken, `{"vote": 0}`)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected OK to delete the vote, but got status: %d", res.Code)
	}
	dec = json.NewDecoder(res.Body)

	// Downvote
	time.Sleep(5000 * time.Millisecond)
	res = POSTRequest(editEndpoint+"/votes", at.AccessToken, `{"vote": -1}`)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected OK to downvote, but got status: %d: %s : %s", res.Code, editEndpoint+"/votes", res.Body.String())
	}

	// Bookmark the post
	res = POSTRequest(editEndpoint+"/bookmarks", at.AccessToken, ``)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected OK to bookmark the post, but got status: %d: %s : %s", res.Code, editEndpoint+"/bookmarks", res.Body.String())
	}

	// Delete bookmark
	res = DELETERequest(editEndpoint+"/bookmarks", at.AccessToken)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected a successfull DELETE of bookmark, but got status: %d", res.Code)
	}

	// Lock the post
	res = POSTRequest(editEndpoint+"/locks", at.AccessToken, ``)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected OK to lock the post, but got status: %d: %s : %s", res.Code, editEndpoint+"/locks", res.Body.String())
	}

	// Unlock the post
	res = DELETERequest(editEndpoint+"/locks", at.AccessToken)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected a successfull DELETE of lock, but got status: %d", res.Code)
	}

	// Post tests END
	// Comment tests BEGIN

	// add a comment to the new post
	endpoint += "/" + pid + "/comments"
	res = POSTRequest(endpoint, at.AccessToken, `{"message": "commento in italiano :DD", "lang": "it"}`)
	if res.Code != http.StatusOK {
		t.Fatalf("Error in POST request: should be authorized to POST "+endpoint+" but got status code: %d and respose: %s", res.Code, res.Body)
	}

	dec = json.NewDecoder(res.Body)
	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %v", err)
	}

	data = mapData["data"].(map[string]interface{})
	if !strings.Contains(data["message"].(string), "italiano :DD") {
		t.Fatalf("Expected a message that contains italiano :DD but got %s\n", data["message"].(string))
	}

	// extraxt the comment id
	hcid := strconv.Itoa(int(data["hcid"].(float64)))

	// Edit the comment
	// Wait 5 second to avoid flood limit (db side)
	time.Sleep(5000 * time.Millisecond)

	editEndpoint = endpoint + "/" + hcid
	res = PUTRequest(editEndpoint, at.AccessToken, `{"message": "english comment", "lang": "en"}`)
	if res.Code == http.StatusUnauthorized {
		t.Fatalf("Error in PUT request: should be authorized to PUT "+editEndpoint+" but got status code: %d", res.Code)
	}

	dec = json.NewDecoder(res.Body)
	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %v", err)
	}

	data = mapData["data"].(map[string]interface{})
	if !strings.Contains(data["message"].(string), "english") {
		t.Fatalf("Expected a message that contains english but got %s\n", data["message"].(string))
	}

	if data["lang"].(string) != "en" {
		t.Fatalf("Language should be updated to 'en', but found %s instead", data["lang"].(string))
	}

	time.Sleep(6000 * time.Millisecond)

	// Upvote the comment
	res = POSTRequest(editEndpoint+"/votes", at.AccessToken, `{"vote": 1}`)
	dec = json.NewDecoder(res.Body)
	if err := dec.Decode(&mapData); err != nil {
		t.Fatalf("Unable to decode received data: %v, %s", err, res.Body.String())
	}
	data = mapData["data"].(map[string]interface{})
	if res.Code != http.StatusOK {
		t.Fatalf("Expected OK to upvote, but got status: %d: %s, %v", res.Code, editEndpoint+"/votes", res.Body.String())
	}

	// Delete the vote
	res = POSTRequest(editEndpoint+"/votes", at.AccessToken, `{"vote": 0}`)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected OK to delete the vote, but got status: %d", res.Code)
	}

	// Downvote
	res = POSTRequest(editEndpoint+"/votes", at.AccessToken, `{"vote": -1}`)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected OK to downvote, but got status: %d", res.Code)
	}

	// Delete the comment
	res = DELETERequest(editEndpoint, at.AccessToken)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected a successfull DELETE, but got status: %d", res.Code)
	}

	// Comment tests END

	// Delete the post
	res = DELETERequest(postEndpoint, at.AccessToken)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected a successfull DELETE, but got status: %d", res.Code)
	}

	cleanUP()
}

func TestPOSTPUTDELETEOnUsersGroup(t *testing.T) {
	at := setUP()
	endpoint := "/v1/users/16/posts/1/lurks"
	// Lurk on post that I have not posted nor commented
	res := POSTRequest(endpoint, at.AccessToken, ``)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected OK to lurks the post, but got status: %d: %s : %s", res.Code, endpoint, res.Body.String())
	}
	// Delete lurk
	res = DELETERequest(endpoint, at.AccessToken)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected a successfull DELETE of lurks, but got status: %d", res.Code)
	}
	cleanUP()
	endpoint = "/v1/users/19/posts"
	postCommentActions(t, endpoint)
}

func TestPOSTPUTDELETEOnMeGroup(t *testing.T) {
	endpoint := "/v1/me/posts"
	postCommentActions(t, endpoint)
}

func TestPOSTPUTDELETEOnProjectsGroup(t *testing.T) {
	at := setUP()
	endpoint := "/v1/projects/7/posts/1/lurks"
	// Lurk on post that I have not posted nor commented
	res := POSTRequest(endpoint, at.AccessToken, ``)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected OK to lurks the post, but got status: %d: %s : %s", res.Code, endpoint, res.Body.String())
	}
	// Delete lurk
	res = DELETERequest(endpoint, at.AccessToken)
	if res.Code != http.StatusOK {
		t.Fatalf("Expected a successfull DELETE of lurks, but got status: %d", res.Code)
	}
	cleanUP()
	endpoint = "/v1/projects/1/posts"
	postCommentActions(t, endpoint)
}
