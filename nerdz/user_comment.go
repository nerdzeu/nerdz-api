package nerdz

import (
	"errors"
	"fmt"
	"html"
	"reflect"
)

// NewUserPostComment initializes a UserPostComment struct
func NewUserPostComment(hcid uint64) (comment *UserPostComment, e error) {
	comment = new(UserPostComment)
	db.First(comment, hcid)

	if comment.Hcid != hcid {
		return nil, errors.New("Invalid hcid")
	}

	return comment, nil
}

// To returns the recipient *User
func (comment *UserPostComment) Recipient() (*User, error) {
	return NewUser(comment.To)
}

// From returns the sender *User
func (comment *UserPostComment) Sender() (*User, error) {
	return NewUser(comment.From)
}

// Post returns the *Post sturct to which the comment is related
func (comment *UserPostComment) Post() (*UserPost, error) {
	return NewUserPost(comment.Hpid)
}

// Implementing NewComment interface

// Set the destination of the post. post can be a *UserPost or the post's id
func (comment *UserPostComment) SetRecipient(post interface{}) error {
	switch post.(type) {
	case uint64:
		comment.Hpid = post.(uint64)
	case *UserPost:
		comment.Hpid = (post.(*UserPost)).Hpid
	default:
		return fmt.Errorf("Invalid comment type: %v. Allowed uint64 and *UserPostComment", reflect.TypeOf(comment))
	}
	return nil
}

// SetMessage set NewComment message and escape html entities. Returns nil on success, error on failure
func (comment *UserPostComment) SetText(message string) error {
	if len(message) == 0 {
		return errors.New("Empty message")
	}

	comment.Message = html.EscapeString(message)
	return nil
}
