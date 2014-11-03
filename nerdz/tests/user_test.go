package nerdz_test

import (
	"fmt"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"testing"
)

var me, blacklisted, withClosedProfile *nerdz.User

func init() {
	var err error
	me, err = nerdz.NewUser(1)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when create existing user, but got: %+v", err))
	}
	blacklisted, _ = nerdz.NewUser(5)
	withClosedProfile, _ = nerdz.NewUser(7)

}

func TestContactInfo(t *testing.T) {
	info := me.ContactInfo()
	if info == nil {
		t.Error("null info")
	}
}

func TestPersonalInfo(t *testing.T) {
	info := me.PersonalInfo()
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

func TestBoardInfo(t *testing.T) {
	info := me.BoardInfo()
	if info == nil {
		t.Error("null info")
	}

	// If whitelist is not empty, the output will be huge (if tested with -v flag)
	fmt.Printf("%+v\n", *info)
	fmt.Printf("Template: %+v", *info.Template)
}

func TestBlackList(t *testing.T) {
	bl := me.Blacklist()
	if len(bl) != 1 {
		t.Error("Expected 1 user in blacklist, but got: %v\n", len(bl))
	}
}

func TestHome(t *testing.T) {
	// At most the last 10 posts from italian users
	userHome := me.UserHome(&nerdz.PostlistOptions{Following: false, Language: "it", N: 10})
	if len(*userHome) != 10 {
		t.Error("Expected 10 posts, but got: %+v\n", len(*userHome))
	}

	fmt.Printf("%+v\n", *userHome)

	// At most the last 10 project posts from italian users
	projectHome := me.ProjectHome(&nerdz.PostlistOptions{Following: false, Language: "it", N: 10})
	if len(*projectHome) != 10 {
		t.Error("Expected 10 posts, but got: %+v\n", len(*projectHome))
	}

	fmt.Printf("%+v\n", *projectHome)

	// At most the last 10 posts from German users
	userHome = me.UserHome(&nerdz.PostlistOptions{Following: false, Language: "de", N: 10})
	if len(*userHome) != 0 {
		t.Error("Expected 0 posts, but got: %+v\n", len(*userHome))
	}

	fmt.Printf("%+v\n", *userHome)

	// At most the last 10 posts to English users from users that "user" is following
	userHome = me.UserHome(&nerdz.PostlistOptions{Following: true, Language: "en", N: 10})

	if len(*userHome) == 0 {
		t.Error("Expected at leat 1 post from an english user the 'user' is following. But 0 found")
	}

	fmt.Printf("%+v\n", *userHome)

	// The single post older (created before) the one with hpid 1000, from some user that 'user' follow and to an english speaking one
	userHome = me.UserHome(&nerdz.PostlistOptions{Following: true, Language: "en", N: 1, Older: 1000})

	if len(*userHome) != 1 {
		t.Errorf("Expeted 1 post, but got: %d", len(*userHome))
	}

	fmt.Printf("THE POST: %+v", (*userHome)[0])

	if (*userHome)[0].Hpid != 36 {
		t.Errorf("Post with hpid 36 expected, but got: %d", (*userHome)[0].Hpid)
	}

	// At most 2 posts in the Homepage formed by my posts and my friends posts
	userHome = me.UserHome(&nerdz.PostlistOptions{Following: true, Followers: true, N: 2})

	if len(*userHome) != 2 {
		t.Errorf("Expeted 2 posts, but got: %d", len(*userHome))
	}

	fmt.Printf("FRIENDZ: %v", *userHome)

	lastFriendPost := (*userHome)[0]

	// Get the (at max 20, in this case only 1) newer posts than the one with the "Newer" hpid, from friends
	userHome = me.UserHome(&nerdz.PostlistOptions{
		Following: true,
		Followers: true,
		Newer:     (*userHome)[1].Hpid})

	if len(*userHome) > 1 || (*userHome)[0].Hpid != lastFriendPost.Hpid {
		t.Errorf("Expected 1 post with hpid %d, but got %d posts and the first post has hpid = %d", lastFriendPost.Hpid, len(*userHome), (*userHome)[0].Hpid)
	}
}

func TestPostlist(t *testing.T) {
	postList := me.Postlist(nil).([]nerdz.UserPost)
	if len(postList) != 20 {
		t.Error("Expected 20  posts, but got: %+v\n", len(postList))
	}

	// Older than 1 (all) and newer than 8000 (no one) -> empty
	postList = me.Postlist(&nerdz.PostlistOptions{
		Older: 1,
		Newer: 80000}).([]nerdz.UserPost)

	if len(postList) != 0 {
		t.Errorf("Expected 0 posts. But got: %d", len(postList))
	}

	// Find posts between 103 and 97 inclusive, in user profile, from everybody.
	postList = me.Postlist(&nerdz.PostlistOptions{
		Older: 103,
		Newer: 97,
	}).([]nerdz.UserPost)

	if len(postList) != 4 {
		t.Errorf("Expected 4 posts. But got: %d", len(postList))
	}
}

func TestAddEditDelete(t *testing.T) {
	var err error
	var hpid uint64
	// New post on my board
	if hpid, err = me.Add(&nerdz.UserPost{Message: "All right"}); err != nil {
		t.Errorf("AddUserPost with *User should work but, got: %v", err)
	}

	if err = me.Delete(&nerdz.UserPost{Hpid: hpid}); err != nil {
		t.Errorf("DelUserPost with hpid %v shoud work, but got error: %v", hpid, err)
	}

	if hpid, err = me.Add(&nerdz.UserPost{Message: "All right2"}); err != nil {
		t.Errorf("AddUserPost with ID should work but, got: %v", err)
	}

	var thisPost *nerdz.UserPost

	if thisPost, err = nerdz.NewUserPost(hpid); err != nil {
		t.Errorf("NewUserPost with hpid %d failed: %s", hpid, err)
	}

	thisPost.Message = "Post updated -> :D\nwow JA JA JA"
	thisPost.Lang = "fu"
	// Language "fu" does not exists, this edit should fail
	if err := me.Edit(thisPost); err == nil {
		t.Errorf("Edit post language and message not failed!", err)
	}

	thisPost.Lang = "de"
	if err := me.Edit(thisPost); err != nil {
		t.Errorf("This edit shold work but got %s1", err)
	}

	// post on the board of a blacklisted user should fail
	if hpid, err = me.Add(&nerdz.UserPost{To: blacklisted.Counter, Message: "<script>alert('I wanna hack u!!!');</script>"}); err == nil {
		t.Errorf("AddUserPost on a blacklisted user should fail. But in this case it succeded :(")
	}

	// Post on a closed board should fail (if I'm not in its whitelist)
	if hpid, err = me.Add(&nerdz.UserPost{Message: "hi!", To: withClosedProfile.Counter}); err == nil {
		t.Errorf("AddUserPost on a closed user's board should fail. But in this case it succeded :(")
	}

	fmt.Printf("AddUserPost on closed user's board failed and returned: %s\n", err.Error())
	// the e.Error() string should be handled in the same way we do in templates
}

func TestAddProjectPost(t *testing.T) {
	var hpid uint64
	var err error

	myProject := me.Projects()[0]
	if hpid, err = me.Add(&nerdz.ProjectPost{To: myProject.Counter, Message: "BEST ADMIN EVER :>\nHello!"}); err != nil {
		t.Errorf("No errors should occur whie adding a post to a project of mine, but got: %v", err)
	}

	if err = me.Delete(&nerdz.ProjectPost{Hpid: hpid}); err != nil {
		t.Errorf("DeleteProjectPost failed with error: %s", err.Error())
	}
}

func TestAddComments(t *testing.T) {
	var err error
	var hcid uint64

	// Add Comment on a post on my profile
	existingPost := me.Postlist(&nerdz.PostlistOptions{N: 1}).([]nerdz.UserPost)[0]
	existingPost.Message = "Nice <html>"
	if hcid, err = me.Add(&existingPost); err != nil {
		t.Errorf("AddUserPostComment failed: %s", err.Error())
	}

	if err = me.Delete(&existingPost); err != nil {
		t.Errorf("DelUserPostComment with hpid %v shoud work, but got error: %v", hcid, err)
	}

	// Add comment on a non existing post should fail
	if hcid, err = me.Add(&nerdz.ProjectPost{To: 103000, Message: "SUPPPA GOMBLODDO\n\n汉语 or 漢語, Hànyǔ)"}); err == nil {
		t.Error("AddProjectPostComment on a non existing post should fail but succeeded")
	}

	myProject := me.Projects()[0]
	// Add comment on an existing project post should work
	projectPost := myProject.Postlist(&nerdz.PostlistOptions{N: 1}).([]nerdz.ProjectPost)[0]
	if hcid, err = me.Add(&nerdz.ProjectPostComment{Hpid: projectPost.Hpid, Message: "lol k"}); err != nil {
		t.Errorf("AddProjectPostComment on an existing post sould work but failed with error: %s", err.Error())
	}
	// get the last comment
	lastComment := projectPost.Comments(1).([]nerdz.ProjectPostComment)[0]
	if hcid != lastComment.Hcid {
		t.Errorf("Fetched the wrong comment. Expected: %v but got %v", hcid, lastComment.Hcid)
	}

	if err := me.Delete(&lastComment); err != nil {
		t.Errorf("DelUserPost with hcid %v shoud work, but got error: %v", hcid, err)
	}

	projectPost.Message = "SUPPPA GOMBLODDO\n\n汉语 or 漢語, Hànyǔ)"
	if hcid, err = me.Add(&projectPost); err != nil {
		t.Errorf("AddProjectPostComment failed: %s", err.Error())
	}

	if err = me.Delete(&projectPost); err != nil {
		t.Errorf("DelProjectPostComment with hpid %v shoud work, but got error: %v", hcid, err)
	}

	// Add comment on a blacklisted profile post should fail
	post := (blacklisted.Postlist(&nerdz.PostlistOptions{N: 1}).([]nerdz.UserPost))[0]

	if hcid, err = me.Add(&post, "THIS SHOULD FAIL"); err == nil {
		t.Errorf("Comment on a blacklisted profile post should fail, but in this case it succeeded")
	}

	fmt.Print(" BLACKLISTED COMMENT RETURN STRING: " + err.Error())
}
