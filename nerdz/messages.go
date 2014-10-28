package nerdz

import "net/url"

// Type definitions for [comment, post, pm]

// newMessage is the interface that wraps methods common to every new mesage
// Implementations: (UserPost, ProjectPost, UserPostComment, ProjectPostComment, Pm)
type newMessage interface {
	SetSender(uint64)
	SetRecipient(uint64)
	SetText(string) error
}

// The existingMessage interface represents a generic existing message
type existingMessage interface {
	Id() uint64
	Sender() (*User, error)
	Recipient() (Board, error)
	Text() string
	IsEditable() bool
	NumericOwners() []uint64
	Owners() []*User
	Modifications() []string
	ModificationsNumber() uint8
	Thumbs() int
}

// exisistingPost is the interface that wraps the methods common to every existing post
type existingPost interface {
	existingMessage
	Comments(...uint) interface{}
	NumericBookmarkers() []uint64
	BookmarkersNumber() uint
	Bookmarkers() []*User
	NumericLurkers() []uint64
	LurkersNumber() uint
	Lurkers() []*User
	URL(*url.URL) *url.URL
}

// existingComment is the interface that wraps the methods common to every existing comment
type existingComment interface {
	existingMessage
	Post() existingPost
}

// Helper functions

// NewMessage is an helper functions. It's used to Init a new message structure
func NewMessage(message newMessage, sender, reference uint64, text string) error {
	message.SetSender(sender)
	message.SetRecipient(reference)
	if err := message.SetText(text); err != nil {
		return err
	}
	return nil
}
