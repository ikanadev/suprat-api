package services

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

// User the user struct
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}
