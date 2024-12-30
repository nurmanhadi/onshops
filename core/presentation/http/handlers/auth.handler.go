package handlers

import (
	"onshops/core/application/dtos"
	"onshops/core/application/services"
	"onshops/pkg"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	AuthRegister(c *fiber.Ctx) error
	AuthLogin(c *fiber.Ctx) error
	AuthLogout(c *fiber.Ctx) error
}
type AuthHandlerImpl struct {
	authServices services.AuthService
}

func NewAuthHandler(authServices *services.AuthService) AuthHandler {
	return &AuthHandlerImpl{authServices: *authServices}
}
func (h *AuthHandlerImpl) AuthRegister(c *fiber.Ctx) error {
	body := new(dtos.AuthRequestDto)
	if err := c.BodyParser(&body); err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "cannot parse to json")
	}
	if err := h.authServices.AuthRegister(*body); err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrBadRequest:
			return pkg.ErrRosponse(c, 400, "bad request", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "register success",
		"links": map[string]string{
			"self": c.OriginalURL(),
		},
	})
}
func (h *AuthHandlerImpl) AuthLogin(c *fiber.Ctx) error {
	body := new(dtos.AuthRequestDto)
	if err := c.BodyParser(&body); err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "cannot parse to json")
	}
	accessToken, err := h.authServices.AuthLogin(*body)
	if err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrBadRequest:
			return pkg.ErrRosponse(c, 400, "bad request", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "login success",
		"data": fiber.Map{
			"access_token": accessToken,
		},
		"links": map[string]string{
			"self": c.OriginalURL(),
		},
	})
}
func (h *AuthHandlerImpl) AuthLogout(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "logout success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
