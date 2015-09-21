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
	SetReference(uint64)
	SetText(string)
	SetLanguage(string) error
}

// Reference represents a reference.
// A comment refers to a user/project post
// A post, refers to a user/project board
type Reference interface {
	ID() uint64
	Language() string
}

// The existingMessage interface represents a generic existing message
type existingMessage interface {
	Reference
	Sender() *User
	NumericSender() uint64
	Reference() Reference
	NumericReference() uint64
	Text() string
	IsEditable() bool
	NumericOwners() []uint64
	Owners() []*User
	Revisions() []string
	RevisionsNumber() uint8
	Thumbs() int
}

// Tge editingMessage interface represents a message while is edited
type editingMessage interface {
	newMessage
	existingMessage
	ClearDefaults()
}

//existingPost is the interface that wraps the methods common to every existing post
type ExistingPost interface {
	existingMessage
	Comments(...uint) interface{}
	CommentsNumber() uint8
	NumericBookmarkers() []uint64
	BookmarkersNumber() uint8
	Bookmarkers() []*User
	NumericLurkers() []uint64
	LurkersNumber() uint8
	Lurkers() []*User
	URL(*url.URL) *url.URL
	//setApiFields(*User)
	IsClosed() bool
}

// existingComment is the interface that wraps the methods common to every existing comment
type existingComment interface {
	existingMessage
	Post() ExistingPost
}

// Helper functions

// createMessage is an helper function. It's used to Init a new message structure
func createMessage(message newMessage, sender, reference uint64, text, language string) error {
	message.SetSender(sender)
	message.SetReference(reference)
	message.SetText(html.EscapeString(text))
	return message.SetLanguage(language)
}

// updateMessage is an helper function. It's used to update a message (requires an editingMessage)
func updateMessage(message editingMessage) error {
	return createMessage(message, message.NumericSender(), message.NumericReference(), message.Text(), message.Language())
}
