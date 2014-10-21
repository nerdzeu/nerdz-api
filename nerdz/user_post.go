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

// New initializes a UserPost struct
func NewUserPost(hpid uint64) (post *UserPost, e error) {
	post = new(UserPost)
	db.First(post, hpid)

	if post.Hpid != hpid {
		return nil, errors.New("Invalid hpid")
	}

	return post, nil
}

// Implementing Message interface

// To returns the recipient *User
func (post *UserPost) Recipient() (Board, error) {
	return NewUser(post.To)
}

// From returns the sender *User
func (post *UserPost) Sender() (*User, error) {
	return NewUser(post.From)
}

// Thumbs returns the post's thumbs value
func (post *UserPost) Thumbs() int {
	type result struct {
		Total int
	}
	var sum result
	db.Model(UserPostThumb{}).Select("COALESCE(sum(vote), 0) as total").Where(&UserPostThumb{Hpid: post.Hpid}).Scan(&sum)
	return sum.Total
}

// Message returns the post message
func (post *UserPost) Text() string {
	return post.Message
}

// Implementing ExistingPost interface

// Comments returns the full comments list, or the selected range of comments
// Comments()  returns the full comments list
// Comments(N) returns at most the last N comments
// Comments(N, X) returns at most N comments, before the last comment + X
func (post *UserPost) Comments(interval ...uint) interface{} {
	var comments []UserPostComment

	switch len(interval) {
	default: //full list
	case 0:
		db.Find(&comments, &UserPostComment{Hpid: post.Hpid})

	case 1: // Get last interval[0] comments [ LIMIT interval[0] ]
		db.Order("hcid DESC").Limit(interval[0]).Find(&comments, &UserPostComment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]UserPostComment)

	case 2: // Get last interval[0] comments, starting from interval[1] [ LIMIT interval[0] OFFSET interval[1] ]
		db.Order("hcid DESC").Limit(interval[0]).Offset(interval[1]).Find(&comments, &UserPostComment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]UserPostComment)
	}

	return comments
}

// Bookmarkers returns a slice of users that bookmarked the post
func (post *UserPost) Bookmarkers() []*User {
	return Users(post.NumericBookmarkers())
}

// BookmarkersNumber returns the number of users that bookmarked the post
func (post *UserPost) BookmarkersNumber() (count uint) {
	db.Model(UserBookmark{}).Where(&UserBookmark{Hpid: post.Hpid}).Count(&count)
	return
}

// Lurkers returns a slice of users that are lurking the post
func (post *UserPost) Lurkers() []*User {
	return Users(post.NumericLurkers())
}

// LurkersNumber returns the number of users that are lurking the post
func (post *UserPost) LurkersNumber() (count uint) {
	db.Model(UserPostLurker{}).Where(&UserPostLurker{Hpid: post.Hpid}).Count(&count)
	return
}

// URL returns the url of the posts, appended to the domain url passed es paremeter.
// Example: post.URL(url.URL{Scheme: "http", Host: "mobile.nerdz.eu"}) returns
// http://mobile.nerdz.eu/ + post.Recipient().Username + "."post.Pid
// If the post is on the board of the "admin" user and has a pid = 44, returns
// http://mobile.nerdz.eu/admin.44
func (post *UserPost) URL(domain *url.URL) *url.URL {
	to, _ := post.Recipient()
	domain.Path = (to.(*User)).Username + "." + strconv.FormatUint(post.Pid, 10)
	return domain
}

// Implementing NewPost interface

// Set the destionation of the post. user can be a user's id or a *User
func (post *UserPost) SetRecipient(user interface{}) error {
	switch user.(type) {
	case uint64:
		post.To = user.(uint64)
	case *User:
		post.To = (user.(*User)).Counter
	default:
		return fmt.Errorf("Invalid user type: %v. Allowed uint64 and *User", reflect.TypeOf(user))
	}
	return nil
}

// SetMessage set NewPost message and escape html entities. Returns nil on success, error on failure
func (post *UserPost) SetText(message string) error {
	if len(message) == 0 {
		return errors.New("Empty message")
	}

	post.Message = html.EscapeString(message)
	return nil
}
