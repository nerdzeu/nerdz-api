package nerdz

import (
	"errors"
	"fmt"
	"github.com/nerdzeu/nerdz-api/utils"
	"html"
	"net/url"
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

// Implementing NewMessage interface

// Set the source of the post (the user ID)
func (post *ProjectPost) SetSender(id uint64) {
	post.From = id
}

// Set the destionation of the post. Project ID
func (post *ProjectPost) SetRecipient(id uint64) {
	post.To = id
}

// SetMessage set NewPost message and escape html entities. Returns nil on success, error on failure
func (post *ProjectPost) SetText(message string) error {
	if len(message) == 0 {
		return errors.New("Empty message")
	}

	post.Message = html.EscapeString(message)

	return nil
}

// Implementing existingPost interface

// Id returns the Project Post ID
func (post *ProjectPost) Id() uint64 {
	return post.Hpid
}

// From returns the sender *User
func (post *ProjectPost) Sender() (*User, error) {
	return NewUser(post.From)
}

// To returns the recipient *Project
func (post *ProjectPost) Recipient() (Board, error) {
	return NewProject(post.To)
}

// Message returns the post message
func (post *ProjectPost) Text() string {
	return post.Message
}

// IsEditable returns true if the ProjectPost is editable
func (post *ProjectPost) IsEditable() bool {
	return true
}

// NumericOwners returns a slice of ids of the owner of the posts (the ones that can perform actions)
func (post *ProjectPost) NumericOwners() (ret []uint64) {
	ret = append(ret, post.From)
	project, _ := NewProject(post.To)
	ret = append(ret, project.NumericOwner())
	ret = append(ret, project.NumericMembers()...)
	return
}

// Owners returns a slice of *User representing the users who own the post

func (post *ProjectPost) Owners() (ret []*User) {
	return Users(post.NumericOwners())
}

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

// Thumbs returns the post's thumbs value
func (post *ProjectPost) Thumbs() int {
	var sum struct {
		Total int
	}

	db.Model(ProjectPostThumb{}).Select("COALESCE(sum(vote), 0) as total").Where(&ProjectPostThumb{Hpid: post.Hpid}).Scan(&sum)
	return sum.Total
}

// NumericBookmarks returns a slice of uint64 representing the ids of the users that bookmarked the post
func (post *ProjectPost) NumericBookmarkers() (bookmarkers []uint64) {
	db.Model(ProjectPostBookmark{}).Where(&ProjectPostBookmark{Hpid: post.Hpid}).Pluck("\"from\"", &bookmarkers)
	return
}

// Bookmarkers returns a slice of users that bookmarked the post
func (post *ProjectPost) Bookmarkers() []*User {
	return Users(post.NumericBookmarkers())
}

// BookmarkersNumber returns the number of users that bookmarked the post
func (post *ProjectPost) BookmarkersNumber() (count uint) {
	db.Model(ProjectPostBookmark{}).Where(&ProjectPostBookmark{Hpid: post.Hpid}).Count(&count)
	return
}

// NumericLurkers returns a slice of uint64 representing the ids of the users that lurked the post
func (post *ProjectPost) NumericLurkers() (lurkers []uint64) {
	db.Model(ProjectPostLurker{}).Where(&ProjectPostLurker{Hpid: post.Hpid}).Pluck("\"from\"", &lurkers)
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
