package orm

import (
	"errors"
)

// New initializes a Comment struct
func (comment *Comment) New(hcid int64) error {
	db.First(comment, hcid)

	if comment.Hcid != hcid {
		return errors.New("Invalid hcid")
	}

	return nil
}

// GetTo returns the recipient *User
func (comment *Comment) GetTo() (*User, error) {
	var to User

	if err := to.New(comment.To); err != nil {
		return nil, err
	}

	return &to, nil
}

// GetFrom returns the sender *User
func (comment *Comment) GetFrom() (*User, error) {
	var from User

	if err := from.New(comment.From); err != nil {
		return nil, err
	}

	return &from, nil
}

// GetPost returns the *Post sturct to which the comment is related
func (comment *Comment) GetPost() (*Post, error) {
	var post Post
	err := post.New(comment.Hpid)
	return &post, err
}
