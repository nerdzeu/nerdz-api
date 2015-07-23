package nerdz

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/nerdzeu/nerdz-api/utils"
)

// NewProjectPost initializes a ProjectPost struct
func NewProjectPost(hpid uint64) (post *ProjectPost, e error) {
	post = new(ProjectPost)
	Db().First(post, hpid)
	if post.Hpid != hpid {
		return nil, errors.New("Invalid hpid")
	}

	return post, nil
}

// Implementing NewMessage interface

// SetSender set the source of the post (the user ID)
func (post *ProjectPost) SetSender(id uint64) {
	post.From = id
}

// SetReference set the destionation of the post. Project ID
func (post *ProjectPost) SetReference(id uint64) {
	post.To = id
}

// SetText set the text of the message
func (post *ProjectPost) SetText(message string) {
	post.Message = message
}

// ClearDefaults set to the go's default values the fields with default sql values
func (post *ProjectPost) ClearDefaults() {
	post.Time = time.Time{}
	post.Pid = 0
}

// Implementing existingPost interface

// ID returns the Project Post ID
func (post *ProjectPost) ID() uint64 {
	return post.Hpid
}

// NumericSender returns the id of the sender user
func (post *ProjectPost) NumericSender() uint64 {
	return post.From
}

// Sender returns the sender *User
func (post *ProjectPost) Sender() *User {
	user, _ := NewUser(post.NumericSender())
	return user
}

// NumericReference returns the id of the recipient project
func (post *ProjectPost) NumericReference() uint64 {
	return post.To
}

// Reference returns the recipient *Project
func (post *ProjectPost) Reference() Reference {
	project, _ := NewProject(post.NumericReference())
	return project
}

// Text returns the post message
func (post *ProjectPost) Text() string {
	return post.Message
}

// IsEditable returns true if the ProjectPost is editable
func (post *ProjectPost) IsEditable() bool {
	return true
}

// IsClosed reuturns ture if the post is closed
func (post *ProjectPost) IsClosed() bool {
	return post.Closed
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

// SetLanguage set the language of the post
func (post *ProjectPost) SetLanguage(language string) error {
	if utils.InSlice(language, Configuration.Languages) {
		post.Lang = language
		return nil
	}
	return fmt.Errorf("Language '%s' is not a valid or supported language", language)
}

// Language returns the message language
func (post *ProjectPost) Language() string {
	return post.Lang
}

// Revisions returns all the revisions of the message
func (post *ProjectPost) Revisions() (modifications []string) {
	Db().Model(ProjectPostRevision{}).Where(&ProjectPostRevision{Hpid: post.Hpid}).Pluck("message", &modifications)
	return
}

// RevisionsNumber returns the number of the revisions
func (post *ProjectPost) RevisionsNumber() (count uint8) {
	Db().Model(ProjectPostRevision{}).Where(&ProjectPostRevision{Hpid: post.Hpid}).Count(&count)
	return
}

// Thumbs returns the post's thumbs value
func (post *ProjectPost) Thumbs() int {
	var sum struct {
		Total int
	}

	Db().Model(ProjectPostThumb{}).Select("COALESCE(sum(vote), 0) as total").Where(&ProjectPostThumb{Hpid: post.Hpid}).Scan(&sum)
	return sum.Total
}

// setApiFields populate the API fileds of the UserPost
func (post *ProjectPost) setApiFields(user *User) {
	if from, e := NewUser(post.From); e == nil {
		post.FromInfo = from.Info()
	}
	if to, e := NewProject(post.To); e == nil {
		post.ToInfo = to.Info()
	}
	post.Rate = post.Thumbs()
	post.RevisionsCount = post.RevisionsNumber()
	post.CommentsCount = post.CommentsNumber()
	post.BookmarkersCount = post.BookmarkersNumber()
	post.LurkersCount = post.LurkersNumber()
	post.Timestamp = post.Time.Unix()
	post.Url = post.URL(Configuration.NERDZURL).String()
	post.CanBookmark = user.canBookmark(post)
	post.CanComment = user.canComment(post)
	post.CanDelete = user.canDelete(post)
	post.CanEdit = user.canEdit(post)
	post.CanLurk = user.canLurk(post)
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
		Db().Find(&comments, &ProjectPostComment{Hpid: post.Hpid})

	case 1: // Get last interval[0] comments [ LIMIT interval[0] ]
		Db().Order("hcid DESC").Limit(interval[0]).Find(&comments, &ProjectPostComment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]ProjectPostComment)

	case 2: // Get last interval[0] comments, starting from interval[1] [ LIMIT interval[0] OFFSET interval[1] ]
		Db().Order("hcid DESC").Limit(interval[0]).Offset(interval[1]).Find(&comments, &ProjectPostComment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]ProjectPostComment)
	}

	return comments
}

// CommentsNumber returns the number of comment's post
func (post *ProjectPost) CommentsNumber() (count uint8) {
	Db().Model(ProjectPostComment{}).Where(&ProjectPostComment{Hpid: post.Hpid}).Count(&count)
	return
}

// NumericBookmarkers returns a slice of uint64 representing the ids of the users that bookmarked the post
func (post *ProjectPost) NumericBookmarkers() (bookmarkers []uint64) {
	Db().Model(ProjectPostBookmark{}).Where(&ProjectPostBookmark{Hpid: post.Hpid}).Pluck("\"from\"", &bookmarkers)
	return
}

// Bookmarkers returns a slice of users that bookmarked the post
func (post *ProjectPost) Bookmarkers() []*User {
	return Users(post.NumericBookmarkers())
}

// BookmarkersNumber returns the number of users that bookmarked the post
func (post *ProjectPost) BookmarkersNumber() (count uint8) {
	Db().Model(ProjectPostBookmark{}).Where(&ProjectPostBookmark{Hpid: post.Hpid}).Count(&count)
	return
}

// NumericLurkers returns a slice of uint64 representing the ids of the users that lurked the post
func (post *ProjectPost) NumericLurkers() (lurkers []uint64) {
	Db().Model(ProjectPostLurker{}).Where(&ProjectPostLurker{Hpid: post.Hpid}).Pluck("\"from\"", &lurkers)
	return
}

// Lurkers returns a slice of users that are lurking the post
func (post *ProjectPost) Lurkers() []*User {
	return Users(post.NumericLurkers())
}

// LurkersNumber returns the number of users that are lurking the post
func (post *ProjectPost) LurkersNumber() (count uint8) {
	Db().Model(ProjectPostLurker{}).Where(&ProjectPostLurker{Hpid: post.Hpid}).Count(&count)
	return
}

// URL returns the url of the posts, appended to the domain url passed es paremeter.
// Example: post.URL(url.URL{Scheme: "http", Host: "mobile.nerdz.eu"}) returns
// http://mobile.nerdz.eu/ + post.Reference().Name + ":"post.Pid
// If the post is on the board of the "admin" project and has a pid = 44, returns
// http://mobile.nerdz.eu/admin:44
func (post *ProjectPost) URL(domain *url.URL) *url.URL {
	domain.Path = (post.Reference().(*Project)).Name + ":" + strconv.FormatUint(post.Pid, 10)
	return domain
}
