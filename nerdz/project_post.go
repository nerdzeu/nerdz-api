package nerdz

import (
	"errors"
	"fmt"
	"github.com/nerdzeu/nerdz-api/utils"
	"html"
	"net/url"
	"reflect"
	"strconv"
)

// NewProjectPost initializes a ProjectPost struct
func NewProjectPost(hpid uint64) (post *ProjectPost, e error) {
	post = new(ProjectPost)
	db.First(post, hpid)
	fmt.Println("parameter: %v\nResult: %v", hpid, post.Hpid)
	if post.Hpid != hpid {
		return nil, errors.New("Invalid hpid")
	}

	return post, nil
}

// Implementing Message interface

// From returns the sender *User
func (post *ProjectPost) Sender() (*User, error) {
	return NewUser(post.From)
}

// To returns the recipient *Project
func (post *ProjectPost) Recipient() (Board, error) {
	return NewProject(post.To)
}

// Thumbs returns the post's thumbs value
func (post *ProjectPost) Thumbs() int {
	var sum struct {
		Total int
	}

	db.Model(ProjectPostThumb{}).Select("COALESCE(sum(vote), 0) as total").Where(&ProjectPostThumb{Hpid: post.Hpid}).Scan(&sum)
	return sum.Total
}

// Message returns the post message
func (post *ProjectPost) Text() string {
	return post.Message
}

// Implementing ExistingPost interface

// Comments returns the full comments list, or the selected range of comments
// Comments()  returns the full comments list
// Comments(N) returns at most the last N comments
// Comments(N, X) returns at most N comments, before the last comment + X
func (post *ProjectPost) Comments(interval ...uint) interface{} {
	var comments []ProjectPostComment

	switch len(interval) {
	default: //full list
	case 0:
		db.Find(&comments, &ProjectPostComment{Hpid: post.Hpid})

	case 1: // Get last interval[0] comments [ LIMIT interval[0] ]
		db.Order("hcid DESC").Limit(interval[0]).Find(&comments, &ProjectPostComment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]ProjectPostComment)

	case 2: // Get last interval[0] comments, starting from interval[1] [ LIMIT interval[0] OFFSET interval[1] ]
		db.Order("hcid DESC").Limit(interval[0]).Offset(interval[1]).Find(&comments, &ProjectPostComment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]ProjectPostComment)
	}

	return comments
}

// Bookmarkers returns a slice of users that bookmarked the post
func (post *ProjectPost) Bookmarkers() []*User {
	return Users(post.NumericBookmarkers())
}

// BookmarkersNumber returns the number of users that bookmarked the post
func (post *ProjectPost) BookmarkersNumber() (count uint) {
	db.Model(ProjectBookmark{}).Where(&ProjectBookmark{Hpid: post.Hpid}).Count(&count)
	return
}

// Lurkers returns a slice of users that are lurking the post
func (post *ProjectPost) Lurkers() []*User {
	return Users(post.NumericLurkers())
}

// LurkersNumber returns the number of users that are lurking the post
func (post *ProjectPost) LurkersNumber() (count uint) {
	db.Model(ProjectPostLurker{}).Where(&ProjectPostLurker{Hpid: post.Hpid}).Count(&count)
	return
}

// URL returns the url of the posts, appended to the domain url passed es paremeter.
// Example: post.URL(url.URL{Scheme: "http", Host: "mobile.nerdz.eu"}) returns
// http://mobile.nerdz.eu/ + post.Recipient().Name + ":"post.Pid
// If the post is on the board of the "admin" project and has a pid = 44, returns
// http://mobile.nerdz.eu/admin:44
func (post *ProjectPost) URL(domain *url.URL) *url.URL {
	to, _ := post.Recipient()
	domain.Path = (to.(*Project)).Name + ":" + strconv.FormatUint(post.Pid, 10)
	return domain
}

// Implementing NewPost interface

// Set the destionation of the post. dest can be a project's id or a *Project.
func (post *ProjectPost) SetRecipient(project interface{}) error {
	switch project.(type) {
	case uint64:
		post.To = project.(uint64)
	case *Project:
		post.To = (project.(*Project)).Counter
	default:
		return fmt.Errorf("Invalid project type: %v. Allowed uint64 and *Project", reflect.TypeOf(project))
	}
	return nil
}

// SetMessage set NewPost message and escape html entities. Returns nil on success, error on failure
func (post *ProjectPost) SetText(message string) error {
	if len(message) == 0 {
		return errors.New("Empty message")
	}

	post.Message = html.EscapeString(message)

	return nil
}
