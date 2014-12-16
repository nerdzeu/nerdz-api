package nerdz

import (
	"errors"
	"fmt"
	"net/mail"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/nerdzeu/nerdz-api/utils"
)

// PersonalInfo is the struct that contains all the personal info of an user
type PersonalInfo struct {
	Id        uint64
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
type Template struct {
	Number uint8
	Name   string
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
	db.Model(UserFollower{}).Where(UserFollower{To: user.Counter}).Pluck("\"from\"", &followers)
	return
}

// NumericFollowing returns a slice containing the IDs of User that user (User *) is following
func (user *User) NumericFollowing() (following []uint64) {
	db.Model(UserFollower{}).Where(&UserFollower{From: user.Counter}).Pluck("\"to\"", &following)
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

// Pms returns a slice of Pm, representing the list of the last messages exchanged with other users
func (user *User) Pms() *[]Pm {
	var pms []Pm
	// TODO: extract last message, make this raw query generic
	db.Raw("SELECT DISTINCT " +
		"EXTRACT(EPOCH FROM MAX(times)) as lasttime, otherid as \"from\", to_read " +
		"FROM ( " +
		"SELECT MAX(\"time\") AS times, \"from\" as otherid, to_read " +
		"FROM pms WHERE \"to\" = ? GROUP BY \"from\", to_read " +
		"UNION " +
		"SELECT MAX(\"time\") AS times, \"to\" as otherid, FALSE AS to_read " +
		"FROM pms WHERE \"from\" = ? GROUP BY \"to\", to_read " +
		") as tmp " +
		"GROUP BY otherid, to_read " +
		"ORDER BY to_read DESC, \"lasttime\" DESC").Scan(&pms)

	return &pms
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
func (user *User) Add(message newMessage) error {
	switch message.(type) {
	case *UserPost:
		post := message.(*UserPost)
		if post.To == 0 {
			post.To = user.Counter
		}
		if err := createMessage(post, user.Counter, post.To, post.Text(), post.Language()); err != nil {
			return err
		}

		return db.Create(post).Error

	case *ProjectPost:
		post := message.(*ProjectPost)
		if err := createMessage(post, user.Counter, post.To, post.Text(), post.Language()); err != nil {
			return err
		}

		return db.Create(post).Error

	case *UserPostComment:
		comment := message.(*UserPostComment)
		if err := createMessage(comment, user.Counter, comment.Hpid, comment.Text(), comment.Language()); err != nil {
			return err
		}

		return db.Create(comment).Error

	case *ProjectPostComment:
		comment := message.(*ProjectPostComment)
		if err := createMessage(comment, user.Counter, comment.Hpid, comment.Text(), comment.Language()); err != nil {
			return err
		}

		return db.Create(comment).Error

	case *Pm:
		pm := message.(*Pm)
		if err := createMessage(pm, user.Counter, pm.To, pm.Text(), pm.Language()); err != nil {
			return err
		}
		return db.Create(pm).Error
	}

	return fmt.Errorf("Invalid parameter type: %s", reflect.TypeOf(message))
}

// An User can delete an existing message
func (user *User) Delete(message existingMessage) error {
	if user.canDelete(message) {
		return db.Delete(message).Error
	}
	return errors.New("You can't delete this message")
}

// An User can edit an existing message
func (user *User) Edit(message editingMessage) error {
	if user.canEdit(message) {
		rollBackText := message.Text() //unencoded
		message.ClearDefaults()
		if err := updateMessage(message); err != nil {
			message.SetText(rollBackText)
			return err
		}
		if err := db.Save(message).Error; err != nil {
			message.SetText(rollBackText)
			return err
		}
		return nil
	}
	return errors.New("You can't edit this message")
}

// Create a new "follow" relationship between the current user
// and another NERDZ board. The board could represent a NERDZ's project
// or another NERDZ's user.
func (user *User) Follow(board Board) error {
	if board == nil {
		return errors.New("Unable to follow an undefined board!")
	}

	switch board.(type) {
	case *User:
		otherUser := board.(*User)
		return db.Create(&UserFollower{From: user.Counter, To: otherUser.Counter}).Error

	case *Project:
		otherProj := board.(*Project)
		return db.Create(&ProjectFollower{From: user.Counter, To: otherProj.Counter}).Error

	}

	return errors.New("Invalid follower type " + reflect.TypeOf(board).String())
}

// Delete a "follow" relationship between the current user
// and another NERDZ board. The board could represent a NERDZ's project
// or another NERDZ's user.
func (user *User) Unfollow(board Board) error {
	if board == nil {
		return errors.New("Unable to unfollow an undefined board!")
	}

	switch board.(type) {
	case *User:
		otherUser := board.(*User)
		return db.Where(&UserFollower{From: user.Counter, To: otherUser.Counter}).Delete(UserFollower{}).Error

	case *Project:
		otherProj := board.(*Project)
		return db.Where(&ProjectFollower{From: user.Counter, To: otherProj.Counter}).Delete(ProjectFollower{}).Error

	}

	return errors.New("Invalid follower type " + reflect.TypeOf(board).String())
}

// Bookmarks the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the
// DBMS
func (user *User) Bookmark(post existingPost) error {
	if post == nil {
		return errors.New("Unable to bookmark undefined post!")
	}

	switch post.(type) {
	case *UserPost:
		userPost := post.(*UserPost)

		return db.Create(&UserPostBookmark{From: user.Counter, Hpid: userPost.Hpid}).Error

	case *ProjectPost:
		projectPost := post.(*ProjectPost)

		return db.Create(&ProjectPostBookmark{From: user.Counter, Hpid: projectPost.Hpid}).Error
	}

	return errors.New("Invalid post type " + reflect.TypeOf(post).String())
}

// Bookmarks the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the
// DBMS
func (user *User) Unbookmark(post existingPost) error {
	if post == nil {
		return errors.New("Unable to unbookmark undefined post!")
	}

	switch post.(type) {
	case *UserPost:
		userPost := post.(*UserPost)

		return db.Where(&UserPostBookmark{From: user.Counter, Hpid: userPost.Hpid}).Delete(UserPostBookmark{}).Error

	case *ProjectPost:
		projectPost := post.(*ProjectPost)

		return db.Where(&ProjectPostBookmark{From: user.Counter, Hpid: projectPost.Hpid}).Delete(ProjectPostBookmark{}).Error
	}

	return errors.New("Invalid post type " + reflect.TypeOf(post).String())
}
