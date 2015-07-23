package nerdz

import (
	"database/sql"
	"net/url"
	"time"
)

// Enrich models structure with unexported types

// boardType represents a board type
type boardType string

const (
	USER    boardType = "user"
	PROJECT boardType = "project"
)

// Info contains the informations common to every board
// Used in API output to give user/project basic informations
type info struct {
	ID            uint64    `json:"id"`
	Owner         *info     `json:"owner"`
	Name          string    `json:"name"`
	Username      string    `json:"username"`
	Website       *url.URL  `json:"-"`
	WebsiteString string    `json:"website"`
	Image         *url.URL  `json:"-"`
	ImageString   string    `json:"image"`
	Closed        bool      `json:"closed"`
	Type          boardType `json:"type"`
	Board         *url.URL  `json:"-"`
	BoardString   string    `json:"board"`
}

type apiPostFields struct {
	FromInfo         *info  `sql:"-" json:"from"`
	ToInfo           *info  `sql:"-" json:"to"`
	Rate             int    `sql:"-" json:"rate"`
	RevisionsCount   uint8  `sql:"-" json:"revisions"`
	CommentsCount    uint8  `sql:"-" json:"comments"`
	BookmarkersCount uint8  `sql:"-" json:"bookmarkers"`
	LurkersCount     uint8  `sql:"-" json:"lurkers"`
	Url              string `sql:"-" json:"url"`
	Timestamp        int64  `sql:"-" json:"timestamp"`
	CanEdit          bool   `sql:"-" json:"canEdit"`
	CanDelete        bool   `sql:"-" json:"canDelete"`
	CanComment       bool   `sql:"-" json:"canComment"`
	CanBookmark      bool   `sql:"-" json:"canBookmark"`
	CanLurk          bool   `sql:"-" json:"canLurk"`
}

// Models

type UserPostsNoNotify struct {
	User    uint64    `json:"user"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (UserPostsNoNotify) TableName() string {
	return "posts_no_notify"
}

type UserPostCommentsNoNotify struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (UserPostCommentsNoNotify) TableName() string {
	return "comments_no_notify"
}

type UserPostCommentsNotify struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (UserPostCommentsNotify) TableName() string {
	return "comments_notify"
}

type Ban struct {
	User       uint64    `json:"user"`
	Motivation string    `json:"motivation"`
	Time       time.Time `sql:"default:NOW()" json:"time"`
	Counter    uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (Ban) TableName() string {
	return "ban"
}

type Blacklist struct {
	From       uint64    `json:"from"`
	To         uint64    `json:"to"`
	Motivation string    `json:"motivation"`
	Time       time.Time `sql:"default:NOW()" json:"time"`
	Counter    uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (Blacklist) TableName() string {
	return "blacklist"
}

type Whitelist struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (Whitelist) TableName() string {
	return "whitelist"
}

type UserFollower struct {
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `sql:"default:NOW()" json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (UserFollower) TableName() string {
	return "followers"
}

type ProjectNotify struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Hpid    uint64    `json:"hpid"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (ProjectNotify) TableName() string {
	return "groups_notify"
}

type ProjectPostsNoNotify struct {
	User    uint64    `json:"user"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (ProjectPostsNoNotify) TableName() string {
	return "groups_posts_no_notify"
}

type ProjectPostCommentsNoNotify struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (ProjectPostCommentsNoNotify) TableName() string {
	return "groups_comments_no_notify"
}

type ProjectPostCommentsNotify struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (ProjectPostCommentsNotify) TableName() string {
	return "groups_comments_notify"
}

type User struct {
	Counter     uint64    `gorm:"primary_key:yes" json:"counter"`
	Last        time.Time `sql:"default:NOW()" json:"last"`
	NotifyStory []byte    `sql:"type:json" json:"notifyStory"`
	Private     bool      `json:"private"`
	Lang        string    `sql:"type:varchar(2)" json:"lang"`
	Username    string    `sql:"type:varchar(90)" json:"username"`
	// Field commented out, to avoid the  possibility to fetch and show the password field
	//	Password    string         `sql:"type:varchar(40)"`
	//	RemoteAddr     string `sql:"type:inet"`
	//	HttpUserAgent  string `sql:"type:text"`
	Email            string    `sql:"type:varchar(350)" json:"-"` // Unexported field in JSON conversion
	Name             string    `sql:"type:varchar(60)" json:"name"`
	Surname          string    `sql:"tyoe:varchar(60)" json:"surname"`
	Gender           bool      `json:"gender"`
	BirthDate        time.Time `sql:"default:NOW()" json:"birthDate"`
	BoardLang        string    `sql:"type:varchar(2)" json:"boardLang"`
	Timezone         string    `sql:"type:varchar(35)" json:"timezone"`
	Viewonline       bool      `json:"viewonline"`
	RegistrationTime time.Time `sql:"default:NOW()" json:"registrationTime"`
	// User struct references Profile with a 1:1 relation
	Profile Profile `json:"profile"`
}

//TableName returns the table name associated with the structure
func (User) TableName() string {
	return "users"
}

type Profile struct {
	Counter        uint64    `gorm:"primary_key:yes" json:"counter"`
	Website        string    `sql:"type:varchar(350)" json:"website"`
	Quotes         string    `sql:"type:text" json:"quotes"`
	Biography      string    `sql:"type:text" json:"biography"`
	Interests      string    `sql:"type:text" json:"interests"`
	Github         string    `sql:"type:varchar(350)" json:"github"`
	Skype          string    `sql:"type:varchar(350)" json:"skype"`
	Jabber         string    `sql:"type:varchar(350)" json:"jabber"`
	Yahoo          string    `sql:"type:varchar(350)" json:"yahoo"`
	Userscript     string    `sql:"type:varchar(128)" json:"userscript"`
	Template       uint8     `json:"template"`
	MobileTemplate uint8     `json:"mobileTemplate"`
	Dateformat     string    `sql:"type:varchar(25)" json:"dateformat"`
	Facebook       string    `sql:"type:varchar(350)" json:"facebook"`
	Twitter        string    `sql:"type:varchar(350)" json:"twitter"`
	Steam          string    `sql:"type:varchar(350)" json:"steam"`
	Push           bool      `json:"push"`
	Pushregtime    time.Time `sql:"default:NOW()" json:"pushregtime"`
	Closed         bool      `json:"closed"`
}

//TableName returns the table name associated with the structure
func (Profile) TableName() string {
	return "profiles"
}

type UserPost struct {
	// Model fields
	Hpid    uint64    `gorm:"primary_key:yes" json:"hpid"`
	From    uint64    `json:"-"`
	To      uint64    `json:"-"`
	Pid     uint64    `sql:"default:0" json:"pid"`
	Message string    `sql:"type:text" json:"message"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Lang    string    `sql:"type:varchar(2)" json:"lang"`
	News    bool      `json:"news"`
	Closed  bool      `json:"closed"`
	// API fields
	apiPostFields
}

//TableName returns the table name associated with the structure
func (UserPost) TableName() string {
	return "posts"
}

type UserPostRevision struct {
	Hpid    uint64    `json:"hpid"`
	Message string    `json:"message"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	RevNo   uint16    `json:"revNo"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (UserPostRevision) TableName() string {
	return "posts_revisions"
}

type UserPostThumb struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Vote    int8      `json:"vote"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (UserPostThumb) TableName() string {
	return "thumbs"
}

type UserPostLurker struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (UserPostLurker) TableName() string {
	return "lurkers"
}

type UserPostComment struct {
	Hcid     uint64    `gorm:"primary_key:yes" json:"hcid"`
	Hpid     uint64    `json:"hpid"`
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Message  string    `sql:"type:text" json:"message"`
	Time     time.Time `sql:"default:NOW()" json:"time"`
	Editable bool      `sql:"default:true" json:"editable"`
}

//TableName returns the table name associated with the structure
func (UserPostComment) TableName() string {
	return "comments"
}

type UserPostCommentRevision struct {
	Hcid    uint64    `json:"hcid"`
	Message string    `json:"message"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	RevNo   int8      `json:"revNo"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (UserPostCommentRevision) TableName() string {
	return "comments_revisions"
}

type UserPostBookmark struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (UserPostBookmark) TableName() string {
	return "bookmarks"
}

type Pm struct {
	Pmid    uint64    `gorm:"primary_key:yes" json:"pmid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Message string    `sql:"type:text" json:"message"`
	ToRead  bool      `json:"toRead"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
}

//TableName returns the table name associated with the structure
func (Pm) TableName() string {
	return "pms"
}

type Project struct {
	Counter      uint64         `gorm:"primary_key:yes" json:"counter"`
	Description  string         `sql:"type:text" json:"description"`
	Name         string         `sql:"type:varchar(30)" json:"name"`
	Private      bool           `json:"private"`
	Photo        sql.NullString `sql:"type:varchar(350)" json:"photo"`
	Website      sql.NullString `sql:"type:varchar(350)" json:"website"`
	Goal         string         `sql:"type:text" json:"goal"`
	Visible      bool           `json:"visible"`
	Open         bool           `json:"open"`
	CreationTime time.Time      `sql:"default:NOW()" json:"creationTime"`
}

//TableName returns the table name associated with the structure
func (Project) TableName() string {
	return "groups"
}

type ProjectMember struct {
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `sql:"default:NOW()" json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (ProjectMember) TableName() string {
	return "groups_members"
}

type ProjectOwner struct {
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `sql:"default:NOW()" json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (ProjectOwner) TableName() string {
	return "groups_owners"
}

type ProjectPost struct {
	// Model fields
	Hpid    uint64    `gorm:"primary_key:yes" json:"hpid"`
	From    uint64    `json:"-"`
	To      uint64    `json:"-"`
	Pid     uint64    `sql:"default:0" json:"pid"`
	Message string    `sql:"type:text" json:"message"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	News    bool      `json:"news"`
	Lang    string    `sql:"type:varchar(2)" json:"lang"`
	Closed  bool      `json:"closed"`
	// API fields
	apiPostFields
}

//TableName returns the table name associated with the structure
func (ProjectPost) TableName() string {
	return "groups_posts"
}

type ProjectPostRevision struct {
	Hpid    uint64    `json:"hpid"`
	Message string    `json:"message"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	RevNo   uint16    `json:"revNo"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (ProjectPostRevision) TableName() string {
	return "groups_posts_revisions"
}

type ProjectPostThumb struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Vote    int8      `json:"vote"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (ProjectPostThumb) TableName() string {
	return "groups_thumbs"
}

type ProjectPostLurker struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (ProjectPostLurker) TableName() string {
	return "groups_lurkers"
}

type ProjectPostComment struct {
	Hcid     uint64    `gorm:"primary_key:yes" json:"hcid"`
	Hpid     uint64    `json:"hpid"`
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Message  string    `sql:"type:text" json:"message"`
	Time     time.Time `sql:"default:NOW()" json:"time"`
	Editable bool      `sql:"default:true" json:"editable"`
}

//TableName returns the table name associated with the structure
func (ProjectPostComment) TableName() string {
	return "groups_comments"
}

type ProjectPostCommentRevision struct {
	Hcid    uint64    `json:"hcid"`
	Message string    `json:"message"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	RevNo   uint16    `json:"revNo"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (ProjectPostCommentRevision) TableName() string {
	return "groups_comments_revisions"
}

type ProjectPostBookmark struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (ProjectPostBookmark) TableName() string {
	return "groups_bookmarks"
}

type ProjectFollower struct {
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `sql:"default:NOW()" json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (ProjectFollower) TableName() string {
	return "groups_followers"
}

type UserPostCommentThumb struct {
	Hcid    uint64 `json:"hcid"`
	User    uint64 `json:"user"`
	Vote    int8   `json:"vote"`
	Counter uint64 `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (UserPostCommentThumb) TableName() string {
	return "comment_thumbs"
}

type ProjectPostCommentThumb struct {
	Hcid    uint64    `json:"hcid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Vote    int8      `json:"vote"`
	Time    time.Time `sql:"default:NOW()" json:"time"`
	Counter uint64    `gorm:"primary_key:yes" json:"counter"`
}

//TableName returns the table name associated with the structure
func (ProjectPostCommentThumb) TableName() string {
	return "groups_comment_thumbs"
}

type DeletedUser struct {
	Counter    uint64    `gorm:"primary_key:yes" json:"counter"`
	Username   string    `json:"username"`
	Time       time.Time `sql:"default:NOW()" json:"time"`
	Motivation string    `json:"motivation"`
}

//TableName returns the table name associated with the structure
func (DeletedUser) TableName() string {
	return "deleted_users"
}

type SpecialUser struct {
	Role    string `gorm:"primary_key:yes" sql:"type:varchar(20)" json:"role"`
	Counter uint64 `json:"counter"`
}

//TableName returns the table name associated with the structure
func (SpecialUser) TableName() string {
	return "special_users"
}

type SpecialProject struct {
	Role    string `gorm:"primary_key:yes" sql:"type:varchar(20)" json:"role"`
	Counter uint64 `json:"counter"`
}

//TableName returns the table name associated with the structure
func (SpecialProject) TableName() string {
	return "special_groups"
}

type PostClassification struct {
	ID    uint64 `gorm:"primary_key:yes" json:"id"`
	UHpid uint64 `json:"uHpid"`
	GHpid uint64 `json:"gHpid"`
	Tag   string `sql:"type:varchar(35)" json:"tag"`
}

//TableName returns the table name associated with the structure
func (PostClassification) TableName() string {
	return "posts_classifications"
}

type Mention struct {
	ID       uint64    `gorm:"primary_key:yes" json:"id"`
	UHpid    uint64    `json:"uHpid"`
	GHpid    uint64    `json:"gHpid"`
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `sql:"default:NOW()" json:"time"`
	ToNotify bool      `json:"toNotify"`
}

//TableName returns the table name associated with the structure
func (Mention) TableName() string {
	return "mentions"
}

// OAuth2Client implements the osin.Client interface
type OAuth2Client struct {
	ID          string `gorm:"primary_key:yes" json:"id"`
	Secret      string `json:"secret"`
	RedirectUri string `json:"redirectUri"`
	UserData    []byte `sql:"type:json" json:"userData"`
}

//TableName returns the table name associated with the structure
func (OAuth2Client) TableName() string {
	return "oauth2_clients"
}

// OAuth2AuthorizeData holds the authorization data for the OAuth2Client
type OAuth2AuthorizeData struct {
	ClientID    string    `json:"clientID"` // OAuth2Client foreign key
	Code        string    `gorm:"primary_key:yes" json:"code"`
	CreatedAt   time.Time `json:"createdAt"`
	ExpiresIn   int32     `json:"expiresIn"`
	RedirectUri string    `json:"redirectUri"`
	Scope       string    `json:"scope"`
	State       string    `json:"state"`
	UserData    []byte    `sql:"type:json" json:"userData"`
}

//TableName returns the table name associated with the structure
func (OAuth2AuthorizeData) TableName() string {
	return "oauth2_authorize_data"
}

type OAuth2AccessData struct {
	ClientID        string    `json:"clientID"`        // OAuth2Client foreign key
	AuthorizeDataID string    `json:"authorizeDataID"` // OAuth2AuthorizeData foreign key
	AccessDataID    string    `json:"accessDataID"`    // Previous access data, for refresh token (can be null)
	AccessToken     string    `gorm:"primary_key:yes" json:"accessToken"`
	RefreshToken    string    `json:"refreshToken"`
	ExpiresIn       int32     `json:"expiresIn"`
	Scope           string    `json:"scope"`
	RedirectUri     string    `json:"redirectUri"`
	CreatedAt       time.Time `json:"createdAt"`
	UserData        []byte    `sql:"type:json" json:"userData"`
}

//TableName returns the table name associated with the structure
func (OAuth2AccessData) TableName() string {
	return "oauth2_access_data"
}
