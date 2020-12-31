package user

import "github.com/vmkevv/suprat-api/ent"

// HandlerActions interface which represents all posibles actions
type HandlerActions interface {
	Save(name, lastName, email, password string) (*ent.User, error)
}

type actions struct{}

// Save saves an user to the database
func (a actions) Save(name, lastName, email, password string) (*ent.User, error) {
	return nil, nil
}
