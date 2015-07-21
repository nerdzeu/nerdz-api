package routes

const (
	MIN_POSTS = 1
	MAX_POSTS = 20
)

// Response represent the response format of the API
type Response struct {
	Data         interface{}
	Message      string
	HumanMessage string
	Status       uint
	Success      bool
}
