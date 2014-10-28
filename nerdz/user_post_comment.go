package nerdz

import (
	"errors"
	"html"
)

// NewUserPostComment initializes a UserPostComment struct
func NewUserPostComment(hcid uint64) (comment *UserPostComment, e error) {
	comment = new(UserPostComment)
	db.First(comment, hcid)

	if comment.Hcid != hcid {
		return nil, errors.New("Invalid hcid")
	}

	return comment, nil
}

// Implementing Message interface

// To returns the recipient *User
func (comment *UserPostComment) Recipient() (Board, error) {
	return NewUser(comment.To)
}

// From returns the sender *User
func (comment *UserPostComment) Sender() (*User, error) {
	return NewUser(comment.From)
}

// Thumbs returns the post's thumbs value
func (comment *UserPostComment) Thumbs() int {
	type result struct {
		Total int
	}
	var sum result
	db.Model(UserPostCommentThumb{}).Select("COALESCE(sum(vote), 0) as total").Where(&UserPostCommentThumb{Hcid: comment.Hcid}).Scan(&sum)
	return sum.Total
}

// Post returns the *Post sturct to which the comment is related
func (comment *UserPostComment) Post() (*UserPost, error) {
	return NewUserPost(comment.Hpid)
}

// Message returns the post message
func (comment *UserPostComment) Text() string {
	return comment.Message
}

// Implementing NewComment interface

// Set the source of the comment (the user ID)
func (comment *UserPostComment) SetSender(id uint64) {
	comment.From = id
}

// Set the destination of the post. post can be a *UserPost or the post's id
func (comment *UserPostComment) SetRecipient(id uint64) {
	comment.Hpid = id
}

// SetMessage set NewComment message and escape html entities. Returns nil on success, error on failure
func (comment *UserPostComment) SetText(message string) error {
	if len(message) == 0 {
		return errors.New("Empty message")
	}

	comment.Message = html.EscapeString(message)
	return nil
}
