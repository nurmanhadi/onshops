package services

import (
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"
	"onshops/core/domain/repositories"
	"onshops/pkg"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type OrderService interface {
	GetOrders(customerId string) ([]entities.Order, error)
	GetOrderById(orderId string) (*entities.Order, error)
	AddOrder(body dtos.OrderRequestDto) error
	DeleteOrder(orderId string) error
	UpdateOrder(orderId string, body dtos.OrderUpdateRequestDto) error
}
type OrderServiceImpl struct {
	orderRepository repositories.OrderRepositories
	validation      *validator.Validate
}

func NewOrderService(orderRepository *repositories.OrderRepositories, validation *validator.Validate) OrderService {
	return &OrderServiceImpl{orderRepository: *orderRepository, validation: validation}
}
func (s *OrderServiceImpl) GetOrders(customerId string) ([]entities.Order, error) {
	orders, err := s.orderRepository.GetOrders(customerId)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
func (s *OrderServiceImpl) GetOrderById(orderId string) (*entities.Order, error) {
	order, err := s.orderRepository.GetOrderById(orderId)
	if err != nil {
		return nil, &pkg.ErrNotFound{Message: "order not found"}
	}
	return order, nil
}
func (s *OrderServiceImpl) AddOrder(body dtos.OrderRequestDto) error {
	if err := s.validation.Struct(&body); err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	id := uuid.New().String()
	order := entities.Order{
		Id:              id,
		CustomerId:      body.CustomerId,
		ShippingAddress: body.ShippingAddress,
		OrderAddress:    body.OrderAddress,
		OrderEmail:      body.OrderEmail,
		Amount:          0,
		OrderDate:       time.Now(),
		OrderStatus:     "pending",
	}
	if err := s.orderRepository.AddOrder(order); err != nil {
		return err
	}
	return nil
}
func (s *OrderServiceImpl) DeleteOrder(orderId string) error {
	count, _ := s.orderRepository.CountOrderById(orderId)
	if count == 0 {
		return &pkg.ErrNotFound{Message: "order not found"}
	}
	return nil
}
func (s *OrderServiceImpl) UpdateOrder(orderId string, body dtos.OrderUpdateRequestDto) error {
	if err := s.validation.Struct(&body); err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	count, _ := s.orderRepository.CountOrderById(orderId)
	if count == 0 {
		return &pkg.ErrNotFound{Message: "order not found"}
	}
	if err := s.orderRepository.UpdateOrder(orderId, body); err != nil {
		return err
	}
	return nil
}
