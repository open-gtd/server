package api

//MessageResponse represents JSON message response body
type MessageResponse struct {
	Message string `json:"message"`
}

//Message returns new message response body, based on error message
func Message(err error) MessageResponse {
	return MessageResponse{Message: err.Error()}
}
