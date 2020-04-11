/*
Copyright (C) 2016-2020 Paolo Galeone <nessuno@nerdz.eu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package nerdz_test

import (
	"fmt"
	"testing"

	"github.com/nerdzeu/nerdz-api/nerdz"
)

var prj *nerdz.Project
var err error

func init() {
	prj, err = nerdz.NewProject(1)
	if err != nil {
		panic(fmt.Sprintf("No error should happen when create existing user, but got: %+v", err))
	}
}

func TestProjectInfo(t *testing.T) {
	info := prj.ProjectInfo()
	if info == nil {
		t.Error("null info")
	}

	t.Logf("Struct: %+v\nMembers:", *info)
	for i, elem := range info.Members {
		t.Logf("%d) %+v\n", i, elem)
	}

	t.Log("Followers\n")
	for i, elem := range info.Followers {
		t.Logf("%d) %+v\n", i, elem)
	}

}

func TestProjectPostlist(t *testing.T) {
	postList := *prj.Postlist(nerdz.PostlistOptions{})
	if len(postList) != 4 {
		t.Fatalf("Expected 4  posts, but got: %+v\n", len(postList))
	}

	t.Logf("%+v\n", postList)
}
