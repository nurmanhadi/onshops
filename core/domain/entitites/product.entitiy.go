package entities

import (
	"time"
)

type Product struct {
	Id                string              `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Sku               string              `gorm:"unique;type:varchar(12);not null" json:"sku"`
	Name              string              `gorm:"type:varchar(100);not null" json:"name"`
	Price             int64               `gorm:"not null" json:"price"`
	Weight            float64             `gorm:"not null" json:"weight"`
	Description       string              `gorm:"type:text;not null" json:"description"`
	Image             string              `gorm:"type:varchar(200)" json:"image"`
	Stock             int                 `gorm:"not null;default:0" json:"stock"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
	ProductCategories []ProductCategories `gorm:"foreignKey:ProductId" json:"product_categories"`
}
