package user

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/suprat-api/ent"
)

// NewUserHandler given a ent client returns a handler
func NewUserHandler(ctx context.Context, client *ent.Client, validator *validator.Validate) Handler {
	return Handler{
		actions:  actions{db: client},
		validate: validator,
		ctx:      ctx,
	}
}

func (uh Handler) ServeHTTP(app fiber.Router) {
	app.Post("/user", uh.register())
}
