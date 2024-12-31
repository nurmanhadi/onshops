package repositories

import (
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"

	"gorm.io/gorm"
)

type OrderRepositories interface {
	GetOrders(customerId string) ([]entities.Order, error)
	GetOrderById(orderId string) (*entities.Order, error)
	CountOrderById(orderId string) (int64, error)
	AddOrder(order entities.Order) error
	UpdateOrder(orderId string, body dtos.OrderUpdateRequestDto) error
	DeleteOrder(orderId string) error
}
type OrderRepositoriesImpl struct {
	db *gorm.DB
}

func NewOrderRepositories(db *gorm.DB) OrderRepositories {
	return &OrderRepositoriesImpl{db: db}
}
func (r *OrderRepositoriesImpl) GetOrders(customerId string) ([]entities.Order, error) {
	var orders []entities.Order
	err := r.db.Where("customer_id = ?", customerId).Find(&orders).Error
	return orders, err
}
func (r *OrderRepositoriesImpl) GetOrderById(orderId string) (*entities.Order, error) {
	var order *entities.Order
	err := r.db.Where("id = ?", orderId).Preload("OrderDetails.Product").Preload("Payments").First(&order).Error
	return order, err
}
func (r *OrderRepositoriesImpl) AddOrder(order entities.Order) error {
	return r.db.Create(&order).Error
}
func (r *OrderRepositoriesImpl) DeleteOrder(orderId string) error {
	var order *entities.Order
	return r.db.Where("id = ?", orderId).Delete(&order).Error
}
func (r *OrderRepositoriesImpl) CountOrderById(orderId string) (int64, error) {
	var count int64
	var order *entities.Order
	err := r.db.Model(&order).Where("id = ?", orderId).Count(&count).Error
	return count, err
}
func (r *OrderRepositoriesImpl) UpdateOrder(orderId string, body dtos.OrderUpdateRequestDto) error {
	var order *entities.Order
	return r.db.Model(&order).Where("id = ?", orderId).Updates(&body).Error
}
