package nerdz_test

import (
	"fmt"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"net/url"
	"testing"
)

var userPost, userPost1 *nerdz.UserPost
var projectPost *nerdz.ProjectPost
var e error

func init() {
	projectPost, e = nerdz.NewProjectPost(3)
	if e != nil {
		panic(fmt.Sprintf("No error should happen when create existing post, but got: %+v", e))
	}

	userPost, e = nerdz.NewUserPost(6)

	if e != nil {
		panic(fmt.Sprintf("No error should happen when create existing post, but got: %+v", e))
	}

	fmt.Printf("** PRJ POST **\n%+v\n**USER POST **\n%+v\n", projectPost, userPost)

	userPost1, _ = nerdz.NewUserPost(20)
}

func TestGetFrom(t *testing.T) {
	from, err := userPost.GetFrom()
	if err != nil {
		t.Errorf("No error should happen when fetching existing user, but got: %+v", err)
	}

	if from.Counter != 1 {
		t.Errorf("Counter should be 1, but go: %d", from.Counter)
	}

	fromPrj, err := projectPost.GetFrom()
	if err != nil {
		t.Errorf("No error should happen when fetching existing user, but got: %+v", err)
	}

	if fromPrj.Counter != 4 {
		t.Errorf("Counter should be 4, but go: %d", fromPrj.Counter)
	}

	fmt.Printf("%+v\n", fromPrj)
}

func TestGetTo(t *testing.T) {
	to, err := userPost.GetTo()
	if err != nil {
		t.Errorf("No error should happen when fetching existing user, but got: %+v", err)
	}

	if to.Counter != 1 {
		t.Errorf("Counter should be 1, but go: %d", to.Counter)
	}

	toPrj, err := projectPost.GetTo()
	if err != nil {
		t.Errorf("No error should happen when fetching existing user, but got: %+v", err)
	}

	if toPrj.Counter != 3 {
		t.Errorf("Counter should be 3, but go: %d", toPrj.Counter)
	}

	fmt.Printf("%+v\n", toPrj)
}

func TestGetComments(t *testing.T) {
	comments := userPost.GetComments().([]nerdz.UserComment)
	if len(comments) == 0 {
		t.Error("No comments found. Expected > 1")
	}

	comments = userPost.GetComments(4).([]nerdz.UserComment)
	if len(comments) != 4 {
		t.Errorf("Expected the last 4 comments, got: %d", len(comments))
	}
	fmt.Printf("%+v\n", comments)

	comment := userPost.GetComments(4, 5).([]nerdz.UserComment)
	if len(comment) != 2 {
		t.Errorf("Expected 2 comments, received: %d", len(comment))
	}
	fmt.Printf("%+v\n", comment)

	prjComments := projectPost.GetComments().([]nerdz.ProjectComment)
	if len(prjComments) == 0 {
		t.Error("No comments found. Expected > 1")
	}

	prjComments = projectPost.GetComments(4).([]nerdz.ProjectComment)
	if len(prjComments) != 1 {
		t.Errorf("Expected the last  comment, got: %d", len(prjComments))
	}
	fmt.Printf("%+v\n", prjComments)

	prjComment := projectPost.GetComments(4, 4).([]nerdz.ProjectComment)
	if len(prjComment) != 0 {
		t.Errorf("Expected no comment, received: %d", len(prjComment))
	}
	fmt.Printf("%+v\n", prjComment)
}

func TestGetThumbs(t *testing.T) {
	num := userPost.GetThumbs()
	if num != -2 {
		t.Errorf("Expected -2, but got %d", num)
	}

	num = projectPost.GetThumbs()
	if num != 1 {
		t.Errorf("Expected 1, but got %d", num)
	}
}

func TestGetBookmarkers(t *testing.T) {
	users := userPost.GetBookmarkers()
	if len(users) != 1 {
		t.Errorf("Expected only 1 users, but got: %d", len(users))
	}

	n := userPost.GetBookmarkersNumber()
	if 1 != n {
		t.Errorf("getBookmarkersNumber retured %d instead of 1", n)
	}

	if users[0].Username != "admin" {
		t.Errorf("Post shoud be bookmarked by 'admin', but got: %v", users[0].Username)
	}

	users = projectPost.GetBookmarkers()
	if len(users) != 1 {
		t.Errorf("Expected only 1 users, but got: %d", len(users))
	}

	n = projectPost.GetBookmarkersNumber()

	if 1 != n {
		t.Errorf("getBookmarkersNumber retured %d instead of 1", n)
	}

	if users[0].Username != "admin" {
		t.Errorf("Post shoud be bookmarked by 'admin', but got: %v", users[0].Username)
	}
}

func TestGetLurkers(t *testing.T) {
	users := userPost1.GetLurkers()

	if len(users) != 1 {
		t.Errorf("Expected only 1 users, but got: %d", len(users))
	}

	n := userPost1.GetLurkersNumber()

	if 1 != n {
		t.Error("getLurkersNumber retured %d instead of 1", n)
	}

	if users[0].Username != "admin" {
		t.Error("Post shoud be lurked by 'admin', but got: %v", users[0].Username)
	}

	users = projectPost.GetLurkers()
	if len(users) != 0 {
		t.Errorf("Expected 0 users, but got: %d", len(users))
	}

	n = projectPost.GetLurkersNumber()
	if 0 != n {
		t.Errorf("getLurkersNumber retured %d instead of 0", n)
	}
}

func TestGetURL(t *testing.T) {
	domain, _ := url.Parse("http://nerdzdoma.in")

	if projectPost.GetURL(domain).String() != "http://nerdzdoma.in/NERDZilla:1" {
		t.Errorf("GetURL returned %s instead of http://nerdzdoma.in/NERDZilla:1", projectPost.GetURL(domain).String())
	}

	if userPost.GetURL(domain).String() != "http://nerdzdoma.in/admin.5" {
		t.Errorf("GetURL returned %s insted of http://nerdzdoma.in/admin.5", userPost.GetURL(domain).String())
	}

}
