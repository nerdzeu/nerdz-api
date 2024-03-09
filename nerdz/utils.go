/*
Copyright (C) 2016-2020 Paolo Galeone <nessuno@nerdz.eu>

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
	"github.com/nerdzeu/nerdz-api/utils"
)

// Users returns a slice of pointer to User, fetched from its Ids
func Users(ids []uint64) []*User {
	var users []*User
	for _, elem := range ids {
		user, _ := NewUser(elem)
		users = append(users, user)
	}
	return users
}

// Projects returns a slice of pointer to Project, fetched from its Ids
func Projects(ids []uint64) []*Project {
	var projects []*Project
	for _, elem := range ids {
		project, _ := NewProject(elem)
		projects = append(projects, project)
	}
	return projects
}

// Infos returns a slice of pointer to Info
func Infos(slice interface{}) []*Info {
	var infos []*Info

	switch boards := slice.(type) {
	case []*User:
		for _, elem := range boards {
			infos = append(infos, elem.Info())
		}
	case []*Project:
		for _, elem := range boards {
			infos = append(infos, elem.Info())
		}
	}
	return infos
}

// AtMostPosts returns a uint8 that's the number of posts to be retrieved
func AtMostPosts(n uint64) uint8 {
	return uint8(utils.AtMost(n, MinPosts, MaxPosts))
}

// AtMostComments returns a uint64 that's the number of comments to be retrieved
func AtMostComments(n uint64) uint8 {
	return uint8(utils.AtMost(n, MinComments, MaxComments))
}

// AtMostPms returns a uint64 that's the number of pms to be retrieved
func AtMostPms(n uint64) uint8 {
	return uint8(utils.AtMost(n, MinPms, MaxPms))
}
