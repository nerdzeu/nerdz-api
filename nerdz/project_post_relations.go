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

// ProjectPostVote: implementing Vote interface

// Value returns the vote's value
func (vote *ProjectPostVote) Value() int8 {
	return vote.Vote
}

// Sender returns the User that casted the vote
func (vote *ProjectPostVote) Sender() (user *User) {
	user, _ = NewUser(vote.From)
	return
}

// NumericSender returns the ID of the Sender
func (vote *ProjectPostVote) NumericSender() uint64 {
	return vote.From
}

// Reference returns the reference of the vote
func (vote *ProjectPostVote) Reference() Reference {
	post, _ := NewProjectPost(vote.Hpid)
	return post
}

// NumericReference returns the numeric ID of the reference
func (vote *ProjectPostVote) NumericReference() uint64 {
	return vote.Hpid
}

// ProjectPostBookmark: implementing Bookmark interface

// Sender returns the User that casted the bookmark
func (bookmark *ProjectPostBookmark) Sender() (user *User) {
	user, _ = NewUser(bookmark.From)
	return
}

// NumericSender returns the ID of the Sender
func (bookmark *ProjectPostBookmark) NumericSender() uint64 {
	return bookmark.From
}

// Reference returns the reference of the bookmark
func (bookmark *ProjectPostBookmark) Reference() Reference {
	post, _ := NewProjectPost(bookmark.Hpid)
	return post
}

// NumericReference returns the numeric ID of the reference
func (bookmark *ProjectPostBookmark) NumericReference() uint64 {
	return bookmark.Hpid
}

// ProjectPostLurk: implementing Lurk interface

// Sender returns the User that casted the lurk
func (lurk *ProjectPostLurk) Sender() (user *User) {
	user, _ = NewUser(lurk.From)
	return
}

// NumericSender returns the ID of the Sender
func (lurk *ProjectPostLurk) NumericSender() uint64 {
	return lurk.From
}

// Reference returns the reference of the lurk
func (lurk *ProjectPostLurk) Reference() Reference {
	post, _ := NewProjectPost(lurk.Hpid)
	return post
}

// NumericReference returns the numeric ID of the reference
func (lurk *ProjectPostLurk) NumericReference() uint64 {
	return lurk.Hpid
}

// ProjectPostLock: implementing Lock interface

// Sender returns the User that casted the lock
func (lock *ProjectPostLock) Sender() (user *User) {
	user, _ = NewUser(lock.User)
	return
}

// NumericSender returns the ID of the Sender
func (lock *ProjectPostLock) NumericSender() uint64 {
	return lock.User
}

// Reference returns the reference of the lurk
func (lock *ProjectPostLock) Reference() Reference {
	post, _ := NewProjectPost(lock.Hpid)
	return post
}

// NumericReference returns the numeric ID of the reference
func (lock *ProjectPostLock) NumericReference() uint64 {
	return lock.Hpid
}

// ProjectPostUserLock: implementing Lock interface

// Sender returns the User that casted the lock
func (lock *ProjectPostUserLock) Sender() (user *User) {
	user, _ = NewUser(lock.From)
	return
}

// NumericSender returns the ID of the Sender
func (lock *ProjectPostUserLock) NumericSender() uint64 {
	return lock.From
}

// Reference returns the reference of the lurk
func (lock *ProjectPostUserLock) Reference() Reference {
	post, _ := NewProjectPost(lock.Hpid)
	return post
}

// NumericReference returns the numeric ID of the reference
func (lock *ProjectPostUserLock) NumericReference() uint64 {
	return lock.Hpid
}
