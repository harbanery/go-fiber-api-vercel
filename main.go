package main

import (
	"gofiber-api-vercel/src/configs"
	"gofiber-api-vercel/src/helpers"
	"gofiber-api-vercel/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.InitDB()
	helpers.Migration()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
