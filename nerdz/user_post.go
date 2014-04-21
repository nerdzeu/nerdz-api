package nerdz

import (
	"errors"
	"github.com/nerdzeu/nerdz-api/utils"
)

// New initializes a Post struct
func NewUserPost(hpid int64) (post *UserPost, e error) {
	post = new(UserPost)
	db.First(post, hpid)

	if post.Hpid != hpid {
		return nil, errors.New("Invalid hpid")
	}

	return post, nil
}

// GetTo returns the recipient *User
func (post *UserPost) GetTo() (*User, error) {
	return NewUser(post.To)
}

// GetFrom returns the sender *User
func (post *UserPost) GetFrom() (*User, error) {
	return NewUser(post.From)
}

// GetThumbs returns the post's thumbs value
func (post *UserPost) GetThumbs() int {
	type result struct {
		Total int
	}
	var sum result
	db.Table("thumbs").Select("COALESCE(sum(vote), 0) as total").Where(&UserPostThumb{Hpid: post.Hpid}).Scan(&sum)
	return sum.Total
}

// GetComments returns the full comments list, or the selected range of comments
// GetComments()  returns the full comments list
// GetComments(N) returns at most the last N comments
// GetComments(N, X) returns at most N comments, before the last comment + X
func (post *UserPost) GetComments(interval ...int) []UserComment {
	var comments []UserComment

	switch len(interval) {
	default: //full list
	case 0:
		db.Find(&comments, &UserComment{Hpid: post.Hpid})

	case 1: // Get last interval[0] comments [ LIMIT interval[0] ]
		db.Order("hcid DESC").Limit(interval[0]).Find(&comments, &UserComment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]UserComment)

	case 2: // Get last interval[0] comments, starting from interval[1] [ LIMIT interval[0] OFFSET interval[1] ]
		db.Order("hcid DESC").Limit(interval[0]).Offset(interval[1]).Find(&comments, &UserComment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]UserComment)
	}

	return comments
}
