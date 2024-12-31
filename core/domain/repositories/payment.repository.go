package repositories

import (
	entities "onshops/core/domain/entitites"

	"gorm.io/gorm"
)

type PaymentRepositories interface {
	AddPayment(payment entities.Payment) error
}
type PaymentRepositoriesImpl struct {
	db *gorm.DB
}

func NewPaymentRepositories(db *gorm.DB) PaymentRepositories {
	return &PaymentRepositoriesImpl{db: db}
}
func (r *PaymentRepositoriesImpl) AddPayment(payment entities.Payment) error {
	return r.db.Create(&payment).Error
}
