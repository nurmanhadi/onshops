package routes

import (
	"onshops/core/presentation/http/handlers"
	"onshops/core/presentation/http/middlewares"

	"github.com/gofiber/fiber/v2"
)

func CustomerRoute(app *fiber.App, h handlers.CustomerHandler) {
	customer := app.Group("api/v1/customers", middlewares.AuthGuaard)
	customer.Patch("/", h.UpdateProfile)
	customer.Get("/", h.GetProfile)
}
