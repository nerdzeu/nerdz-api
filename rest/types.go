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

package rest

import (
	"github.com/nerdzeu/nerdz-api/nerdz"
)

// Response represents the response format of the API
//
// swagger:response apiResponse
type Response struct {
	// The API response data
	// in: body
	Data interface{} `json:"data"`
	// The API generated message
	Message string `json:"message"`
	// The human generated message, easy to understand
	HumanMessage string `json:"humanMessage"`
	// Status Code of the request
	Status uint `json:"status"`
	// Success indicates if the requested succeded
	Success bool `json:"success"`
}

// NewMessage represents a new message from the current user
//
// swagger:parameters addMessage
type NewMessage struct {
	// a Message is a string that contains a new message
	//
	// in: body
	Message string `json:"message"`
	Lang    string `json:"lang, omitempty"`
}

// NewVote represent a new vote from the current user
//
// swagger:parameters addVote
type NewVote struct {
	// a Vote has value in [-1, 0, 1]
	//
	// in: body
	Vote int8 `json:"vote"`
}

// UserInformations represents the user information
//
// swagger:response userInfo
type UserInfo struct {
	// The general information of the User
	// in: body
	Info *nerdz.InfoTO `json:"info"`
	// The contact information of the User
	// in: body
	Contacts *nerdz.ContactInfoTO `json:"contacts"`
	// The personal information of the user
	// in: body
	Personal *nerdz.PersonalInfoTO `json:"personal"`
}

// ProjectInfo represents the project information
//
// swagger:response projectInfo
type ProjectInfo struct {
	// The general information of the project
	// in: body
	Info *nerdz.InfoTO `json:"info"`
}

// PostID is the post ID swagger parameter
//
// swagger:parameters GetUserPost DeleteUserPost EditUserPost GetUserPostComments GetUserPostComment NewUserPostComment EditUserPostComment DeleteUserPostComment GetUserPostVotes NewUserPostVote GetUserPostCommentsVotes NewUserPostCommentVote GetUserPostBookmarks NewUserPostBookmark DeleteUserPostBookmark GetUserPostLurks NewUserPostLurk DeleteUserPostLurk GetUserPostLock NewUserPostLock DeleteUserPostLock NewUserNewPostUserLock DeleteUserPostUserLock EditMePost DeleteMePostComment DeleteMePost EditMeComment NewMePostComment GetMePostVotes NewMePostVote GetMePostCommentsVotes NewMePostCommentVote GetMePostBookmarks NewMePostBookmark DeleteMePostBookmark GetMePostLurks NewMePostLurk DeleteMePostLurk GetMePostLock NewMePostLock DeleteMePostLock NewMeNewPostUserLock DeleteMePostUserLock getProjectPost DeleteProjectPost EditProjectPost getProjectPostComments GetProjectPostComment NewProjectPostComment EditProjectPostComment DeleteProjectPostComment GetProjectPostVotes NewProjectPostVote GetProjectPostCommentsVotes NewProjectPostCommentVote GetProjectPostBookmarks NewProjectPostBookmark DeleteProjectPostBookmark GetProjectPostLurks NewProjectPostLurk DeleteProjectPostLurk GetProjectPostLock NewProjectPostLock DeleteProjectPostLock NewUserNewPostProjectLock DeleteProjectPostUserLock
type PostID struct {
	// a Pid is the post id
	//
	// in: path
	// required: true
	Pid uint64 `json:"pid"`
}

// ID is the ID of the referenced board, user or project
//
// swagger:parameters GetUserPosts GetUserPost NewUserPost DeleteUserPost EditUserPost GetUserPostComments GetUserPostComment NewUserPostComment EditUserPostComment DeleteUserPostComment GetUserInfo GetUserFriends GetUserFollowers GetUserFollowing GetProjectFollowing GetWhitelist GetWhitelisting GetBlacklist GetBlacklisting GetUserPostVotes NewUserPostVote GetUserPostCommentsVotes NewUserPostCommentVote GetUserPostBookmarks NewUserPostBookmark DeleteUserPostBookmark GetUserPostLurks NewUserPostLurk DeleteUserPostLurk GetUserPostLock NewUserPostLock DeleteUserPostLock NewUserNewPostUserLock DeleteUserPostUserLock getProjectPosts getProjectPost NewProjectPost DeleteProjectPost EditProjectPost getProjectPostComments GetProjectPostComment NewProjectPostComment EditProjectPostComment DeleteProjectPostComment getProjectInfo getProjectMembers getProjectFollowers GetProjectPostVotes NewProjectPostVote GetProjectPostCommentsVotes NewProjectPostCommentVote GetProjectPostBookmarks NewProjectPostBookmark DeleteProjectPostBookmark GetProjectPostLurks NewProjectPostLurk DeleteProjectPostLurk GetProjectPostLock NewProjectPostLock DeleteProjectPostLock NewUserNewPostProjectLock DeleteProjectPostUserLock
type ID struct {
	// a ID is the User ID
	//
	// in: path
	// required: true
	ID uint64 `json:"id"`
}

// Target is the ID of the User referenced by the operation
//
// swagger:parameters NewUserNewPostUserLock DeleteUserPostUserLock NewMeNewPostUserLock DeleteMePostUserLock NewMeFollowing DeleteMeFollowing NewProjectFollowing DeleteProjectFollowing NewWhitelisted DeleteWhitelisted NewBlacklisted DeleteBlacklisted NewUserNewPostProjectLock DeleteProjectPostUserLock
type Target struct {
	// a Target is the ID of the User referenced by the operation
	//
	// in: path
	// required: true
	Target uint64 `json:"target"`
}

// CommentID is the ID of the comment
//
// swagger:parameters  GetUserPostComment EditUserPostComment DeleteUserPostComment GetUserPostCommentsVotes NewUserPostCommentVote DeleteMePostComment DeleteMePost EditMeComment GetMePostCommentsVotes NewMePostCommentVote GetProjectPostComment EditProjectPostComment DeleteProjectPostComment GetProjectPostCommentsVotes NewProjectPostCommentVote
type CommentID struct {
	// a Cid is the Comment ID
	//
	// in: path
	// required: true
	Cid uint64 `json:"cid"`
}

// PmID is the ID of the PM
//
// swagger:parameters GetMePm EditMePm DeleteMePm
type PmID struct {
	// a PMID is the ID of the PM
	//
	// in:path
	// required:true
	Pmid uint64 `json:"pmid"`
}

// OtherID is the ID of the other user
//
// swagger:parameters getMeConversation DeleteMePms GetMePm NewMePm EditMePm DeleteMePm
type OtherID struct {
	// Other is the ID of the other user
	//
	// in:path
	// required:true
	Other uint64 `json:"other"`
}
