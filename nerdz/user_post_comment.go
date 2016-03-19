/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

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
	"time"

	"github.com/nerdzeu/nerdz-api/utils"
)

// NewUserPostComment initializes a UserPostComment struct
func NewUserPostComment(hcid uint64) (comment *UserPostComment, e error) {
	comment = new(UserPostComment)
	e = Db().First(comment, hcid)
	return
}

// Implementing Message interface

// NumericSender returns the id of the sender user
func (comment *UserPostComment) NumericSender() uint64 {
	return comment.From
}

// Sender returns the sender *User
func (comment *UserPostComment) Sender() *User {
	user, _ := NewUser(comment.NumericSender())
	return user
}

// NumericReference returns the id of the recipient Post
func (comment *UserPostComment) NumericReference() uint64 {
	return comment.Hpid
}

// Reference returns the recipient *Post
func (comment *UserPostComment) Reference() Reference {
	post, _ := NewUserPost(comment.NumericReference())
	return post
}

// Thumbs returns the post's thumbs value
func (comment *UserPostComment) Thumbs() (sum int) {
	Db().Model(UserPostCommentThumb{}).Select("COALESCE(sum(vote), 0)").Where(&UserPostCommentThumb{Hcid: comment.Hcid}).Scan(&sum)
	return
}

// Post returns the ExistingPost struct to which the comment is related
func (comment *UserPostComment) Post() (ExistingPost, error) {
	var post *UserPost
	var err error
	if post, err = NewUserPost(comment.Hpid); err != nil {
		return nil, err
	}
	return ExistingPost(post), nil
}

// Text returns the post message
func (comment *UserPostComment) Text() string {
	return comment.Message
}

// ID returns the UserPostComment ID
func (comment *UserPostComment) ID() uint64 {
	return comment.Hcid
}

// Implementing NewComment interface

// SetSender sets the source of the comment (the user ID)
func (comment *UserPostComment) SetSender(id uint64) {
	comment.From = id
}

// SetReference sets the destination of the comment (the post ID)
func (comment *UserPostComment) SetReference(id uint64) {
	comment.Hpid = id
}

// SetText set the text of the message
func (comment *UserPostComment) SetText(message string) {
	comment.Message = message
}

// ClearDefaults set to the go's default values the fields with default sql values
func (comment *UserPostComment) ClearDefaults() {
	comment.Time = time.Time{}
	comment.Editable = true
}

// SetLanguage set the language of the comment (TODO: add db side column)
func (comment *UserPostComment) SetLanguage(language string) error {
	if utils.InSlice(language, Configuration.Languages) {
		//post.Lang = language
		return nil
	}
	return fmt.Errorf("Language '%s' is not a valid or supported language", language)
}

// Language returns the message language
func (comment *UserPostComment) Language() string {
	return comment.Reference().(Reference).Language()
}

// IsEditable returns true if the comment is editable
func (comment *UserPostComment) IsEditable() bool {
	return comment.Editable
}

// NumericOwners returns a slice of ids of the owner of the comment (the ones that can perform actions)
func (comment *UserPostComment) NumericOwners() []uint64 {
	return []uint64{comment.From, comment.To}
}

// Owners returns a slice of *User representing the users who own the comment
func (comment *UserPostComment) Owners() []*User {
	return Users(comment.NumericOwners())
}

// Revisions returns all the revisions of the message
func (comment *UserPostComment) Revisions() (modifications []string) {
	Db().Model(UserPostCommentRevision{}).Where(&UserPostCommentRevision{Hcid: comment.Hcid}).Pluck("message", &modifications)
	return
}

// RevisionsNumber returns the number of the revisions
func (comment *UserPostComment) RevisionsNumber() (count uint8) {
	Db().Model(UserPostCommentRevision{}).Where(&UserPostCommentRevision{Hcid: comment.Hcid}).Count(&count)
	return
}
