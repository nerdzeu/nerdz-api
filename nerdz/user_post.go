package nerdz

import (
	"errors"
	"github.com/nerdzeu/nerdz-api/utils"
	"html"
	"net/url"
	"strconv"
)

// New initializes a UserPost struct
func NewUserPost(hpid int64) (post *UserPost, e error) {
	post = new(UserPost)
	db.First(post, hpid)

	if post.Hpid != hpid {
		return nil, errors.New("Invalid hpid")
	}

	return post, nil
}

// Implementing ExistingPost interface

// GetTo returns the recipient *User
func (post *UserPost) GetTo() (Board, error) {
	return NewUser(post.To)
}

// GetFrom returns the sender *User
func (post *UserPost) GetFrom() (*User, error) {
	return NewUser(post.From)
}

// GetThumbs returns the post's thumbs value
func (post *UserPost) GetThumbs() int {
	type result struct {
		Total int
	}
	var sum result
	db.Model(UserPostThumb{}).Select("COALESCE(sum(vote), 0) as total").Where(&UserPostThumb{Hpid: post.Hpid}).Scan(&sum)
	return sum.Total
}

// GetComments returns the full comments list, or the selected range of comments
// GetComments()  returns the full comments list
// GetComments(N) returns at most the last N comments
// GetComments(N, X) returns at most N comments, before the last comment + X
func (post *UserPost) GetComments(interval ...int) interface{} {
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

// GetBookmarkers returns a slice of users that bookmarked the post
func (post *UserPost) GetBookmarkers() []*User {
	return getUsers(post.getNumericBookmarkers())
}

// GetBookmarkersNumber returns the number of users that bookmarked the post
func (post *UserPost) GetBookmarkersNumber() int {
	var count int
	db.Model(UserBookmark{}).Where(&UserBookmark{Hpid: post.Hpid}).Count(&count)
	return count
}

// GetLurkers returns a slice of users that are lurking the post
func (post *UserPost) GetLurkers() []*User {
	return getUsers(post.getNumericLurkers())
}

// GetLurkersNumber returns the number of users that are lurking the post
func (post *UserPost) GetLurkersNumber() int {
	var count int
	db.Model(UserPostLurker{}).Where(&UserPostLurker{Hpid: post.Hpid}).Count(&count)
	return count
}

// GetURL returns the url of the posts, appended to the domain url passed es paremeter.
// Example: post.GetURL(url.URL{Scheme: "http", Host: "mobile.nerdz.eu"}) returns
// http://mobile.nerdz.eu/ + post.GetTo().Username + "."post.Pid
// If the post is on the board of the "admin" user and has a pid = 44, returns
// http://mobile.nerdz.eu/admin.44
func (post *UserPost) GetURL(domain *url.URL) *url.URL {
	to, _ := post.GetTo()
	domain.Path = (to.(*User)).Username + "." + strconv.FormatInt(post.Pid, 10)
	return domain
}

// GetMessage returns the post message
func (post *UserPost) GetMessage() string {
	return post.Message
}

// Implementing NewPost interface

// Set the destionation of the post. user can be a user's id or a *User
func (post *UserPost) SetTo(user interface{}) error {
	switch user.(type) {
	case int:
		post.To = int64(user.(int))
	case *User:
		post.To = (user.(*User)).Counter
	default:
		return errors.New("Invalid user type. Allowed int and *User")
	}
	return nil
}

// SetMessage set NewPost message and escape html entities. Returns nil on success, error on failure
func (post *UserPost) SetMessage(message string) error {
	if len(message) == 0 {
		return errors.New("Empty message")
	}

	post.Message = html.EscapeString(message)
	return nil
}
