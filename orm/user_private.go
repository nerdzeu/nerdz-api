package orm

import (
	"github.com/jinzhu/gorm"
)

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

// homeQueryBuilder returns the same pointer passed as first argument, with new specified options setted
func (user *User) homeQueryBuilder(query *gorm.DB, options *PostlistOptions) *gorm.DB {
	if options.N > 0 && options.N <= 20 {
		query = query.Limit(options.N)
	} else {
		query = query.Limit(20)
	}

	if options.Following {
		following := user.getNumericFollowing()
		if len(following) != 0 {
			query = query.Where("\"from\" IN (?)", user.getNumericFollowing())
		}
	}

	if options.Language != "" {
		query = query.Where(&User{Lang: options.Language})
	}

	if options.After != 0 {
		query = query.Where("hpid < ?", options.After)
	}

	return query
}
