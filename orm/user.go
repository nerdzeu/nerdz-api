package orm

import (
	"errors"
	"net/url"
	"time"
)

// JSON

type Info struct {
	Id       int64
	IsOnline bool
	Nation   string
	Timezone string
	Username string
	Name     string
	Surname  string
	Gender   bool
	Birthday time.Time
	Gravatar url.URL
}

func (user *User) New(id int64) error {

	db.First(user, id)
	db.Find(&user.Profile, id)

	if user.Counter != id || user.Profile.Counter != id {
		return errors.New("Invalid id")
	}

	return nil
}

func (user *User) GetInfo() *Info {

	return &Info{
		Id:       user.Counter,
		IsOnline: user.Viewonline && user.Last.Add(time.Duration(5)*time.Minute).After(time.Now()),
		Nation:   user.Lang,
		Timezone: user.Timezone,
		Name:     user.Name,
		Surname:  user.Surname,
		Gender:   user.Gender,
		Birthday: user.BirthDate,
		Gravatar: getGravatar(user.Email)}
}
