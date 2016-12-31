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

// Implementing Vote interface

// Value returns the vote's value
func (vote *ProjectPostCommentVote) Value() int8 {
	return vote.Vote
}

// Sender returns the User that casted the vote
func (vote *ProjectPostCommentVote) Sender() (user *User) {
	user, _ = NewUser(vote.From)
	return
}

// NumericSender returns the ID of the Sender
func (vote *ProjectPostCommentVote) NumericSender() uint64 {
	return vote.From
}

// Reference returns the reference of the vote
func (vote *ProjectPostCommentVote) Reference() Reference {
	post, _ := NewProjectPostComment(vote.Hcid)
	return post
}

// NumericReference returns the numeric ID of the reference
func (vote *ProjectPostCommentVote) NumericReference() uint64 {
	return vote.Hcid
}
