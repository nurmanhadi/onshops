package main

import (
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
	app.Listen(":3000")
}
