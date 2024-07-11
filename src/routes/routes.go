package routes

import (
	"gofiber-api-vercel/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/users", controllers.GetUsers)
	api.Post("/users", controllers.RegisterUser)
}
