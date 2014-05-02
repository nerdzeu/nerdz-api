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
	if len(bl) != 1 {
		t.Error("Expected 1 user in blacklist, but got: %v\n", len(bl))
	}
}

func TestGetHome(t *testing.T) {

	// At most the last 10 posts from italian users
	userHome := user.GetUserHome(&nerdz.PostlistOptions{Following: false, Language: "it", N: 10})
	if len(*userHome) != 10 {
		t.Error("Expected 10 posts, but got: %+v\n", len(*userHome))
	}

	fmt.Printf("%+v\n", *userHome)

	// At most the last 10 project posts from italian users
	projectHome := user.GetProjectHome(&nerdz.PostlistOptions{Following: false, Language: "it", N: 10})
	if len(*projectHome) != 10 {
		t.Error("Expected 10 posts, but got: %+v\n", len(*projectHome))
	}

	fmt.Printf("%+v\n", *projectHome)

	// At most the last 10 posts from German users
	userHome = user.GetUserHome(&nerdz.PostlistOptions{Following: false, Language: "de", N: 10})
	if len(*userHome) != 0 {
		t.Error("Expected 0 posts, but got: %+v\n", len(*userHome))
	}

	fmt.Printf("%+v\n", *userHome)

	// At most the last 10 posts to English users from users that "user" is following
	userHome = user.GetUserHome(&nerdz.PostlistOptions{Following: true, Language: "en", N: 10})

	if len(*userHome) == 0 {
		t.Error("Expected at leat 1 post from an english user the 'user' is following. But 0 found")
	}

	fmt.Printf("%+v\n", *userHome)

	// The single post older (created before) the one with hpid 1000, from some user that 'user' follow and to an english speaking one
	userHome = user.GetUserHome(&nerdz.PostlistOptions{Following: true, Language: "en", N: 1, Older: 1000})

	if len(*userHome) != 1 {
		t.Errorf("Expeted 1 post, but got: %d", len(*userHome))
	}

	fmt.Printf("THE POST: %+v", (*userHome)[0])

	if (*userHome)[0].Hpid != 26 {
		t.Errorf("Post with hpid 26 expected, but got: %d", (*userHome)[0].Hpid)
	}

	// At most 2 posts in the Homepage formed by my posts and my friends posts
	userHome = user.GetUserHome(&nerdz.PostlistOptions{Following: true, Followers: true, N: 2})

	if len(*userHome) != 2 {
		t.Errorf("Expeted 2 posts, but got: %d", len(*userHome))
	}

	fmt.Printf("FRIENDZ: %v", *userHome)

	lastFriendPost := (*userHome)[0]

	// Get the (at max 20, in this case only 1) newer posts than the one with the "Newer" hpid, from friends
	userHome = user.GetUserHome(&nerdz.PostlistOptions{
		Following: true,
		Followers: true,
		Newer:     (*userHome)[1].Hpid})

	if len(*userHome) > 1 || (*userHome)[0].Hpid != lastFriendPost.Hpid {
		t.Errorf("Expected 1 post with hpid %d, but got %d posts and the first post has hpid = %d", lastFriendPost.Hpid, len(*userHome), (*userHome)[0].Hpid)
	}
}

func TestGetPostlist(t *testing.T) {
	postList := user.GetPostlist(nil).([]nerdz.UserPost)
	if len(postList) != 20 {
		t.Error("Expected 20  posts, but got: %+v\n", len(postList))
	}

	// Older than 1 (all) and newer than 8000 (no one) -> empty
	postList = user.GetPostlist(&nerdz.PostlistOptions{
		Older: 1,
		Newer: 80000}).([]nerdz.UserPost)

	if len(postList) != 0 {
		t.Errorf("Expected 0 posts. But got: %d", len(postList))
	}

	// Find posts between 103 and 97 inclusive, in user profile, from everybody.
	postList = user.GetPostlist(&nerdz.PostlistOptions{
		Older: 103,
		Newer: 97,
	}).([]nerdz.UserPost)

	if len(postList) != 4 {
		t.Errorf("Expected 4 posts. But got: %d", len(postList))
	}
}

func TestAddUserPost(t *testing.T) {
    var e error
	// New post on my board
	if e = user.AddUserPost(user, "All right"); e != nil {
		t.Errorf("AddUserPost with *User should work but, got: %v", e)
	}

	if e = user.AddUserPost(1, "All right"); e != nil {
		t.Errorf("AddUserPost with ID should work but, got: %v", e)
	}

	// post on the board of a blacklisted user should fail
	if e = user.AddUserPost(int8(5), "<script>alert('I wanna hack u!!!');</script>"); e == nil {
		t.Errorf("AddUserPost on a blacklisted user should fail. But in this case it succed :(")
	}

	// Post on a closed board should fail (if I'm not in its whitelist)
	if e = user.AddUserPost(7, "hi!"); e == nil {
		t.Errorf("AddUserPost on a closed user's board should fail. But in this case it succed :(")
	}

    fmt.Printf("AddUserPost on closed user's board failed and returned: %s\n", e.Error())
    // the e.Error() string should be handled in the same way we do in templates
}

func TestAddProjectPost(t *testing.T) {
	// New post on a project of mine
	myProject := user.GetProjects()[0]

	if e := user.AddProjectPost(myProject, "BEST ADMIN EVER :>\nHello!"); e != nil {
		t.Errorf("No errors should occur whie adding a post to a project of mine, but got: %v", e)
	}

}
