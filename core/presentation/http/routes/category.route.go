package routes

import (
	"onshops/core/presentation/http/handlers"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoute(app *fiber.App, h handlers.CategoryHandler) {
	categories := app.Group("/api/v1/categories")
	categories.Get("/", h.GetCategories)
	categories.Get("/:categoryId", h.GetCategoryById)
	categories.Post("/", h.AddCategory)
	categories.Patch("/:categoryId", h.UpdateCategory)
	categories.Delete("/:categoryId", h.DeleteCategory)
}
