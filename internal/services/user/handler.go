package user

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/suprat-api/internal/services"
)

// HandlerActions interface which represents all posibles actions
type HandlerActions interface {
	Save(ctx context.Context, name, lastName, email, password string) (services.User, error)
}

// Handler handle all the http logic related to user
type Handler struct {
	actions  HandlerActions
	validate *validator.Validate
	ctx      context.Context
}

func (h Handler) register() fiber.Handler {
	type request struct {
		Name     string `json:"name" validate:"required,gte=2,lte=30"`
		LastName string `json:"lastName" validate:"required,gte=2,lte=30"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=4,lte=80"`
	}
	type response struct {
		User services.User `json:"user"`
	}
	return func(c *fiber.Ctx) error {
		req := request{}
		start := time.Now()
		if err := c.BodyParser(&req); err != nil {
			return err
		}
		fmt.Printf("Parse: %s\n", time.Since(start))
		start = time.Now()
		err := h.validate.Struct(req)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				if err != nil {
					return services.NewSurpErr(fiber.StatusBadRequest, "datos incorrectos", err.Error())
				}
			}
		}
		fmt.Printf("Validate: %s\n", time.Since(start))
		start = time.Now()
		savedUser, err := h.actions.Save(h.ctx, req.Name, req.LastName, req.Email, req.Password)
		fmt.Printf("Save: %s\n", time.Since(start))
		if err != nil {
			return err
		}
		resp := response{
			User: savedUser,
		}
		c.JSON(services.NewResponse("Registrado correctamente", resp))
		return nil
	}
}
