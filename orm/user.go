package orm

import (
	"errors"
	"net/mail"
	"net/url"
	"strings"
	"time"
)

type PersonalInfo struct {
	Id        int64
	IsOnline  bool
	Nation    string
	Timezone  string
	Username  string
	Name      string
	Surname   string
	Gender    bool
	Birthday  time.Time
	Gravatar  *url.URL
	Interests []string
	Quotes    []string
	Biography string
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

type Template struct {
	Number int16
	Name   string
}

type BoardInfo struct {
	Language       string
	Template       *Template
	MobileTemplate *Template
	Dateformat     string
	IsClosed       bool
	Private        bool
	WhiteList      []*User
	UserScript     *url.URL
}

// NewUser initializes a User struct
func NewUser(id int64) (user *User, e error) {
	user = new(User)
	db.First(user, id)
	db.Find(&user.Profile, id)
	if user.Counter != id || user.Profile.Counter != id {
		return nil, errors.New("Invalid id")
	}

	return user, nil
}

// GetPersonalInfo returns a *PersonalInfo struct
func (user *User) GetPersonalInfo() *PersonalInfo {
	return &PersonalInfo{
		Id:        user.Counter,
		Username:  user.Username,
		IsOnline:  user.Viewonline && user.Last.Add(time.Duration(5)*time.Minute).After(time.Now()),
		Nation:    user.Lang.String,
		Timezone:  user.Timezone,
		Name:      user.Name,
		Surname:   user.Surname,
		Gender:    user.Gender,
		Birthday:  user.BirthDate,
		Gravatar:  getGravatar(user.Email),
		Interests: strings.Split(user.Profile.Interests, "\n"),
		Quotes:    strings.Split(user.Profile.Quotes, "\n"),
		Biography: user.Profile.Biography}
}

// GetContactInfo returns a *ContactInfo struct
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

// GetBoardInfo returns a BoardInfo struct
func (user *User) GetBoardInfo() *BoardInfo {
	defaultTemplate := Template{
		Name:   "", //TODO: find a way to get template name -> unfortunately isn't stored in the database at the moment
		Number: user.Profile.Template}

	mobileTemplate := Template{
		Name:   "", //TODO: find a way to get template name -> unfortunately isn't stored in the database at the moment
		Number: user.Profile.MobileTemplate}

	var closedProfile ClosedProfile
	db.First(&closedProfile, user.Counter)
	closed := closedProfile.Counter == user.Counter

	var whiteList []*User

	if closed {
		var wl []Whitelist
		db.Find(&wl, Whitelist{From: user.Counter})
		for _, elem := range wl {
			user, _ := NewUser(elem.To)
			whiteList = append(whiteList, user)
		}
	}

	return &BoardInfo{
		Language:       user.BoardLang.String,
		Template:       &defaultTemplate,
		MobileTemplate: &mobileTemplate,
		Dateformat:     user.Profile.Dateformat,
		IsClosed:       closed,
		Private:        user.Private,
		WhiteList:      whiteList}
}

// GetBlacklist returns a []*User that user (*User) put in his blacklist
func (user *User) GetBlacklist() []*User {
	var blacklist []*User
	var bl []Blacklist

	db.Find(&bl, Blacklist{From: user.Counter})
	for _, elem := range bl {
		user, _ := NewUser(elem.To)
		blacklist = append(blacklist, user)
	}

	return blacklist
}

//Implements Board interface

//GetInfo returns a *Info struct
func (user *User) GetInfo() *Info {

	website, _ := url.Parse(user.Profile.Website)

	var followers []*User
	var fl []UserFollow

	db.Find(&fl, UserFollow{To: user.Counter})
	for _, elem := range fl {
		user, _ := NewUser(elem.From)
		followers = append(followers, user)
	}

	return &Info{
		Id:        user.Counter,
		Owner:     user,
		Followers: followers,
		Name:      user.Name,
		Website:   website,
		Image:     getGravatar(user.Email)}
}
