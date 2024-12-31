package repositories

import (
	entities "onshops/core/domain/entitites"

	"gorm.io/gorm"
)

type OrderDetailsRepositories interface {
	AddOrderDetails(orderDetail entities.OrderDetails) error
	DeleteOrderDetail(orderDetailId uint) error
}
type OrderDetailsRepositoriesImpl struct {
	db *gorm.DB
}

func NewOrderDetailsRepositories(db *gorm.DB) OrderDetailsRepositories {
	return &OrderDetailsRepositoriesImpl{db: db}
}
func (r *OrderDetailsRepositoriesImpl) AddOrderDetails(orderDetail entities.OrderDetails) error {
	return r.db.Create(&orderDetail).Error
}
func (r *OrderDetailsRepositoriesImpl) DeleteOrderDetail(orderDetailId uint) error {
	var orderDetail *entities.OrderDetails
	return r.db.Where("id = ? ", orderDetail).Delete(&orderDetail).Error
}
