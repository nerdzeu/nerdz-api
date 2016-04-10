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
	"time"
)

const (
	// USER_POST constaint is the numeric identifier of a user post
	// when considered as a generic message
	USER_POST = 1
	// PROJECT_POST constaint is the numeric identifier of a project post
	// when considered as a generic message
	PROJECT_POST = 0
)

// UserPostsNoNotifyTO represents the TO of UserPostsNoNotify
type UserPostsNoNotifyTO struct {
	User    *InfoTO   `json:"user"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

// UserPostCommentsNoNotifyTO represents the TO of UserPostCommentsNoNotify
type UserPostCommentsNoNotifyTO struct {
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Hpid     uint64    `json:"hpid"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// UserPostCommentsNotifyTO  represents the TO of UserPostCommentsNotify
type UserPostCommentsNotifyTO struct {
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Hpid     uint64    `json:"hpid"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// BanTO represents the TO of Ban
type BanTO struct {
	User       *InfoTO   `json:"user"`
	Motivation string    `json:"motivation"`
	Time       time.Time `json:"time"`
	Counter    uint64    `json:"counter"`
}

// BlacklistTO represens the TO of Blacklist
type BlacklistTO struct {
	FromInfo   *InfoTO   `json:"from"`
	ToInfo     *InfoTO   `json:"to"`
	Motivation string    `json:"motivation"`
	Time       time.Time `json:"time"`
	Counter    uint64    `json:"counter"`
}

// WhitelistTO represents the TO of Whitelist
type WhitelistTO struct {
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

//UserFollowerTO represents the TO of UserFollower
type UserFollowerTO struct {
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

// ProjectNotifyTO represents the TO of ProjectNotify
type ProjectNotifyTO struct {
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	Hpid     uint64    `json:"hpid"`
	Counter  uint64    `json:"counter"`
}

// ProjectPostsNoNotifyTO represents the TO of ProjectPostsNoNotify
type ProjectPostsNoNotifyTO struct {
	User    *InfoTO   `json:"user"`
	Hpid    uint64    `json:"hpid"`
	Time    time.Time `json:"time"`
	Counter uint64    `json:"counter"`
}

// ProjectPostCommentsNoNotifyTO represents the TO of ProjectPostCommentsNoNotify
type ProjectPostCommentsNoNotifyTO struct {
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Hpid     uint64    `json:"hpid"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// ProjectPostCommentsNotifyTO represents the TO of ProjectPostCommentsNotify
type ProjectPostCommentsNotifyTO struct {
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Hpid     uint64    `json:"hpid"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// UserTO represents the TO of User
type UserTO struct {
	Counter          uint64    `json:"counter"`
	Last             time.Time `json:"last"`
	NotifyStory      igor.JSON `json:"notifyStory"`
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
	Profile          ProfileTO
}

// ProfileTO represents the TO of Profile
type ProfileTO struct {
	Counter        uint64    `json:"counter"`
	Website        string    `json:"website"`
	Quotes         []string  `json:"quotes"`
	Biography      string    `json:"biography"`
	Interests      []string  `json:"interests"`
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

// PostTO is the Transfor Object of Post.
// It represents the common fields presents in a Post
type PostTO struct {
	Hpid             uint64    `json:"hpid"`
	Pid              uint64    `json:"pid"`
	Message          string    `json:"message"`
	Time             time.Time `json:"time"`
	Lang             string    `json:"lang"`
	News             bool      `json:"news"`
	Closed           bool      `json:"closed"`
	FromInfo         *InfoTO   `json:"from"`
	ToInfo           *InfoTO   `json:"to"`
	Rate             int       `json:"rate"`
	RevisionsCount   uint8     `json:"revisions"`
	CommentsCount    uint8     `json:"comments"`
	BookmarkersCount uint8     `json:"bookmarkers"`
	LurkersCount     uint8     `json:"lurkers"`
	URL              string    `json:"url"`
	Timestamp        int64     `json:"timestamp"`
	CanEdit          bool      `json:"canEdit"`
	CanDelete        bool      `json:"canDelete"`
	CanComment       bool      `json:"canComment"`
	CanBookmark      bool      `json:"canBookmark"`
	CanLurk          bool      `json:"canLurk"`
	Type             boardType `json:"type"`
}

// UserPostRevisionTO represents the TO of UserPostRevision
type UserPostRevisionTO struct {
	Hpid    uint64    `json:"hpid"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	RevNo   uint16    `json:"revNo"`
	Counter uint64    `json:"counter"`
}

// UserPostThumbTO represents the TO of UserPostThumb
type UserPostThumbTO struct {
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Vote     int8      `json:"vote"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// UserPostLurkerTO represents the TO of UserPostLurker
type UserPostLurkerTO struct {
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// UserPostCommentTO represents the TO of UserPostComment
type UserPostCommentTO struct {
	Hcid     uint64    `json:"hcid"`
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	Editable bool      `json:"editable"`
}

// UserPostCommentRevisionTO represents the TO of UserPostCommentRevision
type UserPostCommentRevisionTO struct {
	Hcid    uint64    `json:"hcid"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	RevNo   int8      `json:"revNo"`
	Counter uint64    `json:"counter"`
}

// UserPostBookmarkTO represents the TO of UserPostBookmark
type UserPostBookmarkTO struct {
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// ConversationTO represents the TO of Conversation
type ConversationTO struct {
	FromInfo    *InfoTO   `json:"from"`
	ToInfo      *InfoTO   `json:"to"`
	LastMessage string    `json:"lastMessage"`
	Time        time.Time `json:"time"`
	ToRead      bool      `json:"toRead"`
}

// PmTO represents the TO of Pm
type PmTO struct {
	Pmid     uint64    `json:"pmid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Message  string    `json:"message"`
	ToRead   bool      `json:"toRead"`
	Time     time.Time `json:"time"`
}

// ProjectTO represents the TO of Project
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

// ProjectMemberTO represents the TO of ProjectMember
type ProjectMemberTO struct {
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

// ProjectOwnerTO represents the TO of ProjectOwner
type ProjectOwnerTO struct {
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

// ProjectPostRevisionTO represents the TO of ProjectPostRevision
type ProjectPostRevisionTO struct {
	Hpid    uint64    `json:"hpid"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	RevNo   uint16    `json:"revNo"`
	Counter uint64    `json:"counter"`
}

// ProjectPostThumbTO represents the TO of ProjectPostThumb
type ProjectPostThumbTO struct {
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	Vote     int8      `json:"vote"`
	Counter  uint64    `json:"counter"`
}

// ProjectPostLurkerTO represents the TO of ProjectPostLurker
type ProjectPostLurkerTO struct {
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// ProjectPostCommentTO represents the TO of ProjectPostComment
type ProjectPostCommentTO struct {
	Hcid     uint64    `json:"hcid"`
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	Editable bool      `json:"editable"`
}

// ProjectPostCommentRevisionTO represents the TO of ProjectPostCommentRevision
type ProjectPostCommentRevisionTO struct {
	Hcid    uint64    `json:"hcid"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	RevNo   uint16    `json:"revNo"`
	Counter uint64    `json:"counter"`
}

// ProjectPostBookmarkTO represents the TO of ProjectPostBookmark
type ProjectPostBookmarkTO struct {
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// ProjectFollowerTO represents the TO of ProjectFollower
type ProjectFollowerTO struct {
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

// UserPostCommentThumbTO represents the TO of UserPostCommentThumb
type UserPostCommentThumbTO struct {
	Hcid    uint64  `json:"hcid"`
	User    *InfoTO `json:"user"`
	Vote    int8    `json:"vote"`
	Counter uint64  `json:"counter"`
}

// ProjectPostCommentThumbTO represents the TO of ProjectPostCommentThumb
type ProjectPostCommentThumbTO struct {
	Hcid     uint64    `json:"hcid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Vote     int8      `json:"vote"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// DeletedUserTO represents the TO of DeletedUserTO
type DeletedUserTO struct {
	Counter    uint64    `json:"counter"`
	Username   string    `json:"username"`
	Time       time.Time `json:"time"`
	Motivation string    `json:"motivation"`
}

// SpecialUserTO represents the TO of SpecialUser
type SpecialUserTO struct {
	Role    string `json:"role"`
	Counter uint64 `json:"counter"`
}

// SpecialProjectTO represents the TO of SpecialProject
type SpecialProjectTO struct {
	Role    string `json:"role"`
	Counter uint64 `json:"counter"`
}

// PostClassificationTO represents the TO of PostClassification
type PostClassificationTO struct {
	ID    uint64 `json:"id"`
	UHpid uint64 `json:"uHpid"`
	GHpid uint64 `json:"gHpid"`
	Tag   string `json:"tag"`
}

// MentionTO represents the TO of Mention
type MentionTO struct {
	ID       uint64    `json:"id"`
	UHpid    uint64    `json:"uHpid"`
	GHpid    uint64    `json:"gHpid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
}

// PersonalInfoTO represents the TO of PersonalInfo
type PersonalInfoTO struct {
	ID        uint64    `json:"id"`
	IsOnline  bool      `json:"online"`
	Nation    string    `json:"nation"`
	Timezone  string    `json:"timezone"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Gender    bool      `json:"gender"`
	Birthday  time.Time `json:"birthday"`
	Gravatar  string    `json:"gravatar"`
	Interests []string  `json:"interests"`
	Quotes    []string  `json:"quotes"`
	Biography string    `json:"biography"`
}

// ContactInfoTO represents the TO of ContactInfo
type ContactInfoTO struct {
	Website  string `json:"website"`
	GitHub   string `json:"github"`
	Skype    string `json:"skype"`
	Jabber   string `json:"jabber"`
	Yahoo    string `json:"yahoo"`
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
	Steam    string `json:"steam"`
}

// InfoTO represents the TO of Info
type InfoTO struct {
	ID          uint64    `json:"id"`
	Owner       *InfoTO   `json:"owner"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Website     string    `json:"website"`
	Image       string    `json:"image"`
	Closed      bool      `json:"closed"`
	Type        boardType `json:"type"`
	BoardString string    `json:"board"`
}
