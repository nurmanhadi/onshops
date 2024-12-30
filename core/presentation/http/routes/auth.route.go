package routes

import (
	"onshops/core/presentation/http/handlers"
	"onshops/core/presentation/http/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App, h handlers.AuthHandler) {
	auth := app.Group("api/v1/auth")
	auth.Post("/login", h.AuthLogin)
	auth.Post("/register", h.AuthRegister)
	auth.Post("/logout", middlewares.AuthGuaard, h.AuthLogout)
}
