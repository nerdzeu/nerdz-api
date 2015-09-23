package api_test

import (
	"encoding/json"
	"fmt"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"net/http"
	"testing"
)

var (
	userID          = "1"
	numPosts        = 10
	mainURL         = fmt.Sprintf("http://localhost:%d", nerdz.Configuration.Port)
	userInfoURL     = fmt.Sprintf("%s/users/%s", mainURL, userID)
	userFriendsURL  = fmt.Sprintf("%s/users/%s/friends", mainURL, userID)
	allUserPostsURL = fmt.Sprintf("%s/users/%s/posts", mainURL, userID)
	nUserPostsURL   = fmt.Sprintf("%s/users/%s/posts?n=%d", mainURL, userID, numPosts)
)

func TestUserInfo(t *testing.T) {
	t.Log("Trying to retrieve User(1)'s information")

	res, err := http.DefaultClient.Get(userInfoURL)

	if err != nil {
		t.Errorf("Error in GET request: %+v", err)
		t.FailNow()
	} else if res.StatusCode != http.StatusOK {
		t.Errorf("Error in GET request: status code=%s", res.Status)
		t.FailNow()
	}

	dec := json.NewDecoder(res.Body)

	mapData := map[string]interface{}{}

	if err := dec.Decode(&mapData); err != nil {
		t.Errorf("Unable to decode received data: %+v", err)
		t.FailNow()
	}

}

func TestUserFriends(t *testing.T) {

	t.Log("Trying to retrieve User(1)'s friends")

	res, err := http.DefaultClient.Get(userFriendsURL)

	if err != nil {
		t.Errorf("Error in GET request: %+v", err)
		t.FailNow()
	} else if res.StatusCode != http.StatusOK {
		t.Errorf("Error in GET request: status code=%s", res.Status)
		t.FailNow()
	}

	dec := json.NewDecoder(res.Body)

	friendsData := map[string]interface{}{}

	if err := dec.Decode(&friendsData); err != nil {
		t.Errorf("Unable to decode received data: %+v", err)
		t.FailNow()
	}

	// User 1 has friends
	if lenData := len(friendsData["data"].(map[string]interface{})); lenData != 3 {
		t.Errorf("Incorrect retrived friends. User(1) has 3 friends, got %d", lenData)
	}

}

func TestAllUserPosts(t *testing.T) {
	t.Log("Trying to retrieve all User(1)'s posts")

	res, err := http.DefaultClient.Get(allUserPostsURL)

	if err != nil {
		t.Errorf("Error in GET request: %+v", err)
		t.FailNow()
	} else if res.StatusCode != http.StatusOK {
		t.Errorf("Error in GET request: status code=%s", res.Status)
		t.FailNow()
	}

	dec := json.NewDecoder(res.Body)

	mapData := map[string]interface{}{}

	if err := dec.Decode(&mapData); err != nil {
		t.Errorf("Unable to decode received data: %+v", err)
		t.FailNow()
	}
}

func TestTenUserPosts(t *testing.T) {
	t.Logf("Trying to retrieve <%d> User(1)'s posts", numPosts)

	res, err := http.DefaultClient.Get(userInfoURL)

	if err != nil {
		t.Errorf("Error in GET request: %+v", err)
		t.FailNow()
	} else if res.StatusCode != http.StatusOK {
		t.Errorf("Error in GET request: status code=%s", res.Status)
		t.FailNow()
	}

	dec := json.NewDecoder(res.Body)

	mapData := map[string]interface{}{}

	if err := dec.Decode(&mapData); err != nil {
		t.Errorf("Unable to decode received data: %+v", err)
		t.FailNow()
	}

	if lenData := len(mapData["data"].(map[string]interface{})); lenData == numPosts {
		t.Errorf("Unable to retrieve correctly posts: lenData=%d > numPosts=%d", lenData, numPosts)
		t.FailNow()
	}
}
