/*
Copyright (C) 2016 Paolo Galeone <nessuno@nerdz.eu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package nerdz

import (
	"github.com/galeone/igor"
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
	ClearDefaults()
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
	igor.DBModel
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

// editingMessage interface represents a message while is edited
type editingMessage interface {
	newMessage
	existingMessage
}

// ExistingPost is the interface that wraps the methods common to every existing post
type ExistingPost interface {
	existingMessage
	Comments(...uint) *[]ExistingComment
	CommentsNumber() uint8
	NumericBookmarkers() []uint64
	BookmarkersNumber() uint8
	Bookmarkers() []*User
	NumericLurkers() []uint64
	LurkersNumber() uint8
	Lurkers() []*User
	URL() *url.URL
	//setApiFields(*User)
	IsClosed() bool
	NumericType() uint8
	Type() string
}

// ExistingComment is the interface that wraps the methods common to every existing comment
type ExistingComment interface {
	existingMessage
	Post() (ExistingPost, error)
}

// Helper functions

// createMessage is an helper function. It's used to Init a new message structure
func createMessage(message newMessage, sender, reference uint64, text, language string) error {
	message.ClearDefaults()
	message.SetSender(sender)
	message.SetReference(reference)
	message.SetText(html.EscapeString(text))
	return message.SetLanguage(language)
}

// updateMessage is an helper function. It's used to update a message (requires an editingMessage)
func updateMessage(message editingMessage) error {
	return createMessage(message, message.NumericSender(), message.NumericReference(), message.Text(), message.Language())
}
