package orm_test

import (
	"fmt"
	"github.com/nerdzeu/nerdz-api/orm"
	"testing"
)

var prj *orm.Project
var err error
func init() {
    prj, err = orm.NewProject(1)
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
