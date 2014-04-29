package nerdz

import "net/url"

// Post is the interface that wraps the methods common to every post
type Post interface {
	GetFrom(int64) (*User, error)
	GetTo(int64) (*Board, error)
	GetThumbs() int
	GetComments(...int) interface{}
	GetBookmarkers() []*User
	GetBookmarkersNumber() int
	GetLurkers() []*User
	GetLurkersNumber() int
	GetURL(url.URL) *url.URL
}
