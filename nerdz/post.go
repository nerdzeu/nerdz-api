package nerdz

type Post interface {
	GetFrom(int64) (*User, error)
	GetTo(int64) (*Board, error)
	GetThumbs() int
	GetComments(...int) interface{}
}
