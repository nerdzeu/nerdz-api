package nerdz

// getNumeriBlacklist returns a slice containing the counters (IDs) of blacklisted user
func (user *User) getNumericBlacklist() []int64 {
	return append(user.getNumericBlacklisted(), user.getNumericBlacklisting()...)
}

// getNumericBlacklisted returns a slice containing the IDs of users that user (*User) put in his blacklist
func (user *User) getNumericBlacklisted() []int64 {
	var bl []Blacklist
	var blacklist []int64

	db.Find(&bl, &Blacklist{From: user.Counter})
	for _, elem := range bl {
		blacklist = append(blacklist, elem.To)
	}

	//    db.Find(&bl, &Blacklist{From: user.Counter}).Pluck("\"to\"",&blacklist)
	//    db.Model([]Blacklist{}).Where("\"from\" = ?", user.Counter) //.Pluck("\"to\"",&blacklist)
	//    db.Table("blacklist").Where("\"from\" = ?", user.Counter).Pluck("\"to\"",&blacklist)
	return blacklist
}

// getNumericBlacklisting returns a slice  containing the IDs of users that puts user (*User) in their blacklist
func (user *User) getNumericBlacklisting() []int64 {
	var bl []Blacklist
	var blacklist []int64

	db.Find(&bl, &Blacklist{To: user.Counter})
	for _, elem := range bl {
		blacklist = append(blacklist, elem.From)
	}

	//    db.Find(&bl, &Blacklist{To: user.Counter}).Pluck("\"from\"",&blacklist)
	//    db.Model([]Blacklist{}).Where("\"to\" = ?", user.Counter) //.Pluck("\"from\"",&blacklist)
	//    db.Table("blacklist").Where("\"to\" = ?", user.Counter).Pluck("\"from\"",&blacklist)
	return blacklist
}

// getNumericFollowers returns a slice containing the IDs of User that are user's followers
func (user *User) getNumericFollowers() []int64 {
	var fl []UserFollow
	var followers []int64

	db.Find(&fl, UserFollow{To: user.Counter})
	for _, elem := range fl {
		followers = append(followers, elem.From)
	}

	return followers
}

// getNumericFollowing returns a slice containing the IDs of User that user (User *) is following
func (user *User) getNumericFollowing() []int64 {
	var fl []UserFollow
	var followers []int64

	db.Find(&fl, UserFollow{From: user.Counter})
	for _, elem := range fl {
		followers = append(followers, elem.To)
	}

	return followers
}
