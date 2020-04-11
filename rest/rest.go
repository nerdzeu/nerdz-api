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

package rest

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/utils"
)

// sf is the recursive function used to build the structure neeeded by SelectFields
func sf(in interface{}, c echo.Context) (interface{}, error) {

	in = reflect.Indirect(reflect.ValueOf(in)).Interface()
	Type := reflect.TypeOf(in)

	switch Type.Kind() {
	case reflect.Struct:
		ret := make(map[string]interface{})
		value := reflect.ValueOf(in)
		if fieldString := c.QueryParam("fields"); fieldString != "" {
			fields := strings.Split(fieldString, ",")
			for _, field := range fields {
				fieldName := utils.UpperFirst(field)
				if structField, ok := Type.FieldByName(fieldName); ok {
					jsonTag := structField.Tag.Get("json")
					if jsonTag != "-" && jsonTag != "" {
						ret[strings.Split(jsonTag, ",")[0]] = value.FieldByName(fieldName).Interface()
					}
				} else {
					// Else check if json field name is different from struct field name (with first letter in uppercase)
					// this is the user expected behaviour, but we prefer the above approach to speed up the process
					var found bool
					for i := 0; i < Type.NumField(); i++ {
						jsonTag := Type.Field(i).Tag.Get("json")
						if strings.Split(jsonTag, ",")[0] == field {
							ret[field] = value.Field(i).Interface()
							found = true
							break
						}
					}
					if !found {
						return nil, fmt.Errorf("Field %s does not exists", field)
					}
				}
			}
		} else {
			for i := 0; i < Type.NumField(); i++ {
				jsonTag := Type.Field(i).Tag.Get("json")
				if jsonTag != "-" && jsonTag != "" {
					ret[strings.Split(jsonTag, ",")[0]] = value.Field(i).Interface()
				}
			}
		}
		return &ret, nil

	case reflect.Slice:
		value := reflect.ValueOf(in)
		ret := make([]interface{}, value.Len())
		for i := 0; i < value.Len(); i++ {
			if m, e := sf(value.Index(i).Elem().Interface(), c); e == nil {
				ret[i] = m
			} else {
				return nil, fmt.Errorf(`Error "%s" on field number %d`, e.Error(), i)
			}
		}
		return &ret, nil
	}

	return nil, errors.New("input parameter is not a struct or a slice of struct")
}

// SelectFields changes the json part of struct tags of in interface{} (that must by a struct or a slice of structs with the right json tags)
// Selecting only specified fields (in the query string "fields" value). If "fields" is not present the input parameter is unchanged
// returns error when there's a problem with some required fileld.
// otherwies returns nil and ends the request, printing the c.JSON of the input value, with its field selected
func SelectFields(in interface{}, c echo.Context) error {
	var ret interface{}
	var e error

	if ret, e = sf(in, c); e != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: e.Error(),
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
		return e
	}
	message := "Success"
	c.JSON(http.StatusOK, &Response{
		Data:         ret,
		HumanMessage: message,
		Message:      message,
		Status:       http.StatusOK,
		Success:      true,
	})
	return nil
}

// IsGranted returns true if the c.Get("scopes") slice contains the scope or
// there's a scompatible scope into the slice
func IsGranted(scope string, c echo.Context) bool {
	parts := strings.Split(scope, ":")
	switch parts[0] {
	case "profile_messages", "project_messages":
		return IsGranted("messages:"+parts[1], c)

	case "profile_comments":
		return IsGranted("profile_messages:"+parts[1], c)

	case "project_comments":
		return IsGranted("project_messages:"+parts[1], c)
	}

	scopes := c.Get("scopes").([]string)
	i := sort.SearchStrings(scopes, scope)
	if i < len(scopes) && scopes[i] == scope {
		return true
	}

	scope = "base:" + parts[1]
	i = sort.SearchStrings(scopes, scope)
	return i < len(scopes) && scopes[i] == scope
}

// InvalidScopeResponse prints a JSON response and returns a error
// asserting that the required scope is missing from the accepted scopes
func InvalidScopeResponse(requiredScope string, c echo.Context) error {
	message := "Required scope (" + requiredScope + ") is missing"
	c.JSON(http.StatusUnauthorized, &Response{
		HumanMessage: message,
		Message:      message,
		Status:       http.StatusBadRequest,
		Success:      false,
	})
	return errors.New(message)
}

// User extract "id" from the url parameter, parse it and returns
// the User if the "me" (in the context) user is allowed to see it.
// Otherwise returns an error
func User(userID string, c echo.Context) (*nerdz.User, error) {
	var id uint64
	var e error
	if id, e = strconv.ParseUint(c.Param(userID), 10, 64); e != nil {
		c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "Invalid user identifier specified",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
		return nil, e
	}

	var user *nerdz.User
	if user, e = nerdz.NewUser(id); e != nil {
		c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "User does not exists",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
		return nil, e
	}

	me := c.Get("me").(*nerdz.User)
	if !me.CanSee(user) {
		message := "You can't see the required profile"
		c.JSON(http.StatusUnauthorized, &Response{
			HumanMessage: message,
			Message:      message,
			Status:       http.StatusUnauthorized,
			Success:      false,
		})
		return nil, errors.New(message)
	}
	return user, nil
}

// Project extract "id" from the url parameter, parse it and returns
// the Project if the "me" (in the context) user is allowed to see it.
// Otherwise returns an error
func Project(projectID string, c echo.Context) (*nerdz.Project, error) {
	var id uint64
	var e error
	if id, e = strconv.ParseUint(c.Param(projectID), 10, 64); e != nil {
		c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "Invalid project identifier specified",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
		return nil, e
	}

	var project *nerdz.Project
	if project, e = nerdz.NewProject(id); e != nil {
		c.JSON(http.StatusBadRequest, &Response{
			HumanMessage: "Project does not exists",
			Message:      e.Error(),
			Status:       http.StatusBadRequest,
			Success:      false,
		})
		return nil, e
	}

	me := c.Get("me").(*nerdz.User)
	if !me.CanSee(project) {
		message := "You can't see the required project"
		c.JSON(http.StatusUnauthorized, &Response{
			HumanMessage: message,
			Message:      message,
			Status:       http.StatusUnauthorized,
			Success:      false,
		})
		return nil, errors.New(message)
	}
	return project, nil
}

// GetUserInfo returns the *UserInfo of the user
func GetUserInfo(user *nerdz.User) *UserInfo {
	var info UserInfo
	info.Info = user.Info().GetTO()
	info.Contacts = user.ContactInfo().GetTO()
	info.Personal = user.PersonalInfo().GetTO()
	return &info
}

// GetUsersInfo returns a slice of *Interfations
func GetUsersInfo(users []*nerdz.User) (usersInfo []*UserInfo) {
	for _, u := range users {
		usersInfo = append(usersInfo, GetUserInfo(u))
	}
	return
}

// GetProjectInfo returns the *ProjectInfo of the project
func GetProjectInfo(project *nerdz.Project) *ProjectInfo {
	var info ProjectInfo
	info.Info = project.Info().GetTO()
	return &info
}

// GetProjectsInfo returns a slice of *Interfations
func GetProjectsInfo(projects []*nerdz.Project) (projectsInfo []*ProjectInfo) {
	for _, u := range projects {
		projectsInfo = append(projectsInfo, GetProjectInfo(u))
	}
	return
}
