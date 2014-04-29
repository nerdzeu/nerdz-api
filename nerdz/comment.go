package nerdz

// Comment is the interface that wraps the methods common to every comment
type Comment interface {
	GetTo() *Board
	GetFrom() *User
	GetPost() *Post
}
