package nerdz

import "net/url"

// Post is the interface that wraps the methods common to every existing post (already posted)
type ExistingPost interface {
	GetFrom() (*User, error)
	GetTo() (Board, error)
	GetThumbs() int
	GetComments(...int) interface{}
	GetBookmarkers() []*User
	GetBookmarkersNumber() int
	GetLurkers() []*User
	GetLurkersNumber() int
	GetURL(*url.URL) *url.URL
	GetMessage() string
}

// Post is the interface that represents a generic post. Wraps the interfaces: ExistingPost and NewMessage
type Post interface {
	ExistingPost
	NewMessage
}
