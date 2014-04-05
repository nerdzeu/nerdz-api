package orm

import (
	"errors"
	"net/mail"
	"net/url"
	"time"
)

type PersonalInfo struct {
	Id       int64
	IsOnline bool
	Nation   string
	Timezone string
	Username string
	Name     string
	Surname  string
	Gender   bool
	Birthday time.Time
	Gravatar *url.URL
}

type ContactInfo struct {
	Email    *mail.Address
	Website  *url.URL
	GitHub   *url.URL
	Skype    string
	Jabber   string
	Yahoo    *mail.Address
	Facebook *url.URL
	Twitter  *url.URL
	Steam    string
}

// New initializes a User struct
func (user *User) New(id int64) error {
	db.First(user, id)
	db.Find(&user.Profile, id)

	if user.Counter != id || user.Profile.Counter != id {
		return errors.New("Invalid id")
	}

	return nil
}

// GetInfo returns a PersonalInfo struct
func (user *User) GetPersonalInfo() *PersonalInfo {
	return &PersonalInfo{
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

// GetContactInfo returns a ContactInfo struct
func (user *User) GetContactInfo() *ContactInfo {
	// Errors should never occurs, since values are stored in db after have been controlled
	email, _ := mail.ParseAddress(user.Email)
	yahoo, _ := mail.ParseAddress(user.Profile.Yahoo)
	website, _ := url.Parse(user.Profile.Website)
	github, _ := url.Parse(user.Profile.Github)
	facebook, _ := url.Parse(user.Profile.Facebook)
	twitter, _ := url.Parse(user.Profile.Twitter)

	// Set Address.Name field
	emailName := user.Surname + " " + user.Name
	// email is always != nil, since an email is always required
	email.Name = emailName
	// yahoo address can be nil
	if yahoo != nil {
		yahoo.Name = emailName
	}

	return &ContactInfo{
		Email:    email,
		Website:  website,
		GitHub:   github,
		Skype:    user.Profile.Skype,
		Jabber:   user.Profile.Jabber,
		Yahoo:    yahoo,
		Facebook: facebook,
		Twitter:  twitter,
		Steam:    user.Profile.Steam}
}
