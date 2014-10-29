package nerdz

import (
	"html"
	"net/url"
)

// Type definitions for [comment, post, pm]

// newMessage is the interface that wraps methods common to every new mesage
// Implementations: (UserPost, ProjectPost, UserPostComment, ProjectPostComment, Pm)
type newMessage interface {
	SetSender(uint64)
	SetRecipient(uint64)
	SetText(string)
	SetLanguage(string) error
}

// The existingMessage interface represents a generic existing message
type existingMessage interface {
	Id() uint64
	Sender() *User
	NumericSender() uint64
	Recipient() Board
	NumericRecipient() uint64
	Text() string
	IsEditable() bool
	NumericOwners() []uint64
	Owners() []*User
	Revisions() []string
	RevisionsNumber() uint8
	Thumbs() int
	Language() string
}

// Tge editingMessage interface represents a message while is edited
type editingMessage interface {
	newMessage
	existingMessage
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

// createMessage is an helper function. It's used to Init a new message structure
func createMessage(message newMessage, sender, reference uint64, text, language string) error {
	message.SetSender(sender)
	message.SetRecipient(reference)
	message.SetText(html.EscapeString(text))
	return message.SetLanguage(language)
}

// updateMessage is an helper function. It's used to update a message (requires an editingMessage)
func updateMessage(message editingMessage) error {
	return createMessage(message, message.NumericSender(), message.NumericRecipient(), message.Text(), message.Language())
}
