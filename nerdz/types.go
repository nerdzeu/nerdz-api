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
	"net/mail"
	"net/url"
	"time"
)

// PersonalInfo is the struct that contains all the personal info of an user
type PersonalInfo struct {
	IsOnline  bool
	Nation    string
	Timezone  string
	Username  string
	Name      string
	Surname   string
	Gender    bool
	Birthday  time.Time
	Gravatar  *url.URL
	Interests []string
	Quotes    []string
	Biography string
}

// GetTO returns its Transfer Object
func (p *PersonalInfo) GetTO() Renderable {
	return &PersonalInfoTO{
		IsOnline:  p.IsOnline,
		Nation:    p.Nation,
		Timezone:  p.Timezone,
		Username:  p.Username,
		Name:      p.Name,
		Surname:   p.Surname,
		Gender:    p.Gender,
		Birthday:  p.Birthday,
		Gravatar:  p.Gravatar.String(),
		Interests: p.Interests,
		Quotes:    p.Quotes,
		Biography: p.Biography,
	}
}

// ContactInfo is the struct that contains all the contact info of an user
type ContactInfo struct {
	Website  *url.URL
	GitHub   *url.URL
	Skype    string
	Jabber   string
	Yahoo    *mail.Address
	Facebook *url.URL
	Twitter  *url.URL
	Steam    string
}

// GetTO returns its Transfer Object
func (c *ContactInfo) GetTO() Renderable {
	var yahoo, facebook, twitter, github, website string
	if c.Website != nil {
		website = c.Website.String()
	}
	if c.GitHub != nil {
		github = c.GitHub.String()
	}
	if c.Yahoo != nil {
		yahoo = c.Yahoo.String()
	}
	if c.Facebook != nil {
		facebook = c.Facebook.String()
	}
	if c.Twitter != nil {
		twitter = c.Twitter.String()
	}
	return &ContactInfoTO{
		Website:  website,
		GitHub:   github,
		Skype:    c.Skype,
		Jabber:   c.Jabber,
		Yahoo:    yahoo,
		Facebook: facebook,
		Twitter:  twitter,
		Steam:    c.Steam,
	}
}

// Template is the representation of a nerdz website template
type Template struct {
	Number uint8
	Name   string
}

// BoardInfo is that struct that contains all the informations related to the user's board
type BoardInfo struct {
	Language       string
	Template       *Template
	MobileTemplate *Template
	Dateformat     string
	IsClosed       bool
	Private        bool
	Whitelist      []*User
	UserScript     *url.URL
}

// Info contains the informations common to every board
// Used in API output to give user/project basic informations
type Info struct {
	ID          uint64
	Owner       *Info
	Name        string
	Username    string
	Website     *url.URL
	Image       *url.URL
	Closed      bool
	Type        boardType
	BoardString string
}

// GetTO returns its Transfer Object
func (i *Info) GetTO() Renderable {
	var owner Renderable
	if i.Owner != nil {
		owner = i.Owner.GetTO()
	}
	var website, image string
	if i.Image != nil {
		image = i.Image.String()
	}
	if i.Website != nil {
		website = i.Website.String()
	}
	return &InfoTO{
		ID:          i.ID,
		Owner:       owner,
		Name:        i.Name,
		Username:    i.Username,
		Website:     website,
		Image:       image,
		Closed:      i.Closed,
		Type:        i.Type,
		BoardString: i.BoardString,
	}
}
