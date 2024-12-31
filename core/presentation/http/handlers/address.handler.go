package handlers

import (
	"onshops/core/application/dtos"
	"onshops/core/application/services"
	"onshops/pkg"

	"github.com/gofiber/fiber/v2"
)

type AddressHandler interface {
	GetAddresses(c *fiber.Ctx) error
	GetAddressById(c *fiber.Ctx) error
	AddAddress(c *fiber.Ctx) error
	UpdateAddress(c *fiber.Ctx) error
	DeleteAddress(c *fiber.Ctx) error
}
type AddressHandlerImpl struct {
	addressService services.AddressService
}

func NewAddressHandler(addressService *services.AddressService) AddressHandler {
	return &AddressHandlerImpl{addressService: *addressService}
}
func (h *AddressHandlerImpl) GetAddresses(c *fiber.Ctx) error {
	customerId := c.Params("customerId")
	address, err := h.addressService.GetAddresses(customerId)
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
		"message": "get address success",
		"data":    address,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *AddressHandlerImpl) GetAddressById(c *fiber.Ctx) error {
	addressId := c.Params("addressId")
	address, err := h.addressService.GetAddressById(addressId)
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
		"message": "get address success",
		"data":    address,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *AddressHandlerImpl) AddAddress(c *fiber.Ctx) error {
	customerId := c.Params("customerId")
	body := new(dtos.AddressAddRequestDto)
	if err := c.BodyParser(&body); err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "cannot parse to json")
	}
	if err := h.addressService.AddAddress(customerId, *body); err != nil {
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
		"message": "add address success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *AddressHandlerImpl) UpdateAddress(c *fiber.Ctx) error {
	addressId := c.Params("addressId")
	customerId := c.Params("customerId")
	body := new(dtos.AddressUpdateRequestDto)
	if err := c.BodyParser(&body); err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "cannot parse to json")
	}
	if err := h.addressService.Updateddress(customerId, addressId, *body); err != nil {
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
		"message": "update address success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *AddressHandlerImpl) DeleteAddress(c *fiber.Ctx) error {
	addressId := c.Params("addressId")
	customerId := c.Params("customerId")
	if err := h.addressService.DeleteAddress(customerId, addressId); err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrNotFound:
			return pkg.ErrRosponse(c, 404, "not found", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "delete address success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
