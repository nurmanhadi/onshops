package handlers

import (
	"log"
	"onshops/core/application/dtos"
	"onshops/core/application/services"
	"onshops/pkg"

	"github.com/gofiber/fiber/v2"
)

type PaymentHandler interface {
	AddTransactions(c *fiber.Ctx) error
	PaymentNotification(c *fiber.Ctx) error
}
type PaymentHandlerImpl struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService *services.PaymentService) PaymentHandler {
	return &PaymentHandlerImpl{paymentService: *paymentService}
}
func (h *PaymentHandlerImpl) AddTransactions(c *fiber.Ctx) error {
	body := new(dtos.PaymentRequestDto)
	if err := c.BodyParser(&body); err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "cannot parse to json")
	}
	snapRes, err := h.paymentService.AddPaymentTransactions(*body)
	if err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrBadRequest:
			return pkg.ErrRosponse(c, 400, "bad request", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}
	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "create transaction success",
		"data":    snapRes,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *PaymentHandlerImpl) PaymentNotification(c *fiber.Ctx) error {
	body := new(dtos.PaymentNotificationDto)
	if err := c.BodyParser(&body); err != nil {
		log.Println(err.Error())
		return pkg.ErrRosponse(c, 400, "bad request", "cannot parse to json")
	}
	if err := h.paymentService.AddPayment(*body); err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrBadRequest:
			return pkg.ErrRosponse(c, 400, "bad request", finalErr.Message)
		case *pkg.ErrNotFound:
			return pkg.ErrRosponse(c, 404, "not found", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}
	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "notification receive success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
