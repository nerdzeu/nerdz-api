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
	"time"

	"github.com/nerdzeu/nerdz-api/utils"
)

// NewProjectPostComment initializes a ProjectPostComment struct
func NewProjectPostComment(hcid uint64) (comment *ProjectPostComment, e error) {
	return NewProjectPostCommentWhere(&ProjectPostComment{Hcid: hcid})
}

// NewProjectPostCommentWhere returns the *ProjectPostComment fetching the first one that matches the description
func NewProjectPostCommentWhere(description *ProjectPostComment) (comment *ProjectPostComment, e error) {
	comment = new(ProjectPostComment)
	if e = Db().Model(ProjectPostComment{}).Where(description).Scan(comment); e != nil {
		return nil, e
	}
	if comment.Hcid == 0 {
		return nil, fmt.Errorf("requested ProjectPostComment does not exist")
	}
	return
}

// Implementing Message interface

// NumericReference returns the id of the recipient Post
func (comment *ProjectPostComment) NumericReference() uint64 {
	return comment.Hpid
}

// Reference returns the recipient *ProjectPost
func (comment *ProjectPostComment) Reference() Reference {
	post, _ := NewProjectPost(comment.NumericReference())
	return post
}

// NumericSender returns the id of the sender user
func (comment *ProjectPostComment) NumericSender() uint64 {
	return comment.From
}

// Sender returns the sender *User
func (comment *ProjectPostComment) Sender() *User {
	user, _ := NewUser(comment.NumericSender())
	return user
}

// Votes returns the post's votes value
func (comment *ProjectPostComment) VotesCount() (sum int) {
	_ = Db().Model(ProjectPostCommentVote{}).Select("COALESCE(sum(vote), 0)").Where(&ProjectPostCommentVote{Hcid: comment.Hcid}).Scan(&sum)
	return
}

// Votes returns a pointer to a slice of Vote
func (comment *ProjectPostComment) Votes() *[]Vote {
	ret := []ProjectPostCommentVote{}
	_ = Db().Model(ProjectPostCommentVote{}).Where(&ProjectPostCommentVote{Hcid: comment.Hcid}).Scan(&ret)
	var retVotes []Vote
	for _, v := range ret {
		vote := v
		retVotes = append(retVotes, Vote(&vote))
	}

	return &retVotes
}

// Text returns the post message
func (comment *ProjectPostComment) Text() string {
	return comment.Message
}

// Implementing ExistingComment interface

// ID returns the comment ID
func (comment *ProjectPostComment) ID() uint64 {
	return comment.Hcid
}

// Post returns the ExistingPost sturct to which the projectComment is related
func (comment *ProjectPostComment) Post() (ExistingPost, error) {
	var post *ProjectPost
	var err error
	if post, err = NewProjectPost(comment.Hpid); err != nil {
		return nil, err
	}
	return ExistingPost(post), nil
}

// Implementing NewComment interface

// SetSender set the source of the comment (the user ID)
func (comment *ProjectPostComment) SetSender(id uint64) {
	comment.From = id
}

// SetReference set the destination of the comment
func (comment *ProjectPostComment) SetReference(hpid uint64) {
	comment.Hpid = hpid
}

// ClearDefaults set to the go's default values the fields with default sql values
func (comment *ProjectPostComment) ClearDefaults() {
	comment.Time = time.Time{}
	comment.Editable = true
}

// SetText set the text of the message
func (comment *ProjectPostComment) SetText(message string) {
	comment.Message = message
}

// SetLanguage set the language of the comment
func (comment *ProjectPostComment) SetLanguage(language string) error {
	if language == "" {
		language = comment.Sender().Language()
	}
	if utils.InSlice(language, Configuration.Languages) {
		comment.Lang = language
		return nil
	}
	return fmt.Errorf("Language '%s' is not a valid or supported language", language)
}

// Language returns the message language
func (comment *ProjectPostComment) Language() string {
	return comment.Lang
}

// IsEditable returns true if the comment is editable
func (comment *ProjectPostComment) IsEditable() bool {
	return comment.Editable
}

// NumericOwners returns a slice of ids of the owner of the comment (the ones that can perform actions)
func (comment *ProjectPostComment) NumericOwners() []uint64 {
	return []uint64{comment.From, comment.To}
}

// Owners returns a slice of *User representing the users who own the comment
func (comment *ProjectPostComment) Owners() []*User {
	return Users(comment.NumericOwners())
}

// Revisions returns all the revisions of the message
func (comment *ProjectPostComment) Revisions() (modifications []string) {
	_ = Db().Model(ProjectPostCommentRevision{}).Where(&ProjectPostCommentRevision{Hcid: comment.Hcid}).Pluck("message", &modifications)
	return
}

// RevisionsNumber returns the number of the revisions
func (comment *ProjectPostComment) RevisionsNumber() (count uint8) {
	_ = Db().Model(ProjectPostCommentRevision{}).Where(&ProjectPostCommentRevision{Hcid: comment.Hcid}).Count(&count)
	return
}
