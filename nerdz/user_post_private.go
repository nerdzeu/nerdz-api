package nerdz

// NuericBookmarkers returns a slice of users' ids that bookmarked the post
func (post *UserPost) NumericBookmarkers() []uint64 {
	var users []uint64
	db.Model(UserBookmark{}).Where(&UserBookmark{Hpid: post.Hpid}).Pluck("\"from\"", &users)
	return users
}

// NumericLurkers returns a slice of users' ids that are lurking the post
func (post *UserPost) NumericLurkers() []uint64 {
	var users []uint64
	db.Model(UserPostLurker{}).Where(&UserPostLurker{Hpid: post.Hpid}).Pluck("\"from\"", &users)
	return users
}
