package orm

import (
	"errors"
)

// New initializes a ProjectComment struct
func (projectComment *ProjectComment) New(hcid int64) error {
	db.First(projectComment, hcid)

	if projectComment.Hcid != hcid {
		return errors.New("Invalid hcid")
	}

	return nil
}

// GetTo returns the recipient *User
func (projectComment *ProjectComment) GetTo() (*User, error) {
	var to User

	if err := to.New(projectComment.To); err != nil {
		return nil, err
	}

	return &to, nil
}

// GetFrom returns the sender *User
func (projectComment *ProjectComment) GetFrom() (*User, error) {
	var from User

	if err := from.New(projectComment.From); err != nil {
		return nil, err
	}

	return &from, nil
}

// GetProjectPost returns the *ProjectPost sturct to which the projectComment is related
func (projectComment *ProjectComment) GetProjectPost() (*ProjectPost, error) {
	var post ProjectPost
	err := post.New(projectComment.Hpid)
	return &post, err
}
