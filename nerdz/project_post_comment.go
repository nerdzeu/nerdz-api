package nerdz

import (
	"errors"
	"fmt"
	"time"

	"github.com/nerdzeu/nerdz-api/utils"
)

// NewProjectPostComment initializes a ProjectPostComment struct
func NewProjectPostComment(hcid uint64) (comment *ProjectPostComment, e error) {
	comment = new(ProjectPostComment)
	Db().First(comment, hcid)

	if comment.Hcid != hcid {
		return nil, errors.New("Invalid hcid")
	}

	return comment, nil
}

// Implementing Message interface

// NumericReference returns the id of the recipient Post
func (comment *ProjectPostComment) NumericReference() uint64 {
	return comment.Hpid
}

// Reference returns the recipient *ProjectPost
func (comment *ProjectPostComment) Reference() Reference {
	post, _ := NewProjectPost(comment.NumericReference())
	return post
}

// NumericSender returns the id of the sender user
func (comment *ProjectPostComment) NumericSender() uint64 {
	return comment.From
}

// Sender returns the sender *User
func (comment *ProjectPostComment) Sender() *User {
	user, _ := NewUser(comment.NumericSender())
	return user
}

// Thumbs returns the post's thumbs value
func (comment *ProjectPostComment) Thumbs() int {
	type result struct {
		Total int
	}
	var sum result
	Db().Model(ProjectPostCommentThumb{}).Select("COALESCE(sum(vote), 0) as total").Where(&ProjectPostCommentThumb{Hcid: comment.Hcid}).Scan(&sum)
	return sum.Total
}

// Text returns the post message
func (comment *ProjectPostComment) Text() string {
	return comment.Message
}

// Implementing ExistingComment interface

// ID returns the comment ID
func (comment *ProjectPostComment) ID() uint64 {
	return comment.Hcid
}

// Post returns the *ProjectPost sturct to which the projectComment is related
func (comment *ProjectPostComment) Post() (*ProjectPost, error) {
	return NewProjectPost(comment.Hpid)
}

// Implementing NewComment interface

// SetSender set the source of the comment (the user ID)
func (comment *ProjectPostComment) SetSender(id uint64) {
	comment.From = id
}

// SetReference set the destination of the comment
func (comment *ProjectPostComment) SetReference(hpid uint64) {
	comment.Hpid = hpid
}

// ClearDefaults set to the go's default values the fields with default sql values
func (comment *ProjectPostComment) ClearDefaults() {
	comment.Time = time.Time{}
}

// SetText set the text of the message
func (comment *ProjectPostComment) SetText(message string) {
	comment.Message = message
}

// SetLanguage set the language of the comment (TODO: add db side column)
func (comment *ProjectPostComment) SetLanguage(language string) error {
	if utils.InSlice(language, Configuration.Languages) {
		//post.Lang = language
		return nil
	}
	return fmt.Errorf("Language '%s' is not a valid or supported language", language)
}

// Language returns the message language
func (comment *ProjectPostComment) Language() string {
	return comment.Reference().(Reference).Language()
}

// IsEditable returns true if the comment is editable
func (comment *ProjectPostComment) IsEditable() bool {
	return comment.Editable
}

// NumericOwners returns a slice of ids of the owner of the comment (the ones that can perform actions)
func (comment *ProjectPostComment) NumericOwners() []uint64 {
	return []uint64{comment.From, comment.To}
}

// Owners returns a slice of *User representing the users who own the comment
func (comment *ProjectPostComment) Owners() []*User {
	return Users(comment.NumericOwners())
}

// Revisions returns all the revisions of the message
func (comment *ProjectPostComment) Revisions() (modifications []string) {
	Db().Model(ProjectPostCommentRevision{}).Where(&ProjectPostCommentRevision{Hcid: comment.Hcid}).Pluck("message", &modifications)
	return
}

// RevisionsNumber returns the number of the revisions
func (comment *ProjectPostComment) RevisionsNumber() (count uint8) {
	Db().Model(ProjectPostCommentRevision{}).Where(&ProjectPostCommentRevision{Hcid: comment.Hcid}).Count(&count)
	return
}
