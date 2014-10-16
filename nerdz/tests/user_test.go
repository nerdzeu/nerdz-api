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

	if (*userHome)[0].Hpid != 36 {
		t.Errorf("Post with hpid 36 expected, but got: %d", (*userHome)[0].Hpid)
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
	var err error
	var hpid uint64
	// New post on my board
	if hpid, err = user.AddUserPost(user, "All right"); err != nil {
		t.Errorf("AddUserPost with *User should work but, got: %v", err)
	}

	if err = user.DeleteUserPost(hpid); err != nil {
		t.Errorf("DelUserPost with hpid %v shoud work, but got error: %v", hpid, err)
	}

	if hpid, err = user.AddUserPost(uint64(1), "All right2"); err != nil {
		t.Errorf("AddUserPost with ID should work but, got: %v", err)
	}

	// post on the board of a blacklisted user should fail
	if hpid, err = user.AddUserPost(uint64(5), "<script>alert('I wanna hack u!!!');</script>"); err == nil {
		t.Errorf("AddUserPost on a blacklisted user should fail. But in this case it succeded :(")
	}

	fmt.Print(hpid)

	// Post on a closed board should fail (if I'm not in its whitelist)
	if hpid, err = user.AddUserPost(uint64(7), "hi!"); err == nil {
		t.Errorf("AddUserPost on a closed user's board should fail. But in this case it succeded :(")
	}

	fmt.Printf("AddUserPost on closed user's board failed and returned: %s\n", err.Error())
	// the e.Error() string should be handled in the same way we do in templates
}

func TestAddProjectPost(t *testing.T) {
	// New post on a project of mine
	myProject := user.GetProjects()[0]

	if _, err := user.AddProjectPost(myProject, "BEST ADMIN EVER :>\nHello!"); err != nil {
		t.Errorf("No errors should occur whie adding a post to a project of mine, but got: %v", err)
	}
}

func TestAddComments(t *testing.T) {
	var err error
	var hcid uint64

	// Add Comment on a post on my profile
	if hcid, err = user.AddUserPostComment(uint64(103), "Nice <html>"); err != nil {
		t.Errorf("AddUserPostComment failed: %s", err.Error())
	}

	if err = user.DeleteUserPostComment(hcid); err != nil {
		t.Errorf("DelUserPostComment with hpid %v shoud work, but got error: %v", hcid, err)
	}

	// Add Cmment on a non existing post should fail
	if hcid, err = user.AddProjectPostComment(uint64(103), "SUPPPA GOMBLODDO\n\n汉语 or 漢語, Hànyǔ)"); err == nil {
		t.Error("Add ProjectPost on a non existing post should fail but succeeded")
	}

	// Add comment on an existing project post should work
	if hcid, err = user.AddProjectPostComment(uint64(11), "SUPPPA GOMBLODDO\n\n汉语 or 漢語, Hànyǔ)"); err != nil {
		t.Errorf("AddProjectPostComment failed: %s", err.Error())
	}

	// Add comment on a blacklisted profile post should fail
	stupid, _ := nerdz.NewUser(5)
	post := (stupid.GetPostlist(&nerdz.PostlistOptions{N: 1}).([]nerdz.UserPost))[0]

	if hcid, err = user.AddUserPostComment(&post, "THIS SHOULD FAIL"); err == nil {
		t.Errorf("Comment on a blacklisted profile post should fail, but in this case it succeeded")
	}

	fmt.Print(" BLACKLISTED COMMENT RETURN STRING: " + err.Error())
}
