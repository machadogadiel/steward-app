package main

import (
	// "api/database"
	"api/database"
	"api/user"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	ctx := context.Background()

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	database := database.GetDatabase()

	userRepo := user.NewRepo(database)
	userService := user.NewService(*userRepo)

	app.Post("/user", user.AddUser(userService, ctx))
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")

}
