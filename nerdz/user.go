package nerdz

import (
	"errors"
	"fmt"
	"net/mail"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/nerdzeu/nerdz-api/utils"
)

// NewUser initializes a User struct
func NewUser(id uint64) (user *User, e error) {
	user = new(User)
	Db().First(user, id)
	Db().First(&user.Profile, id)
	if user.Counter != id || user.Profile.Counter != id {
		return nil, errors.New("Invalid id")
	}

	return user, nil
}

// Login initializes a User struct if login (id | email | username) and password are correct
func Login(login, password string) (*User, error) {
	var email *mail.Address
	var username string
	var id uint64
	var e error

	type Us struct {
		Username string
	}
	var usernameS Us

	if email, e = mail.ParseAddress(login); e == nil { // is a mail
		if e = Db().Model(User{}).Select("username").Where(&User{Email: email.Address}).Scan(&usernameS); e != nil {
			return nil, e
		}
		username = usernameS.Username
	} else if id, e = strconv.ParseUint(login, 10, 64); e == nil { // if login the user ID
		if e = Db().Model(User{}).Select("username").Where(&User{Counter: id}).Scan(&usernameS); e != nil {
			return nil, e
		}
		username = usernameS.Username
	} else { // otherwise is the username
		username = login
	}

	type loginFields struct {
		Logged  bool
		Counter uint64
	}

	var result loginFields
	if e = Db().Model(User{}).Select("login(?, ?) AS logged, counter", username, password).Where("LOWER(username) = ?", username).Scan(&result); e != nil {
		return nil, errors.New("wrong username or password")
	}

	if !result.Logged {
		return nil, errors.New("wrong username or password")
	}

	return NewUser(result.Counter)
}

// Begin *Numeric* Methods

// NumericBlacklist returns a slice containing the counters (IDs) of blacklisted user
func (user *User) NumericBlacklist() (blacklist []uint64) {
	Db().Model(Blacklist{}).Where(&Blacklist{From: user.Counter}).Pluck("\"to\"", &blacklist)
	return
}

// NumericBlacklisting returns a slice  containing the IDs of users that puts user (*User) in their blacklist
func (user *User) NumericBlacklisting() (blacklist []uint64) {
	Db().Model(Blacklist{}).Where(&Blacklist{To: user.Counter}).Pluck("\"from\"", &blacklist)
	return
}

// NumericFollowers returns a slice containing the IDs of User that are user's followers
func (user *User) NumericFollowers() (followers []uint64) {
	Db().Model(UserFollower{}).Where(UserFollower{To: user.Counter}).Pluck("\"from\"", &followers)
	return
}

// NumericFollowing returns a slice containing the IDs of User that user (User *) is following
func (user *User) NumericFollowing() (following []uint64) {
	Db().Model(UserFollower{}).Where(&UserFollower{From: user.Counter}).Pluck("\"to\"", &following)
	return
}

// NumericWhitelist returns a slice containing the IDs of users that are in user whitelist
func (user *User) NumericWhitelist() []uint64 {
	var whitelist []uint64
	Db().Model(Whitelist{}).Where(Whitelist{From: user.Counter}).Pluck("\"to\"", &whitelist)
	return append(whitelist, user.Counter)
}

// NumericWhitelisting returns a slice containing thr IDs of users that whitelisted the user
func (user *User) NumericWhitelisting() (whitelisting []uint64) {
	Db().Model(Whitelist{}).Where(Whitelist{To: user.Counter}).Pluck("\"from\"", &whitelisting)
	return
}

// NumericProjects returns a slice containing the IDs of the projects owned by user
func (user *User) NumericProjects() (projects []uint64) {
	Db().Model(ProjectOwner{}).Where(ProjectOwner{From: user.Counter}).Pluck("\"to\"", &projects)
	return
}

// End *Numeric* Methods

// Interests returns a []string of user interests
func (user *User) Interests() (interests []string) {
	Db().Model(Interest{}).Where(Interest{From: user.Counter}).Pluck("\"value\"", &interests)
	return
}

// PersonalInfo returns a *PersonalInfo struct
func (user *User) PersonalInfo() *PersonalInfo {
	return &PersonalInfo{
		Username:  user.Username,
		IsOnline:  user.Viewonline && user.Last.Add(time.Duration(5)*time.Minute).After(time.Now()),
		Nation:    user.Lang,
		Timezone:  user.Timezone,
		Name:      user.Name,
		Surname:   user.Surname,
		Gender:    user.Gender,
		Birthday:  user.BirthDate,
		Gravatar:  utils.Gravatar(user.Email),
		Interests: user.Interests(),
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
	emailName := user.Name + " " + user.Surname
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
		Whitelist:      user.Whitelist()}
}

// Whitelist returns a slice of users that are in the user whitelist
func (user *User) Whitelist() []*User {
	return Users(user.NumericWhitelist())
}

// Whitelisting returns a slice of users that whitelisted the user
func (user *User) Whitelisting() []*User {
	return Users(user.NumericWhitelisting())
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

	query := Db().Model(projectPost).Order("hpid DESC").
		Joins("JOIN " + users + " ON " + users + ".counter = " + posts + ".from " +
			"JOIN " + projects + " ON " + projects + ".counter = " + posts + ".to " +
			"JOIN " + owners + " ON " + owners + ".to = " + posts + ".to")

	query = query.Where("("+posts+".\"from\" NOT IN (SELECT \"to\" FROM blacklist WHERE \"from\" = ?))", user.Counter)
	query = query.Where("( visible IS TRUE OR "+owners+".from = ? OR ( ? IN (SELECT \"from\" FROM "+members+" WHERE \"to\" = "+posts+".to) ) )", user.Counter, user.Counter)

	if options != nil {
		options.User = false
	} else {
		options = new(PostlistOptions)
		options.User = false
	}

	query = postlistQueryBuilder(query, options, user)

	var projectPosts []ProjectPost
	query.Scan(&projectPosts)

	return &projectPosts
}

// UserHome returns a slice of UserPost specified by options
func (user *User) UserHome(options *PostlistOptions) *[]UserPost {
	var userPost UserPost

	query := Db().Model(userPost).Order("hpid DESC")
	query = query.Where("("+UserPost{}.TableName()+".\"to\" NOT IN (SELECT \"to\" FROM blacklist WHERE \"from\" = ?))", user.Counter)

	if options != nil {
		options.User = true
	} else {
		options = new(PostlistOptions)
		options.User = true
	}

	query = postlistQueryBuilder(query, options, user)

	var posts []UserPost
	query.Scan(&posts)

	return &posts
}

// Pms returns a slice of Pm, representing the list of the last messages exchanged with other users
func (user *User) Pms(otherUser uint64, options *PmConfig) (*[]Pm, error) {
	var pms []Pm

	if options.Offset != 0 {
		Db().Offset(int(options.Offset))
	}

	if options.Limit != 0 {
		Db().Limit(int(options.Limit))
	}

	// Checks if is required ascendant or descendant order of visualization
	if options.DescOrder {
		Db().Order("pmid DESC")
	} else {
		Db().Order("pmid ASC")
	}

	err := Db().Model(Pm{}).Where("((\"from\" = ? AND \"to\" = ?) OR (\"from\" = ? AND \"to\" = ?)) ",
		user.Counter, otherUser, otherUser, user.Counter).Scan(&pms)

	return &pms, err
}

// ThumbUp express a positive preference for a post or comment
func (user *User) ThumbUp(message existingMessage) error {
	switch message.(type) {
	case *UserPost:
		post := message.(*UserPost)
		return Db().Create(&UserPostThumb{Hpid: post.Hpid, From: user.Counter, To: post.To, Vote: 1})

	case *ProjectPost:
		post := message.(*ProjectPost)
		return Db().Create(&ProjectPostThumb{Hpid: post.Hpid, From: user.Counter, To: post.To, Vote: 1})

	case *UserPostComment:
		comment := message.(*UserPostComment)
		return Db().Create(&UserPostCommentThumb{Hcid: comment.Hcid, User: user.Counter, Vote: 1})

	case *ProjectPostComment:
		comment := message.(*ProjectPostComment)
		return Db().Create(&ProjectPostCommentThumb{Hcid: comment.Hcid, From: user.Counter, To: comment.To, Vote: 1})

	case *Pm:
		return fmt.Errorf("TODO(galeone): No preference for private message")
	}

	return fmt.Errorf("Invalid parameter type: %s", reflect.TypeOf(message))
}

// ThumbDown  express a negative preference for a post or comment
func (user *User) ThumbDown(message existingMessage) error {
	switch message.(type) {
	case *UserPost:
		post := message.(*UserPost)
		return Db().Create(&UserPostThumb{Hpid: post.Hpid, From: user.Counter, To: post.To, Vote: -1})

	case *ProjectPost:
		post := message.(*ProjectPost)
		return Db().Create(&ProjectPostThumb{Hpid: post.Hpid, From: user.Counter, To: post.To, Vote: -1})

	case *UserPostComment:
		comment := message.(*UserPostComment)
		return Db().Create(&UserPostCommentThumb{Hcid: comment.Hcid, User: user.Counter, Vote: -1})

	case *ProjectPostComment:
		comment := message.(*ProjectPostComment)
		return Db().Create(&ProjectPostCommentThumb{Hcid: comment.Hcid, From: user.Counter, To: comment.To, Vote: -1})

	case *Pm:
		return fmt.Errorf("TODO(galeone): No preference for private message")
	}

	return fmt.Errorf("Invalid parameter type: %s", reflect.TypeOf(message))
}

// Conversations returns all the private conversations done by the user
func (user *User) Conversations() (*[]Conversation, error) {
	query := "SELECT DISTINCT otherid as \"from\", MAX(times) as time, to_read " +
		"FROM (" +
		"(SELECT MAX(\"time\") AS times, \"from\" as otherid, to_read FROM pms WHERE \"to\" = ? GROUP BY \"from\", to_read)" +
		" UNION " +
		"(SELECT MAX(\"time\") AS times, \"to\" as otherid, FALSE AS to_read FROM pms WHERE \"from\" = ? GROUP BY \"to\", to_read)" +
		") AS tmp GROUP BY otherid, to_read ORDER BY to_read DESC, \"time\" DESC"

	var convList []Conversation

	err := Db().Raw(query, user.Counter, user.Counter).Scan(&convList)

	return &convList, err
}

//Implements Board interface

//Info returns a *info struct
func (user *User) Info() *Info {
	website, _ := url.Parse(user.Profile.Website)
	gravaURL := utils.Gravatar(user.Email)
	boardURL, _ := url.Parse(Configuration.NERDZHost)
	boardURL.Path = user.Username + "."

	return &Info{
		ID:          user.Counter,
		Owner:       nil,
		Name:        user.Name,
		Username:    user.Username,
		Website:     website,
		Image:       gravaURL,
		Closed:      user.Profile.Closed,
		BoardString: boardURL.String(),
		Type:        USER}
}

//Postlist returns the specified slice of post on the user board
func (user *User) Postlist(options *PostlistOptions) *[]ExistingPost {
	users := new(User).TableName()
	posts := new(UserPost).TableName()

	query := Db().Model(UserPost{}).Order("hpid DESC").
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
	query.Scan(&userPosts)

	var retPosts []ExistingPost

	for _, p := range userPosts {
		userPost := p
		retPosts = append(retPosts, ExistingPost(&userPost))
	}

	return &retPosts
}

// User actions

// Add a newMessage
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

		return Db().Create(post)

	case *ProjectPost:
		post := message.(*ProjectPost)
		if err := createMessage(post, user.Counter, post.To, post.Text(), post.Language()); err != nil {
			return err
		}

		return Db().Create(post)

	case *UserPostComment:
		comment := message.(*UserPostComment)
		if err := createMessage(comment, user.Counter, comment.Hpid, comment.Text(), comment.Language()); err != nil {
			return err
		}

		return Db().Create(comment)

	case *ProjectPostComment:
		comment := message.(*ProjectPostComment)
		if err := createMessage(comment, user.Counter, comment.Hpid, comment.Text(), comment.Language()); err != nil {
			return err
		}

		return Db().Create(comment)

	case *Pm:
		pm := message.(*Pm)
		if err := createMessage(pm, user.Counter, pm.To, pm.Text(), pm.Language()); err != nil {
			return err
		}
		return Db().Create(pm)
	}

	return fmt.Errorf("Invalid parameter type: %s", reflect.TypeOf(message))
}

// Delete an existing message
func (user *User) Delete(message existingMessage) error {
	if user.canDelete(message) {
		return Db().Delete(message)
	}
	return errors.New("You can't delete this message")
}

// Edit an existing message
func (user *User) Edit(message editingMessage) error {
	if user.canEdit(message) {
		rollBackText := message.Text() //unencoded
		if err := updateMessage(message); err != nil {
			message.SetText(rollBackText)
			return err
		}
		if err := Db().Updates(message); err != nil {
			message.SetText(rollBackText)
			return err
		}
		return nil
	}
	return errors.New("You can't edit this message")
}

// Follow creates a new "follow" relationship between the current user
// and another NERDZ board. The board could represent a NERDZ's project
// or another NERDZ's user.
func (user *User) Follow(board Board) error {
	if board == nil {
		return errors.New("Unable to follow an undefined board!")
	}

	switch board.(type) {
	case *User:
		otherUser := board.(*User)
		return Db().Create(&UserFollower{From: user.Counter, To: otherUser.Counter})

	case *Project:
		otherProj := board.(*Project)
		return Db().Create(&ProjectFollower{From: user.Counter, To: otherProj.Counter})

	}

	return errors.New("Invalid follower type " + reflect.TypeOf(board).String())
}

// Unfollow delete a "follow" relationship between the current user
// and another NERDZ board. The board could represent a NERDZ's project
// or another NERDZ's user.
func (user *User) Unfollow(board Board) error {
	if board == nil {
		return errors.New("Unable to unfollow an undefined board!")
	}

	switch board.(type) {
	case *User:
		otherUser := board.(*User)
		return Db().Where(&UserFollower{From: user.Counter, To: otherUser.Counter}).Delete(UserFollower{})

	case *Project:
		otherProj := board.(*Project)
		return Db().Where(&ProjectFollower{From: user.Counter, To: otherProj.Counter}).Delete(ProjectFollower{})

	}

	return errors.New("Invalid follower type " + reflect.TypeOf(board).String())
}

// Bookmark bookmarks the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the
// DBMS
func (user *User) Bookmark(post ExistingPost) error {
	if post == nil {
		return errors.New("Unable to bookmark undefined post!")
	}

	switch post.(type) {
	case *UserPost:
		userPost := post.(*UserPost)

		return Db().Create(&UserPostBookmark{From: user.Counter, Hpid: userPost.Hpid})

	case *ProjectPost:
		projectPost := post.(*ProjectPost)

		return Db().Create(&ProjectPostBookmark{From: user.Counter, Hpid: projectPost.Hpid})
	}

	return errors.New("Invalid post type " + reflect.TypeOf(post).String())
}

// Unbookmark the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the DBMS
func (user *User) Unbookmark(post ExistingPost) error {
	if post == nil {
		return errors.New("Unable to unbookmark undefined post!")
	}

	switch post.(type) {
	case *UserPost:
		userPost := post.(*UserPost)

		return Db().Where(&UserPostBookmark{From: user.Counter, Hpid: userPost.Hpid}).Delete(UserPostBookmark{})

	case *ProjectPost:
		projectPost := post.(*ProjectPost)

		return Db().Where(&ProjectPostBookmark{From: user.Counter, Hpid: projectPost.Hpid}).Delete(ProjectPostBookmark{})
	}

	return errors.New("Invalid post type " + reflect.TypeOf(post).String())
}

// AddInterest adds the specified interest. An error is returned if the
// interests already exists or some DBMS contraint is violated
func (user *User) AddInterest(interest *Interest) error {
	interest.From = user.Counter
	if interest.Value == "" {
		return errors.New("Invalid interest value: (empty)")
	}
	return Db().Create(interest)
}

// DeleteInterest removes the specified interest (by its ID or its Value).
func (user *User) DeleteInterest(interest *Interest) error {
	var toDelete Interest
	if interest.ID <= 0 {
		if interest.Value == "" {
			return errors.New("Invalid interest ID and empty interest")
		} else {
			toDelete.Value = interest.Value
		}
	} else {
		toDelete.ID = interest.ID
	}

	if interest.From != user.Counter {
		return errors.New("You can't remove other user interests")
	}

	toDelete.From = interest.From

	return Db().Where(&toDelete).Delete(Interest{})
}

// Friends returns the current user's friends
func (user *User) Friends() *[]User {
	var friends []User
	var follow UserFollower

	Db().Table(follow.TableName() + " f").Joins(", " + follow.TableName() + " f1" + ", " + user.TableName() + " u").
		Where(fmt.Sprintf("f.from = '%d' AND f1.to = '%d' AND f.to = f1.from AND u.counter = f.to",
			user.Counter, user.Counter)).Scan(&friends)

	return &friends

}

// Implements Reference interface

// ID returns the user ID
func (user *User) ID() uint64 {
	return user.Counter
}

// Language returns the user language
func (user *User) Language() string {
	return user.Lang
}
