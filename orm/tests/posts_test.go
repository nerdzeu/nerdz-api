package orm_test

import (
	"fmt"
	"github.com/nerdzeu/nerdz-api/orm"
	"testing"
)

var userPost *orm.UserPost
var projectPost *orm.ProjectPost
var e error

func init() {
	projectPost, e = orm.NewProjectPost(3)
	if e != nil {
		panic(fmt.Sprintf("No error should happen when create existing post, but got: %+v", e))
	}

	userPost, e = orm.NewUserPost(2)

	if e != nil {
		panic(fmt.Sprintf("No error should happen when create existing post, but got: %+v", e))
	}

	fmt.Printf("** PRJ POST **\n%+v\n**USER POST **\n%+v\n", projectPost, userPost)
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

	if fromPrj.Counter != 1 {
		t.Errorf("Counter should be 1, but go: %d", from.Counter)
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

	if toPrj.Counter != 1 {
		t.Errorf("Counter should be 1, but go: %d", toPrj.Counter)
	}

	fmt.Printf("%+v\n", toPrj)
}

func TestGetComments(t *testing.T) {
	comments := userPost.GetComments()
	if len(comments) == 0 {
		t.Error("No comments found. Expected > 1")
	}

	comments = userPost.GetComments(4)
	if len(comments) != 4 {
		t.Errorf("Expected the last 4 comments, got: %d", len(comments))
	}
	fmt.Printf("%+v\n", comments)

	comment := userPost.GetComments(4, 11)
	if len(comment) != 1 {
		t.Errorf("Expected 1 comment, received: %d", len(comment))
	}
	fmt.Printf("%+v\n", comment)

	prjComments := projectPost.GetComments()
	if len(prjComments) == 0 {
		t.Error("No comments found. Expected > 1")
	}

	prjComments = projectPost.GetComments(4)
	if len(prjComments) != 4 {
		t.Errorf("Expected the last 4 comments, got: %d", len(prjComments))
	}
	fmt.Printf("%+v\n", prjComments)

	prjComment := projectPost.GetComments(4, 4)
	if len(prjComment) != 1 {
		t.Errorf("Expected 1 comment, received: %d", len(prjComment))
	}
	fmt.Printf("%+v\n", prjComment)
}

func TestGetThumbs(t *testing.T) {
	num := userPost.GetThumbs()
	if num != 1 {
		t.Errorf("Expected 1, but got %d", num)
	}

	num = projectPost.GetThumbs()
	if num != 2 {
		t.Errorf("Expected 2, but got %d", num)
	}
}
