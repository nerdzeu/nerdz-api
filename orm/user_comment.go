package orm

import (
	"errors"
)

// NewUserComment initializes a UserComment struct
func NewUserComment(hcid int64) ( comment *UserComment, e error ) {
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
