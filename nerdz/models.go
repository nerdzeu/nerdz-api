/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

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
	"database/sql"
	"github.com/galeone/igor"
	"strings"
	"time"
)

// Enrich models structure with unexported types

// boardType represents a board type
type boardType string

const (
	// UserBoardID constant (of type boardType) makes possible to distinguish a User
	// board from a Project board
	UserBoardID boardType = "user"
	// ProjectBoardID constant (of type boardType) makes possible to distinguish a PROJECT
	// board from a User board
	ProjectBoardID boardType = "project"
)

// Transferable represents a common interface for all the
// types defined by the backend that are able to generate
// a data structure that can be returned by the API
type Transferable interface {
	// GetTO returns a proper data structure for the API
	// users (max len 1) is the current user viewing the API (the one who grants
	// the proprer access to the client application)
	// Is optional because not all the TO are different (ore have different content)
	// if the user is a specific user
	GetTO(users ...*User) interface{}
}

// Models

// UserPostsNoNotify is the model for the relation posts_no_notify
type UserPostsNoNotify struct {
	User    uint64
	Hpid    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (u *UserPostsNoNotify) GetTO(users ...*User) *UserPostsNoNotifyTO {
	var userInfo *InfoTO
	if user, e := NewUser(u.User); e == nil {
		userInfo = user.Info().GetTO()
	}
	return &UserPostsNoNotifyTO{
		original: u,
		User:     userInfo,
		Hpid:     u.Hpid,
		Time:     u.Time,
		Counter:  u.Counter,
	}
}

// TableName returns the table name associated with the structure
func (UserPostsNoNotify) TableName() string {
	return "posts_no_notify"
}

// UserPostCommentsNoNotify is the model for the relation comments_no_notify
type UserPostCommentsNoNotify struct {
	From    uint64
	To      uint64
	Hpid    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserPostCommentsNoNotify) TableName() string {
	return "comments_no_notify"
}

// GetTO returns its Transfer Object
func (u *UserPostCommentsNoNotify) GetTO(users ...*User) *UserPostCommentsNoNotifyTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(u.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(u.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &UserPostCommentsNoNotifyTO{
		original: u,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Hpid:     u.Hpid,
		Time:     u.Time,
		Counter:  u.Counter,
	}
}

// UserPostCommentsNotify is the model for the relation comments_notify
type UserPostCommentsNotify struct {
	From    uint64
	To      uint64
	Hpid    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (u *UserPostCommentsNotify) GetTO(users ...*User) *UserPostCommentsNotifyTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(u.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(u.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &UserPostCommentsNotifyTO{
		original: u,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Hpid:     u.Hpid,
		Time:     u.Time,
		Counter:  u.Counter,
	}
}

// TableName returns the table name associated with the structure
func (UserPostCommentsNotify) TableName() string {
	return "comments_notify"
}

// Ban is the model for the relation ban
type Ban struct {
	User       uint64
	Motivation string
	Time       time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter    uint64    `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (b *Ban) GetTO(users ...*User) *BanTO {
	var userInfo *InfoTO
	if user, e := NewUser(b.User); e == nil {
		userInfo = user.Info().GetTO()
	}
	return &BanTO{
		original:   b,
		User:       userInfo,
		Motivation: b.Motivation,
		Time:       b.Time,
		Counter:    b.Counter,
	}
}

// TableName returns the table name associated with the structure
func (Ban) TableName() string {
	return "ban"
}

// Blacklist is the model for the relation blacklist
type Blacklist struct {
	From       uint64
	To         uint64
	Motivation string
	Time       time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter    uint64    `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (b *Blacklist) GetTO(users ...*User) *BlacklistTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(b.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(b.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &BlacklistTO{
		original:   b,
		FromInfo:   fromInfo,
		ToInfo:     toInfo,
		Motivation: b.Motivation,
		Time:       b.Time,
		Counter:    b.Counter,
	}
}

// TableName returns the table name associated with the structure
func (Blacklist) TableName() string {
	return "blacklist"
}

// Whitelist is the model for the relation whitelist
type Whitelist struct {
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (w *Whitelist) GetTO(users ...*User) *WhitelistTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(w.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(w.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &WhitelistTO{
		original: w,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Time:     w.Time,
		Counter:  w.Counter,
	}
}

// TableName returns the table name associated with the structure
func (Whitelist) TableName() string {
	return "whitelist"
}

// UserFollower is the model for the relation followers
type UserFollower struct {
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	ToNotify bool
	Counter  uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserFollower) TableName() string {
	return "followers"
}

// GetTO returns its Transfer Object
func (u *UserFollower) GetTO(users ...*User) *UserFollowerTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(u.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(u.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &UserFollowerTO{
		original: u,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Time:     u.Time,
		ToNotify: u.ToNotify,
		Counter:  u.Counter,
	}
}

// ProjectNotify is the model for the relation groups_notify
type ProjectNotify struct {
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Hpid    uint64
	Counter uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectNotify) TableName() string {
	return "groups_notify"
}

// GetTO returns its Transfer Object
func (p *ProjectNotify) GetTO(users ...*User) *ProjectNotifyTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewProject(p.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(p.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &ProjectNotifyTO{
		original: p,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Time:     p.Time,
		Hpid:     p.Hpid,
		Counter:  p.Counter,
	}
}

// ProjectPostsNoNotify is the model for the relation groups_posts_no_notify
type ProjectPostsNoNotify struct {
	User    uint64
	Hpid    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (p *ProjectPostsNoNotify) GetTO(users ...*User) *ProjectPostsNoNotifyTO {
	var userInfo *InfoTO
	if user, e := NewUser(p.User); e == nil {
		userInfo = user.Info().GetTO()
	}
	return &ProjectPostsNoNotifyTO{
		original: p,
		User:     userInfo,
		Hpid:     p.Hpid,
		Time:     p.Time,
		Counter:  p.Counter,
	}
}

// TableName returns the table name associated with the structure
func (ProjectPostsNoNotify) TableName() string {
	return "groups_posts_no_notify"
}

// ProjectPostCommentsNoNotify is the model for the relation groups_comments_no_notify
type ProjectPostCommentsNoNotify struct {
	From    uint64
	To      uint64
	Hpid    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (p *ProjectPostCommentsNoNotify) GetTO(users ...*User) *ProjectPostCommentsNoNotifyTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(p.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(p.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &ProjectPostCommentsNoNotifyTO{
		original: p,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Hpid:     p.Hpid,
		Time:     p.Time,
		Counter:  p.Counter,
	}
}

// TableName returns the table name associated with the structure
func (ProjectPostCommentsNoNotify) TableName() string {
	return "groups_comments_no_notify"
}

// ProjectPostCommentsNotify is the model for the relation groups_comments_notify
type ProjectPostCommentsNotify struct {
	From    uint64
	To      uint64
	Hpid    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (p *ProjectPostCommentsNotify) GetTO(users ...*User) *ProjectPostCommentsNotifyTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(p.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(p.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &ProjectPostCommentsNotifyTO{
		original: p,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Hpid:     p.Hpid,
		Time:     p.Time,
		Counter:  p.Counter,
	}
}

// TableName returns the table name associated with the structure
func (ProjectPostCommentsNotify) TableName() string {
	return "groups_comments_notify"
}

// User is the model for the relation users
type User struct {
	Counter          uint64    `igor:"primary_key"`
	Last             time.Time `sql:"default:(now() at time zone 'utc')"`
	NotifyStory      igor.JSON `sql:"default:'{}'::jsonb"`
	Private          bool
	Lang             string
	Username         string
	Password         string
	RemoteAddr       string
	HTTPUserAgent    string `igor:"column:http_user_agent"`
	Email            string
	Name             string
	Surname          string
	Gender           bool
	BirthDate        time.Time `sql:"default:(now() at time zone 'utc')"`
	BoardLang        string
	Timezone         string
	Viewonline       bool
	RegistrationTime time.Time `sql:"default:(now() at time zone 'utc')"`
	// Relation. Manually fill the field when required
	Profile Profile `sql:"-"`
}

// TableName returns the table name associated with the structure
func (User) TableName() string {
	return "users"
}

// GetTO returns its Transfer Object: *User GetTO embeds *Profile GetTO
func (u *User) GetTO(users ...*User) *UserTO {
	return &UserTO{
		original:         u,
		Counter:          u.Counter,
		Last:             u.Last,
		NotifyStory:      u.NotifyStory,
		Private:          u.Private,
		Lang:             u.Lang,
		Username:         u.Username,
		Name:             u.Name,
		Surname:          u.Surname,
		Gender:           u.Gender,
		BirthDate:        u.BirthDate,
		BoardLang:        u.BoardLang,
		Timezone:         u.Timezone,
		Viewonline:       u.Viewonline,
		RegistrationTime: u.RegistrationTime,
		Profile: ProfileTO{
			Counter:        u.Profile.Counter,
			Website:        u.Profile.Website,
			Quotes:         strings.Split(u.Profile.Quotes, "\n"),
			Biography:      u.Profile.Biography,
			Interests:      u.Interests(),
			Github:         u.Profile.Github,
			Skype:          u.Profile.Skype,
			Jabber:         u.Profile.Jabber,
			Yahoo:          u.Profile.Yahoo,
			Userscript:     u.Profile.Userscript,
			Template:       u.Profile.Template,
			MobileTemplate: u.Profile.MobileTemplate,
			Dateformat:     u.Profile.Dateformat,
			Facebook:       u.Profile.Facebook,
			Twitter:        u.Profile.Twitter,
			Steam:          u.Profile.Steam,
			Push:           u.Profile.Push,
			Pushregtime:    u.Profile.Pushregtime,
			Closed:         u.Profile.Closed,
		},
	}
}

// Profile is the model for the relation profiles
type Profile struct {
	Counter        uint64 `igor:"primary_key"`
	Website        string
	Quotes         string
	Biography      string
	Github         string
	Skype          string
	Jabber         string
	Yahoo          string
	Userscript     string
	Template       uint8
	MobileTemplate uint8
	Dateformat     string
	Facebook       string
	Twitter        string
	Steam          string
	Push           bool
	Pushregtime    time.Time `sql:"default:(now() at time zone 'utc')"`
	Closed         bool
}

// TableName returns the table name associated with the structure
func (Profile) TableName() string {
	return "profiles"
}

// Interest is the model for the relation interests
type Interest struct {
	ID    uint64 `igor:"primary_key"`
	From  uint64
	Value string
	Time  time.Time `sql:"default:(now() at time zone 'utc')"`
}

// TableName returns the table name associated with the structure
func (Interest) TableName() string {
	return "interests"
}

// Post is the type of a generic post
type Post struct {
	Hpid    uint64 `igor:"primary_key"`
	From    uint64
	To      uint64
	Pid     uint64 `sql:"default:0"`
	Message string
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Lang    string
	News    bool
	Closed  bool
}

// UserPost converts the Post to a UserPost
func (p *Post) UserPost() *UserPost {
	return &UserPost{
		Post: *p,
	}
}

// ProjectPost converts the Post to ProjectPost
func (p *Post) ProjectPost() *ProjectPost {
	return &ProjectPost{
		Post: *p,
	}
}

// GetTO returns its Transfer Object
func (p *Post) GetTO(users ...*User) *PostTO {
	return &PostTO{
		original: p,
		Hpid:     p.Hpid,
		Pid:      p.Pid,
		Message:  p.Message,
		Time:     p.Time,
		Lang:     p.Lang,
		News:     p.News,
		Closed:   p.Closed,
	}
}

// UserPost is the model for the relation posts
type UserPost struct {
	Post
}

// GetTO returns its Transfer Object
func (p *UserPost) GetTO(users ...*User) *PostTO {
	if len(users) != 1 {
		panic("UserPost.GetTO requires a user parameter")
	}
	user := users[0]
	postTO := p.Post.GetTO()
	postTO.Type = UserBoardID

	if from, e := NewUser(p.From); e == nil {
		postTO.FromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(p.To); e == nil {
		postTO.ToInfo = to.Info().GetTO()
	}

	postTO.Rate = p.Thumbs()
	postTO.RevisionsCount = p.RevisionsNumber()
	postTO.CommentsCount = p.CommentsNumber()
	postTO.BookmarkersCount = p.BookmarkersNumber()
	postTO.LurkersCount = p.LurkersNumber()
	postTO.Timestamp = p.Time.Unix()
	postTO.URL = p.URL().String()
	postTO.CanBookmark = user.CanBookmark(p)
	postTO.CanComment = user.CanComment(p)
	postTO.CanDelete = user.CanDelete(p)
	postTO.CanEdit = user.CanEdit(p)
	postTO.CanLurk = user.CanLurk(p)
	return postTO
}

// TableName returns the table name associated with the structure
func (UserPost) TableName() string {
	return "posts"
}

// UserPostRevision is the model for the relation posts_revisions
type UserPostRevision struct {
	Hpid    uint64
	Message string
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	RevNo   uint16
	Counter uint64 `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (p *UserPostRevision) GetTO(users ...*User) *UserPostRevisionTO {
	return &UserPostRevisionTO{
		original: p,
		Hpid:     p.Hpid,
		Message:  p.Message,
		Time:     p.Time,
		RevNo:    p.RevNo,
		Counter:  p.Counter,
	}
}

// TableName returns the table name associated with the structure
func (UserPostRevision) TableName() string {
	return "posts_revisions"
}

// UserPostThumb is the model for the relation thumbs
type UserPostThumb struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Vote    int8
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserPostThumb) TableName() string {
	return "thumbs"
}

// GetTO returns its Transfer Object
func (t *UserPostThumb) GetTO(users ...*User) *UserPostThumbTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(t.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(t.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &UserPostThumbTO{
		original: t,
		Hpid:     t.Hpid,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Vote:     t.Vote,
		Time:     t.Time,
		Counter:  t.Counter,
	}
}

// UserPostLurker is the model for the relation lurkers
type UserPostLurker struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (l *UserPostLurker) GetTO(users ...*User) *UserPostLurkerTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(l.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(l.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &UserPostLurkerTO{
		original: l,
		Hpid:     l.Hpid,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Time:     l.Time,
		Counter:  l.Counter,
	}
}

// TableName returns the table name associated with the structure
func (UserPostLurker) TableName() string {
	return "lurkers"
}

// UserPostComment is the model for the relation comments
type UserPostComment struct {
	Hcid     uint64 `igor:"primary_key"`
	Hpid     uint64
	From     uint64
	To       uint64
	Message  string
	Lang     string
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	Editable bool      `sql:"default:true"`
}

// GetTO returns its Transfer Object
func (c *UserPostComment) GetTO(users ...*User) *UserPostCommentTO {
	if len(users) != 1 {
		panic("UserPostComment.GetTO requires a user parameter")
	}
	user := users[0]

	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(c.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(c.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &UserPostCommentTO{
		original:  c,
		Hcid:      c.Hcid,
		Hpid:      c.Hpid,
		FromInfo:  fromInfo,
		ToInfo:    toInfo,
		Message:   c.Message,
		Time:      c.Time,
		CanEdit:   user.CanEdit(c),
		CanDelete: user.CanDelete(c),
	}
}

// TableName returns the table name associated with the structure
func (UserPostComment) TableName() string {
	return "comments"
}

// UserPostCommentRevision is the model for the relation comments_revisions
type UserPostCommentRevision struct {
	Hcid    uint64
	Message string
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	RevNo   int8
	Counter uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserPostCommentRevision) TableName() string {
	return "comments_revisions"
}

// GetTO returns its Transfer Object
func (c *UserPostCommentRevision) GetTO(users ...*User) *UserPostCommentRevisionTO {
	return &UserPostCommentRevisionTO{
		original: c,
		Hcid:     c.Hcid,
		Message:  c.Message,
		Time:     c.Time,
		RevNo:    c.RevNo,
		Counter:  c.Counter,
	}
}

// UserPostBookmark is the model for the relation bookmarks
type UserPostBookmark struct {
	Hpid    uint64
	From    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (UserPostBookmark) TableName() string {
	return "bookmarks"
}

// GetTO returns its Transfer Object
func (b *UserPostBookmark) GetTO(users ...*User) *UserPostBookmarkTO {
	var fromInfo *InfoTO
	if from, e := NewUser(b.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	return &UserPostBookmarkTO{
		original: b,
		Hpid:     b.Hpid,
		FromInfo: fromInfo,
		Time:     b.Time,
		Counter:  b.Counter,
	}
}

// Pm is the model for the relation pms
type Pm struct {
	Pmid    uint64 `igor:"primary_key"`
	From    uint64
	To      uint64
	Message string
	Lang    string
	ToRead  bool
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
}

// GetTO returns its Transfer Object
func (p *Pm) GetTO(users ...*User) *PmTO {
	if len(users) != 1 {
		panic("Pm.GetTO requires a user parameter")
	}
	user := users[0]

	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(p.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(p.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &PmTO{
		original:  p,
		Pmid:      p.Pmid,
		FromInfo:  fromInfo,
		ToInfo:    toInfo,
		Message:   p.Message,
		ToRead:    p.ToRead,
		Time:      p.Time,
		CanDelete: user.CanDelete(p),
		CanEdit:   user.CanEdit(p),
	}
}

// TableName returns the table name associated with the structure
func (Pm) TableName() string {
	return "pms"
}

// Project is the model for the relation groups
type Project struct {
	Counter      uint64 `igor:"primary_key"`
	Description  string
	Name         string
	Private      bool
	Photo        sql.NullString
	Website      sql.NullString
	Goal         string
	Visible      bool
	Open         bool
	CreationTime time.Time `sql:"default:(now() at time zone 'utc')"`
}

// GetTO returns its Transfer Object
func (p *Project) GetTO(users ...*User) *ProjectTO {
	return &ProjectTO{
		original:     p,
		Counter:      p.Counter,
		Description:  p.Description,
		Name:         p.Name,
		Private:      p.Private,
		Photo:        p.Photo,
		Website:      p.Website,
		Goal:         p.Goal,
		Visible:      p.Visible,
		Open:         p.Open,
		CreationTime: p.CreationTime,
	}
}

// TableName returns the table name associated with the structure
func (Project) TableName() string {
	return "groups"
}

// ProjectMember is the model for the relation groups_members
type ProjectMember struct {
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	ToNotify bool
	Counter  uint64 `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (m *ProjectMember) GetTO(users ...*User) *ProjectMemberTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(m.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewProject(m.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &ProjectMemberTO{
		original: m,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Time:     m.Time,
		ToNotify: m.ToNotify,
		Counter:  m.Counter,
	}
}

// TableName returns the table name associated with the structure
func (ProjectMember) TableName() string {
	return "groups_members"
}

// ProjectOwner is the model for the relation groups_owners
type ProjectOwner struct {
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	ToNotify bool
	Counter  uint64 `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (o *ProjectOwner) GetTO(users ...*User) *ProjectOwnerTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(o.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewProject(o.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &ProjectOwnerTO{
		original: o,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Time:     o.Time,
		ToNotify: o.ToNotify,
		Counter:  o.Counter,
	}

}

// TableName returns the table name associated with the structure
func (ProjectOwner) TableName() string {
	return "groups_owners"
}

// ProjectPost is the model for the relation groups_posts
type ProjectPost struct {
	Post
}

// TableName returns the table name associated with the structure
func (ProjectPost) TableName() string {
	return "groups_posts"
}

// GetTO returns its Transfer Object
func (p *ProjectPost) GetTO(users ...*User) *PostTO {
	if len(users) != 1 {
		panic("ProjectPost.GetTO requires a user parameter")
	}
	user := users[0]
	postTO := p.Post.GetTO()
	postTO.Type = ProjectBoardID

	if from, e := NewUser(p.From); e == nil {
		postTO.FromInfo = from.Info().GetTO()
	}
	if to, e := NewProject(p.To); e == nil {
		postTO.ToInfo = to.Info().GetTO()
	}

	postTO.Rate = p.Thumbs()
	postTO.RevisionsCount = p.RevisionsNumber()
	postTO.CommentsCount = p.CommentsNumber()
	postTO.BookmarkersCount = p.BookmarkersNumber()
	postTO.LurkersCount = p.LurkersNumber()
	postTO.Timestamp = p.Time.Unix()
	postTO.URL = p.URL().String()
	postTO.CanBookmark = user.CanBookmark(p)
	postTO.CanComment = user.CanComment(p)
	postTO.CanDelete = user.CanDelete(p)
	postTO.CanEdit = user.CanEdit(p)
	postTO.CanLurk = user.CanLurk(p)
	return postTO
}

// ProjectPostRevision is the model for the relation groups_posts_revisions
type ProjectPostRevision struct {
	Hpid    uint64
	Message string
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	RevNo   uint16
	Counter uint64 `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (p *ProjectPostRevision) GetTO(users ...*User) *ProjectPostRevisionTO {
	return &ProjectPostRevisionTO{
		original: p,
		Hpid:     p.Hpid,
		Message:  p.Message,
		Time:     p.Time,
		RevNo:    p.RevNo,
		Counter:  p.Counter,
	}
}

// TableName returns the table name associated with the structure
func (ProjectPostRevision) TableName() string {
	return "groups_posts_revisions"
}

// ProjectPostThumb is the model for the relation groups_thumbs
type ProjectPostThumb struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Vote    int8
	Counter uint64 `igor:"primary_key"`
}

// TableName returns the table name associated with the structure
func (ProjectPostThumb) TableName() string {
	return "groups_thumbs"
}

// GetTO returns its Transfer Object
func (t *ProjectPostThumb) GetTO(users ...*User) *ProjectPostThumbTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(t.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(t.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &ProjectPostThumbTO{
		original: t,
		Hpid:     t.Hpid,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Time:     t.Time,
		Vote:     t.Vote,
		Counter:  t.Counter,
	}
}

// ProjectPostLurker is the model for the relation groups_lurkers
type ProjectPostLurker struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (l *ProjectPostLurker) GetTO(users ...*User) *ProjectPostLurkerTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(l.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewProject(l.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &ProjectPostLurkerTO{
		original: l,
		Hpid:     l.Hpid,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Time:     l.Time,
		Counter:  l.Counter,
	}
}

// TableName returns the table name associated with the structure
func (ProjectPostLurker) TableName() string {
	return "groups_lurkers"
}

// ProjectPostComment is the model for the relation groups_comments
type ProjectPostComment struct {
	Hcid     uint64 `igor:"primary_key"`
	Hpid     uint64
	From     uint64
	To       uint64
	Message  string
	Lang     string
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	Editable bool      `sql:"default:true"`
}

// GetTO returns its Transfer Object
func (c *ProjectPostComment) GetTO(users ...*User) *ProjectPostCommentTO {
	if len(users) != 1 {
		panic("ProjectPostComment.GetTO requires a user parameter")
	}
	user := users[0]

	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(c.From); e == nil {
		fromInfo = from.Info().GetTO()
	}

	if to, e := NewProject(c.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &ProjectPostCommentTO{
		original:  c,
		Hcid:      c.Hcid,
		Hpid:      c.Hpid,
		FromInfo:  fromInfo,
		ToInfo:    toInfo,
		Message:   c.Message,
		Time:      c.Time,
		CanDelete: user.CanDelete(c),
		CanEdit:   user.CanEdit(c),
	}
}

// TableName returns the table name associated with the structure
func (ProjectPostComment) TableName() string {
	return "groups_comments"
}

// ProjectPostCommentRevision is the model for the relation groups_comments_revisions
type ProjectPostCommentRevision struct {
	Hcid    uint64
	Message string
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	RevNo   uint16
	Counter uint64 `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (r *ProjectPostCommentRevision) GetTO(users ...*User) *ProjectPostCommentRevisionTO {
	return &ProjectPostCommentRevisionTO{
		original: r,
		Hcid:     r.Hcid,
		Message:  r.Message,
		Time:     r.Time,
		RevNo:    r.RevNo,
		Counter:  r.Counter,
	}
}

// TableName returns the table name associated with the structure
func (ProjectPostCommentRevision) TableName() string {
	return "groups_comments_revisions"
}

// ProjectPostBookmark is the model for the relation groups_bookmarks
type ProjectPostBookmark struct {
	Hpid    uint64
	From    uint64
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (b *ProjectPostBookmark) GetTO(users ...*User) *ProjectPostBookmarkTO {
	var fromInfo *InfoTO
	if from, e := NewUser(b.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	return &ProjectPostBookmarkTO{
		original: b,
		Hpid:     b.Hpid,
		FromInfo: fromInfo,
		Time:     b.Time,
		Counter:  b.Counter,
	}
}

// TableName returns the table name associated with the structure
func (ProjectPostBookmark) TableName() string {
	return "groups_bookmarks"
}

// ProjectFollower is the model for the relation groups_followers
type ProjectFollower struct {
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	ToNotify bool
	Counter  uint64 `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (p *ProjectFollower) GetTO(users ...*User) *ProjectFollowerTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(p.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewProject(p.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &ProjectFollowerTO{
		original: p,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Time:     p.Time,
		ToNotify: p.ToNotify,
		Counter:  p.Counter,
	}
}

// TableName returns the table name associated with the structure
func (ProjectFollower) TableName() string {
	return "groups_followers"
}

// UserPostCommentThumb is the model for the relation groups_comment_thumbs
type UserPostCommentThumb struct {
	Hcid    uint64
	User    uint64
	Vote    int8
	Counter uint64 `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (t *UserPostCommentThumb) GetTO(users ...*User) *UserPostCommentThumbTO {
	var userInfo *InfoTO
	if user, e := NewUser(t.User); e == nil {
		userInfo = user.Info().GetTO()
	}
	return &UserPostCommentThumbTO{
		original: t,
		Hcid:     t.Hcid,
		User:     userInfo,
		Vote:     t.Vote,
		Counter:  t.Counter,
	}
}

// TableName returns the table name associated with the structure
func (UserPostCommentThumb) TableName() string {
	return "comment_thumbs"
}

// ProjectPostCommentThumb is the model for the relation groups_comment_thumbs
type ProjectPostCommentThumb struct {
	Hcid    uint64
	From    uint64
	To      uint64
	Vote    int8
	Time    time.Time `sql:"default:(now() at time zone 'utc')"`
	Counter uint64    `igor:"primary_key"`
}

// GetTO returns its Transfer Object
func (t *ProjectPostCommentThumb) GetTO(users ...*User) *ProjectPostCommentThumbTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(t.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(t.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &ProjectPostCommentThumbTO{
		original: t,
		Hcid:     t.Hcid,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Vote:     t.Vote,
		Time:     t.Time,
		Counter:  t.Counter,
	}
}

// TableName returns the table name associated with the structure
func (ProjectPostCommentThumb) TableName() string {
	return "groups_comment_thumbs"
}

// DeletedUser is the model for the relation deleted_users
type DeletedUser struct {
	Counter    uint64 `igor:"primary_key"`
	Username   string
	Time       time.Time `sql:"default:(now() at time zone 'utc')"`
	Motivation string
}

// TableName returns the table name associated with the structure
func (DeletedUser) TableName() string {
	return "deleted_users"
}

// GetTO returns its Transfer Object
func (u *DeletedUser) GetTO(users ...*User) *DeletedUserTO {
	return &DeletedUserTO{
		original:   u,
		Counter:    u.Counter,
		Username:   u.Username,
		Time:       u.Time,
		Motivation: u.Motivation,
	}
}

// SpecialUser is the model for the relation special_users
type SpecialUser struct {
	Role    string `igor:"primary_key"`
	Counter uint64
}

// GetTO returns its Transfer Object
func (u *SpecialUser) GetTO(users ...*User) *SpecialUserTO {
	return &SpecialUserTO{
		original: u,
		Role:     u.Role,
		Counter:  u.Counter,
	}
}

// TableName returns the table name associated with the structure
func (SpecialUser) TableName() string {
	return "special_users"
}

// SpecialProject is the model for the relation special_groups
type SpecialProject struct {
	Role    string `igor:"primary_key"`
	Counter uint64
}

// GetTO returns its Transfer Object
func (p *SpecialProject) GetTO(users ...*User) *SpecialProjectTO {
	return &SpecialProjectTO{
		original: p,
		Role:     p.Role,
		Counter:  p.Counter,
	}
}

// TableName returns the table name associated with the structure
func (SpecialProject) TableName() string {
	return "special_groups"
}

// PostClassification is the model for the relation posts_classifications
type PostClassification struct {
	ID    uint64 `igor:"primary_key"`
	UHpid uint64
	GHpid uint64
	Tag   string
}

// GetTO returns its Transfer Object
func (p *PostClassification) GetTO(users ...*User) *PostClassificationTO {
	return &PostClassificationTO{
		original: p,
		ID:       p.ID,
		UHpid:    p.UHpid,
		GHpid:    p.GHpid,
		Tag:      p.Tag,
	}
}

// TableName returns the table name associated with the structure
func (PostClassification) TableName() string {
	return "posts_classifications"
}

// Mention is the model for the relation mentions
type Mention struct {
	ID       uint64 `igor:"primary_key"`
	UHpid    uint64
	GHpid    uint64
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:(now() at time zone 'utc')"`
	ToNotify bool
}

// GetTO returns its Transfer Object
func (m *Mention) GetTO(users ...*User) *MentionTO {
	var fromInfo, toInfo *InfoTO
	if from, e := NewUser(m.From); e == nil {
		fromInfo = from.Info().GetTO()
	}
	if to, e := NewUser(m.To); e == nil {
		toInfo = to.Info().GetTO()
	}
	return &MentionTO{
		original: m,
		ID:       m.ID,
		UHpid:    m.UHpid,
		GHpid:    m.GHpid,
		FromInfo: fromInfo,
		ToInfo:   toInfo,
		Time:     m.Time,
		ToNotify: m.ToNotify,
	}
}

// TableName returns the table name associated with the structure
func (Mention) TableName() string {
	return "mentions"
}

// Message is the model for the view message
type Message struct {
	Post
	Type uint8
}

// TableName returns the table name associated with the structure
func (Message) TableName() string {
	return "messages"
}

// GetTO returns its Transfer Object
func (p *Message) GetTO(users ...*User) *PostTO {
	if p.Type == UserPostID {
		return p.Post.UserPost().GetTO(users...)
	}
	return p.Post.ProjectPost().GetTO(users...)
}

// OAuth2Client implements the osin.Client interface
type OAuth2Client struct {
	// Surrogated key
	ID uint64 `igor:"primary_key"`
	// Real Primary Key. Application (client) name
	Name string `sql:"UNIQUE"`
	// Secret is the unique secret associated with a client
	Secret string `sql:"UNIQUE"`
	// RedirectURI is the valid redirection URI associated with a client
	RedirectURI string
	// UserID references User that created this client
	UserID uint64
}

// TableName returns the table name associated with the structure
func (OAuth2Client) TableName() string {
	return "oauth2_clients"
}

// OAuth2AuthorizeData is the model for the relation oauth2_authorize
// that represents the authorization granted to to the client
type OAuth2AuthorizeData struct {
	// Surrogated key
	ID uint64 `igor:"primary_key"`
	// ClientID references the client that created this token
	ClientID uint64
	// Code is the Authorization code
	Code string
	// CreatedAt is the instant of creation of the OAuth2AuthorizeToken
	CreatedAt time.Time `sql:"default:(now() at time zone 'utc')"`
	// ExpiresIn is the seconds from CreatedAt before this token expires
	ExpiresIn uint64
	// State data from request
	//State string, [!] we dont't store state variables
	// Scope is the requested scope
	Scope string
	// RedirectUri is the RedirectUri associated with the token
	RedirectURI string
	// UserID is references the User that created the authorization request and thus the AuthorizeData
	UserID uint64
}

// TableName returns the table name associated with the structure
func (OAuth2AuthorizeData) TableName() string {
	return "oauth2_authorize"
}

// OAuth2AccessData is the OAuth2 access data
type OAuth2AccessData struct {
	ID uint64 `igor:"primary_key"`
	// ClientID references the client that created this token
	ClientID uint64
	// CreatedAt is the instant of creation of the OAuth2AccessToken
	CreatedAt time.Time `sql:"default:(now() at time zone 'utc')"`
	// ExpiresIn is the seconds from CreatedAt before this token expires
	ExpiresIn uint64
	// RedirectUri is the RedirectUri associated with the token
	RedirectURI string
	// AuthorizeDataID references the AuthorizationData that authorizated this token. Can be null
	AuthorizeDataID sql.NullInt64 `igor:"column:oauth2_authorize_id"` // Annotation required, since the column name does not follow igor conventions
	// AccessDataID references the Access Data, for refresh token. Can be null
	AccessDataID sql.NullInt64 `igor:"column:oauth2_access_id"` // Annotation required, since the column name does not follow igor conventions
	// RefreshTokenID is the value by which this token can be renewed. Can be null
	RefreshTokenID sql.NullInt64
	// AccessToken is the main value of this tructure, represents the access token
	AccessToken string
	// Scope is the requested scope
	Scope string
	// UserID is references the User that created The access request and thus the AccessData
	UserID uint64
}

// TableName returns the table name associated with the structure
func (OAuth2AccessData) TableName() string {
	return "oauth2_access"
}

// OAuth2RefreshToken is the model for the relation oauth2_refresh
type OAuth2RefreshToken struct {
	ID    uint64 `igor:"primary_key"`
	Token string `sql:"UNIQUE"`
}

// TableName returns the table name associated with the structure
func (OAuth2RefreshToken) TableName() string {
	return "oauth2_refresh"
}
