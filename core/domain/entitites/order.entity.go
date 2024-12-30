package entities

import (
	"time"
)

type Order struct {
	Id              string         `gorm:"primaryKey;type:varchar(36)" json:"id"`
	CustomerId      string         `gorm:"type:uuid;not null" json:"customer_id"`
	Amount          int64          `gorm:"not null" json:"amount"`
	ShippingAddress string         `gorm:"type:text;not null" json:"shipping_address"`
	OrderAddress    string         `gorm:"type:text;not null" json:"order_address"`
	OrderEmail      string         `gorm:"type:varchar(100);not null" json:"order_email"`
	OrderDate       time.Time      `gorm:"type:date;not null" json:"order_date"`
	OrderStatus     string         `gorm:"type:varchar(12);default:pending" json:"order_status"`
	OrderDetails    []OrderDetails `gorm:"foreignKey:OrderId" json:"order_details"`
	Customer        Customer       `gorm:"foreignKey:CustomerId" json:"customer"`
}
