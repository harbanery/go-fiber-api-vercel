package main

import (
	"gofiber-api-vercel/src/configs"
	"gofiber-api-vercel/src/helpers"
	"gofiber-api-vercel/src/routes"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	app := fiber.New()

	configs.InitDB()
	helpers.Migration()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app := fiber.New()

	configs.InitDB()
	helpers.Migration()
	routes.SetupRoutes(app)

	adaptor.FiberApp(app)(w, r)
}
