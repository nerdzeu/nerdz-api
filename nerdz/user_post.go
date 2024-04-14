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

package nerdz

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/nerdzeu/nerdz-api/utils"
)

// NewUserPost returns the *UserPost with id hpid if exists. Returns error otherwise
func NewUserPost(hpid uint64) (*UserPost, error) {
	return NewUserPostWhere(&UserPost{Post{Hpid: hpid}})
}

// NewUserPostWhere returns the *UserPost fetching the first one that matches the description
func NewUserPostWhere(description *UserPost) (post *UserPost, e error) {
	post = new(UserPost)
	if e = Db().Model(UserPost{}).Where(description).Scan(post); e != nil {
		return nil, e
	}
	if post.ID() == 0 {
		return nil, fmt.Errorf("requested UserPost does not exist")
	}
	return
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

// IsClosed reuturns ture if the post is closed
func (post *UserPost) IsClosed() bool {
	return post.Closed
}

// NumericOwners returns a slice of ids of the owner of the posts (the ones that can perform actions)
func (post *UserPost) NumericOwners() []uint64 {
	if post.To != post.From {
		return []uint64{post.To, post.From}
	}
	return []uint64{post.To}
}

// Owners returns a slice of *User representing the users who own the post
func (post *UserPost) Owners() (ret []*User) {
	return Users(post.NumericOwners())
}

// Votes returns the post's votes value
func (post *UserPost) VotesCount() (sum int) {
	_ = Db().Model(UserPostVote{}).Select("COALESCE(sum(vote), 0)").Where(&UserPostVote{Hpid: post.ID()}).Scan(&sum)
	return
}

// Votes returns a pointer to a slice of Vote
func (post *UserPost) Votes() *[]Vote {
	ret := []UserPostVote{}
	_ = Db().Model(UserPostVote{}).Where(&UserPostVote{Hpid: post.ID()}).Scan(&ret)
	var retVotes []Vote
	for _, v := range ret {
		vote := v
		retVotes = append(retVotes, Vote(&vote))
	}

	return &retVotes
}

// Bookmarks returns a pointer to a slice of Bookmark
func (post *UserPost) Bookmarks() *[]Bookmark {
	ret := []UserPostBookmark{}
	_ = Db().Model(UserPostBookmark{}).Where(&UserPostBookmark{Hpid: post.ID()}).Scan(&ret)
	var retBookmarks []Bookmark
	for _, b := range ret {
		bookmark := b
		retBookmarks = append(retBookmarks, Bookmark(&bookmark))
	}

	return &retBookmarks
}

// Lurks returns a pointer to a slice of Lurk
func (post *UserPost) Lurks() *[]Lurk {
	ret := []UserPostLurk{}
	_ = Db().Model(UserPostLurk{}).Where(&UserPostLurk{Hpid: post.ID()}).Scan(&ret)
	var retLurkers []Lurk
	for _, l := range ret {
		lurker := l
		retLurkers = append(retLurkers, Lurk(&lurker))
	}

	return &retLurkers
}

// Locks returns a pointer to a slice of Lock
func (post *UserPost) Locks() *[]Lock {
	ret := []UserPostLock{}
	_ = Db().Model(UserPostLock{}).Where(&UserPostLock{Hpid: post.ID()}).Scan(&ret)
	var retLockers []Lock
	for _, l := range ret {
		locker := l
		retLockers = append(retLockers, Lock(&locker))
	}

	return &retLockers
}

// SetLanguage set the language of the post
func (post *UserPost) SetLanguage(language string) error {
	if language == "" {
		language = post.Sender().Language()
	}
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
	_ = Db().Model(UserPostRevision{}).Where(&UserPostRevision{Hpid: post.ID()}).Pluck("message", &modifications)
	return
}

// RevisionsNumber returns the number of the revisions
func (post *UserPost) RevisionsNumber() (count uint8) {
	_ = Db().Model(UserPostRevision{}).Where(&UserPostRevision{Hpid: post.ID()}).Count(&count)
	return
}

// Comments returns the full comments list, or the selected range of comments
// Comments(options)  returns the comment list, using selected options
func (post *UserPost) Comments(options CommentlistOptions) *[]ExistingComment {
	var comments []UserPostComment

	query := Db().Model(UserPostComment{})
	query = commentlistQueryBuilder(query, options).Where(&UserPostComment{Hpid: post.ID()})
	if err := query.Scan(&comments); err != nil {
		log.Errorf("(UserPost::Comments) Error in query.Scan: %s", err)
	}

	comments = utils.ReverseSlice(comments).([]UserPostComment)

	var ret []ExistingComment
	for _, c := range comments {
		comment := c
		ret = append(ret, ExistingComment(&comment))
	}
	return &ret
}

// CommentsCount returns the number of comment's post
func (post *UserPost) CommentsCount() (count uint8) {
	_ = Db().Model(UserPostComment{}).Where(&UserPostComment{Hpid: post.ID()}).Count(&count)
	return
}

// NumericType returns the numeric type of the post
func (*UserPost) NumericType() uint8 {
	return 1
}

// Type returns a string representing the post type
func (*UserPost) Type() string {
	return "profile"
}

// NumericBookmarkers returns a slice of uint64 representing the ids of the users that bookmarked the post
func (post *UserPost) NumericBookmarkers() (bookmarkers []uint64) {
	_ = Db().Model(UserPostBookmark{}).Where(&UserPostBookmark{Hpid: post.ID()}).Pluck(`"from"`, &bookmarkers)
	return
}

// Bookmarks returns a slice of users that bookmarked the post
func (post *UserPost) Bookmarkers() []*User {
	return Users(post.NumericBookmarkers())
}

// BookmarksCount returns the number of users that bookmarked the post
func (post *UserPost) BookmarksCount() (count uint8) {
	_ = Db().Model(UserPostBookmark{}).Where(&UserPostBookmark{Hpid: post.ID()}).Count(&count)
	return
}

// NumericLurkers returns a slice of uint64 representing the ids of the users that lurked the post
func (post *UserPost) NumericLurkers() (lurkers []uint64) {
	_ = Db().Model(UserPostLurk{}).Where(&UserPostLurk{Hpid: post.ID()}).Pluck(`"from"`, &lurkers)
	return
}

// Lurkers returns a slice of users that are lurking the post
func (post *UserPost) Lurkers() []*User {
	return Users(post.NumericLurkers())
}

// LurkersCount returns the number of users that are lurking the post
func (post *UserPost) LurkersCount() (count uint8) {
	_ = Db().Model(UserPostLurk{}).Where(&UserPostLurk{Hpid: post.ID()}).Count(&count)
	return
}

// URL returns the url of the post
func (post *UserPost) URL() *url.URL {
	return &url.URL{
		Host: Configuration.NERDZHost,
		Path: (post.Reference().(*User)).Username + "." + strconv.FormatUint(post.Pid, 10),
	}
}
