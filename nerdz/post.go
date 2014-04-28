package nerdz

import "net/url"

type Post interface {
	GetFrom(int64) (*User, error)
	GetTo(int64) (*Board, error)
	GetThumbs() int
	GetComments(...int) interface{}
	GetBookmarkers() []*User
	GetBookmarkersNumber() int
	GetLurkers() []*User
	GetLurkersNumber() int
	GetURL() *url.URL
}
