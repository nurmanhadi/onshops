package routes

import (
	"onshops/core/presentation/http/handlers"

	"github.com/gofiber/fiber/v2"
)

func ProductRoute(app *fiber.App, p handlers.ProductHandler) {
	product := app.Group("api/v1/products")
	product.Static("/images", "./core/presentation/resources/img/products")
	product.Get("/", p.GetProducts)
	product.Get("/:productId", p.GetProductById)
	product.Post("/", p.CreateProduct)
	product.Patch("/:productId", p.UpdateProduct)
	product.Delete("/:productId", p.DeleteProduct)
	product.Post("/:productId/images", p.UploadFileImage)
	product.Put("/:productId/images", p.UpdateFileImage)
}
