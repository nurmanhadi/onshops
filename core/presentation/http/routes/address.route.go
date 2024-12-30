package routes

import (
	"onshops/core/presentation/http/handlers"
	"onshops/core/presentation/http/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AddressRoute(app *fiber.App, h handlers.AddressHandler) {
	address := app.Group("api/v1/customers/:customerId/address", middlewares.AuthGuaard)
	address.Get("/", h.GetAddresses)
	address.Get("/:addressId", h.GetAddressById)
	address.Post("/", h.AddAddress)
	address.Patch("/:addressId", h.UpdateAddress)
	address.Delete("/:addressId", h.DeleteAddress)

}
