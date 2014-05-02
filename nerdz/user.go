package nerdz

import (
	"errors"
	"github.com/nerdzeu/nerdz-api/utils"
	"net/mail"
	"net/url"
	"strings"
	"time"
)

// PersonalInfo is the struct that contains all the personal info of an user
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

// ContactInfo is the struct that contains all the contact info of an user
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

// Template is the representation of a nerdz website template
// Note: Template.Name is unimplemented at the moment and is always ""
type Template struct {
	Number int16
	Name   string //TODO
}

// BoardInfo is that struct that contains all the informations related to the user's board
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

// Begin *Numeric* Methods

// GetNumericBlacklist returns a slice containing the counters (IDs) of blacklisted user
func (user *User) GetNumericBlacklist() []int64 {
	return append(user.GetNumericBlacklisted(), user.GetNumericBlacklisting()...)
}

// GetNumericBlacklisted returns a slice containing the IDs of users that user (*User) put in his blacklist
func (user *User) GetNumericBlacklisted() []int64 {
	var blacklist []int64
	db.Model(Blacklist{}).Where(&Blacklist{From: user.Counter}).Pluck("\"to\"", &blacklist)
	return blacklist
}

// GetNumericBlacklisting returns a slice  containing the IDs of users that puts user (*User) in their blacklist
func (user *User) GetNumericBlacklisting() []int64 {
	var blacklist []int64
	db.Model(Blacklist{}).Where(&Blacklist{To: user.Counter}).Pluck("\"from\"", &blacklist)
	return blacklist
}

// GetNumericFollowers returns a slice containing the IDs of User that are user's followers
func (user *User) GetNumericFollowers() []int64 {
	var followers []int64
	db.Model(UserFollow{}).Where(UserFollow{To: user.Counter}).Pluck("\"from\"", &followers)
	return followers
}

// GetNumericFollowing returns a slice containing the IDs of User that user (User *) is following
func (user *User) GetNumericFollowing() []int64 {
	var following []int64
	db.Model(UserFollow{}).Where(&UserFollow{From: user.Counter}).Pluck("\"to\"", &following)
	return following
}

// GetNumericWhitelist returns a slice containing the IDs of users that are in user whitelist
func (user *User) GetNumericWhitelist() []int64 {
	var whitelist []int64
	db.Model(Whitelist{}).Where(Whitelist{From: user.Counter}).Pluck("\"to\"", &whitelist)
	return append(whitelist, user.Counter)
}

// GetNumericProjects returns a slice containng the IDs of the projects owned by user
func (user *User) GetNumericProjects() []int64 {
	var projects []int64
	db.Model(Project{}).Where(Project{Owner: user.Counter}).Pluck("counter", &projects)
	return projects
}

// End *Numeric* Methods

// GetPersonalInfo returns a *PersonalInfo struct
func (user *User) GetPersonalInfo() *PersonalInfo {
	return &PersonalInfo{
		Id:        user.Counter,
		Username:  user.Username,
		IsOnline:  user.Viewonline && user.Last.Add(time.Duration(5)*time.Minute).After(time.Now()),
		Nation:    user.Lang,
		Timezone:  user.Timezone,
		Name:      user.Name,
		Surname:   user.Surname,
		Gender:    user.Gender,
		Birthday:  user.BirthDate,
		Gravatar:  utils.GetGravatar(user.Email),
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
		Name:   "", //TODO: find a way to Get template name -> unfortunately isn't stored in the database at the moment
		Number: user.Profile.Template}

	mobileTemplate := Template{
		Name:   "", //TODO: find a way to Get template name -> unfortunately isn't stored in the database at the moment
		Number: user.Profile.MobileTemplate}

	return &BoardInfo{
		Language:       user.BoardLang,
		Template:       &defaultTemplate,
		MobileTemplate: &mobileTemplate,
		Dateformat:     user.Profile.Dateformat,
		IsClosed:       user.IsClosed(),
		Private:        user.Private,
		WhiteList:      user.GetWhitelist()}
}

// GetWhitelist returns a slice of users that are in the user whitelist
func (user *User) GetWhitelist() []*User {
	return getUsers(user.GetNumericWhitelist())
}

// GetFollowers returns a slice of User that are user's followers
func (user *User) GetFollowers() []*User {
	return getUsers(user.GetNumericFollowers())
}

// GetFollowing returns a slice of User that user (User *) is following
func (user *User) GetFollowing() []*User {
	return getUsers(user.GetNumericFollowing())
}

// GetBlacklisted returns a slice of users that user (*User) put in his blacklist
func (user *User) GetBlacklisted() []*User {
	return getUsers(user.GetNumericBlacklisted())
}

// GetBlacklisting returns a slice of users that puts user (*User) in their blacklist
func (user *User) GetBlacklisting() []*User {
	return getUsers(user.GetNumericBlacklisting())
}

// GetBlacklist returns the union of Blacklisted and Blacklisting users
func (user *User) GetBlacklist() []*User {
	return append(user.GetBlacklisted(), user.GetBlacklisting()...)
}

// GetProjects returns a slice of projects owned by the user
func (user *User) GetProjects() []*Project {
	return getProjects(user.GetNumericProjects())
}

// GetProjectHome returns a slice of ProjectPost selected by options
func (user *User) GetProjectHome(options *PostlistOptions) *[]ProjectPost {
	var projectPost ProjectPost
	projectPosts := projectPost.TableName()
	users := new(User).TableName()
	projects := new(Project).TableName()
	projectMembers := new(ProjectMember).TableName()

	query := db.Model(projectPost).
		Order("hpid DESC").
		// Pre-parsing options is not required for project, since
		Joins("JOIN " + users + " ON " + users + ".counter = " + projectPosts + ".from JOIN " + projects + " ON " + projects + ".counter = " + projectPosts + ".to")
	blacklist := user.GetNumericBlacklist()
	if len(blacklist) != 0 {
		query = query.Not("from", blacklist)
	}
	query = query.Where("( visible IS TRUE OR owner = ? OR ( ? IN (SELECT \"user\" FROM "+projectMembers+" WHERE \"group\" = "+projectPosts+".to) ) )", user.Counter, user.Counter)

	query = postlistQueryBuilder(query, options, user)

	var posts []ProjectPost
	query.Find(&posts)
	return &posts
}

// GetUsertHome returns a slice of UserPost specified by options
func (user *User) GetUserHome(options *PostlistOptions) *[]UserPost {
	var userPost UserPost
	users := new(User).TableName()
	query := db.Model(userPost).Order("hpid DESC")

	//Pre-parsing options to determinate fields to join
	join := "JOIN " + users + " ON " + users + ".counter = " + userPost.TableName() + "."
	if options != nil && (options.Following || options.Followers) {
		//Join with "from" user, since we need to know the language of who's posting
		join += "from"
	} else {
		// Join with "to" user, since we don't need to know the language of who's posting (general homepage postlist in a specified language - or without language)
		join += "to"
	}
	query = query.Joins(join)

	blacklist := user.GetNumericBlacklist()
	if len(blacklist) != 0 {
		query = query.Where("(\"to\" NOT IN (?) OR \"from\" NOT IN (?))", blacklist, blacklist)
	}

	query = postlistQueryBuilder(query, options, user)

	var posts []UserPost
	query.Find(&posts)
	return &posts
}

//Implements Board interface

//GetInfo returns a *Info struct
func (user *User) GetInfo() *Info {
	website, _ := url.Parse(user.Profile.Website)

	return &Info{
		Id:        user.Counter,
		Owner:     user,
		Followers: user.GetFollowers(),
		Name:      user.Name,
		Website:   website,
		Image:     utils.GetGravatar(user.Email)}
}

// GetPostlist returns the specified slice of post on the user board
func (user *User) GetPostlist(options *PostlistOptions) interface{} {
	var posts []UserPost
	var userPost UserPost
	users := new(User).TableName()
	query := db.Model(userPost).Order("hpid DESC").
		Joins("JOIN "+users+" ON "+users+".counter = "+userPost.TableName()+".to"). //PostListOptions.Language support
		Where("(\"to\" = ?)", user.Counter)
	query = postlistQueryBuilder(query, options, user)
	query.Find(&posts)
	return posts
}

// IsClosed returns a boolean indicating if the board is closed
func (user *User) IsClosed() bool {
	var closedProfile ClosedProfile
	db.First(&closedProfile, user.Counter)
	return closedProfile.Counter == user.Counter
}

// User actions

// User can add a post on the board of an other user
// The paremeter other can be a *User or an id
func (user *User) AddUserPost(other interface{}, message string) error {
	var e error

	post := new(UserPost)

	if e = post.SetMessage(message); e != nil {
		return e
	}

	if _ , e = post.SetTo(other); e != nil {
		return e
	}

	post.From = user.Counter

	return db.Save(post).Error
}

// User can add a post on a project.
// The paremeter other can be a *Prject or its id. The news parameter is optional and
// if present and equals to true, the post will be marked as news
func (user *User) AddProjectPost(other interface{}, message string, news ...bool) error {
	var e error

	post := new(ProjectPost)

	if e = post.SetMessage(message); e != nil {
		return e
	}

	_ , e = post.SetTo(other)
	if e != nil {
		return e
	}

	post.News = len(news) > 0 && news[0]
	post.From = user.Counter

	return db.Save(post).Error

}
