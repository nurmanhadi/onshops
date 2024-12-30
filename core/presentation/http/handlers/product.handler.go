package handlers

import (
	"onshops/core/application/dtos"
	"onshops/core/application/services"
	"onshops/pkg"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	GetProducts(c *fiber.Ctx) error
	GetProductById(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
	UploadFileImage(c *fiber.Ctx) error
	UpdateFileImage(c *fiber.Ctx) error
}
type ProductHandlerImpl struct {
	productService services.ProductService
}

func NewProductHandler(productService *services.ProductService) ProductHandler {
	return &ProductHandlerImpl{productService: *productService}
}
func (h *ProductHandlerImpl) GetProducts(c *fiber.Ctx) error {
	products, err := h.productService.GetProducts()
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
		"message": "get products success",
		"data":    products,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *ProductHandlerImpl) GetProductById(c *fiber.Ctx) error {
	productId := c.Params("productId")
	product, err := h.productService.GetProductById(productId)
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
		"message": "get product success",
		"data":    product,
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}

func (h *ProductHandlerImpl) CreateProduct(c *fiber.Ctx) error {
	body := new(dtos.ProductCreateRequestDto)
	if err := c.BodyParser(&body); err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "cannot parse to json")
	}
	if err := h.productService.CreateProduct(*body); err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrBadRequest:
			return pkg.ErrRosponse(c, 400, "bad request", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}
	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "create product success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *ProductHandlerImpl) UpdateProduct(c *fiber.Ctx) error {
	productId := c.Params("productId")
	body := new(dtos.ProductUpdateRequestDto)
	if err := c.BodyParser(&body); err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "cannot parse json")
	}
	if err := h.productService.UpdateProduct(productId, *body); err != nil {
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
		"message": "update product success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *ProductHandlerImpl) DeleteProduct(c *fiber.Ctx) error {
	productId := c.Params("productId")
	if err := h.productService.DeleteProduct(productId); err != nil {
		switch finalErr := err.(type) {
		case *pkg.ErrNotFound:
			return pkg.ErrRosponse(c, 404, "not found", finalErr.Message)
		default:
			return pkg.ErrRosponse(c, 500, "internal server error", err.Error())
		}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "delete product success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *ProductHandlerImpl) UploadFileImage(c *fiber.Ctx) error {
	productId := c.Params("productId")
	image, err := c.FormFile("image")
	if err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "file required")
	}
	if err := h.productService.UploadFileImage(productId, image); err != nil {
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
		"message": "upload file product success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
func (h *ProductHandlerImpl) UpdateFileImage(c *fiber.Ctx) error {
	productId := c.Params("productId")
	image, err := c.FormFile("image")
	if err != nil {
		return pkg.ErrRosponse(c, 400, "bad request", "file required")
	}
	if err := h.productService.UpdateFileImage(productId, image); err != nil {
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
		"message": "update file product success",
		"links": fiber.Map{
			"self": c.OriginalURL(),
		},
	})
}
