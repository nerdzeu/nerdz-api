package orm

import (
	"github.com/jinzhu/gorm"
)

// getNumeriBlacklist returns a slice containing the counters (IDs) of blacklisted user
func (user *User) getNumericBlacklist() []int64 {
	bl := user.GetBlacklist()
	var blacklist []int64
	for _, u := range bl {
		blacklist = append(blacklist, u.Counter)
	}

	return blacklist
}

// homeQueryBuilder returns the same pointer passed as first argument, with new specified options setted
func (user *User) homeQueryBuilder(query *gorm.DB, options *PostlistOptions) *gorm.DB {
	if options.N > 0 && options.N <= 20 {
		query = query.Limit(options.N)
	} else {
		query = query.Limit(20)
	}

	if options.Following {
		fl := user.GetFollowing()
		var following []int64
		for _, u := range fl {
			following = append(following, u.Counter)
		}
		query = query.Where("\"from\" IN (?)", following)
	}

	if options.Language != "" {
		query = query.Where(&User{Lang: options.Language})
	}

	if options.After != 0 {
		query = query.Where("hpid < ?", options.After)
	}

	return query
}
