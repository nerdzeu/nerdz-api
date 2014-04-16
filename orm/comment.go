package orm

type Comment interface {
    GetTo() *Board
    GetFrom() *User
    GetPost() *Post
}
