/*
Copyright (C) 2016-2020 Paolo Galeone <nessuno@nerdz.eu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

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

	"github.com/labstack/gommon/log"
	"github.com/nerdzeu/nerdz-api/utils"
)

// NewUser returns the user with the specified id
func NewUser(id uint64) (*User, error) {
	return NewUserWhere(&User{Counter: id})
}

// NewUserWhere returns the first user that matches the description
func NewUserWhere(description *User) (user *User, e error) {
	user = new(User)
	if e = Db().Where(description).Scan(user); e != nil {
		return
	}

	if e = Db().First(&user.Profile, user.ID()); e != nil {
		return
	}

	return
}

// Login initializes a User struct if login (id | email | username) and password are correct
func Login(login, password string) (*User, error) {
	var email *mail.Address
	var username string
	var id uint64
	var e error

	if email, e = mail.ParseAddress(login); e == nil { // is a mail
		if e = Db().Model(User{}).Select("username").Where(&User{Email: email.Address}).Scan(&username); e != nil {
			return nil, e
		}
	} else if id, e = strconv.ParseUint(login, 10, 64); e == nil { // if login the user ID
		if e = Db().Model(User{}).Select("username").Where(&User{Counter: id}).Scan(&username); e != nil {
			return nil, e
		}
	} else { // otherwise is the username
		username = login
	}

	var logged bool
	var counter uint64

	if e = Db().Model(User{}).Select("login(?, ?) AS logged, counter", username, password).Where("LOWER(username) = ?", username).Scan(&logged, &counter); e != nil {
		return nil, e
	}

	if !logged {
		return nil, errors.New("wrong username or password")
	}

	return NewUser(counter)
}

// Begin *Numeric* Methods

// NumericBlacklist returns a slice containing the counters (IDs) of blacklisted user
func (user *User) NumericBlacklist() (blacklist []uint64) {
	_ = Db().Model(Blacklist{}).Where(&Blacklist{From: user.ID()}).Pluck(`"to"`, &blacklist)
	return
}

// NumericBlacklisting returns a slice  containing the IDs of users that puts user (*User) in their blacklist
func (user *User) NumericBlacklisting() (blacklist []uint64) {
	_ = Db().Model(Blacklist{}).Where(&Blacklist{To: user.ID()}).Pluck(`"from"`, &blacklist)
	return
}

// NumericFollowers returns a slice containing the IDs of User that are user's followers
func (user *User) NumericFollowers() (followers []uint64) {
	_ = Db().Model(UserFollower{}).Where(UserFollower{To: user.ID()}).Pluck(`"from"`, &followers)
	return
}

// NumericUserFollowing returns a slice containing the IDs of User that user (User *) is following
func (user *User) NumericUserFollowing() (following []uint64) {
	_ = Db().Model(UserFollower{}).Where(&UserFollower{From: user.ID()}).Pluck(`"to"`, &following)
	return
}

// NumericProjectFollowing returns a slice containing the IDs of Project that user (User *) is following
func (user *User) NumericProjectFollowing() (following []uint64) {
	_ = Db().Model(ProjectFollower{}).Where(&ProjectFollower{From: user.ID()}).Pluck(`"to"`, &following)
	return
}

// NumericFriends returns a slice containing the IDs of Users that are user's friends (follows each other)
func (user *User) NumericFriends() (friends []uint64) {
	_ = Db().Raw(`SELECT "to" FROM (
		select "to" from followers where "from" = ?) as f
		inner join
		(select "from" from followers where "to" = ?) as e
		on f.to = e.from
		inner join users u on u.counter = f.to`, user.ID(), user.ID()).Scan(&friends)
	return
}

// NumericWhitelist returns a slice containing the IDs of users that are in user whitelist
func (user *User) NumericWhitelist() []uint64 {
	var whitelist []uint64
	_ = Db().Model(Whitelist{}).Where(Whitelist{From: user.ID()}).Pluck(`"to"`, &whitelist)
	return append(whitelist, user.ID())
}

// NumericWhitelisting returns a slice containing thr IDs of users that whitelisted the user
func (user *User) NumericWhitelisting() (whitelisting []uint64) {
	_ = Db().Model(Whitelist{}).Where(Whitelist{To: user.ID()}).Pluck(`"from"`, &whitelisting)
	return
}

// NumericProjects returns a slice containing the IDs of the projects owned by user
func (user *User) NumericProjects() (projects []uint64) {
	_ = Db().Model(ProjectOwner{}).Where(ProjectOwner{From: user.ID()}).Pluck(`"to"`, &projects)
	return
}

// End *Numeric* Methods

// Interests returns a []string of user interests
func (user *User) Interests() (interests []string) {
	_ = Db().Model(Interest{}).Where(Interest{From: user.ID()}).Pluck(`"value"`, &interests)
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
		Birthday:  user.BirthDate,
		Gravatar:  utils.Gravatar(user.Email),
		Interests: user.Interests(),
		Quotes:    strings.Split(user.Profile.Quotes, "\n"),
		Biography: user.Profile.Biography}
}

// ContactInfo returns a *ContactInfo struct
func (user *User) ContactInfo() *ContactInfo {
	// Errors should never occurs, since values are stored in db after have been controlled
	telegram, _ := url.Parse(user.Profile.Telegram)
	website, _ := url.Parse(user.Profile.Website)
	github, _ := url.Parse(user.Profile.Github)
	facebook, _ := url.Parse(user.Profile.Facebook)
	twitter, _ := url.Parse(user.Profile.Twitter)

	return &ContactInfo{
		Website:  website,
		GitHub:   github,
		Skype:    user.Profile.Skype,
		Jabber:   user.Profile.Jabber,
		Telegram: telegram,
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

// UserFollowing returns a slice of User that user (User *) is following
func (user *User) UserFollowing() []*User {
	return Users(user.NumericUserFollowing())
}

// ProjectFollowing returns a slice of Project that user (User *) is following
func (user *User) ProjectFollowing() []*Project {
	return Projects(user.NumericProjectFollowing())
}

// Blacklist returns a slice of users that user (*Project) put in his blacklist
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
func (user *User) ProjectHome(options PostlistOptions) *[]ProjectPost {
	var projectPost ProjectPost

	query := Db().Model(projectPost).Order("hpid DESC")
	query = projectPostlistConditions(query, user)

	options.Model = projectPost
	query = postlistQueryBuilder(query, options, user)

	var projectPosts []ProjectPost
	if err := query.Scan(&projectPosts); err != nil {
		log.Errorf("(Postlist) Error in query.Scan: %s", err)
	}

	return &projectPosts
}

// UserHome returns a slice of UserPost specified by options
func (user *User) UserHome(options PostlistOptions) *[]UserPost {
	var userPost UserPost

	query := Db().Model(userPost).Order("hpid DESC")
	query = query.Where("("+UserPost{}.TableName()+`."to" NOT IN (SELECT "to" FROM blacklist WHERE "from" = ?))`, user.ID())

	options.Model = userPost
	query = postlistQueryBuilder(query, options, user)

	var posts []UserPost
	if err := query.Scan(&posts); err != nil {
		log.Errorf("(UserHome) Error in query.Scan: %s", err)
	}
	return &posts
}

// Home returns a slice of Post representing the user home. Posts are
// filtered by specified options.
func (user *User) Home(options PostlistOptions) *[]Message {
	var message Message
	query := Db().
		CTE(`WITH blist AS (SELECT "to" FROM blacklist WHERE "from" = ?)`, user.ID()). // WITH cte
		Table(message.TableName()).                                                    // select * from messages
		Where(`"from" NOT IN (SELECT * FROM blist) AND
		CASE type
		WHEN 1 THEN "to" NOT IN (SELECT * FROM blist)
		ELSE ( -- groups conditions
			TRUE IN (SELECT visible FROM groups g WHERE g.counter = "to")
			OR
			(? IN (
				SELECT "from" FROM groups_members gm WHERE gm."to" = "to"
				UNION ALL
				SELECT "from" FROM groups_owners go WHERE go."to" = "to")
			)
		)
		END`, user.ID()).
		Order("time DESC")

	options.Model = message
	query = postlistQueryBuilder(query, options, user) // handle following, followers, language, newer, older, between...
	var posts []Message
	if err := query.Scan(&posts); err != nil {
		log.Errorf("(Home) Error in query.Scan: %s", err)
	}
	return &posts
}

// Pms returns a slice of Pm, representing the list of the last messages exchanged with other users
func (user *User) Pms(otherUser uint64, options PmsOptions) (*[]Pm, error) {
	var pms []Pm

	query := Db().Model(Pm{}).Where(
		`("from" = ? AND "to" = ?) OR ("from" = ? AND "to" = ?)`,
		user.ID(), otherUser, otherUser, user.ID())
	// build query in function of parameters
	query = pmsQueryBuilder(query, options)

	e := query.Scan(&pms)
	return &pms, e
}

// Vote express a positive/negative preference for a post or comment.
// Returns the vote if everything went ok
func (user *User) Vote(message existingMessage, vote int8) (Vote, error) {
	method := Db().Create
	if vote > 0 {
		vote = 1
	} else if vote == 0 {
		vote = 0
		method = Db().Delete
	} else {
		vote = -1
	}
	var err error
	switch message := message.(type) {
	case *UserPost:
		dbVote := UserPostVote{Hpid: message.ID(), From: user.ID(), To: message.To, Vote: vote}
		err = method(&dbVote)
		return &dbVote, err

	case *ProjectPost:
		dbVote := ProjectPostVote{Hpid: message.ID(), From: user.ID(), To: message.To, Vote: vote}
		err = method(&dbVote)
		return &dbVote, err

	case *UserPostComment:
		dbVote := UserPostCommentVote{Hcid: message.Hcid, From: user.ID(), Vote: vote}
		err = method(&dbVote)
		return &dbVote, err

	case *ProjectPostComment:
		dbVote := ProjectPostCommentVote{Hcid: message.Hcid, From: user.ID(), To: message.To, Vote: vote}
		err = method(&dbVote)
		return &dbVote, err

	case *Pm:
		return nil, fmt.Errorf("TODO(galeone): No preference for private message")
	}

	return nil, fmt.Errorf("Invalid parameter type: %s", reflect.TypeOf(message))
}

// Conversations returns all the private conversations done by the user
func (user *User) Conversations() (*[]Conversation, error) {
	var convList []Conversation
	err := Db().Raw(`WITH conversations_with_duplicates AS (
		SELECT DISTINCT ?::bigint AS me, otherid, MAX(times) as "time", to_read FROM (
			SELECT MAX("time") AS times, "from" as otherid, to_read FROM pms WHERE "to" = ? GROUP BY "from", to_read
			UNION
			SELECT MAX("time") AS times, "to" as otherid, FALSE AS to_read FROM pms WHERE "from" = ? GROUP BY "to", to_read
		) AS tmp GROUP BY otherid, to_read
	)
	SELECT c.me, c.otherid, p.message, MAX(c."time") AS t, c.to_read
	FROM conversations_with_duplicates c
	INNER JOIN pms p
	ON c."time" = p."time" AND (
		(c.me = p."from" AND c.otherid = p."to")
		OR
		(c.me = p."to" AND c.otherid = p."from")
	)
	GROUP BY c.me, c.otherid, p.message, c.to_read
	ORDER BY to_read DESC, t DESC`, user.ID(), user.ID(), user.ID()).Scan(&convList)
	return &convList, err
}

// DeleteConversation deletes the conversation of user with other user
func (user *User) DeleteConversation(other uint64) error {
	return Db().Where(`("from" = ? AND "to" = ?) OR ("from" = ? AND "to" = ?)`, user.ID(), other, other, user.ID()).Delete(&Pm{})
}

//Implements Board interface

// Info returns a *info struct
func (user *User) Info() *Info {
	website, _ := url.Parse(user.Profile.Website)
	gravaURL := utils.Gravatar(user.Email)
	boardURL, _ := url.Parse(Configuration.NERDZHost)
	boardURL.Path = user.Username + "."

	return &Info{
		ID:          user.ID(),
		Owner:       nil,
		Name:        user.Name,
		Username:    user.Username,
		Website:     website,
		Image:       gravaURL,
		Closed:      user.Profile.Closed,
		BoardString: boardURL.String(),
		Type:        UserBoardID}
}

// Postlist returns the specified slice of post on the user board
func (user *User) Postlist(options PostlistOptions) *[]ExistingPost {
	users := User{}.TableName()
	var post UserPost

	query := Db().Model(UserPost{}).Order("hpid DESC").
		Joins("JOIN "+users+" ON "+users+".counter = "+post.TableName()+".to").
		Where(`"to" = ?`, user.ID())

	options.Model = post

	var userPosts []UserPost
	query = postlistQueryBuilder(query, options, user)
	if err := query.Scan(&userPosts); err != nil {
		log.Errorf("(User::Postlist) Error in query.Scan: %s", err)
	}

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
	switch message := message.(type) {
	case *UserPost:
		if message.To == 0 {
			message.To = user.ID()
		}
		if err := createMessage(message, user.ID(), message.To, message.Text(), message.Language()); err != nil {
			return err
		}

		return Db().Create(message)

	case *ProjectPost:
		if err := createMessage(message, user.ID(), message.To, message.Text(), message.Language()); err != nil {
			return err
		}

		return Db().Create(message)

	case *UserPostComment:
		if err := createMessage(message, user.ID(), message.Hpid, message.Text(), message.Language()); err != nil {
			return err
		}

		return Db().Create(message)

	case *ProjectPostComment:
		if err := createMessage(message, user.ID(), message.Hpid, message.Text(), message.Language()); err != nil {
			return err
		}

		return Db().Create(message)

	case *Pm:
		if err := createMessage(message, user.ID(), message.To, message.Text(), message.Language()); err != nil {
			return err
		}
		return Db().Create(message)
	}

	return fmt.Errorf("Invalid parameter type: %s", reflect.TypeOf(message))
}

// Delete an existing message
func (user *User) Delete(message existingMessage) error {
	if user.CanDelete(message) {
		return Db().Delete(message)
	}
	return errors.New("You can't delete this message")
}

// Edit an existing message
func (user *User) Edit(message editingMessage) error {
	if user.CanEdit(message) {
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
	return errors.New("you can't edit this message")
}

// Follow creates a new "follow" relationship between the current user
// and another NERDZ board. The board could represent a NERDZ's project
// or another NERDZ's user.
func (user *User) Follow(board Board) error {
	if board == nil {
		return errors.New("unable to follow an undefined board!")
	}

	switch board := board.(type) {
	case *User:
		return Db().Create(&UserFollower{From: user.ID(), To: board.ID()})

	case *Project:
		return Db().Create(&ProjectFollower{From: user.ID(), To: board.ID()})

	}

	return errors.New("Invalid follower type " + reflect.TypeOf(board).String())
}

// WhitelistUser add other user to the user whitelist
func (user *User) WhitelistUser(other *User) error {
	if other == nil {
		return errors.New("other user should be a vaid user")
	}
	return Db().Create(&Whitelist{From: user.ID(), To: other.ID()})
}

// UnwhitelistUser removes other user to the user whitelist
func (user *User) UnwhitelistUser(other *User) error {
	if other == nil {
		return errors.New("other user should be a vaid user")
	}
	return Db().Where(&Whitelist{From: user.ID(), To: other.ID()}).Delete(Whitelist{})
}

// BlacklistUser add other user to the user blacklist
func (user *User) BlacklistUser(other *User, motivation string) error {
	if other == nil {
		return errors.New("other user should be a vaid user")
	}
	return Db().Create(&Blacklist{From: user.ID(), To: other.ID(), Motivation: motivation})
}

// UnblacklistUser removes other user to the user blacklist
func (user *User) UnblacklistUser(other *User) error {
	if other == nil {
		return errors.New("other user should be a vaid user")
	}
	return Db().Where(&Blacklist{From: user.ID(), To: other.ID()}).Delete(Blacklist{})
}

// Unfollow delete a "follow" relationship between the current user
// and another NERDZ board. The board could represent a NERDZ's project
// or another NERDZ's user.
func (user *User) Unfollow(board Board) error {
	if board == nil {
		return errors.New("unable to unfollow an undefined board!")
	}

	switch board := board.(type) {
	case *User:
		return Db().Where(&UserFollower{From: user.ID(), To: board.ID()}).Delete(UserFollower{})

	case *Project:
		return Db().Where(&ProjectFollower{From: user.ID(), To: board.ID()}).Delete(ProjectFollower{})

	}

	return errors.New("Invalid follower type " + reflect.TypeOf(board).String())
}

// Bookmark bookmarks the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the
// DBMS
func (user *User) Bookmark(post ExistingPost) (Bookmark, error) {
	if post == nil {
		return nil, errors.New("unable to bookmark undefined post")
	}

	switch post := post.(type) {
	case *UserPost:
		bookmark := UserPostBookmark{From: user.ID(), Hpid: post.ID()}
		err := Db().Create(&bookmark)
		return &bookmark, err

	case *ProjectPost:
		bookmark := ProjectPostBookmark{From: user.ID(), Hpid: post.ID()}
		err := Db().Create(&bookmark)
		return &bookmark, err
	}

	return nil, errors.New("Invalid post type " + reflect.TypeOf(post).String())
}

// Unbookmark the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the DBMS
func (user *User) Unbookmark(post ExistingPost) error {
	if post == nil {
		return errors.New("unable to unbookmark undefined post")
	}

	switch post := post.(type) {
	case *UserPost:
		return Db().Where(&UserPostBookmark{From: user.ID(), Hpid: post.ID()}).Delete(UserPostBookmark{})

	case *ProjectPost:
		return Db().Where(&ProjectPostBookmark{From: user.ID(), Hpid: post.ID()}).Delete(ProjectPostBookmark{})
	}

	return errors.New("Invalid post type " + reflect.TypeOf(post).String())
}

// Lurk lurkes the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the
// DBMS
func (user *User) Lurk(post ExistingPost) (Lurk, error) {
	if post == nil {
		return nil, errors.New("unable to lurk undefined post")
	}

	switch post := post.(type) {
	case *UserPost:
		lurk := UserPostLurk{From: user.ID(), Hpid: post.ID()}
		err := Db().Create(&lurk)
		return &lurk, err

	case *ProjectPost:
		lurk := ProjectPostLurk{From: user.ID(), Hpid: post.ID()}
		err := Db().Create(&lurk)
		return &lurk, err
	}

	return nil, errors.New("Invalid post type " + reflect.TypeOf(post).String())
}

// Unlurk the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the DBMS
func (user *User) Unlurk(post ExistingPost) error {
	if post == nil {
		return errors.New("unable to unlurk undefined post")
	}

	switch post := post.(type) {
	case *UserPost:
		return Db().Where(&UserPostLurk{From: user.ID(), Hpid: post.ID()}).Delete(UserPostLurk{})

	case *ProjectPost:
		return Db().Where(&ProjectPostLurk{From: user.ID(), Hpid: post.ID()}).Delete(ProjectPostLurk{})
	}

	return errors.New("Invalid post type " + reflect.TypeOf(post).String())
}

// Lock lockes the specified post. If users are present, indiidual notifications
// are disabled from the user presents in the users list.
func (user *User) Lock(post ExistingPost, users ...*User) (*[]Lock, error) {
	if post == nil {
		return nil, errors.New("unable to lurk undefined post")
	}

	switch post := post.(type) {
	case *UserPost:
		if len(users) == 0 {
			lock := UserPostLock{User: user.ID(), Hpid: post.ID()}
			err := Db().Create(&lock)
			return &[]Lock{&lock}, err
		}
		var locks []Lock
		for _, other := range users {
			lock := UserPostUserLock{From: user.ID(), To: other.ID(), Hpid: post.ID()}
			if err := Db().Create(&lock); err != nil {
				return nil, err
			}
			locks = append(locks, Lock(&lock))
		}
		return &locks, nil

	case *ProjectPost:
		if len(users) == 0 {
			lock := ProjectPostLock{User: user.ID(), Hpid: post.ID()}
			err := Db().Create(&lock)
			return &[]Lock{&lock}, err
		}
		var locks []Lock
		for _, other := range users {
			lock := ProjectPostUserLock{From: user.ID(), To: other.ID(), Hpid: post.ID()}
			if err := Db().Create(&lock); err != nil {
				return nil, err
			}
			locks = append(locks, Lock(&lock))
		}
		return &locks, nil
	}

	return nil, errors.New("Invalid post type " + reflect.TypeOf(post).String())
}

// Unlock the specified post by a specific user. An error is returned if the
// post isn't defined or if there are other errors returned by the DBMS
func (user *User) Unlock(post ExistingPost, users ...*User) error {
	if post == nil {
		return errors.New("unable to unlock undefined post")
	}

	switch post := post.(type) {
	case *UserPost:
		if len(users) == 0 {
			return Db().Where(&UserPostLock{User: user.ID(), Hpid: post.ID()}).Delete(UserPostLock{})
		}
		for _, other := range users {
			err := Db().Where(&UserPostUserLock{From: user.ID(), To: other.ID(), Hpid: post.ID()}).Delete(UserPostUserLock{})
			if err != nil {
				return err
			}
		}
		return nil

	case *ProjectPost:
		if len(users) == 0 {
			return Db().Where(&ProjectPostLock{User: user.ID(), Hpid: post.ID()}).Delete(ProjectPostLock{})
		}
		for _, other := range users {
			err := Db().Where(&ProjectPostUserLock{From: user.ID(), To: other.ID(), Hpid: post.ID()}).Delete(ProjectPostUserLock{})
			if err != nil {
				return err
			}
		}
		return nil
	}

	return errors.New("Invalid post type " + reflect.TypeOf(post).String())
}

// AddInterest adds the specified interest. An error is returned if the
// interests already exists or some DBMS contraint is violated
func (user *User) AddInterest(interest *Interest) error {
	interest.From = user.ID()
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
		}
		toDelete.Value = interest.Value
	} else {
		toDelete.ID = interest.ID
	}

	if interest.From != user.ID() {
		return errors.New("You can't remove other user interests")
	}

	toDelete.From = interest.From

	return Db().Where(&toDelete).Delete(Interest{})
}

// Friends returns the current user's friends
func (user *User) Friends() []*User {
	return Users(user.NumericFriends())
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

// Can* methods

// CanEdit returns true if user can edit the editingMessage
func (user *User) CanEdit(message editingMessage) bool {
	return message.ID() > 0 && message.IsEditable() && utils.InSlice(user.ID(), message.NumericOwners())
}

// CanDelete returns true if user can delete the existingMessage
func (user *User) CanDelete(message existingMessage) bool {
	return message.ID() > 0 && utils.InSlice(user.ID(), message.NumericOwners())
}

// CanBookmark returns true if user haven't bookamrked to existingPost yet
func (user *User) CanBookmark(message ExistingPost) bool {
	return message.ID() > 0 && !utils.InSlice(user.ID(), message.NumericBookmarkers())
}

// CanLurk returns true if the user haven't lurked the existingPost yet
func (user *User) CanLurk(message ExistingPost) bool {
	return message.ID() > 0 && !utils.InSlice(user.ID(), message.NumericLurkers())
}

// CanComment returns true if the user can comment to the existingPost
func (user *User) CanComment(message ExistingPost) bool {
	return !utils.InSlice(user.ID(), message.Sender().NumericBlacklist()) && message.ID() > 0 && !message.IsClosed()
}

// CanSee returns true if the user can see the Board content
func (user *User) CanSee(board Board) bool {
	switch board := board.(type) {
	case *User:
		return !utils.InSlice(user.ID(), board.NumericBlacklist())

	case *Project:
		if board.Visible {
			return true
		}

		return user.ID() == board.NumericOwner() || utils.InSlice(user.ID(), board.NumericMembers())
	}
	return false
}
