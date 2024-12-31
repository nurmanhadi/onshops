package services

import (
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"
	"onshops/core/domain/repositories"
	"onshops/pkg"

	"github.com/go-playground/validator/v10"
)

type OrderDetailsService interface {
	AddOrderDetail(body dtos.OrderDetailRequestDto) error
	// DeleteOrderDetail(orderDetailId string) error
}
type OrderDetailsServiceImpl struct {
	orderDetailsRepository repositories.OrderDetailsRepositories
	orderRepository        repositories.OrderRepositories
	vaidation              *validator.Validate
}

func NewOrderDetailsService(orderDetailsRepository *repositories.OrderDetailsRepositories, orderRepository *repositories.OrderRepositories, vaidation *validator.Validate) OrderDetailsService {
	return &OrderDetailsServiceImpl{orderDetailsRepository: *orderDetailsRepository, orderRepository: *orderRepository, vaidation: vaidation}
}
func (s *OrderDetailsServiceImpl) AddOrderDetail(body dtos.OrderDetailRequestDto) error {
	if err := s.vaidation.Struct(&body); err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	orderDetails := entities.OrderDetails{
		OrderId:     body.OrderId,
		ProductId:   body.ProductId,
		Price:       body.Price,
		Sku:         body.Sku,
		Quantity:    body.Quantity,
		GrossAmount: body.GrossAmount,
	}
	if err := s.orderDetailsRepository.AddOrderDetails(orderDetails); err != nil {
		return err
	}
	updateAmount := dtos.OrderUpdateRequestDto{
		Amount: &body.GrossAmount,
	}
	if err := s.orderRepository.UpdateOrder(body.OrderId, updateAmount); err != nil {
		return err
	}
	return nil
}
