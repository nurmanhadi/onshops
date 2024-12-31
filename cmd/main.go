package main

import (
	"log"
	"onshops/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New()
	app.Use(logger.New())
	config.Initialized(app)

	if err := app.Listen(":5000"); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
