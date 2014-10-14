package nerdz

// getNuericBookmarkers returns a slice of users' ids that bookmarked the post
func (post *ProjectPost) getNumericBookmarkers() []int64 {
	var users []int64
	db.Model(ProjectBookmark{}).Where(&ProjectBookmark{Hpid: post.Hpid}).Pluck("\"from\"", &users)
	return users
}

// getNumericLurkers returns a slice of users' ids that are lurking the post
func (post *ProjectPost) getNumericLurkers() []int64 {
	var users []int64
	db.Model(ProjectPostLurker{}).Where(&ProjectPostLurker{Hpid: post.Hpid}).Pluck("\"from\"", &users)
	return users
}
