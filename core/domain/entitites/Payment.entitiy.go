package entities

import "time"

type Payment struct {
	Id              string    `gorm:"primaryKey;type:varchar(36)" json:"id"`
	OrderId         string    `gorm:"type:varchar(36);not null" json:"order_id"`
	Amount          int64     `gorm:"not null" json:"amount"`
	PaymentMethod   string    `gorm:"type:varchar(50);not null" json:"payment_method"`
	PaymentStatus   string    `gorm:"type:varchar(20);not null" json:"payment_status"`
	TransactionId   string    `gorm:"type:varchar(50);unique" json:"transaction_id"`
	TransactionTime time.Time `gorm:"autoCreateTime" json:"transaction_time"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Order           Order     `gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE" json:"order"`
}
