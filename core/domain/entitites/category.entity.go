package entities

import "time"

type Categories struct {
	Id                uint                `gorm:"primaryKey" json:"id"`
	Name              string              `gorm:"type:varchar(50);not null" json:"name"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
	ProductCategories []ProductCategories `gorm:"foreignKey:CategoryId" json:"product_categories"`
}
