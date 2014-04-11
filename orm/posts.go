package orm

import (
	"errors"
	"time"
)

type iPost struct {
	Hpid    int64
	From    *User
	Pid     int64
	Message string
	Time    time.Time
}

type UserPost struct {
	iPost
	To *User
}

type ProjectPost struct {
	iPost
	To *Group
}

// New initializes a Post struct
func (this *UserPost) New(hpid int64) error {
	var post Post
	db.First(&post, hpid)

	if post.Hpid != hpid {
		return errors.New("Invalid id")
	}

	var from, to User

	if err := from.New(post.From); err != nil {
		return err
	}

	if err := to.New(post.To); err != nil {
		return err
	}

	this.Hpid = hpid
	this.Pid = post.Hpid
	this.Message = post.Message
	this.Time = post.Time
	this.From = &from
	this.To = &to

	return nil
}

// New initializes a ProjectPost struct
func (this *ProjectPost) New(hpid int64) error {
	var post GroupsPost
	db.First(&post, hpid)

	if post.Hpid != hpid {
		return errors.New("Invalid id")
	}

	var from User
	var to Group

	if err := from.New(post.From); err != nil {
		return err
	}

	if err := to.New(post.To); err != nil {
		return err
	}

	this.Hpid = hpid
	this.Pid = post.Hpid
	this.Message = post.Message
	this.Time = post.Time
	this.From = &from
	this.To = &to

	return nil
}
