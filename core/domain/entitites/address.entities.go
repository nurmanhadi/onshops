package entities

import (
	"time"
)

type Address struct {
	Id            uint      `gorm:"primaryKey" json:"id"`
	CustomerId    string    `gorm:"type:varchar(36);not null" json:"customer_id"`
	RecipientName string    `gorm:"type:varchar(100);not null" json:"recipient_name"`
	Phone         string    `gorm:"type:varchar(20);not null" json:"phone"`
	Street        string    `gorm:"type:varchar(255);not null" json:"street"`
	City          string    `gorm:"type:varchar(50);not null" json:"city"`
	State         string    `gorm:"type:varchar(50);not null" json:"state"`
	PrtalCode     string    `gorm:"type:varchar(20);not null" json:"prtal_code"`
	Country       string    `gorm:"type:varchar(50);not null" json:"country"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Customer      Customer  `gorm:"foreignKey:CustomerId" json:"customer"`
}
