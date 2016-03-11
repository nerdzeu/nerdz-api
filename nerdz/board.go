package nerdz

import "github.com/galeone/igor"

// PostlistOptions is used to specify the options of a list of posts.
// The 4 fields are documented and can be combined.
//
// If Following = Followers = true -> show posts FROM user that I follow that follow me back (friends)
// If Older != 0 && Newer != 0 -> find posts BETWEEN this 2 posts
//
// For example:
// - user.UserHome(&PostlistOptions{Followed: true, Language: "en"})
// returns at most the last 20 posts from the english speaking users that I follow.
// - user.UserHome(&PostlistOptions{Followed: true, Following: true, Language: "it", Older: 90, Newer: 50, N: 10})
// returns at most 10 posts, from user's friends, speaking italian, between the posts with hpid 90 and 50
type PostlistOptions struct {
	User      bool   // true -> options for a User post list (false is project post list)
	Following bool   // true -> show posts only FROM following
	Followers bool   // true -> show posts only FROM followers
	Language  string // if Language is a valid 2 characters identifier, show posts from users (users selected enabling/disabling following & folowers) speaking that Language
	N         uint8  // number of post to return (min 1, max 20)
	Older     uint64 // if specified, tells to the function using this struct to return N posts OLDER (created before) than the post with the specified "Older" ID
	Newer     uint64 // if specified, tells to the function using this struct to return N posts NEWER (created after) the post with the specified "Newer" ID
}

// Board is the interface that wraps the methods common to every board.
// Every board has its own Informations and Postlist
type Board interface {
	Info() *Info
	// The return value type of Postlist must be changed by type assertion.
	Postlist(*PostlistOptions) *[]ExistingPost
}

// postlistQueryBuilder returns the same pointer passed as first argument, with new specified options setted
// If the user parameter is present, it's intentend to be the user browsing the website.
// So it will be used to fetch the following list -> so we can easily find the posts on a bord/project/home/ecc made by the users that "user" is following
func postlistQueryBuilder(query *igor.Database, options *PostlistOptions, user ...*User) *igor.Database {
	if options == nil {
		return query.Limit(20)
	}

	if options.N > 0 && options.N < 20 {
		query = query.Limit(int(options.N))
	} else {
		query = query.Limit(20)
	}

	userOK := len(user) == 1 && user[0] != nil
	followersTable := UserFollower{}.TableName()

	if !options.Followers && options.Following && userOK { // from following + me
		query = query.Where(`"from" IN (SELECT "to" FROM `+followersTable+` WHERE "from" = ? UNION SELECT ?)`, user[0].Counter, user[0].Counter)
	} else if !options.Following && options.Followers && userOK { //from followers + me
		query = query.Where(`"from" IN (SELECT "from" FROM `+followersTable+` WHERE "to" = ? UNION SELECT ?)`, user[0].Counter, user[0].Counter)
	} else if options.Following && options.Followers && userOK { //from friends + me
		query = query.Where(`"from" IN ( SELECT ? UNION  (SELECT "to" FROM (SELECT "to" FROM `+
			followersTable+
			` WHERE "from" = ?) AS f INNER JOIN (SELECT "from" FROM `+
			followersTable+
			` WHERE "to" = ?) AS e on f.to = e.from) )`, user[0].Counter, user[0].Counter, user[0].Counter)
	}

	if options.Language != "" {
		if options.User {
			query = query.Where(UserPost{}.TableName()+".lang = ?", options.Language)
		} else {
			query = query.Where(ProjectPost{}.TableName()+".lang = ?", options.Language)
		}
	}

	if options.Older != 0 && options.Newer != 0 {
		query = query.Where("hpid BETWEEN ? AND ?", options.Newer, options.Older)
	} else if options.Older != 0 {
		query = query.Where("hpid < ?", options.Older)
	} else if options.Newer != 0 {
		query = query.Where("hpid > ?", options.Newer)
	}

	return query
}
