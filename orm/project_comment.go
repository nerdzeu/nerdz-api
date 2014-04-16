package orm

import (
	"errors"
)

// NewProjectComment initializes a ProjectComment struct
func NewProjectComment(hcid int64) (comment *ProjectComment, e error) {
    comment = new(ProjectComment)
	db.First(comment, hcid)

	if comment.Hcid != hcid {
		return nil, errors.New("Invalid hcid")
	}

	return comment, nil
}

// GetTo returns the recipient *Project
func (comment *ProjectComment) GetTo() (*Project, error) {
    return NewProject(comment.To)
}

// GetFrom returns the sender *User
func (comment *ProjectComment) GetFrom() (*User, error) {
    return NewUser(comment.From)
}

// GetProjectPost returns the *ProjectPost sturct to which the projectComment is related
func (comment *ProjectComment) GetPost() (*ProjectPost, error) {
    return NewProjectPost(comment.Hpid)
}
