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
	"github.com/nerdzeu/nerdz-api/utils"
)

// canEdit returns true if user can edit the editingMessage
func (user *User) canEdit(message editingMessage) bool {
	return message.ID() > 0 && message.IsEditable() && utils.InSlice(user.Counter, message.NumericOwners())
}

// canDelete returns true if user can delete the existingMessage
func (user *User) canDelete(message existingMessage) bool {
	return message.ID() > 0 && utils.InSlice(user.Counter, message.NumericOwners())
}

// canBookmark returns true if user haven't bookamrked to existingPost yet
func (user *User) canBookmark(message ExistingPost) bool {
	return message.ID() > 0 && !utils.InSlice(user.Counter, message.NumericBookmarkers())
}

// canLurk returns true if the user haven't lurked the existingPost yet
func (user *User) canLurk(message ExistingPost) bool {
	return message.ID() > 0 && !utils.InSlice(user.Counter, message.NumericLurkers())
}

// canComment returns true if the user can comment to the existingPost
func (user *User) canComment(message ExistingPost) bool {
	return !utils.InSlice(user.Counter, message.Sender().NumericBlacklist()) && message.ID() > 0 && !message.IsClosed()
}
