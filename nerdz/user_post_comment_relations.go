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

// Implementing Vote interface

// Value returns the vote's value
func (vote *UserPostCommentVote) Value() int8 {
	return vote.Vote
}

// Sender returns the User that casted the vote
func (vote *UserPostCommentVote) Sender() (user *User) {
	user, _ = NewUser(vote.From)
	return
}

// NumericSender returns the ID of the Sender
func (vote *UserPostCommentVote) NumericSender() uint64 {
	return vote.From
}

// Reference returns the reference of the vote
func (vote *UserPostCommentVote) Reference() Reference {
	post, _ := NewUserPostComment(vote.Hcid)
	return post
}

// NumericReference returns the numeric ID of the reference
func (vote *UserPostCommentVote) NumericReference() uint64 {
	return vote.Hcid
}
