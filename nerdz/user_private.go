package nerdz

import (
	"github.com/nerdzeu/nerdz-api/utils"
	"log"
)

func (user *User) canEdit(message editingMessage) bool {
	return message.Id() > 0 && message.IsEditable() && utils.InSlice(user.Counter, message.NumericOwners())
}

func (user *User) canDelete(message existingMessage) bool {
	return message.Id() > 0 && utils.InSlice(user.Counter, message.NumericOwners())
}
