package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tyange/pian-fiber/handlers"
)

func SetBurgerRoutes(app *fiber.App) {
	app.Get("/burger", handlers.GetAllBurger)
	app.Get("/burger/:id", handlers.GetBurger)
	app.Put("/burger/:id", handlers.EditBurger)
}
