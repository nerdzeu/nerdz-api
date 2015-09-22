package api

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/utils"
)

// NewPostlistOptions creates a *nerdz.PostlistOptions from a *http.Request
func NewPostlistOptions(c *echo.Context) (*nerdz.PostlistOptions, error) {
	var postsN uint64
	var following bool
	var followers bool
	var language string
	var older uint64
	var newer uint64
	var e error

	// legal parameters
	n := c.Query("n")
	fing := c.Query("fing")
	fers := c.Query("fers")
	lang := c.Query("lang")
	old := c.Query("older")
	new := c.Query("newer")

	if n == "" {
		postsN = MaxPosts
	} else {
		if postsN, e = strconv.ParseUint(n, 10, 8); e != nil {
			postsN = MinPosts
		} else {
			if postsN > MaxPosts {
				postsN = MaxPosts
			}
		}
	}

	if fing == "" {
		following = false
	} else {
		following = true
	}

	if fers == "" {
		followers = false
	} else {
		followers = true
	}

	if lang == "" {
		language = ""
	} else {
		// TODO: check if lang is a valid language.
		language = lang
	}

	if old == "" {
		older = 0
	} else {
		if older, e = strconv.ParseUint(old, 10, 64); e != nil {
			older = 0
		}
	}

	if new == "" {
		newer = 0
	} else {
		if newer, e = strconv.ParseUint(new, 10, 64); e != nil {
			newer = 0
		}
	}

	return &nerdz.PostlistOptions{
		Following: following,
		Followers: followers,
		Language:  language,
		N:         uint8(postsN),
		Older:     older,
		Newer:     newer,
	}, nil
}

// SelectFields changes the json part of struct tags of in interface{} (that must by a struct or a slice of structs with the right json tags)
// Selecting only specified fields (in *http.Request "fields" value). If "fields" is not present the input parameter is unchanged
func SelectFields(in interface{}, c *echo.Context) (*map[string]interface{}, error) {
	ret := make(map[string]interface{})
	Type := reflect.TypeOf(in)

	switch reflect.TypeOf(in).Kind() {
	case reflect.Struct:
		value := reflect.ValueOf(in)
		if fieldString := c.Query("fields"); fieldString != "" {
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
			value := reflect.ValueOf(in)
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
			if m, e := SelectFields(value.Index(i).Elem().Interface(), c); e == nil {
				ret[strconv.Itoa(i)] = m
			} else {
				return nil, errors.New(e.Error() + " On field number: " + strconv.Itoa(i))
			}
		}
		return &ret, nil
	}

	return nil, errors.New("input parameter is not a struct or a slice of struct")
}
