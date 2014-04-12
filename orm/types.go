package orm

import (
	"database/sql"
	"time"
)

// Specify struct name that respect gorm's conventions, for tables that does not

type PostsNoNotify struct {
	User int64
	Hpid int64
	Time time.Time
}

func (x PostsNoNotify) TableName() string {
	return "posts_no_notify"
}

type CommentsNoNotify struct {
	From int64
	To   int64
	Hpid int64
	Time time.Time
}

func (x CommentsNoNotify) TableName() string {
	return "comments_no_notify"
}

type CommentsNotify struct {
	From int64
	To   int64
	Hpid int64
	Time time.Time
}

func (x CommentsNotify) TableName() string {
	return "comments_notify"
}

type Ban struct {
	User       int64
	Motivation string
}

func (x Ban) TableName() string {
	return "ban"
}

type Blacklist struct {
	From       int64
	To         int64
	Motivation string
}

func (x Blacklist) TableName() string {
	return "blacklist"
}

type Whitelist struct {
	From int64
	To   int64
}

func (x Whitelist) TableName() string {
	return "whitelist"
}

type Follow struct {
	From     int64
	To       int64
	Time     time.Time
	Notified bool
}

func (x Follow) TableName() string {
	return "follow"
}

type ProjectNotify struct {
	Group int64
	To    int64
	Time  time.Time
}

func (x ProjectNotify) TableName() string {
	return "groups_notify"
}

type ProjectPostsNoNotify struct {
	User int64
	Hpid int64
	Time time.Time
}

func (x ProjectPostsNoNotify) TableName() string {
	return "groups_posts_no_notify"
}

type ProjectCommentsNoNotify struct {
	From int64
	To   int64
	Hpid int64
	Time time.Time
}

func (x ProjectCommentsNoNotify) TableName() string {
	return "groups_comments_no_notify"
}

type ProjectCommentsNotify struct {
	From int64
	To   int64
	Hpid int64
	Time time.Time
}

func (x ProjectCommentsNotify) TableName() string {
	return "groups_comments_notify"
}

// Begin structures with table name that respect conventions
// In this cas we don't need to map struct with table manually with TableName

type User struct {
	Counter     int64 `primaryKey:"yes"`
	Last        time.Time
	NotifyStory string `sql:"type:json"`
	Private     bool
	Lang        sql.NullString `sql:"type:varchar(2)"`
	Username    string         `sql:"type:varchar(90)"`
	// Field commented out, to avoid the  possibility to fetch and show the password field
	//	Password    string         `sql:"type:varchar(40)"`
	Name       string `sql:"type:varchar(60)"`
	Surname    string `sql:"tyoe:varchar(60)"`
	Email      string `sql:"type:varchar(350)"`
	Gender     bool
	BirthDate  time.Time
	BoardLang  sql.NullString `sql:"type:varchar(2)"`
	Timezone   string         `sql:"type:varchar(35)"`
	Viewonline bool
	// User struct references Profile with a 1:1 relation
	Profile Profile
}

type Profile struct {
	Counter int64 `primaryKey:"yes"`
	// Field commented out, to avoid the  possibility to fetch and show the IP Address and User Agent field
	//	RemoteAddr     string `sql:"type:inet"`
	//	HttpUserAgent  string `sql:"type:text"`
	Website        string `sql:"type:varchar(350)"`
	Quotes         string `sql:"type:text"`
	Biography      string `sql:"type:text"`
	Interests      string `sql:"type:text"`
	Github         string `sql:"type:varchar(350)"`
	Skype          string `sql:"type:varchar(350)"`
	Jabber         string `sql:"type:varchar(350)"`
	Yahoo          string `sql:"type:varchar(350)"`
	Userscript     string `sql:"type:varchar(128)"`
	Template       int16
	MobileTemplate int16
	Dateformat     string `sql:"type:varchar(25)"`
	Facebook       string `sql:"type:varchar(350)"`
	Twitter        string `sql:"type:varchar(350)"`
	Steam          string `sql:"type:varchar(350)"`
	Push           bool
	Pushregtime    time.Time
}

type ClosedProfile struct {
	Counter int64 `primaryKey:"yes"`
}

type Post struct {
	Hpid    int64 `primaryKey:"yes"`
	From    int64
	To      int64
	Pid     int64
	Message string `sql:"type:text"`
	Notify  bool
	Time    time.Time
}

type PostThumb struct {
	Hpid int64
	User int64
	Vote int16
}

func (x PostThumb) TableName() string {
	return "thumbs"
}

type Lurker struct {
	User int64
	Post int64
	Time time.Time
}

type Comment struct {
	Hcid    int64 `primaryKey:"yes"`
	Hpid    int64
	From    int64
	To      int64
	Message string `sql:"type:text"`
	Time    time.Time
}

type Bookmark struct {
	Hpid int64
	From int64
	Time time.Time
}

type Pm struct {
	Pmid    int64 `primaryKey:"yes"`
	From    int64
	To      int64
	Pid     int64
	Message string `sql:"type:text"`
	Read    bool
	Time    time.Time
}

type Project struct {
	Counter     int64  `primaryKey:"yes"`
	Description string `sql:"type:text"`
	Owner       int64
	Name        string `sql:"type:varchar(30)"`
	Private     bool
	Photo       sql.NullString `sql:"type:varchar(350)"`
	Website     string         `sql:"type:varchar(350)"`
	Goal        string         `sql:"type:text"`
	Visible     bool
	Open        bool
}

func (x Project) TableName() string {
	return "groups"
}

type ProjectMember struct {
	Group int64
	User  int64
}

func (x ProjectMember) TableName() string {
	return "groups_members"
}

type ProjectPost struct {
	Hpid    int64 `primaryKey:"yes"`
	From    int64
	To      int64
	Pid     int64
	Message string `sql:"type:text"`
	News    bool
	Time    time.Time
}

func (x ProjectPost) TableName() string {
	return "groups_posts"
}

type ProjectPostThumb struct {
	Hpid int64
	User int64
	Vote int16
}

func (x ProjectPostThumb) TableName() string {
	return "groups_thumb"
}

type ProjectPostLurker struct {
	User int64
	Post int64
	Time time.Time
}

func (x ProjectPostLurker) TableName() string {
	return "groups_lurkers"
}

type ProjectComment struct {
	Hcid    int64 `primaryKey:"yes"`
	Hpid    int64
	From    int64
	To      int64
	Message string `sql:"type:text"`
	Time    time.Time
}

func (x ProjectComment) TableName() string {
	return "groups_comments"
}

type ProjectBookmark struct {
	Hpid int64
	From int64
	Time time.Time
}

func (x ProjectBookmark) TableName() string {
	return "groups_bookmarks"
}

type ProjectFollower struct {
	Group int64
	User  int64
}

func (x ProjectFollower) TableName() string {
	return "groups_followers"
}

type CommentThumb struct {
	Hcid int64
	User int64
	Vote int16
}

type ProjectCommentThumb struct {
	Hcid int64
	User int64
	Vote int16
}

func (x ProjectCommentThumb) TableName() string {
	return "groups_comment_thumbs"
}
