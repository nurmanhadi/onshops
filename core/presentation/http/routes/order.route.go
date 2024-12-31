package routes

import (
	"onshops/core/presentation/http/handlers"
	"onshops/core/presentation/http/middlewares"

	"github.com/gofiber/fiber/v2"
)

func OrderRoute(app *fiber.App, h handlers.OrderHandler) {
	api := app.Group("/api/v1")
	customer := api.Group("/customers/:customerId")
	order := customer.Group("/orders", middlewares.AuthGuaard)
	order.Get("/", h.GetOrders)
	order.Get("/:orderId", h.GetOrderById)
	order.Post("/", h.AddOrder)
	order.Patch("/:orderId", h.UpdateOrder)
	order.Delete("/:orderId", h.DeleteOrderById)
}
