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

// NewProjectPost initializes a ProjectPost struct
func NewProjectPost(hpid uint64) (*ProjectPost, error) {
	return NewProjectPostWhere(&ProjectPost{Post{Hpid: hpid}})
}

// NewProjectPostWhere returns the *ProjectPost fetching the first one that matches the description
func NewProjectPostWhere(description *ProjectPost) (post *ProjectPost, e error) {
	post = new(ProjectPost)
	if e = Db().Model(ProjectPost{}).Where(description).Scan(post); e != nil {
		return nil, e
	}
	if post.ID() == 0 {
		return nil, fmt.Errorf("requested ProjectPost does not exist")
	}
	return
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
func (post *ProjectPost) Language() string {
	return post.Lang
}

// Revisions returns all the revisions of the message
func (post *ProjectPost) Revisions() (modifications []string) {
	_ = Db().Model(ProjectPostRevision{}).Where(&ProjectPostRevision{Hpid: post.ID()}).Pluck("message", &modifications)
	return
}

// RevisionsNumber returns the number of the revisions
func (post *ProjectPost) RevisionsNumber() (count uint8) {
	_ = Db().Model(ProjectPostRevision{}).Where(&ProjectPostRevision{Hpid: post.ID()}).Count(&count)
	return
}

// Votes returns the post's votes value
func (post *ProjectPost) VotesCount() (sum int) {
	_ = Db().Model(ProjectPostVote{}).Select("COALESCE(sum(vote), 0)").Where(&ProjectPostVote{Hpid: post.ID()}).Scan(&sum)
	return
}

// Votes returns a pointer to a slice of Vote
func (post *ProjectPost) Votes() *[]Vote {
	ret := []ProjectPostVote{}
	_ = Db().Model(ProjectPostVote{}).Where(&ProjectPostVote{Hpid: post.ID()}).Scan(&ret)
	var retVotes []Vote
	for _, v := range ret {
		vote := v
		retVotes = append(retVotes, Vote(&vote))
	}

	return &retVotes
}

// Bookmarks returns a pointer to a slice of Bookmark
func (post *ProjectPost) Bookmarks() *[]Bookmark {
	ret := []ProjectPostBookmark{}
	_ = Db().Model(ProjectPostBookmark{}).Where(&ProjectPostBookmark{Hpid: post.ID()}).Scan(&ret)
	var retBookmarks []Bookmark
	for _, b := range ret {
		bookmark := b
		retBookmarks = append(retBookmarks, Bookmark(&bookmark))
	}

	return &retBookmarks
}

// Lurks returns a pointer to a slice of Lurk
func (post *ProjectPost) Lurks() *[]Lurk {
	ret := []ProjectPostLurk{}
	_ = Db().Model(ProjectPostLurk{}).Where(&ProjectPostLurk{Hpid: post.ID()}).Scan(&ret)
	var retLurkers []Lurk
	for _, l := range ret {
		lurker := l
		retLurkers = append(retLurkers, Lurk(&lurker))
	}

	return &retLurkers
}

// Locks returns a pointer to a slice of Lock
func (post *ProjectPost) Locks() *[]Lock {
	ret := []ProjectPostLock{}
	_ = Db().Model(ProjectPostLock{}).Where(&ProjectPostLock{Hpid: post.ID()}).Scan(&ret)
	var retLockers []Lock
	for _, l := range ret {
		locker := l
		retLockers = append(retLockers, Lock(&locker))
	}

	return &retLockers
}

// Comments returns the full comments list, or the selected range of comments
// Comments(options)  returns the comment list, using selected options
func (post *ProjectPost) Comments(options CommentlistOptions) *[]ExistingComment {
	var comments []ProjectPostComment

	query := Db().Model(ProjectPostComment{})
	query = commentlistQueryBuilder(query, options).Where(&ProjectPostComment{Hpid: post.ID()})
	if err := query.Scan(&comments); err != nil {
		log.Errorf("(Postlist) Error in query.Scan: %s", err)
	}

	comments = utils.ReverseSlice(comments).([]ProjectPostComment)

	var ret []ExistingComment
	for _, c := range comments {
		comment := c
		ret = append(ret, ExistingComment(&comment))
	}
	return &ret
}

// CommentsCount returns the number of comment's post
func (post *ProjectPost) CommentsCount() (count uint8) {
	_ = Db().Model(ProjectPostComment{}).Where(&ProjectPostComment{Hpid: post.ID()}).Count(&count)
	return
}

// NumericType returns the numeric type of the post
func (*ProjectPost) NumericType() uint8 {
	return 0
}

// Type returns a string representing the post type
func (*ProjectPost) Type() string {
	return "project"
}

// NumericBookmarkers returns a slice of uint64 representing the ids of the users that bookmarked the post
func (post *ProjectPost) NumericBookmarkers() (bookmarkers []uint64) {
	_ = Db().Model(ProjectPostBookmark{}).Where(&ProjectPostBookmark{Hpid: post.ID()}).Pluck(`"from"`, &bookmarkers)
	return
}

// Bookmarks returns a slice of users that bookmarked the post
func (post *ProjectPost) Bookmarkers() []*User {
	return Users(post.NumericBookmarkers())
}

// BookmarksCount returns the number of users that bookmarked the post
func (post *ProjectPost) BookmarksCount() (count uint8) {
	_ = Db().Model(ProjectPostBookmark{}).Where(&ProjectPostBookmark{Hpid: post.ID()}).Count(&count)
	return
}

// NumericLurkers returns a slice of uint64 representing the ids of the users that lurked the post
func (post *ProjectPost) NumericLurkers() (lurkers []uint64) {
	_ = Db().Model(ProjectPostLurk{}).Where(&ProjectPostLurk{Hpid: post.ID()}).Pluck(`"from"`, &lurkers)
	return
}

// Lurkers returns a slice of users that are lurking the post
func (post *ProjectPost) Lurkers() []*User {
	return Users(post.NumericLurkers())
}

// LurkersCount returns the number of users that are lurking the post
func (post *ProjectPost) LurkersCount() (count uint8) {
	_ = Db().Model(ProjectPostLurk{}).Where(&ProjectPostLurk{Hpid: post.ID()}).Count(&count)
	return
}

// URL returns the url of the post
func (post *ProjectPost) URL() *url.URL {
	return &url.URL{
		Host: Configuration.NERDZHost,
		Path: (post.Reference().(*Project)).Name + ":" + strconv.FormatUint(post.Pid, 10),
	}
}
