package routes

import (
	"onshops/core/presentation/http/handlers"

	"github.com/gofiber/fiber/v2"
)

func ProductCategoriesRoute(app *fiber.App, h handlers.ProductCategoriesHandler) {
	api := app.Group("api/v1")
	products := api.Group("/products")
	products.Post("/:productId/categories", h.AddProductCategories)
	products.Delete("/:productId/categories/:categoryId", h.DeleteProductCategories)
}
