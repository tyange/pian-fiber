package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/tyange/pian-fiber/database"
	"github.com/tyange/pian-fiber/routes"
)

var SessionStore *session.Store

func setUpRoutes(app *fiber.App) {
	routes.SetBurgerRoutes(app)
	routes.SetUserRoutes(app)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Use(cors.New())

	setUpRoutes(app)

	app.Listen(":8080")
}
