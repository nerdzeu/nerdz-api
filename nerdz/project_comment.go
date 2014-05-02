package nerdz

import (
	"errors"
	"html"
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

// Implementing NewComment interface

// Set the destination of the post. post can be a *ProjectPost or the post's id
func (comment *ProjectComment) SetTo(post interface{}) error {
	switch post.(type) {
	case int:
		comment.Hpid = int64(post.(int))
	case *ProjectPost:
		comment.Hpid = (post.(*ProjectPost)).Hpid
	default:
		return errors.New("Invalid comment type. Allowed int and *UserPost")
	}
	return nil
}

// SetMessage set NewComment message and escape html entities. Returns nil on success, error on failure
func (comment *ProjectComment) SetMessage(message string) error {
	if len(message) == 0 {
		return errors.New("Empty message")
	}

	comment.Message = html.EscapeString(message)
	return nil
}
