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
	Number uint8
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
	return blacklist
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
func (user *User) NumericProjects() []uint64 {
	var projects []uint64
	db.Model(ProjectOwner{}).Where(ProjectOwner{From: user.Counter}).Pluck("\"to\"", &projects)
	return projects
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

// BoardInfo returns a BoardInfo struct
func (user *User) BoardInfo() *BoardInfo {
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

// User can add a post on the board of an other user
// The paremeter other can be a *User or an id. The news parameter is optional and
// if present and equals to true, the post will be marked as news
func (user *User) AddUserPost(other interface{}, message string, news ...bool) (uint64, error) {
	post := new(UserPost)
	if err := NewMessageInit(post, other, message); err != nil {
		return 0, err
	}

	post.News = len(news) > 0 && news[0]
	post.From = user.Counter

	err := db.Save(post).Error
	fmt.Printf("HPID: %d\n", post.Hpid)
	return post.Hpid, err
}

func (user *User) EditUserPost(post *UserPost) (err error) {
	if user.canEdit(post) {
		fmt.Println("CAN EDIT THIS SHIT")
		return db.Save(post).Error
	}
	fmt.Println("CANT EDIT THIS SHIT")
	return errors.New("You can't edit this post")
}

// User cam remove a post (if he has the right permissions)
// The parameter post can be a *UserPost or an uint64 (representing the post hpid)
func (user *User) DeleteUserPost(post interface{}) error {
	var hpid uint64

	switch post.(type) {
	case uint64:
		hpid = post.(uint64)
	case *UserPost:
		hpid = (post.(*UserPost)).Hpid
	default:
		return fmt.Errorf("Invalid post type: %v. Allowed uint64 and *UserPost", reflect.TypeOf(post))
	}

	return db.Where(UserPost{Hpid: hpid}).Delete(UserPost{}).Error
}

// User can add a post on a project.
// The paremeter other can be a *Prject or its id. The news parameter is optional and
// if present and equals to true, the post will be marked as news
func (user *User) AddProjectPost(other interface{}, message string, news ...bool) (uint64, error) {
	post := new(ProjectPost)

	if err := NewMessageInit(post, other, message); err != nil {
		return 0, err
	}

	post.News = len(news) > 0 && news[0]
	post.From = user.Counter

	err := db.Save(post).Error
	return post.Hpid, err
}

// User can remove a post (if he has the right permissions)
// The parameter post can be a *ProjectPost or an uint64 (representing the post hpid)
func (user *User) DeleteProjectPost(post interface{}) error {
	var hpid uint64

	switch post.(type) {
	case uint64:
		hpid = post.(uint64)
	case *ProjectPost:
		hpid = (post.(*ProjectPost)).Hpid
	default:
		return fmt.Errorf("Invalid post type: %v. Allowed uint64 and *ProjectPost", reflect.TypeOf(post))
	}

	return db.Where(ProjectPost{Hpid: hpid}).Delete(ProjectPost{}).Error
}

// User can comment posts on profile
// The parameter other can be a *UserPost or its id.
func (user *User) AddUserPostComment(other interface{}, message string) (uint64, error) {
	comment := new(UserPostComment)

	if err := NewMessageInit(comment, other, message); err != nil {
		return 0, err
	}

	comment.From = user.Counter
	var to struct {
		To uint64
	}
	db.Select("\"to\"").Model(UserPost{}).Where(&UserPost{Hpid: comment.Hpid}).Scan(&to)
	comment.To = to.To

	err := db.Save(comment).Error
	return comment.Hcid, err
}

// User can remov a comment (if he hash the right permissions)
// The comment parameter can be a *UserPostComment or an uint64 (hidden commen id: hcid)
func (user *User) DeleteUserPostComment(comment interface{}) error {
	var hcid uint64

	switch comment.(type) {
	case uint64:
		hcid = comment.(uint64)
	case *UserPostComment:
		hcid = (comment.(*UserPostComment)).Hcid
	default:
		return fmt.Errorf("Invalid comment type: %v. Allowed uint64 and *UserPostComment", reflect.TypeOf(comment))
	}

	return db.Where(UserPostComment{Hcid: hcid}).Delete(UserPostComment{}).Error
}

// User can comment posts on profile
// The parameter other can be a *UserPost or its id.
func (user *User) AddProjectPostComment(other interface{}, message string) (uint64, error) {
	comment := new(ProjectPostComment)

	if err := NewMessageInit(comment, other, message); err != nil {
		return 0, err
	}

	comment.From = user.Counter
	var to struct {
		To uint64
	}
	db.Select("\"to\"").Model(ProjectPost{}).Where(&ProjectPost{Hpid: comment.Hpid}).Scan(&to)
	comment.To = to.To

	err := db.Save(comment).Error
	return comment.Hcid, err
}

// User can remov a comment (if he hash the right permissions)
// The comment parameter can be a *UserPostComment or an uint64 (hidden commen id: hcid)
func (user *User) DeleteProjectPostComment(comment interface{}) error {
	var hcid uint64

	switch comment.(type) {
	case uint64:
		hcid = comment.(uint64)
	case *ProjectPostComment:
		hcid = (comment.(*ProjectPostComment)).Hcid
	default:
		return fmt.Errorf("Invalid comment type: %v. Allowed uint64 and *ProjectPostComment", reflect.TypeOf(comment))
	}

	return db.Where(ProjectPostComment{Hcid: hcid}).Delete(ProjectPostComment{}).Error
}
