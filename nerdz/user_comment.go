package nerdz

import (
	"errors"
	"html"
)

// NewUserComment initializes a UserComment struct
func NewUserComment(hcid int64) (comment *UserComment, e error) {
	comment = new(UserComment)
	db.First(comment, hcid)

	if comment.Hcid != hcid {
		return nil, errors.New("Invalid hcid")
	}

	return comment, nil
}

// GetTo returns the recipient *User
func (comment *UserComment) GetTo() (*User, error) {
	return NewUser(comment.To)
}

// GetFrom returns the sender *User
func (comment *UserComment) GetFrom() (*User, error) {
	return NewUser(comment.From)
}

// GetPost returns the *Post sturct to which the comment is related
func (comment *UserComment) GetPost() (*UserPost, error) {
	return NewUserPost(comment.Hpid)
}

// Implementing NewComment interface

// Set the destination of the post. post can be a *UserPost or the post's id
func (comment *UserComment) SetTo(post interface{}) error {
	switch post.(type) {
	case int:
		comment.Hpid = int64(post.(int))
	case *UserPost:
		comment.Hpid = (post.(*UserPost)).Hpid
	default:
		return errors.New("Invalid comment type. Allowed int and *UserPost")
	}
	return nil
}

// SetMessage set NewComment message and escape html entities. Returns nil on success, error on failure
func (comment *UserComment) SetMessage(message string) error {
	if len(message) == 0 {
		return errors.New("Empty message")
	}

	comment.Message = html.EscapeString(message)
	return nil
}
