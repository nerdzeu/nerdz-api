package nerdz

import (
	"errors"
	"fmt"
	"github.com/nerdzeu/nerdz-api/utils"
	"html"
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

// Implementing Message interface

// To returns the recipient *Project
func (comment *ProjectPostComment) Reference() Reference {
	project, _ := NewProject(comment.To)
	return project
}

// From returns the sender *User
func (comment *ProjectPostComment) Sender() *User {
	user, _ := NewUser(comment.From)
	return user
}

// Thumbs returns the post's thumbs value
func (comment *ProjectPostComment) Thumbs() int {
	type result struct {
		Total int
	}
	var sum result
	db.Model(ProjectPostCommentThumb{}).Select("COALESCE(sum(vote), 0) as total").Where(&ProjectPostCommentThumb{Hcid: comment.Hcid}).Scan(&sum)
	return sum.Total
}

// Message returns the post message
func (comment *ProjectPostComment) Text() string {
	return comment.Message
}

// Implementing ExistingComment interface

// Id returns the comment ID
func (comment *ProjectPostComment) Id() uint64 {
	return comment.Hcid
}

// ProjectPost returns the *ProjectPost sturct to which the projectComment is related
func (comment *ProjectPostComment) Post() (*ProjectPost, error) {
	return NewProjectPost(comment.Hpid)
}

// Implementing NewComment interface

// Set the source of the comment (the user ID)
func (comment *ProjectPostComment) SetSender(id uint64) {
	comment.From = id
}

// Set the destination of the post
func (comment *ProjectPostComment) SetReference(hpid uint64) {
	comment.Hpid = hpid
}

// SetText set the text of the message
func (comment *ProjectPostComment) SetText(message string) {
	comment.Message = html.EscapeString(message)
}

// SetLanguage set the language of the comment (TODO: add db side column)
func (comment *ProjectPostComment) SetLanguage(language string) error {
	if utils.InSlice(language, Configuration.Languages) {
		//post.Lang = language
		return nil
	}
	return fmt.Errorf("Language '%s' is not valid a supported language", language)
}

// Lanaugage returns the message language
func (comment *ProjectPostComment) Language() string {
	return comment.Reference().(Reference).Language()
}
