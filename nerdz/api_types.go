package nerdz

import (
	"database/sql"
	"time"
)

const (
	MinPosts = 1
	MaxPosts = 20
)

// Response represent the response format of the API
type Response struct {
	Data         interface{} `json:"data"`
	Message      string      `json:"message"`
	HumanMessage string      `json:"humanMessage"`
	Status       uint        `json:"status"`
	Success      bool        `json:"success"`
}

type Renderable interface {
	Render() string
}

// Info contains the informations common to every board
// Used in API output to give user/project basic informations
type info struct {
	ID            uint64    `json:"id"`
	Owner         *info     `json:"owner"`
	Name          string    `json:"name"`
	Username      string    `json:"username"`
	WebsiteString string    `json:"website"`
	ImageString   string    `json:"image"`
	Closed        bool      `json:"closed"`
	Type          boardType `json:"type"`
	BoardString   string    `json:"board"`
}

type PostFields struct {
	FromInfo         *info  `json:"from"`
	ToInfo           *info  `json:"to"`
	Rate             int    `json:"rate"`
	RevisionsCount   uint8  `json:"revisions"`
	CommentsCount    uint8  `json:"comments"`
	BookmarkersCount uint8  `json:"bookmarkers"`
	LurkersCount     uint8  `json:"lurkers"`
	URL              string `json:"url"`
	Timestamp        int64  `json:"timestamp"`
	CanEdit          bool   `json:"canEdit"`
	CanDelete        bool   `json:"canDelete"`
	CanComment       bool   `json:"canComment"`
	CanBookmark      bool   `json:"canBookmark"`
	CanLurk          bool   `json:"canLurk"`
}

/*
// setApiFields populate the API fileds of the UserPost
func (post *ProjectPost) setApiFields(user *User) {
	if from, e := NewUser(post.From); e == nil {
		post.FromInfo = from.Info()
	}
	if to, e := NewProject(post.To); e == nil {
		post.ToInfo = to.Info()
	}
	post.Rate = post.Thumbs()
	post.RevisionsCount = post.RevisionsNumber()
	post.CommentsCount = post.CommentsNumber()
	post.BookmarkersCount = post.BookmarkersNumber()
	post.LurkersCount = post.LurkersNumber()
	post.Timestamp = post.Time.Unix()
	post.Url = post.URL(Configuration.NERDZURL).String()
	post.CanBookmark = user.canBookmark(post)
	post.CanComment = user.canComment(post)
	post.CanDelete = user.canDelete(post)
	post.CanEdit = user.canEdit(post)
	post.CanLurk = user.canLurk(post)
}

// setApiFields populate the API fileds of the UserPost
func (post *UserPost) setApiFields(user *User) {
	if from, e := NewUser(post.From); e == nil {
		post.FromInfo = from.Info()
	}
	if to, e := NewUser(post.To); e == nil {
		post.ToInfo = to.Info()
	}
	post.Rate = post.Thumbs()
	post.RevisionsCount = post.RevisionsNumber()
	post.CommentsCount = post.CommentsNumber()
	post.BookmarkersCount = post.BookmarkersNumber()
	post.LurkersCount = post.LurkersNumber()
	post.Timestamp = post.Time.Unix()
	post.Url = post.URL(Configuration.NERDZURL).String()
	post.CanBookmark = user.canBookmark(post)
	post.CanComment = user.canComment(post)
	post.CanDelete = user.canDelete(post)
	post.CanEdit = user.canEdit(post)
	post.CanLurk = user.canLurk(post)
}

*/

type UserPostsNoNotifyTO struct {
	User    uint64    `json:"user"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (u UserPostsNoNotifyTO) Render() string {
	return "UserPostsNoNotify"
}

type UserPostCommentsNoNotifyTO struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (u UserPostCommentsNoNotifyTO) Render() string {
	return "UserPostCommentsNoNotify"
}

type UserPostCommentsNotifyTO struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (u UserPostCommentsNotifyTO) Render() string {
	return "UserPostCommentsNotify"
}

type BanTO struct {
	User       uint64    `json:"user"`
	Motivation string    `json:"motivation"`
	Time       time.Time `json:"time"`
	Counter    uint64    `json:"counter"`
}

func (BanTO) Render() string {
	return "Ban"
}

type BlacklistTO struct {
	From       uint64    `json:"from"`
	To         uint64    `json:"to"`
	Motivation string    `json:"motivation"`
	Time       time.Time `json:"time"`
	Counter    uint64    `json:"counter"`
}

func (BlacklistTO) Render() string {
	return "Blacklist"
}

type WhitelistTO struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (WhitelistTO) Render() string {
	return "WhiteList"
}

type UserFollowerTO struct {
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

func (UserFollowerTO) Render() string {
	return "UserFollower"
}

type ProjectNotifyTO struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `json:"time"`
	Hpid    uint64    `json:"hpid"`
	Counter uint64    `json:"counter"`
}

func (ProjectNotifyTO) Render() string {
	return "ProjectNotify"
}

type ProjectPostsNoNotifyTO struct {
	User    uint64    `json:"user"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostsNoNotifyTO) Render() string {
	return "ProjectPostsNoNotify"
}

type ProjectPostCommentsNoNotifyTO struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostCommentsNoNotifyTO) Render() string {
	return "ProjectPostCommentsNoNotify"
}

type ProjectPostCommentsNotifyTO struct {
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostCommentsNotifyTO) Render() string {
	return "ProjectPostCommentsNotifyTO"
}

type UserTO struct {
	Counter          uint64    `json:"counter"`
	Last             time.Time `json:"last"`
	NotifyStory      []byte    `json:"notifyStory"`
	Private          bool      `json:"private"`
	Lang             string    `json:"lang"`
	Username         string    `json:"username"`
	Name             string    `json:"name"`
	Surname          string    `json:"surname"`
	Gender           bool      `json:"gender"`
	BirthDate        time.Time `json:"birthDate"`
	BoardLang        string    `json:"boardLang"`
	Timezone         string    `json:"timezone"`
	Viewonline       bool      `json:"viewonline"`
	RegistrationTime time.Time `json:"registrationTime"`
	Profile          Profile   `json:"profile"`
}

func (UserTO) Render() string {
	return "User"
}

type ProfileTO struct {
	Counter        uint64    `json:"counter"`
	Website        string    `json:"website"`
	Quotes         string    `json:"quotes"`
	Biography      string    `json:"biography"`
	Interests      string    `json:"interests"`
	Github         string    `json:"github"`
	Skype          string    `json:"skype"`
	Jabber         string    `json:"jabber"`
	Yahoo          string    `json:"yahoo"`
	Userscript     string    `json:"userscript"`     // ?API?
	Template       uint8     `json:"template"`       // ?API?
	MobileTemplate uint8     `json:"mobileTemplate"` // ?API?
	Dateformat     string    `json:"dateformat"`     // ?API?
	Facebook       string    `json:"facebook"`
	Twitter        string    `json:"twitter"`
	Steam          string    `json:"steam"`
	Push           bool      `json:"push"`        // ?API?
	Pushregtime    time.Time `json:"pushregtime"` // ?API?
	Closed         bool      `json:"closed"`
}

func (ProfileTO) Render() string {
	return "Profile"
}

type UserPostTO struct {
	Hpid     uint64    `json:"hpid"`
	Pid      uint64    `json:"pid"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	Lang     string    `json:"lang"`
	News     bool      `json:"news"`
	Closed   bool      `json:"closed"`
	PostInfo PostFields
}

func (UserPostTO) Render() string {
	return "UserPost"
}

type UserPostRevisionTO struct {
	Hpid    uint64    `json:"hpid"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	RevNo   uint16    `json:"revNo"`
	Counter uint64    `json:"counter"`
}

func (UserPostRevisionTO) Render() string {
	return "UserPostRevision"
}

type UserPostThumbTO struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Vote    int8      `json:"vote"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (UserPostThumbTO) Render() string {
	return "UserPostThumb"
}

type UserPostLurkerTO struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (UserPostLurkerTO) Render() string {
	return "UserPostLurker"
}

type UserPostCommentTO struct {
	Hcid     uint64    `json:"hcid"`
	Hpid     uint64    `json:"hpid"`
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	Editable bool      `json:"editable"`
}

func (UserPostCommentTO) Render() string {
	return "UserPostComment"
}

type UserPostCommentRevisionTO struct {
	Hcid    uint64    `json:"hcid"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	RevNo   int8      `json:"revNo"`
	Counter uint64    `json:"counter"`
}

func (UserPostCommentRevisionTO) Render() string {
	return "UserPostCommentRevision"
}

type UserPostBookmarkTO struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (UserPostBookmarkTO) Render() string {
	return "UserPostBookmarkTO"
}

type PmTO struct {
	Pmid    uint64    `json:"pmid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Message string    `json:"message"`
	ToRead  bool      `json:"toRead"`
	Time    time.Time `json:"time"`
}

func (PmTO) Render() string {
	return "Pm"
}

type ProjectTO struct {
	Counter      uint64         `json:"counter"`
	Description  string         `json:"description"`
	Name         string         `json:"name"`
	Private      bool           `json:"private"`
	Photo        sql.NullString `json:"photo"`
	Website      sql.NullString `json:"website"`
	Goal         string         `json:"goal"`
	Visible      bool           `json:"visible"`
	Open         bool           `json:"open"`
	CreationTime time.Time      `json:"creationTime"`
}

func (ProjectTO) Render() string {
	return "Project"
}

type ProjectMemberTO struct {
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

func (ProjectMemberTO) Render() string {
	return "ProjectMember"
}

type ProjectOwnerTO struct {
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

func (ProjectOwnerTO) Render() string {
	return "ProjectOwner"
}

type ProjectPostTO struct {
	Hpid     uint64    `json:"hpid"`
	Pid      uint64    `json:"pid"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	News     bool      `json:"news"`
	Lang     string    `json:"lang"`
	Closed   bool      `json:"closed"`
	PostInfo PostFields
}

func (ProjectPostTO) Render() string {
	return "ProjectPost"
}

type ProjectPostRevisionTO struct {
	Hpid    uint64    `json:"hpid"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	RevNo   uint16    `json:"revNo"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostRevisionTO) Render() string {
	return "ProjectPost"
}

type ProjectPostThumbTO struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `json:"time"`
	Vote    int8      `json:"vote"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostThumbTO) Render() string {
	return "ProjectPostThumb"
}

type ProjectPostLurkerTO struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostLurkerTO) Render() string {
	return "ProjectPostLurker"
}

type ProjectPostCommentTO struct {
	Hcid     uint64    `json:"hcid"`
	Hpid     uint64    `json:"hpid"`
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	Editable bool      `json:"editable"`
}

func (ProjectPostCommentTO) Render() string {
	return "ProjectPostComment"
}

type ProjectPostCommentRevisionTO struct {
	Hcid    uint64    `json:"hcid"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	RevNo   uint16    `json:"revNo"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostCommentRevisionTO) Render() string {
	return "ProjectPostCommentRevision"
}

type ProjectPostBookmarkTO struct {
	Hpid    uint64    `json:"hpid"`
	From    uint64    `json:"from"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostBookmarkTO) Render() string {
	return "ProjectPostBookmark"
}

type ProjectFollowerTO struct {
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

func (ProjectFollowerTO) Render() string {
	return "ProjectFollower"
}

type UserPostCommentThumbTO struct {
	Hcid    uint64 `json:"hcid"`
	User    uint64 `json:"user"`
	Vote    int8   `json:"vote"`
	Counter uint64 `json:"counter"`
}

func (UserPostCommentThumbTO) Render() string {
	return "UserPostCommentThumb"
}

type ProjectPostCommentThumbTO struct {
	Hcid    uint64    `json:"hcid"`
	From    uint64    `json:"from"`
	To      uint64    `json:"to"`
	Vote    int8      `json:"vote"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

func (ProjectPostCommentThumbTO) Render() string {
	return "ProjectPostCommentThumb"
}

type DeletedUserTO struct {
	Counter    uint64    `gorm:"primary_key:yes" json:"counter"`
	Username   string    `json:"username"`
	Time       time.Time `sql:"default:NOW()" json:"time"`
	Motivation string    `json:"motivation"`
}

func (DeletedUserTO) Render() string {
	return "DeletedUser"
}

type SpecialUserTO struct {
	Role    string `json:"role"`
	Counter uint64 `json:"counter"`
}

func (SpecialUserTO) Render() string {
	return "SpecialUser"
}

type SpecialProjectTO struct {
	Role    string `json:"role"`
	Counter uint64 `json:"counter"`
}

func (SpecialProjectTO) Render() string {
	return "SpecialProject"
}

type PostClassificationTO struct {
	ID    uint64 `json:"id"`
	UHpid uint64 `json:"uHpid"`
	GHpid uint64 `json:"gHpid"`
	Tag   string `json:"tag"`
}

func (PostClassificationTO) Render() string {
	return "PostClassification"
}

type MentionTO struct {
	ID       uint64    `json:"id"`
	UHpid    uint64    `json:"uHpid"`
	GHpid    uint64    `json:"gHpid"`
	From     uint64    `json:"from"`
	To       uint64    `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
}

func (MentionTO) Render() string {
	return "Mention"
}
