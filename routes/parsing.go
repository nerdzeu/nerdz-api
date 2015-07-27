package routes

import (
	"github.com/nerdzeu/nerdz-api/nerdz"
	"net/http"
	"reflect"
	"strconv"
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

// SelectFields changes the json part of struct tags of in interface{} (that must be of type struct)
// Selecting only specified fields (in *http.Request "fields" value). If "fields" is not present the input parameter is unchanged
func SelectFields(in interface{}, r *http.Request) {
	// TODO
	Value := reflect.ValueOf(in)
	Type := Value.Type()
	if Type != reflect.Struct {
		return
	}
	elem := reflect.Indirect(reflect.New(Type))

}
