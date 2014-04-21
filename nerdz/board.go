package nerdz

import (
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

// PostlistOptions is used to specify homepage options. The 3 fields are documented and can be combined.
// For example: GetUserHome(&HomepageOptions{Followed: true, Language: "en"}) returns the last 20 posts from the english speaking users that I follow.
// The only mandatory field is Kind
type PostlistOptions struct {
	Following bool   // true -> show posts only from following
	Language  string // if Language is a valid 2 characters identifier, show posts from users speaking Language
	N         int    // number of post to return (min 1, max 20)
	After     int    // if specified, tells to the function using this struct to return N posts after the post with the specified "After" (that can be an hpid, or hcid)
}

type Board interface {
	GetInfo() *Info
	GetPostlist(*PostlistOptions) []Post
}
