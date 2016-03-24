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
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/utils"
)

// sf is the recursive function used to build the structure neeeded by selectFields
func sf(in interface{}, c echo.Context) (*map[string]interface{}, error) {
	ret := make(map[string]interface{})
	in = reflect.Indirect(reflect.ValueOf(in)).Interface()
	Type := reflect.TypeOf(in)

	switch Type.Kind() {
	case reflect.Struct:
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
					return nil, errors.New(field + " does not exists")
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
		for i := 0; i < value.Len(); i++ {
			if m, e := sf(value.Index(i).Elem().Interface(), c); e == nil {
				ret[strconv.Itoa(i)] = m
			} else {
				return nil, errors.New(e.Error() + " On field number: " + strconv.Itoa(i))
			}
		}
		return &ret, nil
	}

	return nil, errors.New("input parameter is not a struct or a slice of struct")
}

// selectFields changes the json part of struct tags of in interface{} (that must by a struct or a slice of structs with the right json tags)
// Selecting only specified fields (in the query string "fields" value). If "fields" is not present the input parameter is unchanged
// returns error when there's a problem with some required fileld.
// otherwies returns nil and ends the request, printing the c.JSON of the input value, with its field selected
func selectFields(in interface{}, c echo.Context) error {
	var ret *map[string]interface{}
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
