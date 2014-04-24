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
	Following bool // true -> show posts only FROM following
	Followers bool // true -> show posts only FROM followers
	// Following = Followers = true -> show posts FROM user that I follow that follow me back
	Language string // if Language is a valid 2 characters identifier, show posts from users (users selected enabling/disabling following & folowers) speaking that Language
	N        int    // number of post to return (min 1, max 20)
	Older    int64  // if specified, tells to the function using this struct to return N posts OLDER (created before) than the post with the specified "Older" ID
	Newer    int64  // if specified, tells to the function using this struct to return N posts NEWER (created after) the post with the specified "Newer"" ID
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

	if options.N > 0 && options.N < 20 {
		query = query.Limit(options.N)
	} else {
		query = query.Limit(20)
	}

	userOK := len(user) == 1 && user[0] != nil

	if !options.Followers && options.Following && userOK { // from following + me
		following := user[0].getNumericFollowing()
		if len(following) != 0 {
			query = query.Where("\"from\" IN (? , ?)", following, user[0].Counter)
		}
	} else if !options.Following && options.Followers && userOK { //from followers + me
		followers := user[0].getNumericFollowers()
		if len(followers) != 0 {
			query = query.Where("\"from\" IN (? , ?)", followers, user[0].Counter)
		}
	} else if options.Following && options.Followers && userOK { //from friends + me
		follows := new(UserFollow).TableName()
		query = query.Where("\"from\" IN ( (SELECT ?) UNION  (SELECT \"to\" FROM (SELECT \"to\" FROM "+follows+" WHERE \"from\" = ?) AS f INNER JOIN (SELECT \"from\" FROM "+follows+" WHERE \"to\" = ?) AS e on f.to = e.from) )", user[0].Counter, user[0].Counter, user[0].Counter)
	}

	if options.Language != "" {
		query = query.Where(&User{Lang: options.Language})
	}

	if options.Older != 0 {
		query = query.Where("hpid < ?", options.Older)
	}

	if options.Newer != 0 {
		query = query.Where("hpid > ?", options.Newer)
	}

	return query
}
