package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vmkevv/suprat-api/ent"
	"github.com/vmkevv/suprat-api/ent/user"
	"github.com/vmkevv/suprat-api/internal/services"
)

// actions is an struct which implements the HandlerActions interface
type actions struct {
	db *ent.Client
}

// Save saves an user to the database
func (a actions) Save(ctx context.Context, name, lastName, email, password string) (services.User, error) {
	var savedUser services.User
	exists, err := a.db.User.Query().Where(user.EmailEQ(email)).Exist(ctx)
	if err != nil {
		return savedUser, err
	}
	if exists {
		return savedUser, services.NewSurpErr(fiber.StatusConflict, "Ya existe una cuenta con ese email", "")
	}
	hashedPasswd, err := HashPassword(password)
	if err != nil {
		return savedUser, err
	}
	entUser, err := a.db.User.Create().SetName(name).SetLastName(lastName).SetEmail(email).SetPassword(hashedPasswd).Save(ctx)
	if err != nil {
		return savedUser, err
	}
	savedUser.Name = entUser.Name
	savedUser.LastName = entUser.LastName
	savedUser.Email = entUser.Email
	savedUser.ID = int64(entUser.ID)
	return savedUser, nil
}
