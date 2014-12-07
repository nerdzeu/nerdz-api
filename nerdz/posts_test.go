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
	if projectPost, e = nerdz.NewProjectPost(uint64(3)); e != nil {
		panic(fmt.Sprintf("No error should happen when create existing post, but got: %+v", e))
	}

	if userPost, e = nerdz.NewUserPost(6); e != nil {
		panic(fmt.Sprintf("No error should happen when create existing post, but got: %+v", e))
	}

	userPost1, _ = nerdz.NewUserPost(20)
}

func TestFrom(t *testing.T) {
	from := userPost.Sender()

	if from.Counter != 1 {
		t.Errorf("Counter should be 1, but go: %d", from.Counter)
	}

	fromPrj := projectPost.Sender()

	if fromPrj.Counter != 4 {
		t.Errorf("Counter should be 4, but go: %d", fromPrj.Counter)
	}

	t.Logf("%+v\n", fromPrj)
}

func TestTo(t *testing.T) {
	to := userPost.Reference()

	user := to.(*nerdz.User)

	if user.Counter != 1 {
		t.Errorf("Counter should be 1, but go: %d", user.Counter)
	}

	to = projectPost.Reference()

	project := to.(*nerdz.Project)

	if project.Counter != 3 {
		t.Errorf("Counter should be 3, but go: %d", project.Counter)
	}

	t.Logf("%+v\n", project)
}

func TestComments(t *testing.T) {
	comments := userPost.Comments().([]nerdz.UserPostComment)
	if len(comments) == 0 {
		t.Error("No comments found. Expected > 1")
	}

	comments = userPost.Comments(4).([]nerdz.UserPostComment)
	if len(comments) != 4 {
		t.Errorf("Expected the last 4 comments, got: %d", len(comments))
	}
	t.Logf("%+v\n", comments)

	comment := userPost.Comments(4, 5).([]nerdz.UserPostComment)
	if len(comment) != 2 {
		t.Errorf("Expected 2 comments, received: %d", len(comment))
	}
	t.Logf("%+v\n", comment)

	prjComments := projectPost.Comments().([]nerdz.ProjectPostComment)
	if len(prjComments) == 0 {
		t.Error("No comments found. Expected > 1")
	}

	prjComments = projectPost.Comments(4).([]nerdz.ProjectPostComment)
	if len(prjComments) != 1 {
		t.Errorf("Expected the last  comment, got: %d", len(prjComments))
	}
	t.Logf("%+v\n", prjComments)

	prjComment := projectPost.Comments(4, 4).([]nerdz.ProjectPostComment)
	if len(prjComment) != 0 {
		t.Errorf("Expected no comment, received: %d", len(prjComment))
	}
	t.Logf("%+v\n", prjComment)
}

func TestThumbs(t *testing.T) {
	num := userPost.Thumbs()
	if num != -2 {
		t.Errorf("Expected -2, but got %d", num)
	}

	num = projectPost.Thumbs()
	if num != 1 {
		t.Errorf("Expected 1, but got %d", num)
	}
}

func TestBookmarkers(t *testing.T) {
	users := userPost.Bookmarkers()
	if len(users) != 1 {
		t.Errorf("Expected only 1 users, but got: %d", len(users))
	}

	n := userPost.BookmarkersNumber()
	if 1 != n {
		t.Errorf("BookmarkersNumber retured %d instead of 1", n)
	}

	if users[0].Username != "admin" {
		t.Errorf("Post shoud be bookmarked by 'admin', but got: %v", users[0].Username)
	}

	users = projectPost.Bookmarkers()
	if len(users) != 1 {
		t.Errorf("Expected only 1 users, but got: %d", len(users))
	}

	n = projectPost.BookmarkersNumber()

	if 1 != n {
		t.Errorf("BookmarkersNumber retured %d instead of 1", n)
	}

	if users[0].Username != "admin" {
		t.Errorf("Post shoud be bookmarked by 'admin', but got: %v", users[0].Username)
	}
}

func TestLurkers(t *testing.T) {
	users := userPost1.Lurkers()

	if len(users) != 1 {
		t.Errorf("Expected only 1 users, but got: %d", len(users))
	}

	n := userPost1.LurkersNumber()

	if 1 != n {
		t.Error("LurkersNumber retured %d instead of 1", n)
	}

	if users[0].Username != "admin" {
		t.Error("Post shoud be lurked by 'admin', but got: %v", users[0].Username)
	}

	users = projectPost.Lurkers()
	if len(users) != 0 {
		t.Errorf("Expected 0 users, but got: %d", len(users))
	}

	n = projectPost.LurkersNumber()
	if 0 != n {
		t.Errorf("LurkersNumber retured %d instead of 0", n)
	}
}

func TestURL(t *testing.T) {
	domain, _ := url.Parse("http://nerdzdoma.in")

	if projectPost.URL(domain).String() != "http://nerdzdoma.in/NERDZilla:1" {
		t.Errorf("URL returned %s instead of http://nerdzdoma.in/NERDZilla:1", projectPost.URL(domain).String())
	}

	if userPost.URL(domain).String() != "http://nerdzdoma.in/admin.5" {
		t.Errorf("URL returned %s insted of http://nerdzdoma.in/admin.5", userPost.URL(domain).String())
	}

}
