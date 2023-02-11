package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tyange/pian-fiber/handlers"
)

func SetBurgerRoutes(app *fiber.App) {
	app.Get("/burger", handlers.AuthMiddleware(), handlers.GetAllBurger)
	app.Get("/burger/:id", handlers.GetBurger)
	app.Post("/burger", handlers.AddBurger)
	app.Put("/burger/:id", handlers.EditBurger)
	app.Delete("/burger/:id", handlers.DeleteBurger)
}
