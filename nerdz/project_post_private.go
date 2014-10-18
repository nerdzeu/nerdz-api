package nerdz

// NuericBookmarkers returns a slice of users' ids that bookmarked the post
func (post *ProjectPost) NumericBookmarkers() []uint64 {
	var users []uint64
	db.Model(ProjectBookmark{}).Where(&ProjectBookmark{Hpid: post.Hpid}).Pluck("\"from\"", &users)
	return users
}

// NumericLurkers returns a slice of users' ids that are lurking the post
func (post *ProjectPost) NumericLurkers() []uint64 {
	var users []uint64
	db.Model(ProjectPostLurker{}).Where(&ProjectPostLurker{Hpid: post.Hpid}).Pluck("\"from\"", &users)
	return users
}
