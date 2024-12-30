package entities

import (
	"time"
)

type ProductCategories struct {
	Id         uint       `gorm:"primaryKey" json:"id"`
	ProductId  string     `gorm:"type:varchar(36);not null" json:"product_id"`
	CategoryId uint       `gorm:"not null" json:"category_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	Product    Product    `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE" json:"product"`
	Categories Categories `gorm:"foreignKey:CategoryId;constraint:OnDelete:CASCADE" json:"categories"`
}
