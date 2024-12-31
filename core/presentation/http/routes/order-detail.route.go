package routes

import (
	"onshops/core/presentation/http/handlers"
	"onshops/core/presentation/http/middlewares"

	"github.com/gofiber/fiber/v2"
)

func OrderDetailsRoute(app *fiber.App, h handlers.OrderDetailsHandler) {
	api := app.Group("/api/v1")
	customer := api.Group("/customers/:customerId")
	order := customer.Group("/orders/:orderId")
	orderDetail := order.Group("/order-details", middlewares.AuthGuaard)
	orderDetail.Post("/", h.AddOrderDetail)
}
