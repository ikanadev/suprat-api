package services

// ResponseWrapper represents a resonse structure
type ResponseWrapper struct {
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

// NewResponse takes message an data and returns an Response struct
func NewResponse(message string, data interface{}) ResponseWrapper {
	return ResponseWrapper{
		Message: message,
		Content: data,
	}
}
