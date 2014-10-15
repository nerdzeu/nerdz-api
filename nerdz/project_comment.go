package nerdz

import (
	"errors"
	"html"
    "reflect"
    "fmt"
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

// GetTo returns the recipient *Project
func (comment *ProjectPostComment) GetTo() (*Project, error) {
	return NewProject(comment.To)
}

// GetFrom returns the sender *User
func (comment *ProjectPostComment) GetFrom() (*User, error) {
	return NewUser(comment.From)
}

// GetProjectPost returns the *ProjectPost sturct to which the projectComment is related
func (comment *ProjectPostComment) GetPost() (*ProjectPost, error) {
	return NewProjectPost(comment.Hpid)
}

// Implementing NewComment interface

// Set the destination of the post. post can be a *ProjectPost or the post's id
func (comment *ProjectPostComment) SetTo(post interface{}) error {
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
func (comment *ProjectPostComment) SetMessage(message string) error {
	if len(message) == 0 {
		return errors.New("Empty message")
	}

	comment.Message = html.EscapeString(message)
	return nil
}
