package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tyange/pian-fiber/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Listen(":3000")
}
