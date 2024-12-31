package entities

import (
	"time"
)

type OrderDetails struct {
	Id          uint      `gorm:"primaryKey" json:"id"`
	OrderId     string    `gorm:"type:varchar(36);not null" json:"order_id"`
	ProductId   string    `gorm:"type:varchar(36);not null" json:"product_id"`
	Price       int64     `gorm:"not null" json:"price"`
	Sku         string    `gorm:"unique;type:varchar(12);not null" json:"sku"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	GrossAmount int64     `gorm:"not null" json:"gross_amount"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Order       Order     `gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE" json:"order"`
	Product     Product   `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE" json:"product"`
}
