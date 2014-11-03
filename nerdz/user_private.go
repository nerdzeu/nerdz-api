package nerdz

func (user *User) canEdit(message editingMessage) bool {
	return message.Id() > 0 && message.IsEditable() && idInSlice(user.Counter, message.NumericOwners())
}

func (user *User) canDelete(message existingMessage) bool {
	return message.Id() > 0 && idInSlice(user.Counter, message.NumericOwners())
}
