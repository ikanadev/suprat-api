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

// SuprError this will represent an espected error
type SuprError struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Code    int    `json:"-"`
}

// Error to implement the error interface
func (se SuprError) Error() string {
	return se.Message
}

// NewSurpErr return a SuprError
func NewSurpErr(code int, msg, detail string) SuprError {
	return SuprError{
		Message: msg,
		Detail:  detail,
		Code:    code,
	}
}
