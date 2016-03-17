package rest

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/nerdzeu/nerdz-api/nerdz"
	"github.com/nerdzeu/nerdz-api/utils"
)

// atMost parses a echo.Context, extract the GET parameter "n" and check if n is
// a valid uint64 value between min and max.
// In that case, it returns "n", otherwise it returns max
func atMost(c echo.Context, min, max uint64) (ret uint64) {
	var e error
	n := c.Query("n")
	if n == "" {
		ret = max
	} else {
		if ret, e = strconv.ParseUint(n, 10, 8); e != nil {
			ret = min
		} else {
			if ret > max {
				ret = max
			}
		}
	}
	return
}

// NewPostlistOptions creates a *nerdz.PostlistOptions from a echo.Context
// handle GET parameters:
// fing: if setted, requires posts from following users
// fers: if setted, requires posts from followers users
// lang: if setted to a supported language (nerdz.Configuration.Languages), requires
//       posts in that language
// older: if setted to an existing hpid, requires posts older than the "older" value
// newer: if setted to an existing hpid, requires posts newer than the "newer" value
func NewPostlistOptions(c echo.Context) (*nerdz.PostlistOptions, error) {
	var following bool
	var followers bool
	var language string
	var older uint64
	var newer uint64

	// legal parameters
	fing := c.Query("fing")
	fers := c.Query("fers")
	lang := c.Query("lang")
	old := c.Query("older")
	new := c.Query("newer")

	postsN := atMost(c, MinPosts, MaxPosts)

	if fing != "" {
		following = true
	}

	if fers != "" {
		followers = true
	}

	if lang == "" {
		language = ""
	} else {
		if !utils.InSlice(lang, nerdz.Configuration.Languages) {
			return nil, errors.New("Not supported language " + lang)
		}
		language = lang
	}

	older, _ = strconv.ParseUint(old, 10, 64)
	newer, _ = strconv.ParseUint(new, 10, 64)

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
func SelectFields(in interface{}, c echo.Context) (*map[string]interface{}, error) {
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
