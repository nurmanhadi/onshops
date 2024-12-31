package config

import (
	"onshops/core/application/services"
	"onshops/core/domain/repositories"
	"onshops/core/infrastructure/database"
	"onshops/core/infrastructure/database/migration"
	"onshops/core/infrastructure/midtrans"
	"onshops/core/presentation/http/handlers"
	"onshops/core/presentation/http/routes"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Initialized(app *fiber.App) {
	midtrans.SetupMidtrans()

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

	orderRepositories := repositories.NewOrderRepositories(db)
	orderService := services.NewOrderService(&orderRepositories, validation)
	orderHandler := handlers.NewOrderHandler(&orderService)
	routes.OrderRoute(app, orderHandler)

	orderDetailRepositories := repositories.NewOrderDetailsRepositories(db)
	orderDetailService := services.NewOrderDetailsService(&orderDetailRepositories, &orderRepositories, validation)
	orderDetailHandler := handlers.NewOrderDetailsHandler(&orderDetailService)
	routes.OrderDetailsRoute(app, orderDetailHandler)

	paymentRepositories := repositories.NewPaymentRepositories(db)
	paymentService := services.NewPaymentService(&paymentRepositories, &productRepository, &orderRepositories, validation)
	paymentHandler := handlers.NewPaymentHandler(&paymentService)
	routes.PaymentRoute(app, paymentHandler)
}
