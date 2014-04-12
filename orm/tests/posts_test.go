package orm_test

import (
	"fmt"
	"github.com/nerdzeu/nerdz-api/orm"
	"testing"
)

var userPost orm.Post
var projectPost orm.ProjectPost

func init() {
	if err := projectPost.New(3); err != nil {
		panic(fmt.Sprintf("No error should happen when create existing post, but got: %+v", err))
	}

	if err := userPost.New(2); err != nil {
		panic(fmt.Sprintf("No error should happen when create existing post, but got: %+v", err))
	}

	fmt.Printf("** PRJ POST **\n%+v\n**USER POST **\n%+v\n", projectPost, userPost)
}

func TestGetFrom(t *testing.T) {
	from, err := userPost.GetFrom()
	if err != nil {
		panic(fmt.Sprintf("No error should happen when fetching existing user, but got: %+v", err))
	}

	if from.Counter != 1 {
		panic(fmt.Sprintf("Counter should be 1, but go: %d", from.Counter))
	}

	fromPrj, err := projectPost.GetFrom()
	if err != nil {
		panic(fmt.Sprintf("No error should happen when fetching existing user, but got: %+v", err))
	}

	if fromPrj.Counter != 1 {
		panic(fmt.Sprintf("Counter should be 1, but go: %d", from.Counter))
	}

	fmt.Printf("%+v\n", fromPrj)
}

func TestGetTo(t *testing.T) {
	to, err := userPost.GetTo()
	if err != nil {
		panic(fmt.Sprintf("No error should happen when fetching existing user, but got: %+v", err))
	}

	if to.Counter != 1 {
		panic(fmt.Sprintf("Counter should be 1, but go: %d", to.Counter))
	}

	toPrj, err := projectPost.GetTo()
	if err != nil {
		panic(fmt.Sprintf("No error should happen when fetching existing user, but got: %+v", err))
	}

	if toPrj.Counter != 1 {
		panic(fmt.Sprintf("Counter should be 1, but go: %d", toPrj.Counter))
	}

	fmt.Printf("%+v\n", toPrj)
}
