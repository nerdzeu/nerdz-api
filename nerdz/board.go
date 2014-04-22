package nerdz

import (
	"github.com/jinzhu/gorm"
	"net/url"
)

// Informations common to all the implementation of Board
type Info struct {
	Id        int64
	Owner     *User
	Followers []*User
	Name      string
	Website   *url.URL
	Image     *url.URL
}

// PostlistOptions is used to specify the options of a list of posts.
// The 4 fields are documented and can be combined.
// For example: GetUserHome(&PostlistOptions{Followed: true, Language: "en"}) returns the last 20 posts from the english speaking users that I follow.
type PostlistOptions struct {
	Following bool   // true -> show posts only from following
	Language  string // if Language is a valid 2 characters identifier, show posts from users speaking Language
	N         int    // number of post to return (min 1, max 20)
	After     int    // if specified, tells to the function using this struct to return N posts after the post with the specified "After" (that can be an hpid, or hcid)
}

// Board is the representation of a generic Board.
// Every board has its own Informations and Postlist
type Board interface {
	GetInfo() *Info
	// The return value type of GetPostlist must be changed by type assertion.
	GetPostlist(*PostlistOptions) interface{}
}

// postlistQueryBuilder returns the same pointer passed as first argument, with new specified options setted
// If the user parameter is present, it's intentend to be the user browsing the website.
// So it will be used to fetch the following list -> so we can easily find the posts on a bord/project/home/ecc made by the users that "user" is following
func postlistQueryBuilder(query *gorm.DB, options *PostlistOptions, user ...*User) *gorm.DB {
	if options == nil {
		return query.Limit(20)
	}

	if options.N > 0 && options.N <= 20 {
		query = query.Limit(options.N)
	} else {
		query = query.Limit(20)
	}

	if options.Following && len(user) == 1 && user[0] != nil {
		following := user[0].getNumericFollowing()
		if len(following) != 0 {
			query = query.Where("\"from\" IN (?)", following)
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
