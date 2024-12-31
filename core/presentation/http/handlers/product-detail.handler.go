package handlers

import (
	"onshops/core/application/dtos"
	"onshops/core/application/services"
	"onshops/pkg"

	"github.com/gofiber/fiber/v2"
)

type OrderDetailsHandler interface {
	AddOrderDetail(c *fiber.Ctx) error
}
type OrderDetailsHandlerImpl struct {
	orderDetailsService services.OrderDetailsService
}

func NewOrderDetailsHandler(orderDetailsService *services.OrderDetailsService) OrderDetailsHandler {
	return &OrderDetailsHandlerImpl{orderDetailsService: *orderDetailsService}
}
func (h *OrderDetailsHandlerImpl) AddOrderDetail(c *fiber.Ctx) error {
	body := new(dtos.OrderDetailRequestDto)
	if err := c.BodyParser(&body); err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "cannot parse to json")
	}
	if err := h.orderDetailsService.AddOrderDetail(*body); err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrBadRequest:
			return pkg.ErrRosponse(c, 400, "bad request", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 400, "bad request", err.Error())
		}
	}
	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "add order detail success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
