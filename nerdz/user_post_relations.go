/*
Copyright (C) 2016-2020 Paolo Galeone <nessuno@nerdz.eu>

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

// UserPostVote: implementing Vote interface

// Value returns the vote's value
func (vote *UserPostVote) Value() int8 {
	return vote.Vote
}

// Sender returns the User that casted the vote
func (vote *UserPostVote) Sender() (user *User) {
	user, _ = NewUser(vote.From)
	return
}

// NumericSender returns the ID of the Sender
func (vote *UserPostVote) NumericSender() uint64 {
	return vote.From
}

// Reference returns the reference of the vote
func (vote *UserPostVote) Reference() Reference {
	post, _ := NewUserPost(vote.Hpid)
	return post
}

// NumericReference returns the numeric ID of the reference
func (vote *UserPostVote) NumericReference() uint64 {
	return vote.Hpid
}

// UserPostBookmark: implementing Bookmark interface

// Sender returns the User that casted the bookmark
func (bookmark *UserPostBookmark) Sender() (user *User) {
	user, _ = NewUser(bookmark.From)
	return
}

// NumericSender returns the ID of the Sender
func (bookmark *UserPostBookmark) NumericSender() uint64 {
	return bookmark.From
}

// Reference returns the reference of the bookmark
func (bookmark *UserPostBookmark) Reference() Reference {
	post, _ := NewUserPost(bookmark.Hpid)
	return post
}

// NumericReference returns the numeric ID of the reference
func (bookmark *UserPostBookmark) NumericReference() uint64 {
	return bookmark.Hpid
}

// UserPostLurk: implementing Lurk interface

// Sender returns the User that casted the lurk
func (lurk *UserPostLurk) Sender() (user *User) {
	user, _ = NewUser(lurk.From)
	return
}

// NumericSender returns the ID of the Sender
func (lurk *UserPostLurk) NumericSender() uint64 {
	return lurk.From
}

// Reference returns the reference of the lurk
func (lurk *UserPostLurk) Reference() Reference {
	post, _ := NewUserPost(lurk.Hpid)
	return post
}

// NumericReference returns the numeric ID of the reference
func (lurk *UserPostLurk) NumericReference() uint64 {
	return lurk.Hpid
}

// UserPostLock: implementing Lock interface

// Sender returns the User that casted the lock
func (lock *UserPostLock) Sender() (user *User) {
	user, _ = NewUser(lock.User)
	return
}

// NumericSender returns the ID of the Sender
func (lock *UserPostLock) NumericSender() uint64 {
	return lock.User
}

// Reference returns the reference of the lurk
func (lock *UserPostLock) Reference() Reference {
	post, _ := NewUserPost(lock.Hpid)
	return post
}

// NumericReference returns the numeric ID of the reference
func (lock *UserPostLock) NumericReference() uint64 {
	return lock.Hpid
}

// UserPostUserLock: implementing Lock interface

// Sender returns the User that casted the lock
func (lock *UserPostUserLock) Sender() (user *User) {
	user, _ = NewUser(lock.From)
	return
}

// NumericSender returns the ID of the Sender
func (lock *UserPostUserLock) NumericSender() uint64 {
	return lock.From
}

// Reference returns the reference of the lurk
func (lock *UserPostUserLock) Reference() Reference {
	post, _ := NewUserPost(lock.Hpid)
	return post
}

// NumericReference returns the numeric ID of the reference
func (lock *UserPostUserLock) NumericReference() uint64 {
	return lock.Hpid
}
