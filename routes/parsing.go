package routes

import (
	"errors"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/utils"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// NewPostlistOptions creates a *nerdz.PostlistOptions from a *http.Request
func NewPostlistOptions(r *http.Request) (*nerdz.PostlistOptions, error) {
	//TODO: fill every field
	var posts_n uint64
	var e error

	n := r.FormValue("n")
	if n == "" {
		posts_n = MAX_POSTS
	} else {
		if posts_n, e = strconv.ParseUint(n, 10, 8); e != nil {
			posts_n = MIN_POSTS
		} else {
			if posts_n > MAX_POSTS {
				posts_n = MAX_POSTS
			}
		}
	}
	return &nerdz.PostlistOptions{
		N: uint8(posts_n),
	}, nil
}

// SelectFields changes the json part of struct tags of in interface{} (that must by a struct or a slice of structs with the right json tags)
// Selecting only specified fields (in *http.Request "fields" value). If "fields" is not present the input parameter is unchanged
func SelectFields(in interface{}, r *http.Request) (*map[string]interface{}, error) {
	ret := make(map[string]interface{})
	Type := reflect.TypeOf(in)

	switch reflect.TypeOf(in).Kind() {
	case reflect.Struct:
		value := reflect.ValueOf(in)
		if field_string := r.FormValue("fields"); field_string != "" {
			fields := strings.Split(field_string, ",")
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
			if m, e := SelectFields(value.Index(i).Interface(), r); e == nil {
				ret[strconv.Itoa(i)] = m
			} else {
				return nil, errors.New(e.Error() + " On field number: " + strconv.Itoa(i))
			}
		}
		return &ret, nil
	}

	return nil, errors.New("input parameter is not a struct or a slice of struct")
}
