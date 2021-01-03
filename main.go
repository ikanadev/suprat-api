package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/vmkevv/suprat-api/ent"
	"github.com/vmkevv/suprat-api/internal/services"
	"github.com/vmkevv/suprat-api/internal/services/user"
)

func start() {
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_NAME", "supratdb")
	os.Setenv("DB_PASSWORD", "12345")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")

	conf, err := GetConfig()
	if err != nil {
		log.Fatalf("Error loading ENV variables: %v", err)
	}

	context := context.Background()

	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.DBName,
		),
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
