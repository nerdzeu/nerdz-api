package nerdz

type Comment interface {
	GetTo() *Board
	GetFrom() *User
	GetPost() *Post
}
