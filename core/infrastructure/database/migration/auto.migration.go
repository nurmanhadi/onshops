package migration

import (
	entities "onshops/core/domain/entitites"

	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(
		&entities.Product{},
		&entities.Categories{},
		&entities.ProductCategories{},
		&entities.Address{},
		&entities.Customer{},
		&entities.Order{},
		&entities.OrderDetails{},
	)
}
