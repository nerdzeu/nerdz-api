package nerdz

import (
	"errors"
	"github.com/nerdzeu/nerdz-api/utils"
)

// NewProjectPost initializes a ProjectPost struct
func NewProjectPost(hpid int64) (post *ProjectPost, e error) {
	post = new(ProjectPost)
	db.First(post, hpid)

	if post.Hpid != hpid {
		return nil, errors.New("Invalid hpid")
	}

	return post, nil
}

// Implementing Board interface

// GetTo returns the recipient *Project
func (post *ProjectPost) GetTo() (*Project, error) {
	return NewProject(post.To)
}

// GetFrom returns the sender *User
func (post *ProjectPost) GetFrom() (*User, error) {
	return NewUser(post.From)
}

// GetThumbs returns the post's thumbs value
func (post *ProjectPost) GetThumbs() int {
	var sum struct {
		Total int
	}

	db.Table("groups_thumbs").Select("COALESCE(sum(vote), 0) as total").Where(&UserPostThumb{Hpid: post.Hpid}).Scan(&sum)
	return sum.Total
}

// GetComments returns the full comments list, or the selected range of comments
// GetComments()  returns the full comments list
// GetComments(N) returns at most the last N comments
// GetComments(N, X) returns at most N comments, before the last comment + X
func (post *ProjectPost) GetComments(interval ...int) interface{} {
	var comments []ProjectComment

	switch len(interval) {
	default: //full list
	case 0:
		db.Find(&comments, &ProjectComment{Hpid: post.Hpid})

	case 1: // Get last interval[0] comments [ LIMIT interval[0] ]
		db.Order("hcid DESC").Limit(interval[0]).Find(&comments, &ProjectComment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]ProjectComment)

	case 2: // Get last interval[0] comments, starting from interval[1] [ LIMIT interval[0] OFFSET interval[1] ]
		db.Order("hcid DESC").Limit(interval[0]).Offset(interval[1]).Find(&comments, &ProjectComment{Hpid: post.Hpid})
		comments = utils.ReverseSlice(comments).([]ProjectComment)
	}

	return comments
}
