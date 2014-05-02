package nerdz

// NewMessage is the interface that wraps methods common to every new mesage (post, comment, pm)
type NewMessage interface {
	SetTo(interface{}) error
	SetMessage(string) error
}
