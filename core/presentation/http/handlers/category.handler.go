package handlers

import (
	"onshops/core/application/dtos"
	"onshops/core/application/services"
	"onshops/pkg"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler interface {
	GetCategories(c *fiber.Ctx) error
	GetCategoryById(c *fiber.Ctx) error
	AddCategory(c *fiber.Ctx) error
	UpdateCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
}
type CategoryHandlerImpl struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(categoryService *services.CategoryService) CategoryHandler {
	return &CategoryHandlerImpl{categoryService: *categoryService}
}
func (h *CategoryHandlerImpl) GetCategories(c *fiber.Ctx) error {
	categories, err := h.categoryService.GetCategories()
	if err != nil {
		return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "get categories success",
		"data":    categories,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *CategoryHandlerImpl) GetCategoryById(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")
	category, err := h.categoryService.GetCategoriByID(categoryId)
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
		"message": "get category success",
		"data":    category,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *CategoryHandlerImpl) AddCategory(c *fiber.Ctx) error {
	body := new(dtos.CategoryRequestDto)
	if err := c.BodyParser(&body); err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "cannot parse to json")
	}
	if err := h.categoryService.AddCategory(*body); err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrBadRequest:
			return pkg.ErrRosponse(c, 400, "bad request", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}
	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "add category success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *CategoryHandlerImpl) UpdateCategory(c *fiber.Ctx) error {
	body := new(dtos.CategoryRequestDto)
	categoryId := c.Params("categoryId")
	if err := c.BodyParser(&body); err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "cannot parse to json")
	}
	if err := h.categoryService.UpdateCategory(categoryId, *body); err != nil {
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
		"message": "update category success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *CategoryHandlerImpl) DeleteCategory(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")
	if err := h.categoryService.DeleteCategory(categoryId); err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrNotFound:
			return pkg.ErrRosponse(c, 404, "not found", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "delete category success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
