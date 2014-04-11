package orm_test

import (
	"fmt"
	"github.com/nerdzeu/nerdz-api/orm"
	"testing"
)

var userPost orm.UserPost
var projectPost orm.ProjectPost

func TestNew(t *testing.T) {
	if err := projectPost.New(3); err != nil {
		panic(fmt.Sprintf("No error should happen when create existing post, but got: %+v", err))
	}

	if err := userPost.New(2); err != nil {
		panic(fmt.Sprintf("No error should happen when create existing post, but got: %+v", err))
	}

	fmt.Printf("** PRJ POST **\n%+v\n**USER POST **\n%+v\n", projectPost, userPost)
}
