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

// NewPost is the interface that wraps methods common to every new post
type NewPost interface {
	SetTo(interface{}) (Board, error)
	SetMessage(string) error
}

// Post is the interface that represents a generic post. Wraps the interfaces: ExistingPost and NewPost
type Post interface {
	ExistingPost
	NewPost
}
