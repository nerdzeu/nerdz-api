package nerdz

import (
	"errors"
	"html"
)

// NewUserPostComment initializes a UserPostComment struct
func NewUserPostComment(hcid int64) (comment *UserPostComment, e error) {
	comment = new(UserPostComment)
	db.First(comment, hcid)

	if comment.Hcid != hcid {
		return nil, errors.New("Invalid hcid")
	}

	return comment, nil
}

// GetTo returns the recipient *User
func (comment *UserPostComment) GetTo() (*User, error) {
	return NewUser(comment.To)
}

// GetFrom returns the sender *User
func (comment *UserPostComment) GetFrom() (*User, error) {
	return NewUser(comment.From)
}

// GetPost returns the *Post sturct to which the comment is related
func (comment *UserPostComment) GetPost() (*UserPost, error) {
	return NewUserPost(comment.Hpid)
}

// Implementing NewComment interface

// Set the destination of the post. post can be a *UserPost or the post's id
func (comment *UserPostComment) SetTo(post interface{}) error {
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
func (comment *UserPostComment) SetMessage(message string) error {
	if len(message) == 0 {
		return errors.New("Empty message")
	}

	comment.Message = html.EscapeString(message)
	return nil
}
