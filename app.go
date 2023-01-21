package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tyange/pian-fiber/database"
	"github.com/tyange/pian-fiber/routes"
)

func setUpRoutes(app *fiber.App) {
	routes.SetBurgerRoutes(app)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Listen(":3000")
}
