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
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"strings"

	"github.com/labstack/echo"
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
	}
	message := "Success"
	return c.JSON(http.StatusOK, &Response{
		Data:         ret,
		HumanMessage: message,
		Message:      message,
		Status:       http.StatusOK,
		Success:      true,
	})
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

// InvalidScopeResponse prints a c.JSON response and returns a error
// asserting that the required scope is missing from the accepted scopes
func InvalidScopeResponse(requiredScope string, c echo.Context) error {
	message := "Required scope (" + requiredScope + ") is missing"
	return c.JSON(http.StatusUnauthorized, &Response{
		HumanMessage: message,
		Message:      message,
		Status:       http.StatusBadRequest,
		Success:      false,
	})
}
