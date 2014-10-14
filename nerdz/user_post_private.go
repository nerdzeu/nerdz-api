package nerdz

// getNuericBookmarkers returns a slice of users' ids that bookmarked the post
func (post *UserPost) getNumericBookmarkers() []int64 {
	var users []int64
	db.Model(UserBookmark{}).Where(&UserBookmark{Hpid: post.Hpid}).Pluck("\"from\"", &users)
	return users
}

// getNumericLurkers returns a slice of users' ids that are lurking the post
func (post *UserPost) getNumericLurkers() []int64 {
	var users []int64
	db.Model(UserPostLurker{}).Where(&UserPostLurker{Hpid: post.Hpid}).Pluck("\"from\"", &users)
	return users
}
