package handlers

import (
	"onshops/core/application/dtos"
	"onshops/core/application/services"
	"onshops/pkg"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	GetOrders(c *fiber.Ctx) error
	GetOrderById(c *fiber.Ctx) error
	AddOrder(c *fiber.Ctx) error
	UpdateOrder(c *fiber.Ctx) error
	DeleteOrderById(c *fiber.Ctx) error
}
type OrderHandlerImpl struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) OrderHandler {
	return &OrderHandlerImpl{orderService: *orderService}
}
func (h *OrderHandlerImpl) GetOrders(c *fiber.Ctx) error {
	customerId := c.Params("customerId")
	orders, err := h.orderService.GetOrders(customerId)
	if err != nil {
		return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "get orders success",
		"data":    orders,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *OrderHandlerImpl) GetOrderById(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	order, err := h.orderService.GetOrderById(orderId)
	if err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrNotFound:
			return pkg.ErrRosponse(c, 404, "not found", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "get order success",
		"data":    order,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *OrderHandlerImpl) AddOrder(c *fiber.Ctx) error {
	body := new(dtos.OrderRequestDto)
	if err := h.orderService.AddOrder(*body); err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrBadRequest:
			return pkg.ErrRosponse(c, 400, "bad request", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}
	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "add order success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *OrderHandlerImpl) UpdateOrder(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	body := new(dtos.OrderUpdateRequestDto)
	if err := h.orderService.UpdateOrder(orderId, *body); err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrBadRequest:
			return pkg.ErrRosponse(c, 400, "bad request", finalErr.Message)
		case *pkg.ErrNotFound:
			return pkg.ErrRosponse(c, 404, "not found", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "update order success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *OrderHandlerImpl) DeleteOrderById(c *fiber.Ctx) error {
	orderId := c.Params("orderId")
	if err := h.orderService.DeleteOrder(orderId); err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrNotFound:
			return pkg.ErrRosponse(c, 404, "not found", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "delete order success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
