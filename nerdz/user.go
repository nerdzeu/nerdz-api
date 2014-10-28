package nerdz

import (
	"errors"
	"fmt"
	"github.com/nerdzeu/nerdz-api/utils"
	"net/mail"
	"net/url"
	"reflect"
	"strings"
	"time"
)

// NewUser initializes a User struct
func NewUser(id uint64) (user *User, e error) {
	user = new(User)
	db.First(user, id)
	db.Find(&user.Profile, id)
	if user.Counter != id || user.Profile.Counter != id {
		return nil, errors.New("Invalid id")
	}

	return user, nil
}

// Begin *Numeric* Methods

// NumericBlacklist returns a slice containing the counters (IDs) of blacklisted user
func (user *User) NumericBlacklist() (blacklist []uint64) {
	db.Model(Blacklist{}).Where(&Blacklist{From: user.Counter}).Pluck("\"to\"", &blacklist)
	return
}

// NumericBlacklisting returns a slice  containing the IDs of users that puts user (*User) in their blacklist
func (user *User) NumericBlacklisting() (blacklist []uint64) {
	db.Model(Blacklist{}).Where(&Blacklist{To: user.Counter}).Pluck("\"from\"", &blacklist)
	return
}

// NumericFollowers returns a slice containing the IDs of User that are user's followers
func (user *User) NumericFollowers() (followers []uint64) {
	db.Model(UserFollow{}).Where(UserFollow{To: user.Counter}).Pluck("\"from\"", &followers)
	return
}

// NumericFollowing returns a slice containing the IDs of User that user (User *) is following
func (user *User) NumericFollowing() (following []uint64) {
	db.Model(UserFollow{}).Where(&UserFollow{From: user.Counter}).Pluck("\"to\"", &following)
	return
}

// NumericWhitelist returns a slice containing the IDs of users that are in user whitelist
func (user *User) NumericWhitelist() []uint64 {
	var whitelist []uint64
	db.Model(Whitelist{}).Where(Whitelist{From: user.Counter}).Pluck("\"to\"", &whitelist)
	return append(whitelist, user.Counter)
}

// NumericProjects returns a slice containng the IDs of the projects owned by user
func (user *User) NumericProjects() (projects []uint64) {
	db.Model(ProjectOwner{}).Where(ProjectOwner{From: user.Counter}).Pluck("\"to\"", &projects)
	return
}

// End *Numeric* Methods

// PersonalInfo returns a *PersonalInfo struct
func (user *User) PersonalInfo() *PersonalInfo {
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
		Gravatar:  utils.Gravatar(user.Email),
		Interests: strings.Split(user.Profile.Interests, "\n"),
		Quotes:    strings.Split(user.Profile.Quotes, "\n"),
		Biography: user.Profile.Biography}
}

// ContactInfo returns a *ContactInfo struct
func (user *User) ContactInfo() *ContactInfo {
	// Errors should never occurs, since values are stored in db after have been controlled
	yahoo, _ := mail.ParseAddress(user.Profile.Yahoo)
	website, _ := url.Parse(user.Profile.Website)
	github, _ := url.Parse(user.Profile.Github)
	facebook, _ := url.Parse(user.Profile.Facebook)
	twitter, _ := url.Parse(user.Profile.Twitter)

	// Set Address.Name field
	emailName := user.Surname + " " + user.Name
	// yahoo address can be nil
	if yahoo != nil {
		yahoo.Name = emailName
	}

	return &ContactInfo{
		Website:  website,
		GitHub:   github,
		Skype:    user.Profile.Skype,
		Jabber:   user.Profile.Jabber,
		Yahoo:    yahoo,
		Facebook: facebook,
		Twitter:  twitter,
		Steam:    user.Profile.Steam}
}

// BoardInfo returns a *BoardInfo struct
func (user *User) BoardInfo() *BoardInfo {

	defaultTemplate := Template{
		Name:   Configuration.Templates[user.Profile.Template],
		Number: user.Profile.Template}

	mobileTemplate := Template{
		Name:   Configuration.Templates[user.Profile.MobileTemplate],
		Number: user.Profile.MobileTemplate}

	return &BoardInfo{
		Language:       user.BoardLang,
		Template:       &defaultTemplate,
		MobileTemplate: &mobileTemplate,
		Dateformat:     user.Profile.Dateformat,
		IsClosed:       user.Profile.Closed,
		Private:        user.Private,
		WhiteList:      user.Whitelist()}
}

// Whitelist returns a slice of users that are in the user whitelist
func (user *User) Whitelist() []*User {
	return Users(user.NumericWhitelist())
}

// Followers returns a slice of User that are user's followers
func (user *User) Followers() []*User {
	return Users(user.NumericFollowers())
}

// Following returns a slice of User that user (User *) is following
func (user *User) Following() []*User {
	return Users(user.NumericFollowing())
}

// Blacklist returns a slice of users that user (*User) put in his blacklist
func (user *User) Blacklist() []*User {
	return Users(user.NumericBlacklist())
}

// Blacklisting returns a slice of users that puts user (*User) in their blacklist
func (user *User) Blacklisting() []*User {
	return Users(user.NumericBlacklisting())
}

// Projects returns a slice of projects owned by the user
func (user *User) Projects() []*Project {
	return Projects(user.NumericProjects())
}

// ProjectHome returns a slice of ProjectPost selected by options
func (user *User) ProjectHome(options *PostlistOptions) *[]ProjectPost {
	var projectPost ProjectPost
	posts := projectPost.TableName()
	users := new(User).TableName()
	projects := new(Project).TableName()
	members := new(ProjectMember).TableName()
	owners := new(ProjectOwner).TableName()

	query := db.Model(projectPost).Select(posts + ".*").Order("hpid DESC").
		Joins("JOIN " + users + " ON " + users + ".counter = " + posts + ".from " +
		"JOIN " + projects + " ON " + projects + ".counter = " + posts + ".to " +
		"JOIN " + owners + " ON " + owners + ".to = " + posts + ".to")

	blacklist := user.NumericBlacklist()
	if len(blacklist) != 0 {
		query = query.Where(posts+".from NOT IN (?)", blacklist)
	}
	query = query.Where("( visible IS TRUE OR "+owners+".from = ? OR ( ? IN (SELECT \"from\" FROM "+members+" WHERE \"to\" = "+posts+".to) ) )", user.Counter, user.Counter)

	if options != nil {
		options.User = false
	} else {
		options = new(PostlistOptions)
		options.User = false
	}

	query = postlistQueryBuilder(query, options, user)

	var projectPosts []ProjectPost
	query.Find(&projectPosts)
	return &projectPosts
}

// UsertHome returns a slice of UserPost specified by options
func (user *User) UserHome(options *PostlistOptions) *[]UserPost {
	var userPost UserPost

	query := db.Model(userPost).Select(userPost.TableName() + ".*").Order("hpid DESC")
	blacklist := user.NumericBlacklist()
	if len(blacklist) != 0 {
		query = query.Where("(\"to\" NOT IN (?))", blacklist)
	}

	if options != nil {
		options.User = true
	} else {
		options = new(PostlistOptions)
		options.User = true
	}

	query = postlistQueryBuilder(query, options, user)

	var posts []UserPost
	query.Find(&posts)
	return &posts
}

//Implements Board interface

//Info returns a *Info struct
func (user *User) Info() *Info {
	website, _ := url.Parse(user.Profile.Website)

	return &Info{
		Id:               user.Counter,
		Owner:            user,
		NumericOwner:     user.Counter,
		Followers:        user.Followers(),
		NumericFollowers: user.NumericFollowers(),
		Name:             user.Name,
		Website:          website,
		Image:            utils.Gravatar(user.Email),
		Closed:           user.Profile.Closed}
}

// Postlist returns the specified slice of post on the user board
func (user *User) Postlist(options *PostlistOptions) interface{} {
	users := new(User).TableName()
	posts := new(UserPost).TableName()

	query := db.Model(UserPost{}).Select(posts+".*").Order("hpid DESC").
		Joins("JOIN "+users+" ON "+users+".counter = "+posts+".to").
		Where("(\"to\" = ?)", user.Counter)
	if options != nil {
		options.User = true
	} else {
		options = new(PostlistOptions)
		options.User = true
	}

	var userPosts []UserPost
	query = postlistQueryBuilder(query, options, user)
	query.Find(&userPosts)
	return userPosts
}

// User actions

// An User can Add a new message
func (user *User) Add(message newMessage) (uint64, error) {
	message.SetSender(user.Counter)
	switch message.(type) {
	case *UserPost:
		post := message.(*UserPost)
		if post.To <= 0 {
			post.To = user.Counter
		}
		if err := NewMessage(post, user.Counter, post.To, post.Text()); err != nil {
			return 0, err
		}

		err := db.Save(post).Error
		return post.Hpid, err

	case *ProjectPost:
		post := message.(*ProjectPost)
		if err := NewMessage(post, user.Counter, post.To, post.Text()); err != nil {
			return 0, err
		}

		err := db.Save(post).Error
		return post.Hpid, err

	case *UserPostComment:
		comment := message.(*UserPostComment)
		if err := NewMessage(comment, user.Counter, comment.Hpid, comment.Text()); err != nil {
			return 0, err
		}

		err := db.Save(comment).Error
		return comment.Hcid, err

	case *ProjectPostComment:
		comment := message.(*ProjectPostComment)
		if err := NewMessage(comment, user.Counter, comment.Hpid, comment.Text()); err != nil {
			return 0, err
		}

		err := db.Save(comment).Error
		return comment.Hcid, err
	}

	return 0, fmt.Errorf("Invalid parameter type: %s", reflect.TypeOf(message))
}

// An User can delete an existing message
func (user *User) Delete(message existingMessage) error {
	if user.canDelete(message) {
		return db.Delete(message).Error
	}
	return errors.New("You can't delete this message")
}

// An User can edit an existing message
func (user *User) Edit(message existingMessage) (err error) {
	if user.canEdit(message) {
		return db.Save(message).Error
	}
	return errors.New("You can't edit this message")
}
