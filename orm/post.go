package orm

import (
	"errors"
	"github.com/nerdzeu/nerdz-api/utils"
)

// New initializes a Post struct
func (post *Post) New(hpid int64) error {
	db.First(post, hpid)

	if post.Hpid != hpid {
		return errors.New("Invalid hpid")
	}

	return nil
}

// GetTo returns the recipient *User
func (post *Post) GetTo() (*User, error) {
	var to User

	if err := to.New(post.To); err != nil {
		return nil, err
	}

	return &to, nil
}

// GetFrom returns the sender *User
func (post *Post) GetFrom() (*User, error) {
	var from User

	if err := from.New(post.From); err != nil {
		return nil, err
	}

	return &from, nil
}

// GetThumbs returns the post's thumbs value
func (post *Post) GetThumbs() int {
	type result struct {
		Total int
	}
	var sum result
	db.Table("thumbs").Select("sum(vote) as total").Where(&PostThumb{Hpid: post.Hpid}).Scan(&sum)
	return sum.Total
}

// GetComments returns the full comments list, or the selected range of comments
// GetComments()  returns the full comments list
// GetComments(N) returns at most the last N comments
// GetComments(N, X) returns at most N comments, before the last comment + X
func (post *Post) GetComments(interval ...int) []Comment {
	var comments []Comment

	switch len(interval) {
	default: //full list
	case 0:
		db.Find(&comments, &Comment{Hpid: post.Hpid})

	case 1: // Get last interval[0] comments [ LIMIT interval[0] ]
		db.Order("hcid DESC").Limit(interval[0]).Find(&comments, &Comment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]Comment)

	case 2: // Get last interval[0] comments, starting from interval[1] [ LIMIT interval[0] OFFSET interval[1] ]
		db.Order("hcid DESC").Limit(interval[0]).Offset(interval[1]).Find(&comments, &Comment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]Comment)
	}

	return comments
}
