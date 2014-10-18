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

// To returns the recipient *Project
func (comment *ProjectPostComment) Recipient() (*Project, error) {
	return NewProject(comment.To)
}

// From returns the sender *User
func (comment *ProjectPostComment) Sender() (*User, error) {
	return NewUser(comment.From)
}

// ProjectPost returns the *ProjectPost sturct to which the projectComment is related
func (comment *ProjectPostComment) Post() (*ProjectPost, error) {
	return NewProjectPost(comment.Hpid)
}

// Implementing NewComment interface

// Set the destination of the post. post can be a *ProjectPost or the post's id
func (comment *ProjectPostComment) SetRecipient(post interface{}) error {
	switch post.(type) {
	case uint64:
		comment.Hpid = post.(uint64)
	case *ProjectPost:
		comment.Hpid = (post.(*ProjectPost)).Hpid
	default:
		return fmt.Errorf("Invalid post type: %v. Allowed uint64 and *ProjectPostComment", reflect.TypeOf(post))
	}
	return nil
}

// SetMessage set NewComment message and escape html entities. Returns nil on success, error on failure
func (comment *ProjectPostComment) SetText(message string) error {
	if len(message) == 0 {
		return errors.New("Empty message")
	}

	comment.Message = html.EscapeString(message)
	return nil
}
