package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/tyange/pian-fiber/database"
	"github.com/tyange/pian-fiber/routes"
	"github.com/tyange/pian-fiber/store"
)

var SessionStore *session.Store

func main() {
	database.ConnectDb()

	app := fiber.New()

	store.SetSession()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))

	routes.SetupRoutes(app)

	app.Listen(":8080")
}
