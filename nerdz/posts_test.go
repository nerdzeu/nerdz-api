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
	"strings"
	"testing"

	"github.com/nerdzeu/nerdz-api/nerdz"
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
		t.Fatalf("Counter should be 1, but go: %d", from.Counter)
	}

	fromPrj := projectPost.Sender()

	if fromPrj.Counter != 4 {
		t.Fatalf("Counter should be 4, but go: %d", fromPrj.Counter)
	}

	t.Logf("%+v\n", fromPrj)
}

func TestTo(t *testing.T) {
	to := userPost.Reference()

	user := to.(*nerdz.User)

	if user.Counter != 1 {
		t.Fatalf("Counter should be 1, but go: %d", user.Counter)
	}

	to = projectPost.Reference()

	project := to.(*nerdz.Project)

	if project.Counter != 3 {
		t.Fatalf("Counter should be 3, but go: %d", project.Counter)
	}

	t.Logf("%+v\n", project)
}

func TestComments(t *testing.T) {
	comments := *userPost.Comments(nerdz.CommentlistOptions{})
	if len(comments) == 0 {
		t.Error("No comments found. Expected > 1")
	}

	comments = *userPost.Comments(nerdz.CommentlistOptions{N: 4})
	if len(comments) != 4 {
		t.Fatalf("Expected the last 4 comments, got: %d", len(comments))
	}

	comments = *userPost.Comments(nerdz.CommentlistOptions{
		// Comments are fetched in inversed temporal order
		Older: comments[0].ID(),
		Newer: comments[3].ID() - 1,
	})
	if len(comments) != 3 {
		t.Fatalf("Expected 3 comments, received: %d", len(comments))
	}
	t.Logf("%+v\n", comments)

	prjComments := *projectPost.Comments(nerdz.CommentlistOptions{})
	if len(prjComments) == 0 {
		t.Error("No comments found. Expected > 1")
	}

	prjComments = *projectPost.Comments(nerdz.CommentlistOptions{N: 4})
	if len(prjComments) != 1 {
		t.Fatalf("Expected the last  comment, got: %d", len(prjComments))
	}
	t.Logf("%+v\n", prjComments)

	prjComments = *projectPost.Comments(nerdz.CommentlistOptions{Newer: 100})
	if len(prjComments) != 0 {
		t.Fatalf("Expected no comment, received: %d", len(prjComments))
	}
	t.Logf("%+v\n", prjComments)
}

func TestVotes(t *testing.T) {
	num := userPost.VotesCount()
	if num != -2 {
		t.Fatalf("Expected -2, but got %d", num)
	}

	num = projectPost.VotesCount()
	if num != 1 {
		t.Fatalf("Expected 1, but got %d", num)
	}
}

func TestBookmarks(t *testing.T) {
	users := userPost.Bookmarkers()
	if len(users) != 1 {
		t.Fatalf("Expected only 1 users, but got: %d", len(users))
	}

	n := userPost.BookmarksCount()
	if 1 != n {
		t.Fatalf("BookmarksCount returned %d instead of 1", n)
	}

	if users[0].Username != "admin" {
		t.Fatalf("Post shoud be bookmarked by 'admin', but got: %v", users[0].Username)
	}

	users = projectPost.Bookmarkers()
	if len(users) != 1 {
		t.Fatalf("Expected only 1 users, but got: %d", len(users))
	}

	n = projectPost.BookmarksCount()

	if 1 != n {
		t.Fatalf("BookmarksCount returned %d instead of 1", n)
	}

	if users[0].Username != "admin" {
		t.Fatalf("Post shoud be bookmarked by 'admin', but got: %v", users[0].Username)
	}
}

func TestLurkers(t *testing.T) {
	users := userPost1.Lurkers()

	if len(users) != 1 {
		t.Fatalf("Expected only 1 users, but got: %d", len(users))
	}

	n := userPost1.LurkersCount()

	if 1 != n {
		t.Fatalf("LurkersCount returned %d instead of 1", n)
	}

	if users[0].Username != "admin" {
		t.Fatalf("Post shoud be lurked by 'admin', but got: %v", users[0].Username)
	}

	users = projectPost.Lurkers()
	if len(users) != 0 {
		t.Fatalf("Expected 0 users, but got: %d", len(users))
	}

	n = projectPost.LurkersCount()
	if 0 != n {
		t.Fatalf("LurkersCount returned %d instead of 0", n)
	}
}

func TestURL(t *testing.T) {
	if !strings.HasSuffix(projectPost.URL().String(), "/NERDZilla:1") {
		t.Fatalf("URL returned %s instead of Configuration.NERDZHost/NERDZilla:1", projectPost.URL().String())
	}

	if !strings.HasSuffix(userPost.URL().String(), "/admin.5") {
		t.Fatalf("URL returned %s insted of Configuration.NERDZHost/admin.5", userPost.URL().String())
	}
}
