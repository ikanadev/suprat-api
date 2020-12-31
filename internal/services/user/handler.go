package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/suprat-api/ent"
)

// Handler handle all the http logic related to user
type Handler struct {
	client  *ent.Client
	actions HandlerActions
}

func (h Handler) register() fiber.Handler {
	type request struct {
		Name     string `json:"name"`
		LastName string `json:"lastName"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		Message string `json:"message"`
	}
	return func(c *fiber.Ctx) error {
		req := request{}
		if err := c.BodyParser(req); err != nil {
			return err
		}
		savedUser, err := h.actions.Save(req.Name, req.LastName, req.Email, req.Password)
		if err != nil {
			return err
		}
		c.JSON(response{Message: "User " + savedUser.Name + " saved"})
		return nil
	}
}
