package nerdz

import "net/url"

// Post is the interface that wraps the methods common to every existing post (already posted)
type ExistingPost interface {
	Sender() (*User, error)
	Recipient() (Board, error)
	Thumbs() int
	Comments(...int) interface{}
	Bookmarkers() []*User
	BookmarkersNumber() int
	Lurkers() []*User
	LurkersNumber() int
	URL(*url.URL) *url.URL
	Text() string
}

// Post is the interface that represents a generic post. Wraps the interfaces: ExistingPost and NewMessage
type Post interface {
	ExistingPost
	NewMessage
}
