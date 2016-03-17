package router_test

// router tests are a TODO: https://github.com/labstack/echo/issues/417
/*

import (
	//"encoding/json"
	"fmt"
	//"github.com/RangelReale/osin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
	//"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/rest"
	//"github.com/nerdzeu/nerdz-api/router"
	"testing"
)

var (
	userInfoURL, userFriendsURL, allUserPostsURL string
	userID                                       = "1"
	numPosts                                     = 10
	e                                            *echo.Echo
)

func init() {
	//e = router.Init(nerdz.Configuration.Port, nerdz.Configuration.EnableLog)
	e = echo.New()
}

/*
func init() {
	var storage *nerdz.OAuth2Storage
	var err error

	create := &osin.DefaultClient{
		Secret:      "secret 1",
		RedirectUri: "http://localhost/",
		UserData:    uint64(1),
	}

	if _, err = storage.CreateClient(create, "App 1"); err != nil {
		panic(fmt.Sprintf("Unable to create application client1: %s\n", err.Error()))
	}

	create2 := &osin.DefaultClient{
		Secret:      "secret 2",
		RedirectUri: "http://localhost/",
		UserData:    uint64(2),
	}

	if _, err = storage.CreateClient(create2, "App 2"); err != nil {
		panic(fmt.Sprintf("Unable to create application client2: %s\n", err.Error()))
	}

	e = router.Init(nerdz.Configuration.Port, nerdz.Configuration.EnableLog)

	userInfoURL = fmt.Sprintf("%s/users/%s", server.URL, userID)
	userFriendsURL = fmt.Sprintf("%s/users/%s/friends", server.URL, userID)
	allUserPostsURL = fmt.Sprintf("%s/users/%s/posts", server.URL, userID)
}


func TestUserInfo(t *testing.T) {
	t.Log("GET /users/1")
	h := rest.UserInfo()
	e.Get("/users/:id", h)
	req := test.NewRequest(echo.GET, "/users/1", nil)
	res := test.NewResponseRecorder()
	c := echo.NewContext(req, res, e)
	fmt.Println("%s", c.ParamNames())
	h.Handle(c)
	panic(res.Body.String())
}

/*
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

	if lenData := len(mapData["data"].(map[string]interface{})); lenData > numPosts {
		t.Errorf("Unable to retrieve correctly posts: lenData=%d > numPosts=%d", lenData, numPosts)
		t.FailNow()
	}

}
*/
