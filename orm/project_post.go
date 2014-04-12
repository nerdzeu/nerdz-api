package orm

import (
	"errors"
)

// New initializes a ProjectPost struct
func (post *ProjectPost) New(hpid int64) error {
	db.First(post, hpid)

	if post.Hpid != hpid {
		return errors.New("Invalid hpid")
	}

	return nil
}

// GetTo returns the recipient *User
func (post *ProjectPost) GetTo() (*Project, error) {
	var to Project

	if err := to.New(post.To); err != nil {
		return nil, err
	}

	return &to, nil
}

// GetFrom returns the sender *User
func (post *ProjectPost) GetFrom() (*User, error) {
	var from User

	if err := from.New(post.From); err != nil {
		return nil, err
	}

	return &from, nil
}
