package nerdz

import (
	"database/sql"
	"time"
)

// Enrich models structure with unexported types

// boardType represents a board type
type boardType string

const (
	// USER constant (of type boardType) makes possible to distinguish a User
	// board from a Project board
	USER boardType = "user"
	// PROJECT constant (of type boardType) makes possible to distinguish a PROJECT
	// board from a User board
	PROJECT boardType = "project"
)

//Transferable represents a common interface for all the
//types defined by the backend that are able to generate
//a data structure that can be returned by the API
type Transferable interface {
	//GetTO returns a proper data structure for the API
	GetTO() Renderable
}

// Models

type UserPostsNoNotify struct {
	User    uint64
	Hpid    uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

func (u *UserPostsNoNotify) GetTO() Renderable {
	return &UserPostsNoNotifyTO{
		User:    u.User,
		Hpid:    u.Hpid,
		Time:    u.Time,
		Counter: u.Counter,
	}
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

func (u *UserPostCommentsNoNotify) GetTO() Renderable {
	return &UserPostCommentsNoNotifyTO{
		From:    u.From,
		To:      u.To,
		Hpid:    u.Hpid,
		Time:    u.Time,
		Counter: u.Counter,
	}
}

type UserPostCommentsNotify struct {
	From    uint64
	To      uint64
	Hpid    uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

func (u *UserPostCommentsNotify) GetTO() Renderable {
	return &UserPostCommentsNotifyTO{
		From:    u.From,
		To:      u.To,
		Hpid:    u.Hpid,
		Time:    u.Time,
		Counter: u.Counter,
	}
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

func (b *Ban) GetTO() Renderable {
	return &BanTO{
		User:       b.User,
		Motivation: b.Motivation,
		Time:       b.Time,
		Counter:    b.Counter,
	}
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

func (b *Blacklist) GetTO() Renderable {
	return &BlacklistTO{
		From:       b.From,
		To:         b.To,
		Motivation: b.Motivation,
		Time:       b.Time,
		Counter:    b.Counter,
	}
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

func (w *Whitelist) GetTO() Renderable {
	return &WhitelistTO{
		From:    w.From,
		To:      w.To,
		Time:    w.Time,
		Counter: w.Counter,
	}
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

func (u *UserFollower) GetTO() Renderable {
	return &UserFollowerTO{
		From:     u.From,
		To:       u.To,
		Time:     u.Time,
		ToNotify: u.ToNotify,
		Counter:  u.Counter,
	}
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

func (p *ProjectNotify) GetTO() Renderable {
	return &ProjectNotifyTO{
		From:    p.From,
		To:      p.To,
		Time:    p.Time,
		Hpid:    p.Hpid,
		Counter: p.Counter,
	}
}

type ProjectPostsNoNotify struct {
	User    uint64
	Hpid    uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

func (p *ProjectPostsNoNotify) GetTO() Renderable {
	return &ProjectPostsNoNotifyTO{
		User:    p.User,
		Hpid:    p.Hpid,
		Time:    p.Time,
		Counter: p.Counter,
	}
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

func (p *ProjectPostCommentsNoNotify) GetTO() Renderable {
	return &ProjectPostCommentsNoNotifyTO{
		From:    p.From,
		To:      p.To,
		Hpid:    p.Hpid,
		Time:    p.Time,
		Counter: p.Counter,
	}
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

func (p *ProjectPostCommentsNotify) GetTO() Renderable {
	return &ProjectPostCommentsNotifyTO{
		From:    p.From,
		To:      p.To,
		Hpid:    p.Hpid,
		Time:    p.Time,
		Counter: p.Counter,
	}
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
	Email            string `sql:"type:varchar(350)"`
	Name             string `sql:"type:varchar(60)"`
	Surname          string `sql:"tyoe:varchar(60)"`
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

func (u *User) GetTO() Renderable {
	return &UserTO{
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
		Profile:          u.Profile,
	}
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

func (p *Profile) GetTO() Renderable {
	return &ProfileTO{
		Counter:        p.Counter,
		Website:        p.Website,
		Quotes:         p.Quotes,
		Biography:      p.Biography,
		Interests:      p.Interests,
		Github:         p.Github,
		Skype:          p.Skype,
		Jabber:         p.Jabber,
		Yahoo:          p.Yahoo,
		Userscript:     p.Userscript,
		Template:       p.Template,
		MobileTemplate: p.MobileTemplate,
		Dateformat:     p.Dateformat,
		Facebook:       p.Facebook,
		Twitter:        p.Twitter,
		Steam:          p.Steam,
		Push:           p.Push,
		Pushregtime:    p.Pushregtime,
		Closed:         p.Closed,
	}
}

//TableName returns the table name associated with the structure
func (Profile) TableName() string {
	return "profiles"
}

type Post struct {
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

type UserPost struct {
	Post
}

func (p *UserPost) GetTO() Renderable {
	user, _ := NewUser(p.From)

	to := UserPostTO{
		Hpid:    p.Hpid,
		Pid:     p.Pid,
		Message: p.Message,
		Time:    p.Time,
		Lang:    p.Lang,
		News:    p.News,
		Closed:  p.Closed,
	}

	to.SetPostFields(user, p)

	return &to
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

func (p *UserPostRevision) GetTO() Renderable {
	return &UserPostRevisionTO{
		Hpid:    p.Hpid,
		Message: p.Message,
		Time:    p.Time,
		RevNo:   p.RevNo,
		Counter: p.Counter,
	}
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

func (t *UserPostThumb) GetTO() Renderable {
	return &UserPostThumbTO{
		Hpid:    t.Hpid,
		From:    t.From,
		To:      t.To,
		Vote:    t.Vote,
		Time:    t.Time,
		Counter: t.Counter,
	}
}

type UserPostLurker struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

func (l *UserPostLurker) GetTO() Renderable {
	return &UserPostLurkerTO{
		Hpid:    l.Hpid,
		From:    l.From,
		To:      l.To,
		Time:    l.Time,
		Counter: l.Counter,
	}
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

func (c *UserPostComment) GetTO() Renderable {
	return &UserPostCommentTO{
		Hcid:     c.Hcid,
		Hpid:     c.Hpid,
		From:     c.From,
		To:       c.To,
		Message:  c.Message,
		Time:     c.Time,
		Editable: c.Editable,
	}
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

func (c *UserPostCommentRevision) GetTO() Renderable {
	return &UserPostCommentRevisionTO{
		Hcid:    c.Hcid,
		Message: c.Message,
		Time:    c.Time,
		RevNo:   c.RevNo,
		Counter: c.Counter,
	}
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

func (b *UserPostBookmark) GetTO() Renderable {
	return &UserPostBookmarkTO{
		Hpid:    b.Hpid,
		From:    b.From,
		Time:    b.Time,
		Counter: b.Counter,
	}
}

type Pm struct {
	Pmid    uint64 `gorm:"primary_key:yes"`
	From    uint64
	To      uint64
	Message string `sql:"type:text"`
	ToRead  bool
	Time    time.Time `sql:"default:NOW()"`
}

func (p *Pm) GetTO() Renderable {
	return &PmTO{
		Pmid:    p.Pmid,
		From:    p.From,
		To:      p.To,
		Message: p.Message,
		ToRead:  p.ToRead,
		Time:    p.Time,
	}
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

func (p *Project) GetTO() Renderable {
	return &ProjectTO{
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

func (m *ProjectMember) GetTO() Renderable {
	return &ProjectMemberTO{
		From:     m.From,
		To:       m.To,
		Time:     m.Time,
		ToNotify: m.ToNotify,
		Counter:  m.Counter,
	}
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

func (o *ProjectOwner) GetTO() Renderable {
	return &ProjectOwnerTO{
		From:     o.From,
		To:       o.To,
		Time:     o.Time,
		ToNotify: o.ToNotify,
		Counter:  o.Counter,
	}

}

//TableName returns the table name associated with the structure
func (ProjectOwner) TableName() string {
	return "groups_owners"
}

type ProjectPost struct {
	Post
}

//TableName returns the table name associated with the structure
func (ProjectPost) TableName() string {
	return "groups_posts"
}

func (p *ProjectPost) GetTO() Renderable {
	user, _ := NewUser(p.From)

	to := ProjectPostTO{
		Hpid:    p.Hpid,
		Pid:     p.Pid,
		Message: p.Message,
		Time:    p.Time,
		News:    p.News,
		Lang:    p.Lang,
		Closed:  p.Closed,
	}

	to.SetPostFields(user, p)

	return &to
}

type ProjectPostRevision struct {
	Hpid    uint64
	Message string
	Time    time.Time `sql:"default:NOW()"`
	RevNo   uint16
	Counter uint64 `gorm:"primary_key:yes"`
}

func (p *ProjectPostRevision) GetTO() Renderable {
	return &ProjectPostRevisionTO{
		Hpid:    p.Hpid,
		Message: p.Message,
		Time:    p.Time,
		RevNo:   p.RevNo,
		Counter: p.Counter,
	}
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

func (t *ProjectPostThumb) GetTO() Renderable {
	return &ProjectPostThumbTO{
		Hpid:    t.Hpid,
		From:    t.From,
		To:      t.To,
		Time:    t.Time,
		Vote:    t.Vote,
		Counter: t.Counter,
	}
}

type ProjectPostLurker struct {
	Hpid    uint64
	From    uint64
	To      uint64
	Time    time.Time `sql:"default:NOW()"`
	Counter uint64    `gorm:"primary_key:yes"`
}

func (l *ProjectPostLurker) GetTO() Renderable {
	return &ProjectPostLurkerTO{
		Hpid:    l.Hpid,
		From:    l.From,
		To:      l.To,
		Time:    l.Time,
		Counter: l.Counter,
	}
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

func (c *ProjectPostComment) GetTO() Renderable {
	return &ProjectPostCommentTO{
		Hcid:     c.Hcid,
		Hpid:     c.Hpid,
		From:     c.From,
		To:       c.To,
		Message:  c.Message,
		Time:     c.Time,
		Editable: c.Editable,
	}
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

func (r *ProjectPostCommentRevision) GetTO() Renderable {
	return &ProjectPostCommentRevisionTO{
		Hcid:    r.Hcid,
		Message: r.Message,
		Time:    r.Time,
		RevNo:   r.RevNo,
		Counter: r.Counter,
	}
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

func (b *ProjectPostBookmark) GetTO() Renderable {
	return &ProjectPostBookmarkTO{
		Hpid:    b.Hpid,
		From:    b.From,
		Time:    b.Time,
		Counter: b.Counter,
	}
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

func (p *ProjectFollower) GetTO() Renderable {
	return &ProjectFollowerTO{
		From:     p.From,
		To:       p.To,
		Time:     p.Time,
		ToNotify: p.ToNotify,
		Counter:  p.Counter,
	}
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

func (t *UserPostCommentThumb) GetTO() Renderable {
	return &UserPostCommentThumbTO{
		Hcid:    t.Hcid,
		User:    t.User,
		Vote:    t.Vote,
		Counter: t.Counter,
	}
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

func (t *ProjectPostCommentThumb) GetTO() Renderable {
	return &ProjectPostCommentThumbTO{
		Hcid:    t.Hcid,
		From:    t.From,
		To:      t.To,
		Vote:    t.Vote,
		Time:    t.Time,
		Counter: t.Counter,
	}
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

func (u *DeletedUser) GetTO() Renderable {
	return &DeletedUserTO{
		Counter:    u.Counter,
		Username:   u.Username,
		Time:       u.Time,
		Motivation: u.Motivation,
	}
}

type SpecialUser struct {
	Role    string `gorm:"primary_key:yes" sql:"type:varchar(20)"`
	Counter uint64
}

func (u *SpecialUser) GetTO() Renderable {
	return &SpecialUserTO{
		Role:    u.Role,
		Counter: u.Counter,
	}
}

//TableName returns the table name associated with the structure
func (SpecialUser) TableName() string {
	return "special_users"
}

type SpecialProject struct {
	Role    string `gorm:"primary_key:yes" sql:"type:varchar(20)"`
	Counter uint64
}

func (p *SpecialProject) GetTO() Renderable {
	return &SpecialProjectTO{
		Role:    p.Role,
		Counter: p.Counter,
	}
}

//TableName returns the table name associated with the structure
func (SpecialProject) TableName() string {
	return "special_groups"
}

type PostClassification struct {
	ID    uint64 `gorm:"primary_key:yes"`
	UHpid uint64
	GHpid uint64
	Tag   string `sql:"type:varchar(35)"`
}

func (p *PostClassification) GetTO() Renderable {
	return &PostClassificationTO{
		ID:    p.ID,
		UHpid: p.UHpid,
		GHpid: p.GHpid,
		Tag:   p.Tag,
	}
}

//TableName returns the table name associated with the structure
func (PostClassification) TableName() string {
	return "posts_classifications"
}

type Mention struct {
	ID       uint64 `gorm:"primary_key:yes"`
	UHpid    uint64
	GHpid    uint64
	From     uint64
	To       uint64
	Time     time.Time `sql:"default:NOW()"`
	ToNotify bool
}

func (m *Mention) GetTO() Renderable {
	return &MentionTO{
		ID:       m.ID,
		UHpid:    m.UHpid,
		GHpid:    m.GHpid,
		From:     m.From,
		To:       m.To,
		Time:     m.Time,
		ToNotify: m.ToNotify,
	}
}

//TableName returns the table name associated with the structure
func (Mention) TableName() string {
	return "mentions"
}

//Message view
type Message struct {
	Post
	Type uint8
}

//TableName returns the table name associated with the structure
func (Message) TableName() string {
	return "messages"
}

// OAuth2Client implements the osin.Client interface
type OAuth2Client struct {
	// Surrogated key, should be a random unique value
	ID uint64 `gorm:"primary_key:yes"`
	// Secret is the unique secret associated with a client
	Secret string `sql:"UNIQUE"`
	// RedirectURI is the valid redirection URI associated with a client
	RedirectURI string
	// UserID references User that created this client
	UserID uint64
	User   *User
}

//TableName returns the table name associated with the structure
func (OAuth2Client) TableName() string {
	return "oauth2_clients"
}

type OAuth2AuthorizeData struct {
	// Surrogated key
	ID uint64 `gorm:"primary_key:yes"`
	// ClientID references the client that created this token
	// gorm 1:1 relation
	Client   OAuth2Client
	ClientID uint64
	// Code is the Authorization code
	Code string
	// CreatedAt is the instant of creation of the OAuth2AuthorizeToken
	CreatedAt time.Time
	// ExpiresIn is the seconds from CreatedAt before this token expires
	ExpiresIn uint64
	// State data from request
	State string
	// Scope is the requested scope
	Scope string
	// RedirectUri is the RedirectUri associated with the token
	RedirectURI string
	// UserID is references the User that created the authorization request and thus the AuthorizeData
	UserID uint64
	User   *User
}

//TableName returns the table name associated with the structure
func (OAuth2AuthorizeData) TableName() string {
	return "oauth2_authorize"
}

// OAuth2AccessData is the OAuth2 access data
type OAuth2AccessData struct {
	ID uint64 `gorm:"primary_key:yes"`
	// ClientID references the client that created this token
	// gorm 1:1 relation
	Client   OAuth2Client
	ClientID uint64
	// CreatedAt is the instant of creation of the OAuth2AccessToken
	CreatedAt time.Time
	// ExpiresIn is the seconds from CreatedAt before this token expires
	ExpiresIn uint64
	// RedirectUri is the RedirectUri associated with the token
	RedirectURI string
	// AuthorizeDataID references the AuthorizationData that authorizated this token
	// gorm 1:1 relation
	AuthorizeDataID uint64
	AuthorizeData   *OAuth2AuthorizeData
	// AccessDataID references the Access Data, for refresh token. Can be null
	AccessDataID sql.NullInt64
	AccessData   *OAuth2AccessData
	// RefreshToken is the value by which this token can be renewed. Can be blank.
	RefreshToken   *OAuth2RefreshToken
	RefreshTokenID sql.NullInt64
	// AccessToken is the main value of this tructure, represents the access token
	AccessToken string
	// Scope is the requested scope
	Scope string
	// UserID is references the User that created The access request and thus the AccessData
	UserID uint64
	User   *User
}

//TableName returns the table name associated with the structure
func (OAuth2AccessData) TableName() string {
	return "oauth2_access"
}

type OAuth2RefreshToken struct {
	ID           uint64 `gorm:"primary_key:yes"`
	Token        string `sql:"UNIQUE"`
	AccessData   OAuth2AccessData
	AccessDataID uint64
}

//TableName returns the table name associated with the structure
func (OAuth2RefreshToken) TableName() string {
	return "oauth2_refresh"
}
