/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package nerdz

import "github.com/galeone/igor"

const (
	// MinPosts represents the minimum posts number that can be required in a postList
	MinPosts uint64 = 1
	// MaxPosts represents the maximum posts number that can be required in a postList
	MaxPosts uint64 = 20

	// MinComments represents the minimum comments number that can be required in a commentList
	MinComments uint64 = 1
	// MaxComments represents the maximum comments number that can be required in a commentList
	MaxComments uint64 = 20
)

// PostlistOptions is used to specify the options for a list of posts.
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
	Model      igor.DBModel // igor.DBModel used to apply filter (like language) to avoid conflics while doing joins
	Following  bool         // true -> show posts only FROM following
	Followers  bool         // true -> show posts only FROM followers
	Language   string       // if Language is a valid 2 characters identifier, show posts from users (users selected enabling/disabling following & folowers) speaking that Language
	N          uint8        // number of posts to return
	Older      uint64       // if specified, tells to the function using this struct to return N posts OLDER (created before) than the post with the specified "Older" ID
	OlderModel igor.DBModel // igor.DBModel required when the older identifier is fetched from a view
	Newer      uint64       // if specified, tells to the function using this struct to return N posts NEWER (created after) the post with the specified "Newer" ID
	NewerModel igor.DBModel // igor.DBModel required when the newer identifier is fetched from a view
}

// CommentlistOptions is used to specify the options for a list of comments
type CommentlistOptions struct {
	N     uint8  // number of comments to return
	Older uint64 // if specified, tells to the function that is using this struct to return N comments OLDER (created before) than the comment with the spefified "Older" ID
	Newer uint64 // if specified, tells to the function that is using this struct to return N comments NEWER (created after) than the comment with the spefified "Newer" ID
}

// Board is the interface that wraps the methods common to every board.
// Every board has its own Informations and Postlist
type Board interface {
	Info() *Info
	// The return value type of Postlist must be changed by type assertion.
	Postlist(PostlistOptions) *[]ExistingPost
}

// postlistQueryBuilder returns the same pointer passed as first argument, with new specified options setted
// If the user parameter is present, it's intentend to be the user browsing the website.
// So it will be used to fetch the following list -> so we can easily find the posts on a bord/project/home/ecc made by the users that "user" is following
func postlistQueryBuilder(query *igor.Database, options PostlistOptions, user ...*User) *igor.Database {
	query = query.Limit(int(AtMostPosts(uint64(options.N))))

	userOK := len(user) == 1 && user[0] != nil
	followersTable := UserFollower{}.TableName()

	if !options.Followers && options.Following && userOK { // from following + me
		query = query.Where(`"from" IN (SELECT "to" FROM `+followersTable+` WHERE "from" = ? UNION ALL SELECT ?)`, user[0].Counter, user[0].Counter)
	} else if !options.Following && options.Followers && userOK { //from followers + me
		query = query.Where(`"from" IN (SELECT "from" FROM `+followersTable+` WHERE "to" = ? UNION ALL SELECT ?)`, user[0].Counter, user[0].Counter)
	} else if options.Following && options.Followers && userOK { //from friends + me
		query = query.Where(`"from" IN (SELECT ? UNION ALL (SELECT "to" FROM (SELECT "to" FROM `+
			followersTable+
			` WHERE "from" = ?) AS f INNER JOIN (SELECT "from" FROM `+
			followersTable+
			` WHERE "to" = ?) AS e on f.to = e.from))`, user[0].Counter, user[0].Counter, user[0].Counter)
	}

	if options.Language != "" {
		query = query.Where(options.Model.TableName()+".lang = ?", options.Language)
	}

	messageTpl := Message{}.TableName()

	if options.Older != 0 && options.Newer != 0 {
		if options.Model.TableName() == messageTpl {
			query = query.Where(
				`"time" BETWEEN
				(SELECT "time" FROM `+options.NewerModel.TableName()+` WHERE hpid = ?)
				AND
				(SELECT "time" FROM `+options.OlderModel.TableName()+` WHERE hpid = ?)
				AND hpid NOT IN (?)`, options.Newer, options.Older, []uint64{options.Newer, options.Older})
		} else {
			query = query.Where("hpid BETWEEN ? AND ?", options.Newer, options.Older)
		}
	} else if options.Older != 0 {
		if options.Model.TableName() == messageTpl {
			query = query.Where(
				`"time" <= (
					SELECT "time" FROM `+options.OlderModel.TableName()+` WHERE hpid = ?
				) AND hpid <> ?`, options.Older, options.Older)
		} else {
			query = query.Where("hpid < ?", options.Older)
		}
	} else if options.Newer != 0 {
		if options.Model.TableName() == messageTpl {
			query = query.Where(
				`"time" >= (
					SELECT "time" FROM `+options.NewerModel.TableName()+` WHERE hpid = ?
				) AND hpid <> ?`, options.Newer, options.Newer)
		} else {
			query = query.Where("hpid > ?", options.Newer)
		}
	}

	return query
}

// projectPostlistConditions returns the same pointer passed as first argumet with the project conditions setted
func projectPostlistConditions(query *igor.Database, user *User) *igor.Database {
	var projectPost ProjectPost
	projectPosts := projectPost.TableName()
	users := new(User).TableName()
	projects := new(Project).TableName()
	members := new(ProjectMember).TableName()
	owners := new(ProjectOwner).TableName()

	query = query.Joins("JOIN " + users + " ON " + users + ".counter = " + projectPosts + ".from " +
		"JOIN " + projects + " ON " + projects + ".counter = " + projectPosts + ".to " +
		"JOIN " + owners + " ON " + owners + ".to = " + projectPosts + ".to")

	query = query.Where(`(`+projectPosts+`."from" NOT IN (SELECT "to" FROM blacklist WHERE "from" = ?))`, user.Counter)
	return query.Where("( visible IS TRUE OR "+owners+`.from = ? OR ( ? IN (SELECT "from" FROM `+members+` WHERE "to" = `+projectPosts+`.to) ) )`, user.Counter, user.Counter)
}

// postlistQueryBuilder returns the same pointer passed as first argument, with new specified options setted
func commentlistQueryBuilder(query *igor.Database, options CommentlistOptions) *igor.Database {
	query = query.Limit(int(AtMostComments(uint64(options.N)))).Order("hcid DESC")

	if options.Older != 0 && options.Newer != 0 {
		query = query.Where("hcid BETWEEN ? AND ?", options.Newer, options.Older)
	} else if options.Older != 0 {
		query = query.Where("hcid < ?", options.Older)
	} else if options.Newer != 0 {
		query = query.Where("hcid > ?", options.Newer)
	}

	return query

}
