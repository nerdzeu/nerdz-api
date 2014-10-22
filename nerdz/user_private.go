package nerdz

func (user *User) canEdit(message Message) bool {
	switch message.(type) {
	case *UserPost, *UserPostComment, *ProjectPostComment:
		var editable bool = true
		switch message.(type) {
		case *UserPostComment:
			editable = message.(*UserPostComment).Editable
		case *ProjectPostComment:
			editable = message.(*ProjectPostComment).Editable
		}

		if recipient, err := message.Recipient(); err == nil {
			return editable && recipient.Info().Id == user.Counter
		}

	case *ProjectPost:
		var project *Project
		var board Board
		var err error
		if board, err = message.Recipient(); err == nil {
			projectInfo := (board.(*Project)).ProjectInfo()

			membersOwner := projectInfo.NumericMembers
			membersOwner = append(membersOwner, project.Info().NumericOwner)
			return idInSlice(user.Counter, membersOwner)
		}
	}
	return false
}
