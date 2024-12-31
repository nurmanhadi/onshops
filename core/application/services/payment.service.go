package services

import (
	"fmt"
	"log"
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"
	"onshops/core/domain/repositories"
	"onshops/pkg"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentService interface {
	AddPaymentTransactions(body dtos.PaymentRequestDto) (*snap.Response, error)
	AddPayment(body dtos.PaymentNotificationDto) error
}
type PaymentServiceImpl struct {
	paymentRepository repositories.PaymentRepositories
	productRepository repositories.ProductRepository
	orderRepository   repositories.OrderRepositories
	validation        *validator.Validate
}

func NewPaymentService(
	paymentRepository *repositories.PaymentRepositories,
	productRepository *repositories.ProductRepository,
	orderRepository *repositories.OrderRepositories,
	validation *validator.Validate) PaymentService {
	return &PaymentServiceImpl{paymentRepository: *paymentRepository, productRepository: *productRepository, orderRepository: *orderRepository, validation: validation}
}
func (s *PaymentServiceImpl) AddPaymentTransactions(body dtos.PaymentRequestDto) (*snap.Response, error) {
	if err := s.validation.Struct(&body); err != nil {
		log.Println(err.Error())
		return nil, &pkg.ErrBadRequest{Message: err.Error()}
	}
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  body.OrderId,
			GrossAmt: body.GrossAmount,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:       body.ItemDetails.ProductId,
				Price:    body.ItemDetails.Price,
				Name:     body.ItemDetails.Name,
				Qty:      int32(body.ItemDetails.Quantity),
				Category: body.ItemDetails.Category,
			},
		},
	}
	snapRes, err := snap.CreateTransaction(snapReq)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Message)
	}
	return snapRes, nil
}
func (s *PaymentServiceImpl) AddPayment(body dtos.PaymentNotificationDto) error {
	if err := s.validation.Struct(&body); err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	order, err := s.orderRepository.GetOrderById(body.OrderID)
	if err != nil {
		return &pkg.ErrNotFound{Message: "order not found"}
	}
	orderStatus := dtos.OrderUpdateRequestDto{
		OrderStatus: &body.TransactionStatus,
	}
	if err := s.orderRepository.UpdateOrder(order.Id, orderStatus); err != nil {
		return err
	}
	stock := order.OrderDetails[0].Product.Stock - order.OrderDetails[0].Quantity
	productId := order.OrderDetails[0].ProductId
	productStock := dtos.ProductUpdateRequestDto{
		Stock: &stock,
	}
	if err := s.productRepository.UpdateProduct(productId, productStock); err != nil {
		return err
	}
	id := uuid.New().String()
	grossAmt, _ := strconv.ParseInt(body.GrossAmount, 10, 64)
	layout := "2023-11-15 18:45:13"
	tranTime, _ := time.Parse(layout, body.TransactionTime)
	payment := entities.Payment{
		Id:              id,
		OrderId:         body.OrderID,
		Amount:          grossAmt,
		PaymentMethod:   body.PaymentType,
		PaymentStatus:   body.TransactionStatus,
		TransactionId:   body.TransactionID,
		TransactionTime: tranTime,
	}
	if err := s.paymentRepository.AddPayment(payment); err != nil {
		return err
	}
	return nil
}
