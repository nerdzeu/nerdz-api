package nerdz

// NewMessage is the interface that wraps methods common to every new mesage (post, comment, pm)
type NewMessage interface {
	SetRecipient(interface{}) error
	SetText(string) error
}

// NewMessage is an helper functions. It's used to Init a new message structure
func NewMessageInit(newMessage NewMessage, other interface{}, message string) error {
	var e error = nil

	if e = newMessage.SetText(message); e != nil {
		return e
	}

	if e = newMessage.SetRecipient(other); e != nil {
		return e
	}

	return e
}
