package services

// User the user struct
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}
