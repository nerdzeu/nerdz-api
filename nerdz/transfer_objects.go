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
	// UserPostID constaint is the numeric identifier of a user post
	// when considered as a generic message
	UserPostID = 1
	// ProjectPostID constaint is the numeric identifier of a project post
	// when considered as a generic message
	ProjectPostID = 0
)

// UserPostsNoNotifyTO represents the TO of UserPostsNoNotify
type UserPostsNoNotifyTO struct {
	original *UserPostsNoNotify
	User     *InfoTO   `json:"user"`
	Hpid     uint64    `json:"hpid"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *UserPostsNoNotifyTO) Original() *UserPostsNoNotify {
	return to.original
}

// UserPostCommentsNoNotifyTO represents the TO of UserPostCommentsNoNotify
type UserPostCommentsNoNotifyTO struct {
	original *UserPostCommentsNoNotify
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Hpid     uint64    `json:"hpid"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *UserPostCommentsNoNotifyTO) Original() *UserPostCommentsNoNotify {
	return to.original
}

// UserPostCommentsNotifyTO  represents the TO of UserPostCommentsNotify
type UserPostCommentsNotifyTO struct {
	original *UserPostCommentsNotify
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Hpid     uint64    `json:"hpid"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *UserPostCommentsNotifyTO) Original() *UserPostCommentsNotify {
	return to.original
}

// BanTO represents the TO of Ban
type BanTO struct {
	original   *Ban
	User       *InfoTO   `json:"user"`
	Motivation string    `json:"motivation"`
	Time       time.Time `json:"time"`
	Counter    uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *BanTO) Original() *Ban {
	return to.original
}

// BlacklistTO represens the TO of Blacklist
type BlacklistTO struct {
	original   *Blacklist
	FromInfo   *InfoTO   `json:"from"`
	ToInfo     *InfoTO   `json:"to"`
	Motivation string    `json:"motivation"`
	Time       time.Time `json:"time"`
	Counter    uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *BlacklistTO) Original() *Blacklist {
	return to.original
}

// WhitelistTO represents the TO of Whitelist
type WhitelistTO struct {
	original *Whitelist
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *WhitelistTO) Original() *Whitelist {
	return to.original
}

//UserFollowerTO represents the TO of UserFollower
type UserFollowerTO struct {
	original *UserFollower
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *UserFollowerTO) Original() *UserFollower {
	return to.original
}

// ProjectNotifyTO represents the TO of ProjectNotify
type ProjectNotifyTO struct {
	original *ProjectNotify
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	Hpid     uint64    `json:"hpid"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *ProjectNotifyTO) Original() *ProjectNotify {
	return to.original
}

// ProjectPostsNoNotifyTO represents the TO of ProjectPostsNoNotify
type ProjectPostsNoNotifyTO struct {
	original *ProjectPostsNoNotify
	User     *InfoTO   `json:"user"`
	Hpid     uint64    `json:"hpid"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *ProjectPostsNoNotifyTO) Original() *ProjectPostsNoNotify {
	return to.original
}

// ProjectPostCommentsNoNotifyTO represents the TO of ProjectPostCommentsNoNotify
type ProjectPostCommentsNoNotifyTO struct {
	original *ProjectPostCommentsNoNotify
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Hpid     uint64    `json:"hpid"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *ProjectPostCommentsNoNotifyTO) Original() *ProjectPostCommentsNoNotify {
	return to.original
}

// ProjectPostCommentsNotifyTO represents the TO of ProjectPostCommentsNotify
type ProjectPostCommentsNotifyTO struct {
	original *ProjectPostCommentsNotify
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Hpid     uint64    `json:"hpid"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *ProjectPostCommentsNotifyTO) Original() *ProjectPostCommentsNotify {
	return to.original
}

// UserTO represents the TO of User
type UserTO struct {
	original         *User
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

// Original returns the original object of the TO
func (to *UserTO) Original() *User {
	return to.original
}

// ProfileTO represents the TO of Profile
type ProfileTO struct {
	original       *Profile
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

// Original returns the original object of the TO
func (to *ProfileTO) Original() *Profile {
	return to.original
}

// PostTO is the Transfor Object of Post.
// It represents the common fields presents in a Post
type PostTO struct {
	original         *Post
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
	Type             boardType `json:"type"`
	CanComment       bool      `json:"canComment"`
	CanBookmark      bool      `json:"canBookmark"`
	CanLurk          bool      `json:"canLurk"`
	CanEdit          bool      `json:"canEdit"`
	CanDelete        bool      `json:"canDelete"`
}

// Original returns the original object of the TO
func (to *PostTO) Original() *Post {
	return to.original
}

// UserPostRevisionTO represents the TO of UserPostRevision
type UserPostRevisionTO struct {
	original *UserPostRevision
	Hpid     uint64    `json:"hpid"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	RevNo    uint16    `json:"revNo"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *UserPostRevisionTO) Original() *UserPostRevision {
	return to.original
}

// UserPostVoteTO represents the TO of UserPostVote
type UserPostVoteTO struct {
	original *UserPostVote
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Vote     int8      `json:"vote"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *UserPostVoteTO) Original() *UserPostVote {
	return to.original
}

// UserPostLurkerTO represents the TO of UserPostLurker
type UserPostLurkerTO struct {
	original *UserPostLurker
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *UserPostLurkerTO) Original() *UserPostLurker {
	return to.original
}

// UserPostCommentTO represents the TO of UserPostComment
type UserPostCommentTO struct {
	original  *UserPostComment
	Hcid      uint64    `json:"hcid"`
	Hpid      uint64    `json:"hpid"`
	FromInfo  *InfoTO   `json:"from"`
	ToInfo    *InfoTO   `json:"to"`
	Message   string    `json:"message"`
	Lang      string    `json:"lang"`
	Time      time.Time `json:"time"`
	CanEdit   bool      `json:"canEdit"`
	CanDelete bool      `json:"canDelete"`
}

// Original returns the original object of the TO
func (to *UserPostCommentTO) Original() *UserPostComment {
	return to.original
}

// UserPostCommentRevisionTO represents the TO of UserPostCommentRevision
type UserPostCommentRevisionTO struct {
	original *UserPostCommentRevision
	Hcid     uint64    `json:"hcid"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	RevNo    int8      `json:"revNo"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *UserPostCommentRevisionTO) Original() *UserPostCommentRevision {
	return to.original
}

// UserPostBookmarkTO represents the TO of UserPostBookmark
type UserPostBookmarkTO struct {
	original *UserPostBookmark
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *UserPostBookmarkTO) Original() *UserPostBookmark {
	return to.original
}

// ConversationTO represents the TO of Conversation
type ConversationTO struct {
	original    *Conversation
	FromInfo    *InfoTO   `json:"from"`
	ToInfo      *InfoTO   `json:"to"`
	LastMessage string    `json:"lastMessage"`
	Time        time.Time `json:"time"`
	ToRead      bool      `json:"toRead"`
}

// Original returns the original object of the TO
func (to *ConversationTO) Original() *Conversation {
	return to.original
}

// PmTO represents the TO of Pm
type PmTO struct {
	original  *Pm
	Pmid      uint64    `json:"pmid"`
	FromInfo  *InfoTO   `json:"from"`
	ToInfo    *InfoTO   `json:"to"`
	Message   string    `json:"message"`
	Lang      string    `json:"lang"`
	ToRead    bool      `json:"toRead"`
	Time      time.Time `json:"time"`
	CanEdit   bool      `json:"canEdit"`
	CanDelete bool      `json:"canDelete"`
}

// Original returns the original object of the TO
func (to *PmTO) Original() *Pm {
	return to.original
}

// ProjectTO represents the TO of Project
type ProjectTO struct {
	original     *Project
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

// Original returns the original object of the TO
func (to *ProjectTO) Original() *Project {
	return to.original
}

// ProjectMemberTO represents the TO of ProjectMember
type ProjectMemberTO struct {
	original *ProjectMember
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *ProjectMemberTO) Original() *ProjectMember {
	return to.original
}

// ProjectOwnerTO represents the TO of ProjectOwner
type ProjectOwnerTO struct {
	original *ProjectOwner
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *ProjectOwnerTO) Original() *ProjectOwner {
	return to.original
}

// ProjectPostRevisionTO represents the TO of ProjectPostRevision
type ProjectPostRevisionTO struct {
	original *ProjectPostRevision
	Hpid     uint64    `json:"hpid"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	RevNo    uint16    `json:"revNo"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *ProjectPostRevisionTO) Original() *ProjectPostRevision {
	return to.original
}

// ProjectPostVoteTO represents the TO of ProjectPostVote
type ProjectPostVoteTO struct {
	original *ProjectPostVote
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	Vote     int8      `json:"vote"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *ProjectPostVoteTO) Original() *ProjectPostVote {
	return to.original
}

// ProjectPostLurkerTO represents the TO of ProjectPostLurker
type ProjectPostLurkerTO struct {
	original *ProjectPostLurker
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *ProjectPostLurkerTO) Original() *ProjectPostLurker {
	return to.original
}

// ProjectPostCommentTO represents the TO of ProjectPostComment
type ProjectPostCommentTO struct {
	original  *ProjectPostComment
	Hcid      uint64    `json:"hcid"`
	Hpid      uint64    `json:"hpid"`
	FromInfo  *InfoTO   `json:"from"`
	ToInfo    *InfoTO   `json:"to"`
	Message   string    `json:"message"`
	Lang      string    `json:"lang"`
	Time      time.Time `json:"time"`
	CanEdit   bool      `json:"canEdit"`
	CanDelete bool      `json:"canDelete"`
}

// Original returns the original object of the TO
func (to *ProjectPostCommentTO) Original() *ProjectPostComment {
	return to.original
}

// ProjectPostCommentRevisionTO represents the TO of ProjectPostCommentRevision
type ProjectPostCommentRevisionTO struct {
	original *ProjectPostCommentRevision
	Hcid     uint64    `json:"hcid"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
	RevNo    uint16    `json:"revNo"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *ProjectPostCommentRevisionTO) Original() *ProjectPostCommentRevision {
	return to.original
}

// ProjectPostBookmarkTO represents the TO of ProjectPostBookmark
type ProjectPostBookmarkTO struct {
	original *ProjectPostBookmark
	Hpid     uint64    `json:"hpid"`
	FromInfo *InfoTO   `json:"from"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *ProjectPostBookmarkTO) Original() *ProjectPostBookmark {
	return to.original
}

// ProjectFollowerTO represents the TO of ProjectFollower
type ProjectFollowerTO struct {
	original *ProjectFollower
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *ProjectFollowerTO) Original() *ProjectFollower {
	return to.original
}

// UserPostCommentVoteTO represents the TO of UserPostCommentVote
type UserPostCommentVoteTO struct {
	original *UserPostCommentVote
	Hcid     uint64  `json:"hcid"`
	User     *InfoTO `json:"user"`
	Vote     int8    `json:"vote"`
	Counter  uint64  `json:"counter"`
}

// Original returns the original object of the TO
func (to *UserPostCommentVoteTO) Original() *UserPostCommentVote {
	return to.original
}

// ProjectPostCommentVoteTO represents the TO of ProjectPostCommentVote
type ProjectPostCommentVoteTO struct {
	original *ProjectPostCommentVote
	Hcid     uint64    `json:"hcid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Vote     int8      `json:"vote"`
	Time     time.Time `json:"time"`
	Counter  uint64    `json:"counter"`
}

// Original returns the original object of the TO
func (to *ProjectPostCommentVoteTO) Original() *ProjectPostCommentVote {
	return to.original
}

// DeletedUserTO represents the TO of DeletedUserTO
type DeletedUserTO struct {
	original   *DeletedUser
	Counter    uint64    `json:"counter"`
	Username   string    `json:"username"`
	Time       time.Time `json:"time"`
	Motivation string    `json:"motivation"`
}

// Original returns the original object of the TO
func (to *DeletedUserTO) Original() *DeletedUser {
	return to.original
}

// SpecialUserTO represents the TO of SpecialUser
type SpecialUserTO struct {
	original *SpecialUser
	Role     string `json:"role"`
	Counter  uint64 `json:"counter"`
}

// Original returns the original object of the TO
func (to *SpecialUserTO) Original() *SpecialUser {
	return to.original
}

// SpecialProjectTO represents the TO of SpecialProject
type SpecialProjectTO struct {
	original *SpecialProject
	Role     string `json:"role"`
	Counter  uint64 `json:"counter"`
}

// Original returns the original object of the TO
func (to *SpecialProjectTO) Original() *SpecialProject {
	return to.original
}

// PostClassificationTO represents the TO of PostClassification
type PostClassificationTO struct {
	original *PostClassification
	ID       uint64 `json:"id"`
	UHpid    uint64 `json:"uHpid"`
	GHpid    uint64 `json:"gHpid"`
	Tag      string `json:"tag"`
}

// Original returns the original object of the TO
func (to *PostClassificationTO) Original() *PostClassification {
	return to.original
}

// MentionTO represents the TO of Mention
type MentionTO struct {
	original *Mention
	ID       uint64    `json:"id"`
	UHpid    uint64    `json:"uHpid"`
	GHpid    uint64    `json:"gHpid"`
	FromInfo *InfoTO   `json:"from"`
	ToInfo   *InfoTO   `json:"to"`
	Time     time.Time `json:"time"`
	ToNotify bool      `json:"toNotify"`
}

// Original returns the original object of the TO
func (to *MentionTO) Original() *Mention {
	return to.original
}

// PersonalInfoTO represents the TO of PersonalInfo
type PersonalInfoTO struct {
	original  *PersonalInfo
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

// Original returns the original object of the TO
func (to *PersonalInfoTO) Original() *PersonalInfo {
	return to.original
}

// ContactInfoTO represents the TO of ContactInfo
type ContactInfoTO struct {
	original *ContactInfo
	Website  string `json:"website"`
	GitHub   string `json:"github"`
	Skype    string `json:"skype"`
	Jabber   string `json:"jabber"`
	Yahoo    string `json:"yahoo"`
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
	Steam    string `json:"steam"`
}

// Original returns the original object of the TO
func (to *ContactInfoTO) Original() *ContactInfo {
	return to.original
}

// InfoTO represents the TO of Info
type InfoTO struct {
	original    *Info
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

// Original returns the original object of the TO
func (to *InfoTO) Original() *Info {
	return to.original
}
