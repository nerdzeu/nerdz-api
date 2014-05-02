package nerdz

// Comment is the interface that wraps the methods common to every existing comment
type ExistingComment interface {
	GetTo() *Board
	GetFrom() *User
	GetPost() *Post
}

// Comment is the interface that represents a generic comment. Wraps the interfaces: ExistingComment and NewMessage
type Comment interface {
	ExistingComment
	NewMessage
}
