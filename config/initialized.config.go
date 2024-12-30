package config

import (
	"onshops/core/application/services"
	"onshops/core/domain/repositories"
	"onshops/core/infrastructure/database"
	"onshops/core/infrastructure/database/migration"
	"onshops/core/presentation/http/handlers"
	"onshops/core/presentation/http/routes"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Initialized(app *fiber.App) {
	db := database.ConnectionDB()
	redis := database.RedisConnection()
	migration.AutoMigration(db)
	validation := validator.New()

	productRepository := repositories.NewProductRepositories(db, redis)
	productService := services.NewProductService(&productRepository, validation)
	productHandler := handlers.NewProductHandler(&productService)
	routes.ProductRoute(app, productHandler)

	categoryRepository := repositories.NewCategoryRepositories(db, redis)
	categoryService := services.NewCategoryService(&categoryRepository, validation)
	categoryHandler := handlers.NewCategoryHandler(&categoryService)
	routes.CategoryRoute(app, categoryHandler)

	productCategoriesRepositories := repositories.NewProductCategoriesRepositories(db, redis)
	productCategoriesService := services.NewProductCategoriesService(&productCategoriesRepositories, validation)
	productCategoriesHandler := handlers.NewProductCategoriesHandler(&productCategoriesService)
	routes.ProductCategoriesRoute(app, productCategoriesHandler)

	authRepositories := repositories.NewAuthRepositories(db)
	authService := services.NewAuthService(&authRepositories, validation)
	authHandler := handlers.NewAuthHandler(&authService)
	routes.AuthRoute(app, authHandler)

	customerRepositories := repositories.NewCustomerRepositories(db, redis)
	customerServices := services.NewCustomerService(&customerRepositories, validation)
	customerHandlers := handlers.NewCustomerHandler(&customerServices)
	routes.CustomerRoute(app, customerHandlers)

	addressRepositories := repositories.NewAddressRepositories(db)
	addressService := services.NewAddressService(&addressRepositories, &customerRepositories, validation)
	addressHandler := handlers.NewAddressHandler(&addressService)
	routes.AddressRoute(app, addressHandler)
}
