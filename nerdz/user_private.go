package nerdz

func (user *User) canEdit(message existingMessage) bool {
	return message.IsEditable() && idInSlice(user.Counter, message.NumericOwners())
}

func (user *User) canDelete(message existingMessage) bool {
	return idInSlice(user.Counter, message.NumericOwners())
}
