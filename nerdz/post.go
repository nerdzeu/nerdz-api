package nerdz

import "net/url"

// Post is the interface that wraps the methods common to every existing post (already posted)
type ExistingPost interface {
	Message
	Comments(...uint) interface{}
	Bookmarkers() []*User
	NumericBookmarkers() []uint64
	BookmarkersNumber() uint
	Lurkers() []*User
	NumericLurkers() []uint64
	LurkersNumber() uint
	URL(*url.URL) *url.URL
}

// Post is the interface that represents a generic post. Wraps the interfaces: ExistingPost and NewMessage
type Post interface {
	ExistingPost
	NewMessage
}
