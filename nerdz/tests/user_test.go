package nerdz_test

import (
	"fmt"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"testing"
)

var user *nerdz.User

func init() {
	var err error
	user, err = nerdz.NewUser(1)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when create existing user, but got: %+v", err))
	}
}

func TestGetContactInfo(t *testing.T) {
	info := user.GetContactInfo()
	if info == nil {
		t.Error("null info")
	}
}

func TestGetPersonalInfo(t *testing.T) {
	info := user.GetPersonalInfo()
	if info == nil {
		t.Error("null info")
	}

	fmt.Printf("Struct: %+v\nINTERESTES:", *info)
	for i, elem := range info.Interests {
		fmt.Printf("%d) %s\n", i, elem)
	}

	fmt.Println("Quotes")
	for i, elem := range info.Quotes {
		fmt.Printf("%d) %s\n", i, elem)
	}

}

func TestGetBoardInfo(t *testing.T) {
	info := user.GetBoardInfo()
	if info == nil {
		t.Error("null info")
	}

	// If whitelist is not empty, the output will be huge (if tested with -v flag)
	fmt.Printf("%+v\n", *info)
}

func TestGetBlackList(t *testing.T) {
	bl := user.GetBlacklist()
	if len(bl) != 2 {
		t.Error("Expected 2 user in blacklist, but got: %v\n", len(bl))
	}
}

func TestGetHome(t *testing.T) {

	//Last 10 posts from italian users
	userHome := user.GetUserHome(&nerdz.PostlistOptions{Following: false, Language: "it", N: 10})
	if len(*userHome) != 10 {
		t.Error("Expected 10 posts, but got: %+v\n", len(*userHome))
	}

	fmt.Printf("%+v\n", *userHome)

	// Last 10 project posts from italian users
	projectHome := user.GetProjectHome(&nerdz.PostlistOptions{Following: false, Language: "it", N: 10})
	if len(*projectHome) != 10 {
		t.Error("Expected 10 posts, but got: %+v\n", len(*projectHome))
	}

	fmt.Printf("%+v\n", *projectHome)

	// Last 10 posts from German users
	userHome = user.GetUserHome(&nerdz.PostlistOptions{Following: false, Language: "de", N: 10})
	if len(*userHome) != 0 {
		t.Error("Expected 0 posts, but got: %+v\n", len(*userHome))
	}

	fmt.Printf("%+v\n", *userHome)

	// Last 10 posts to English users from users that "user" is following
	userHome = user.GetUserHome(&nerdz.PostlistOptions{Following: true, Language: "en", N: 10})

	if len(*userHome) == 0 {
		t.Error("Expected at leat 1 post from an english user the 'user' is following. But 0 found")
	}

	fmt.Printf("%+v\n", *userHome)

	// The single post after the one with hpid 86421, from some user that 'user' follow and to an english speaking one
	userHome = user.GetUserHome(&nerdz.PostlistOptions{Following: true, Language: "en", N: 1, After: 86421})

	if len(*userHome) != 1 {
		t.Errorf("Expeted 1 post, but got: %d", len(*userHome))
	}

	if (*userHome)[0].Hpid != 75657 {
		t.Errorf("Post with hpid 78475 expected, but got: %d", (*userHome)[0].Hpid)
	}

	// Homepage formed by my posts and my friends posts
	userHome = user.GetUserHome(&nerdz.PostlistOptions{Following: true, Followers: true, N: 2})

	if len(*userHome) != 2 {
		t.Errorf("Expeted 2 posts, but got: %d", len(*userHome))
	}

	fmt.Printf("FRIENDZ: %v", *userHome)
}

func TestGetPostlist(t *testing.T) {
	postList := user.GetPostlist(nil).([]nerdz.UserPost)
	if len(postList) != 20 {
		t.Error("Expected 20  posts, but got: %+v\n", len(postList))
	}

	//postList = user.GetPostlist(&nerdz.PostlistOptions{Following:

	fmt.Printf("%+v\n", postList)
}
