package helpers

import (
	"gofiber-api-vercel/src/configs"
	"gofiber-api-vercel/src/models"
	"log"
)

func Migration() {
	err := configs.DB.AutoMigrate(&models.User{})

	if err != nil {
		// Handle the error, e.g., log it or panic
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}
