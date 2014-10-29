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

// NumericSender returns the id of the sender user
func (comment *UserPostComment) NumericSender() uint64 {
	return post.From
}

// From returns the sender *User
func (comment *UserPostComment) Sender() *User {
	user, _ := NewUser(post.NumericSender())
	return user
}

// NumericRecipient returns the id of the recipient user
func (comment *UserPostComment) NumericRecipient() uint64 {
	return post.To
}

// To returns the recipient *User
func (comment *UserPostComment) Recipient() (Board, error) {
	user, _ := NewUser(post.NumericRecipient())
	return user
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

// SetText set the text of the message
func (comment *UserPostComment) SetText(message string) {
	comment.Message = html.EscapeString(message)
}
