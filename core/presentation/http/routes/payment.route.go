package routes

import (
	"onshops/core/presentation/http/handlers"
	"onshops/core/presentation/http/middlewares"

	"github.com/gofiber/fiber/v2"
)

func PaymentRoute(app *fiber.App, h handlers.PaymentHandler) {
	api := app.Group("/api/v1")
	customer := api.Group("/customers/:customerId")
	order := customer.Group("/orders/:orderId")
	payment := order.Group("/payments", middlewares.AuthGuaard)
	payment.Post("/", h.AddTransactions)
	api.Post("/notifications/payments", h.PaymentNotification)
}
