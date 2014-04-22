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
	userHome := user.GetUserHome(&nerdz.PostlistOptions{Following: false, Language: "it", N: 10})
	if len(*userHome) != 10 {
		t.Error("Expected 10 posts, but got: %+v\n", len(*userHome))
	}

	fmt.Printf("%+v\n", *userHome)

	projectHome := user.GetProjectHome(&nerdz.PostlistOptions{Following: false, Language: "it", N: 10})
	if len(*projectHome) != 10 {
		t.Error("Expected 10 posts, but got: %+v\n", len(*projectHome))
	}

	fmt.Printf("%+v\n", *projectHome)

	userHome = user.GetUserHome(&nerdz.PostlistOptions{Following: false, Language: "de", N: 10})
	if len(*userHome) != 0 {
		t.Error("Expected 0 posts, but got: %+v\n", len(*userHome))
	}

	fmt.Printf("%+v\n", *userHome)

	userHome = user.GetUserHome(&nerdz.PostlistOptions{Following: true, Language: "en", N: 10})

	fmt.Printf("%+v\n", *userHome)
}

func TestGetPostlist(t *testing.T) {
	postList := user.GetPostlist(nil).([]nerdz.UserPost)
	if len(postList) != 20 {
		t.Error("Expected 20  posts, but got: %+v\n", len(postList))
	}

	fmt.Printf("%+v\n", postList)
}
