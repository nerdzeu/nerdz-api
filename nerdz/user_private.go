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
func (user *User) canBookmark(message existingPost) bool {
	return message.ID() > 0 && !utils.InSlice(user.Counter, message.NumericBookmarkers())
}

// canLurk returns true if the user haven't lurked the existingPost yet
func (user *User) canLurk(message existingPost) bool {
	return message.ID() > 0 && !utils.InSlice(user.Counter, message.NumericLurkers())
}

// canComment returns true if the user can comment to the existingPost
func (user *User) canComment(message existingPost) bool {
	return !utils.InSlice(user.Counter, message.Sender().NumericBlacklist()) && message.ID() > 0 && !message.IsClosed()
}
