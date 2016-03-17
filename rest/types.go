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

// UserInfo represent the userinformation returned by the API
type UserInformations struct {
	Info     *nerdz.InfoTO         `json:"info"`
	Contacts *nerdz.ContactInfoTO  `json:"contacts"`
	Personal *nerdz.PersonalInfoTO `json:"personal"`
}
