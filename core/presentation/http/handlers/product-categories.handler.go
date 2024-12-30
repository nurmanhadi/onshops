package handlers

import (
	"onshops/core/application/dtos"
	"onshops/core/application/services"
	"onshops/pkg"

	"github.com/gofiber/fiber/v2"
)

type ProductCategoriesHandler interface {
	AddProductCategories(c *fiber.Ctx) error
	DeleteProductCategories(c *fiber.Ctx) error
}
type ProductCategoriesHandlerImpl struct {
	productCategoriesService services.ProductCategoriesService
}

func NewProductCategoriesHandler(productCategoriesService *services.ProductCategoriesService) ProductCategoriesHandler {
	return &ProductCategoriesHandlerImpl{productCategoriesService: *productCategoriesService}
}
func (h *ProductCategoriesHandlerImpl) AddProductCategories(c *fiber.Ctx) error {
	productId := c.Params("productId")
	body := new(dtos.ProductCategoriesRequestDto)
	if err := c.BodyParser(&body); err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "cannot parse to japan")
	}
	if err := h.productCategoriesService.AddProductCategories(productId, *body); err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrBadRequest:
			return pkg.ErrRosponse(c, 400, "bad request", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "add product category success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *ProductCategoriesHandlerImpl) DeleteProductCategories(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")
	productId := c.Params("productId")
	if err := h.productCategoriesService.DeleteProductCategories(productId, categoryId); err != nil {
		return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "delete product category success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
