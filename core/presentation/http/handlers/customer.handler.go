package handlers

import (
	"onshops/core/application/dtos"
	"onshops/core/application/services"
	"onshops/pkg"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler interface {
	UpdateProfile(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
}
type CustomerHandlerImpl struct {
	customerService services.CustomerService
}

func NewCustomerHandler(customerService *services.CustomerService) CustomerHandler {
	return &CustomerHandlerImpl{customerService: *customerService}
}
func (h *CustomerHandlerImpl) UpdateProfile(c *fiber.Ctx) error {
	customerId := c.Locals("customer_id").(string)
	body := new(dtos.CustomerRequestDto)
	if err := c.BodyParser(&body); err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "cannot parse to json")
	}
	if err := h.customerService.UpdateProfile(customerId, *body); err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrNotFound:
			return pkg.ErrRosponse(c, 404, "not found", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internla server error", err.Error())
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "update profile success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *CustomerHandlerImpl) GetProfile(c *fiber.Ctx) error {
	customerId := c.Locals("customer_id").(string)
	customer, err := h.customerService.GetCustomerById(customerId)
	if err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrNotFound:
			return pkg.ErrRosponse(c, 404, "not found", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internla server error", err.Error())
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "update profile success",
		"data":    customer,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
