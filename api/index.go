package handler

import (
	"gofiber-api-vercel/src/configs"
	"gofiber-api-vercel/src/helpers"
	"gofiber-api-vercel/src/routes"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {
	app := fiber.New()

	configs.InitDB()
	helpers.Migration()
	routes.SetupRoutes(app)

	return adaptor.FiberApp(app)
}
