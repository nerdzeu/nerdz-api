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

// getUserInfo returns the *UserInformations of the user
func getUserInfo(user *nerdz.User) *UserInformations {
	var info UserInformations
	info.Info = user.Info().GetTO().(*nerdz.InfoTO)
	info.Contacts = user.ContactInfo().GetTO().(*nerdz.ContactInfoTO)
	info.Personal = user.PersonalInfo().GetTO().(*nerdz.PersonalInfoTO)
	return &info
}
