package nerdz

import (
	"database/sql"
	"time"
)

type UserPostsNoNotify struct {
	User    uint64
	Hpid    uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (UserPostsNoNotify) TableName() string {
	return "posts_no_notify"
}

type UserPostCommentsNoNotify struct {
	From    uint64
	To      uint64
	Hpid    uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (UserPostCommentsNoNotify) TableName() string {
	return "comments_no_notify"
}

type UserPostCommentsNotify struct {
	From    uint64
	To      uint64
	Hpid    uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (UserPostCommentsNotify) TableName() string {
	return "comments_notify"
}

type Ban struct {
	User       uint64
	Motivation string
	Time       time.Time `sql:"default:NOW()"`
	Counter    uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (Ban) TableName() string {
	return "ban"
}

type Blacklist struct {
	From       uint64
	To         uint64
	Motivation string
	Time       time.Time `sql:"default:NOW()"`
	Counter    uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (Blacklist) TableName() string {
	return "blacklist"
}

type Whitelist struct {
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (Whitelist) TableName() string {
	return "whitelist"
}

type UserFollower struct {
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:NOW()"`
	ToNotify bool
	Counter  uint64 `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (UserFollower) TableName() string {
	return "followers"
}

type ProjectNotify struct {
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:NOW()"`
	Hpid    uint64
	Counter uint64 `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (ProjectNotify) TableName() string {
	return "groups_notify"
}

type ProjectPostsNoNotify struct {
	User    uint64
	Hpid    uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (ProjectPostsNoNotify) TableName() string {
	return "groups_posts_no_notify"
}

type ProjectPostCommentsNoNotify struct {
	From    uint64
	To      uint64
	Hpid    uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (ProjectPostCommentsNoNotify) TableName() string {
	return "groups_comments_no_notify"
}

type ProjectPostCommentsNotify struct {
	From    uint64
	To      uint64
	Hpid    uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (ProjectPostCommentsNotify) TableName() string {
	return "groups_comments_notify"
}

type User struct {
	Counter     uint64    `gorm:"primary_key:yes"`
	Last        time.Time `sql:"default:NOW()"`
	NotifyStory []byte    `sql:"type:json"`
	Private     bool
	Lang        string `sql:"type:varchar(2)"`
	Username    string `sql:"type:varchar(90)"`
	// Field commented out, to avoid the  possibility to fetch and show the password field
	//	Password    string         `sql:"type:varchar(40)"`
	//	RemoteAddr     string `sql:"type:inet"`
	//	HttpUserAgent  string `sql:"type:text"`
	Name             string `sql:"type:varchar(60)"`
	Surname          string `sql:"tyoe:varchar(60)"`
	Email            string `sql:"type:varchar(350)"`
	Gender           bool
	BirthDate        time.Time `sql:"default:NOW()"`
	BoardLang        string    `sql:"type:varchar(2)"`
	Timezone         string    `sql:"type:varchar(35)"`
	Viewonline       bool
	RegistrationTime time.Time `sql:"default:NOW()"`
	// User struct references Profile with a 1:1 relation
	Profile Profile
}

//TableName returns the table name associated with the structure
func (User) TableName() string {
	return "users"
}

type Profile struct {
	Counter        uint64 `gorm:"primary_key:yes"`
	Website        string `sql:"type:varchar(350)"`
	Quotes         string `sql:"type:text"`
	Biography      string `sql:"type:text"`
	Interests      string `sql:"type:text"`
	Github         string `sql:"type:varchar(350)"`
	Skype          string `sql:"type:varchar(350)"`
	Jabber         string `sql:"type:varchar(350)"`
	Yahoo          string `sql:"type:varchar(350)"`
	Userscript     string `sql:"type:varchar(128)"`
	Template       uint8
	MobileTemplate uint8
	Dateformat     string `sql:"type:varchar(25)"`
	Facebook       string `sql:"type:varchar(350)"`
	Twitter        string `sql:"type:varchar(350)"`
	Steam          string `sql:"type:varchar(350)"`
	Push           bool
	Pushregtime    time.Time `sql:"default:NOW()"`
	Closed         bool
}

//TableName returns the table name associated with the structure
func (Profile) TableName() string {
	return "profiles"
}

type UserPost struct {
	Hpid    uint64 `gorm:"primary_key:yes"`
	From    uint64
	To      uint64
	Pid     uint64    `sql:"default:0"`
	Message string    `sql:"type:text"`
	Time    time.Time `sql:"default:NOW()"`
	Lang    string    `sql:"type:varchar(2)"`
	News    bool
	Closed  bool
}

//TableName returns the table name associated with the structure
func (UserPost) TableName() string {
	return "posts"
}

type UserPostRevision struct {
	Hpid    uint64
	Message string
	Time    time.Time `sql:"default:NOW()"`
	RevNo   uint16
	Counter uint64 `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (UserPostRevision) TableName() string {
	return "posts_revisions"
}

type UserPostThumb struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Vote    int8
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (UserPostThumb) TableName() string {
	return "thumbs"
}

type UserPostLurker struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (UserPostLurker) TableName() string {
	return "lurkers"
}

type UserPostComment struct {
	Hcid     uint64 `gorm:"primary_key:yes"`
	Hpid     uint64
	From     uint64
	To       uint64
	Message  string    `sql:"type:text"`
	Time     time.Time `sql:"default:NOW()"`
	Editable bool      `sql:"default:true"`
}

//TableName returns the table name associated with the structure
func (UserPostComment) TableName() string {
	return "comments"
}

type UserPostCommentRevision struct {
	Hcid    uint64
	Message string
	Time    time.Time `sql:"default:NOW()"`
	RevNo   int8
	Counter uint64 `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (UserPostCommentRevision) TableName() string {
	return "comments_revisions"
}

type UserPostBookmark struct {
	Hpid    uint64
	From    uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (UserPostBookmark) TableName() string {
	return "bookmarks"
}

type Pm struct {
	Pmid    uint64 `gorm:"primary_key:yes"`
	From    uint64
	To      uint64
	Message string `sql:"type:text"`
	ToRead  bool
	Time    time.Time `sql:"default:NOW()"`
}

type PmConfig struct {
	// TRUE: PM messages ordered in descending order using timestamp
	// FALSE: PM messages ordered in ascending order using timestamp
	DescOrder bool
	// number of messages returned (default: 0 - all the pms messages)
	Limit uint64
	// used in combination with Limit grant the possibility to return
	// a fraction of the whole pms
	Offset uint64
	// TRUE: Returns PM messages that should be read
	// FALSE: Returns PM messages that have already read
	ToRead bool
}

// Detail about a single private conversation between two users
type Conversation struct {
	From   string
	Time   time.Time
	ToRead bool
}

//TableName returns the table name associated with the structure
func (Pm) TableName() string {
	return "pms"
}

type Project struct {
	Counter      uint64 `gorm:"primary_key:yes"`
	Description  string `sql:"type:text"`
	Name         string `sql:"type:varchar(30)"`
	Private      bool
	Photo        sql.NullString `sql:"type:varchar(350)"`
	Website      sql.NullString `sql:"type:varchar(350)"`
	Goal         string         `sql:"type:text"`
	Visible      bool
	Open         bool
	CreationTime time.Time `sql:"default:NOW()"`
}

//TableName returns the table name associated with the structure
func (Project) TableName() string {
	return "groups"
}

type ProjectMember struct {
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:NOW()"`
	ToNotify bool
	Counter  uint64 `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (ProjectMember) TableName() string {
	return "groups_members"
}

type ProjectOwner struct {
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:NOW()"`
	ToNotify bool
	Counter  uint64 `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (ProjectOwner) TableName() string {
	return "groups_owners"
}

type ProjectPost struct {
	Hpid    uint64 `gorm:"primary_key:yes"`
	From    uint64
	To      uint64
	Pid     uint64    `sql:"default:0"`
	Message string    `sql:"type:text"`
	Time    time.Time `sql:"default:NOW()"`
	News    bool
	Lang    string `sql:"type:varchar(2)"`
	Closed  bool
}

//TableName returns the table name associated with the structure
func (ProjectPost) TableName() string {
	return "groups_posts"
}

type ProjectPostRevision struct {
	Hpid    uint64
	Message string
	Time    time.Time `sql:"default:NOW()"`
	RevNo   uint16
	Counter uint64 `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (ProjectPostRevision) TableName() string {
	return "groups_posts_revisions"
}

type ProjectPostThumb struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:NOW()"`
	Vote    int8
	Counter uint64 `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (ProjectPostThumb) TableName() string {
	return "groups_thumbs"
}

type ProjectPostLurker struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (ProjectPostLurker) TableName() string {
	return "groups_lurkers"
}

type ProjectPostComment struct {
	Hcid     uint64 `gorm:"primary_key:yes"`
	Hpid     uint64
	From     uint64
	To       uint64
	Message  string    `sql:"type:text"`
	Time     time.Time `sql:"default:NOW()"`
	Editable bool      `sql:"default:true"`
}

//TableName returns the table name associated with the structure
func (ProjectPostComment) TableName() string {
	return "groups_comments"
}

type ProjectPostCommentRevision struct {
	Hcid    uint64
	Message string
	Time    time.Time `sql:"default:NOW()"`
	RevNo   uint16
	Counter uint64 `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (ProjectPostCommentRevision) TableName() string {
	return "groups_comments_revisions"
}

type ProjectPostBookmark struct {
	Hpid    uint64
	From    uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (ProjectPostBookmark) TableName() string {
	return "groups_bookmarks"
}

type ProjectFollower struct {
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:NOW()"`
	ToNotify bool
	Counter  uint64 `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (ProjectFollower) TableName() string {
	return "groups_followers"
}

type UserPostCommentThumb struct {
	Hcid    uint64
	User    uint64
	Vote    int8
	Counter uint64 `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (UserPostCommentThumb) TableName() string {
	return "comment_thumbs"
}

type ProjectPostCommentThumb struct {
	Hcid    uint64
	From    uint64
	To      uint64
	Vote    int8
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

//TableName returns the table name associated with the structure
func (ProjectPostCommentThumb) TableName() string {
	return "groups_comment_thumbs"
}

type DeletedUser struct {
	Counter    uint64 `gorm:"primary_key:yes"`
	Username   string
	Time       time.Time `sql:"default:NOW()"`
	Motivation string
}

//TableName returns the table name associated with the structure
func (DeletedUser) TableName() string {
	return "deleted_users"
}

type SpecialUser struct {
	Role    string `gorm:"primary_key:yes"; sql:"type:varchar(20)"`
	Counter uint64
}

//TableName returns the table name associated with the structure
func (SpecialUser) TableName() string {
	return "special_users"
}

type SpecialProject struct {
	Role    string `gorm:"primary_key:yes"; sql:"type:varchar(20)"`
	Counter uint64
}

//TableName returns the table name associated with the structure
func (SpecialProject) TableName() string {
	return "special_groups"
}

type PostClassification struct {
	Id    uint64 `gorm:"primary_key:yes"`
	UHpid uint64
	GHpid uint64
	Tag   string `sql:"type:varchar(35)"`
}

//TableName returns the table name associated with the structure
func (PostClassification) TableName() string {
	return "posts_classifications"
}

type Mention struct {
	Id       uint64 `gorm:"primary_key:yes"`
	UHpid    uint64
	GHpid    uint64
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:NOW()"`
	ToNotify bool
}

//TableName returns the table name associated with the structure
func (Mention) TableName() string {
	return "mentions"
}
