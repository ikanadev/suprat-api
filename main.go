package main

import (
	"context"
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/vmkevv/suprat-api/config"
	"github.com/vmkevv/suprat-api/ent"
	"github.com/vmkevv/suprat-api/internal/services"
	"github.com/vmkevv/suprat-api/internal/services/user"
)

func start() {
	config.SetEnvs()

	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error loading ENV variables: %v", err)
	}

	context := context.Background()

	client, err := ent.Open(
		"postgres",
		conf.PostgresConn(),
	)
	if err != nil {
		log.Fatalf("Failed postgres connection: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context); err != nil {
		log.Fatalf("Problem creating database schemas: %v", err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			if err, ok := e.(services.SuprError); ok {
				c.Status(err.Code)
				return c.JSON(err)
			}
			if err != nil {
				c.Status(fiber.StatusInternalServerError)
				return c.JSON(services.SuprError{
					Message: "Server error",
					Detail:  e.Error(),
				})
			}
			return nil
		},
	})
	appV1 := app.Group("/api/v1")

	validator := validator.New()
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	user.NewUserHandler(context, client, validator).ServeHTTP(appV1)

	app.Listen(":8000")
}

func main() {
	start()
}
