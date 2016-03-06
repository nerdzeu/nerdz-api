package nerdz

import (
	"errors"
	"time"
)

// PmConfig represent the configuration used to fetch a Pm list
type PmConfig struct {
	// TRUE: PM messages ordered in descending order using timestamp
	// FALSE: PM messages ordered in ascending order using timestamp
	DescOrder bool `json:"descOrder"`
	// number of messages returned (default: 0 - all the pms messages)
	Limit uint64 `json:"limit"`
	// used in combination with Limit grant the possibility to return
	// a fraction of the whole pms
	Offset uint64 `json:"offset"`
	// TRUE: Returns PM messages that should be read
	// FALSE: Returns PM messages that have already read
	ToRead bool `json:"toRead"`
}

// NewPmConfig creates a new PmConfig struct
func NewPmConfig() *PmConfig {
	return &PmConfig{}
}

// WithDescOrder sets the descendant order to PmConfig
func (pmConf *PmConfig) WithDescOrder(descOrder bool) *PmConfig {
	pmConf.DescOrder = descOrder
	return pmConf
}

// WithLimit adds the offset to PmConfig
func (pmConf *PmConfig) WithLimit(limit uint64) *PmConfig {
	pmConf.Limit = limit
	return pmConf
}

// WithOffset add the offset to PmConfig
func (pmConf *PmConfig) WithOffset(offset uint64) *PmConfig {
	pmConf.Offset = offset
	return pmConf
}

// WithToRead add the toRead flag to PmConfig
func (pmConf *PmConfig) WithToRead(toRead bool) *PmConfig {
	pmConf.ToRead = toRead
	return pmConf
}

// Conversation represents the details about a single private conversation between two users
type Conversation struct {
	From   string    `json:"from"`
	Time   time.Time `json:"time"`
	ToRead bool      `json:"toRead"`
}

// NewPm initializes a Pm struct
func NewPm(pmid uint64) (pm *Pm, e error) {
	pm = new(Pm)
	Db().First(pm, pmid)

	if pm.Pmid != pmid {
		return nil, errors.New("Invalid pmid")
	}

	return pm, nil
}

// Implementing newMessage interface

// SetSender sets the source of the pm (the user ID)
func (pm *Pm) SetSender(id uint64) {
	pm.From = id
}

// SetReference sets the destionation of the pm: user ID
func (pm *Pm) SetReference(id uint64) {
	pm.To = id
}

// SetText set the text of the message
func (pm *Pm) SetText(message string) {
	pm.Message = message
}

// SetLanguage set the language of the pm (useless)
func (pm *Pm) SetLanguage(language string) error {
	return nil
}

// ClearDefaults set to the go's default values the fields with default sql values
func (pm *Pm) ClearDefaults() {
	pm.Time = time.Time{}
}

// Implementing existingMessage interface

// ID returns the User Post ID
func (pm *Pm) ID() uint64 {
	return pm.Pmid
}

// Language returns the message language
func (pm *Pm) Language() (lang string) {
	return lang
}

// NumericSender returns the id of the sender user
func (pm *Pm) NumericSender() uint64 {
	return pm.From
}

// Sender returns the sender *User
func (pm *Pm) Sender() *User {
	user, _ := NewUser(pm.NumericSender())
	return user
}

// NumericReference returns the id of the recipient user
func (pm *Pm) NumericReference() uint64 {
	return pm.To
}

// Reference returns the recipient *User
func (pm *Pm) Reference() Reference {
	user, _ := NewUser(pm.NumericReference())
	return user
}

// Text returns the pm message
func (pm *Pm) Text() string {
	return pm.Message
}

// IsEditable returns true if the pm is editable
func (pm *Pm) IsEditable() bool {
	return false
}

// NumericOwners returns a slice of ids of the owner of the pms (the ones that can perform actions)
func (pm *Pm) NumericOwners() []uint64 {
	return []uint64{pm.To, pm.From}
}

// Owners returns a slice of *User representing the users who own the pm
func (pm *Pm) Owners() (ret []*User) {
	return Users(pm.NumericOwners())
}

// Revisions returns all the revisions of the message
func (pm *Pm) Revisions() (modifications []string) {
	return
}

// RevisionsNumber returns the number of the revisions
func (pm *Pm) RevisionsNumber() uint8 {
	return 0
}

// Thumbs returns the pm's thumbs value
func (pm *Pm) Thumbs() int {
	return 0
}
