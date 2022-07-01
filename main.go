package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func handleUser(c *fiber.Ctx) error {

	user := User{
		FirstName: "John",
		LastName:  "Doe",
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {

	user := User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.ID = uuid.New().String()

	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {

	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(requestid.New())
	userGroup := app.Group("/users")

	userGroup.Get("/", handleUser)
	userGroup.Post("/", handleCreateUser)

	app.Listen(":3000")
}
