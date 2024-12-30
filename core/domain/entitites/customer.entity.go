package entities

import (
	"time"
)

type Customer struct {
	Id        string    `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Name      *string   `gorm:"type:varchar(100)" json:"name"`
	Email     string    `gorm:"unique;type:varchar(100)" json:"email"`
	Password  string    `gorm:"type:varchar(100)" json:"password"`
	Country   *string   `gorm:"type:varchar(50)" json:"country"`
	Address   []Address `gorm:"foreignKey:CustomerId" json:"address"`
	Phone     *string   `gorm:"type:varchar(20)" json:"phone"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Orders    []Order   `gorm:"foreignKey:CustomerId" json:"orders"`
}
