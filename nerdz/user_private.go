package nerdz

// getNumeriBlacklist returns a slice containing the counters (IDs) of blacklisted user
func (user *User) getNumericBlacklist() []int64 {
	return append(user.getNumericBlacklisted(), user.getNumericBlacklisting()...)
}

// getNumericBlacklisted returns a slice containing the IDs of users that user (*User) put in his blacklist
func (user *User) getNumericBlacklisted() []int64 {
	var blacklist []int64
	db.Model(Blacklist{}).Where(&Blacklist{From: user.Counter}).Pluck("\"to\"", &blacklist)
	return blacklist
}

// getNumericBlacklisting returns a slice  containing the IDs of users that puts user (*User) in their blacklist
func (user *User) getNumericBlacklisting() []int64 {
	var blacklist []int64
	db.Model(Blacklist{}).Where(&Blacklist{To: user.Counter}).Pluck("\"from\"", &blacklist)
	return blacklist
}

// getNumericFollowers returns a slice containing the IDs of User that are user's followers
func (user *User) getNumericFollowers() []int64 {
	var followers []int64
	db.Model(UserFollow{}).Where(UserFollow{To: user.Counter}).Pluck("\"from\"", &followers)
	return followers
}

// getNumericFollowing returns a slice containing the IDs of User that user (User *) is following
func (user *User) getNumericFollowing() []int64 {
	var following []int64
	db.Model(UserFollow{}).Where(&UserFollow{From: user.Counter}).Pluck("\"to\"", &following)
	return following
}
