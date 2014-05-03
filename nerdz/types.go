package nerdz

import (
	"database/sql"
	"time"
)

type UserPostsNoNotify struct {
	User int64
	Hpid int64
	Time time.Time
}

//TableName returns the table name associated with the structure
func (UserPostsNoNotify) TableName() string {
	return "posts_no_notify"
}

type UserPostCommentsNoNotify struct {
	From int64
	To   int64
	Hpid int64
	Time time.Time
}

//TableName returns the table name associated with the structure
func (UserPostCommentsNoNotify) TableName() string {
	return "comments_no_notify"
}

type UserPostCommentsNotify struct {
	From int64
	To   int64
	Hpid int64
	Time time.Time
}

//TableName returns the table name associated with the structure
func (UserPostCommentsNotify) TableName() string {
	return "comments_notify"
}

type Ban struct {
	User       int64
	Motivation string
}

//TableName returns the table name associated with the structure
func (Ban) TableName() string {
	return "ban"
}

type Blacklist struct {
	From       int64
	To         int64
	Motivation string
}

//TableName returns the table name associated with the structure
func (Blacklist) TableName() string {
	return "blacklist"
}

type Whitelist struct {
	From int64
	To   int64
}

//TableName returns the table name associated with the structure
func (Whitelist) TableName() string {
	return "whitelist"
}

type UserFollow struct {
	From     int64
	To       int64
	Time     time.Time
	Notified bool
}

//TableName returns the table name associated with the structure
func (UserFollow) TableName() string {
	return "follow"
}

type ProjectNotify struct {
	Group int64
	To    int64
	Time  time.Time
}

//TableName returns the table name associated with the structure
func (ProjectNotify) TableName() string {
	return "groups_notify"
}

type ProjectPostsNoNotify struct {
	User int64
	Hpid int64
	Time time.Time
}

//TableName returns the table name associated with the structure
func (ProjectPostsNoNotify) TableName() string {
	return "groups_posts_no_notify"
}

type ProjectPostCommentsNoNotify struct {
	From int64
	To   int64
	Hpid int64
	Time time.Time
}

//TableName returns the table name associated with the structure
func (ProjectPostCommentsNoNotify) TableName() string {
	return "groups_comments_no_notify"
}

type ProjectPostCommentsNotify struct {
	From int64
	To   int64
	Hpid int64
	Time time.Time
}

//TableName returns the table name associated with the structure
func (ProjectPostCommentsNotify) TableName() string {
	return "groups_comments_notify"
}

type User struct {
	Counter     int64 `primaryKey:"yes"`
	Last        time.Time
	NotifyStory []byte `sql:"type:json"`
	Private     bool
	Lang        string `sql:"type:varchar(2)"`
	Username    string `sql:"type:varchar(90)"`
	// Field commented out, to avoid the  possibility to fetch and show the password field
	//	Password    string         `sql:"type:varchar(40)"`
	Name       string `sql:"type:varchar(60)"`
	Surname    string `sql:"tyoe:varchar(60)"`
	Email      string `sql:"type:varchar(350)"`
	Gender     bool
	BirthDate  time.Time
	BoardLang  string `sql:"type:varchar(2)"`
	Timezone   string `sql:"type:varchar(35)"`
	Viewonline bool
	// User struct references Profile with a 1:1 relation
	Profile Profile
}

//TableName returns the table name associated with the structure
func (User) TableName() string {
	return "users"
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

//TableName returns the table name associated with the structure
func (Profile) TableName() string {
	return "profiles"
}

type ClosedProfile struct {
	Counter int64 `primaryKey:"yes"`
}

//TableName returns the table name associated with the structure
func (ClosedProfile) TableName() string {
	return "closed_profiles"
}

type UserPost struct {
	Hpid    int64 `primaryKey:"yes"`
	From    int64
	To      int64
	Pid     int64
	Message string `sql:"type:text"`
	Notify  bool
	Time    time.Time
}

//TableName returns the table name associated with the structure
func (UserPost) TableName() string {
	return "posts"
}

type UserPostThumb struct {
	Hpid int64
	User int64
	Vote int16
}

//TableName returns the table name associated with the structure
func (UserPostThumb) TableName() string {
	return "thumbs"
}

type UserPostLurker struct {
	User int64
	Post int64
	Time time.Time
}

//TableName returns the table name associated with the structure
func (UserPostLurker) TableName() string {
	return "lurkers"
}

type UserPostComment struct {
	Hcid    int64 `primaryKey:"yes"`
	Hpid    int64
	From    int64
	To      int64
	Message string `sql:"type:text"`
	Time    time.Time
}

//TableName returns the table name associated with the structure
func (UserPostComment) TableName() string {
	return "comments"
}

type UserBookmark struct {
	Hpid int64
	From int64
	Time time.Time
}

//TableName returns the table name associated with the structure
func (UserBookmark) TableName() string {
	return "bookmarks"
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

//TableName returns the table name associated with the structure
func (Pm) TableName() string {
	return "pms"
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

//TableName returns the table name associated with the structure
func (Project) TableName() string {
	return "groups"
}

type ProjectMember struct {
	Group int64
	User  int64
}

//TableName returns the table name associated with the structure
func (ProjectMember) TableName() string {
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

//TableName returns the table name associated with the structure
func (ProjectPost) TableName() string {
	return "groups_posts"
}

type ProjectPostThumb struct {
	Hpid int64
	User int64
	Vote int16
}

//TableName returns the table name associated with the structure
func (ProjectPostThumb) TableName() string {
	return "groups_thumbs"
}

type ProjectPostLurker struct {
	User int64
	Post int64
	Time time.Time
}

//TableName returns the table name associated with the structure
func (ProjectPostLurker) TableName() string {
	return "groups_lurkers"
}

type ProjectPostComment struct {
	Hcid    int64 `primaryKey:"yes"`
	Hpid    int64
	From    int64
	To      int64
	Message string `sql:"type:text"`
	Time    time.Time
}

//TableName returns the table name associated with the structure
func (ProjectPostComment) TableName() string {
	return "groups_comments"
}

type ProjectBookmark struct {
	Hpid int64
	From int64
	Time time.Time
}

//TableName returns the table name associated with the structure
func (ProjectBookmark) TableName() string {
	return "groups_bookmarks"
}

type ProjectFollower struct {
	Group int64
	User  int64
}

//TableName returns the table name associated with the structure
func (ProjectFollower) TableName() string {
	return "groups_followers"
}

type UserPostCommentThumb struct {
	Hcid int64
	User int64
	Vote int16
}

//TableName returns the table name associated with the structure
func (UserPostCommentThumb) TableName() string {
	return "comment_thumbs"
}

type ProjectPostCommentThumb struct {
	Hcid int64
	User int64
	Vote int16
}

//TableName returns the table name associated with the structure
func (ProjectPostCommentThumb) TableName() string {
	return "groups_comment_thumbs"
}
