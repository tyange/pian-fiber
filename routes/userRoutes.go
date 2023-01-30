package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tyange/pian-fiber/handlers"
)

func SetUserRoutes(app *fiber.App) {
	app.Post("/user/google-auth", handlers.VerifyingCredential)
}
