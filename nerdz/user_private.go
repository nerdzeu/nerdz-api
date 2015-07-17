package nerdz

import (
	"github.com/nerdzeu/nerdz-api/utils"
)

// canEdit returns true if user can edit the editingMessage
func (user *User) canEdit(message editingMessage) bool {
	return message.ID() > 0 && message.IsEditable() && utils.InSlice(user.Counter, message.NumericOwners())
}

// canDelete returns trhe if user can delet ethe existingMessage
func (user *User) canDelete(message existingMessage) bool {
	return message.ID() > 0 && utils.InSlice(user.Counter, message.NumericOwners())
}
