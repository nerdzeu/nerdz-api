package nerdz

// The Message interface represents a generic message
type Message interface {
    Sender() (*User, error)
	Recipient() (Board, error)
	Thumbs() int
    Text() string
}
