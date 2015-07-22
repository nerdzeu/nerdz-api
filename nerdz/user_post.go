package nerdz

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/nerdzeu/nerdz-api/utils"
)

// NewUserPost initializes a UserPost struct
func NewUserPost(hpid uint64) (post *UserPost, e error) {
	post = new(UserPost)
	Db().First(post, hpid)

	if post.Hpid != hpid {
		return nil, errors.New("Invalid hpid")
	}

	return post, nil
}

// Implementing NewPost interface

// SetSender sets the source of the post (the user ID)
func (post *UserPost) SetSender(id uint64) {
	post.From = id
}

// SetReference sets the destionation of the post: user ID
func (post *UserPost) SetReference(id uint64) {
	post.To = id
}

// SetText set the text of the message
func (post *UserPost) SetText(message string) {
	post.Message = message
}

// ClearDefaults set to the go's default values the fields with default sql values
func (post *UserPost) ClearDefaults() {
	post.Time = time.Time{}
	post.Pid = 0
}

// Implementing existingPost interface

// ID returns the User Post ID
func (post *UserPost) ID() uint64 {
	return post.Hpid
}

// NumericSender returns the id of the sender user
func (post *UserPost) NumericSender() uint64 {
	return post.From
}

// Sender returns the sender *User
func (post *UserPost) Sender() *User {
	user, _ := NewUser(post.NumericSender())
	return user
}

// NumericReference returns the id of the recipient user
func (post *UserPost) NumericReference() uint64 {
	return post.To
}

// Reference returns the recipient *User
func (post *UserPost) Reference() Reference {
	user, _ := NewUser(post.NumericReference())
	return user
}

// Text returns the post message
func (post *UserPost) Text() string {
	return post.Message
}

// IsEditable returns true if the post is editable
func (post *UserPost) IsEditable() bool {
	return true
}

// NumericOwners returns a slice of ids of the owner of the posts (the ones that can perform actions)
func (post *UserPost) NumericOwners() []uint64 {
	return []uint64{post.To, post.From}
}

// Owners returns a slice of *User representing the users who own the post
func (post *UserPost) Owners() (ret []*User) {
	return Users(post.NumericOwners())
}

// Thumbs returns the post's thumbs value
func (post *UserPost) Thumbs() int {
	type result struct {
		Total int
	}
	var sum result
	Db().Model(UserPostThumb{}).Select("COALESCE(sum(vote), 0) as total").Where(&UserPostThumb{Hpid: post.Hpid}).Scan(&sum)
	return sum.Total
}

// SetLanguage set the language of the post
func (post *UserPost) SetLanguage(language string) error {
	if utils.InSlice(language, Configuration.Languages) {
		post.Lang = language
		return nil
	}
	return fmt.Errorf("Language '%s' is not a valid or supported language", language)
}

// Language returns the message language
func (post *UserPost) Language() string {
	return post.Lang
}

// Revisions returns all the revisions of the message
func (post *UserPost) Revisions() (modifications []string) {
	Db().Model(UserPostRevision{}).Where(&UserPostRevision{Hpid: post.Hpid}).Pluck("message", &modifications)
	return
}

// RevisionsNumber returns the number of the revisions
func (post *UserPost) RevisionsNumber() (count uint8) {
	Db().Model(UserPostRevision{}).Where(&UserPostRevision{Hpid: post.Hpid}).Count(&count)
	return
}

// setApiFields populate the API fileds of the UserPost
func (post *UserPost) setApiFields() {
	if from, e := NewUser(post.From); e == nil {
		post.FromInfo = from.Info()
	}
	if to, e := NewUser(post.To); e == nil {
		post.ToInfo = to.Info()
	}
	post.OwnersInfo = Infos(post.Owners())
	post.Rate = post.Thumbs()
	post.RevisionsCount = post.RevisionsNumber()
	post.CommentsCount = post.CommentsNumber()
	post.BookmarkersCount = post.BookmarkersNumber()
	post.LurkersCount = post.LurkersNumber()
}

// Comments returns the full comments list, or the selected range of comments
// Comments()  returns the full comments list
// Comments(N) returns at most the last N comments
// Comments(N, X) returns at most N comments, before the last comment + X
func (post *UserPost) Comments(interval ...uint) interface{} {
	var comments []UserPostComment

	switch len(interval) {
	default: //full list
	case 0:
		Db().Find(&comments, &UserPostComment{Hpid: post.Hpid})

	case 1: // Get last interval[0] comments [ LIMIT interval[0] ]
		Db().Order("hcid DESC").Limit(interval[0]).Find(&comments, &UserPostComment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]UserPostComment)

	case 2: // Get last interval[0] comments, starting from interval[1] [ LIMIT interval[0] OFFSET interval[1] ]
		Db().Order("hcid DESC").Limit(interval[0]).Offset(interval[1]).Find(&comments, &UserPostComment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]UserPostComment)
	}

	return comments
}

// CommentsNumber returns the number of comment's post
func (post *UserPost) CommentsNumber() (count uint8) {
	Db().Model(UserPostComment{}).Where(&UserPostComment{Hpid: post.Hpid}).Count(&count)
	return
}

// NumericBookmarkers returns a slice of uint64 representing the ids of the users that bookmarked the post
func (post *UserPost) NumericBookmarkers() (bookmarkers []uint64) {
	Db().Model(UserPostBookmark{}).Where(&UserPostBookmark{Hpid: post.Hpid}).Pluck("\"from\"", &bookmarkers)
	return
}

// Bookmarkers returns a slice of users that bookmarked the post
func (post *UserPost) Bookmarkers() []*User {
	return Users(post.NumericBookmarkers())
}

// BookmarkersNumber returns the number of users that bookmarked the post
func (post *UserPost) BookmarkersNumber() (count uint8) {
	Db().Model(UserPostBookmark{}).Where(&UserPostBookmark{Hpid: post.Hpid}).Count(&count)
	return
}

// NumericLurkers returns a slice of uint64 representing the ids of the users that lurked the post
func (post *UserPost) NumericLurkers() (lurkers []uint64) {
	Db().Model(UserPostLurker{}).Where(&UserPostLurker{Hpid: post.Hpid}).Pluck("\"from\"", &lurkers)
	return
}

// Lurkers returns a slice of users that are lurking the post
func (post *UserPost) Lurkers() []*User {
	return Users(post.NumericLurkers())
}

// LurkersNumber returns the number of users that are lurking the post
func (post *UserPost) LurkersNumber() (count uint8) {
	Db().Model(UserPostLurker{}).Where(&UserPostLurker{Hpid: post.Hpid}).Count(&count)
	return
}

// URL returns the url of the posts, appended to the domain url passed es paremeter.
// Example: post.URL(url.URL{Scheme: "http", Host: "mobile.nerdz.eu"}) returns
// http://mobile.nerdz.eu/ + post.Reference().Username + "."post.Pid
// If the post is on the board of the "admin" user and has a pid = 44, returns
// http://mobile.nerdz.eu/admin.44
func (post *UserPost) URL(domain *url.URL) *url.URL {
	domain.Path = (post.Reference().(*User)).Username + "." + strconv.FormatUint(post.Pid, 10)
	return domain
}
