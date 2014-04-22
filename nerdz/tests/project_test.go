package nerdz_test

import (
	"fmt"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"testing"
)

var prj *nerdz.Project
var err error

func init() {
	prj, err = nerdz.NewProject(1)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when create existing user, but got: %+v", err))
	}
}

func TestGetProjectInfo(t *testing.T) {
	info := prj.GetProjectInfo()
	if info == nil {
		t.Error("null info")
	}

	fmt.Printf("Struct: %+v\nMembers:", *info)
	for i, elem := range info.Members {
		fmt.Printf("%d) %+v\n", i, elem)
	}

	fmt.Println("Followers")
	for i, elem := range info.Followers {
		fmt.Printf("%d) %+v\n", i, elem)
	}

}

func TestGetPostlist(t *testing.T) {
	postList := prj.GetPostlist(nil).([]nerdz.ProjectPost)
	if len(postList) != 20 {
		t.Error("Expected 20  posts, but got: %+v\n", len(postList))
	}

	fmt.Printf("%+v\n", postList)
}
