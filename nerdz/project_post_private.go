package nerdz

// getNuericBookmarkers returns a slice of users' ids that bookmarked the post
func (post *ProjectPost) getNumericBookmarkers() []uint64 {
	var users []uint64
	db.Model(ProjectBookmark{}).Where(&ProjectBookmark{Hpid: post.Hpid}).Pluck("\"from\"", &users)
	return users
}

// getNumericLurkers returns a slice of users' ids that are lurking the post
func (post *ProjectPost) getNumericLurkers() []uint64 {
	var users []uint64
	db.Model(ProjectPostLurker{}).Where(&ProjectPostLurker{Hpid: post.Hpid}).Pluck("\"from\"", &users)
	return users
}
