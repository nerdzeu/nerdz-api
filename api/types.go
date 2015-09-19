package api

const (
	MinPosts = 1
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
