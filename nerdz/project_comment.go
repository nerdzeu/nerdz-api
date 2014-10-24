package nerdz

import (
	"errors"
	"fmt"
	"html"
	"reflect"
)

// NewProjectPostComment initializes a ProjectPostComment struct
func NewProjectPostComment(hcid uint64) (comment *ProjectPostComment, e error) {
	comment = new(ProjectPostComment)
	db.First(comment, hcid)

	if comment.Hcid != hcid {
		return nil, errors.New("Invalid hcid")
	}

	return comment, nil
}

// Implementing Message interface

// To returns the recipient *Project
func (comment *ProjectPostComment) Recipient() (Board, error) {
	return NewProject(comment.To)
}

// From returns the sender *User
func (comment *ProjectPostComment) Sender() (*User, error) {
	return NewUser(comment.From)
}

// Thumbs returns the post's thumbs value
func (comment *ProjectPostComment) Thumbs() int {
	type result struct {
		Total int
	}
	var sum result
	db.Model(ProjectPostCommentThumb{}).Select("COALESCE(sum(vote), 0) as total").Where(&ProjectPostCommentThumb{Hcid: comment.Hcid}).Scan(&sum)
	return sum.Total
}

// Message returns the post message
func (comment *ProjectPostComment) Text() string {
	return comment.Message
}

// Implementing ExistingComment interface

// ProjectPost returns the *ProjectPost sturct to which the projectComment is related
func (comment *ProjectPostComment) Post() (*ProjectPost, error) {
	return NewProjectPost(comment.Hpid)
}

// Implementing NewComment interface

// Set the source of the comment (the user ID)
func (comment *ProjectPostComment) SetSender(id uint64) {
	comment.From = id
}

// Set the destination of the post
func (comment *ProjectPostComment) SetRecipient(hpid uint64) {
	comment.Hpid = hpid
}

// SetMessage set NewComment message and escape html entities. Returns nil on success, error on failure
func (comment *ProjectPostComment) SetText(message string) error {
	if len(message) == 0 {
		return errors.New("Empty message")
	}

	comment.Message = html.EscapeString(message)
	return nil
}
