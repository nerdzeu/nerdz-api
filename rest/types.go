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

package rest

import (
	"github.com/nerdzeu/nerdz-api/nerdz"
)

const (
	// MinPosts represents the minimum posts number that can be required in a postList
	MinPosts = 1
	// MaxPosts represents the minimum posts number that can be required in a postList
	MaxPosts = 20
)

// Response represent the response format of the API
type Response struct {
	Data         interface{} `json:"data"`
	Message      string      `json:"message"`
	HumanMessage string      `json:"humanMessage"`
	Status       uint        `json:"status"`
	Success      bool        `json:"success"`
}

// UserInformations represent the userinformation returned by the API
type UserInformations struct {
	Info     *nerdz.InfoTO         `json:"info"`
	Contacts *nerdz.ContactInfoTO  `json:"contacts"`
	Personal *nerdz.PersonalInfoTO `json:"personal"`
}
