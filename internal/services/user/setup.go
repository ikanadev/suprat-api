package user

import "github.com/vmkevv/suprat-api/ent"

// NewUserHandler given a ent client returns a handler
func NewUserHandler(client *ent.Client) Handler {
	return Handler{client: client}
}

func (uh Handler) ServeHTTP() {}
