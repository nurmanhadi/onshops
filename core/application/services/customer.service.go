package services

import (
	"encoding/json"
	"log"
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"
	"onshops/core/domain/repositories"
	"onshops/pkg"

	"github.com/go-playground/validator/v10"
)

type CustomerService interface {
	UpdateProfile(customerId string, body dtos.CustomerRequestDto) error
	GetCustomerById(customerId string) (*entities.Customer, error)
}
type CustomerServiceImpl struct {
	customerRepository repositories.CustomerRepositories
	validation         *validator.Validate
}

func NewCustomerService(customerRepository *repositories.CustomerRepositories, validation *validator.Validate) CustomerService {
	return &CustomerServiceImpl{customerRepository: *customerRepository, validation: validation}
}
func (s *CustomerServiceImpl) UpdateProfile(customerId string, body dtos.CustomerRequestDto) error {
	if err := s.validation.Struct(&body); err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	customer, _ := s.customerRepository.CountCustomerById(customerId)
	if customer == 0 {
		return &pkg.ErrNotFound{Message: "customer not found"}
	}
	if err := s.customerRepository.UpdateCustomer(customerId, body); err != nil {
		return err
	}
	key := "customer:" + customerId
	value, _ := s.customerRepository.RedisRemoveCustomerById(key)
	log.Printf("remove %s has %v", key, value)
	return nil
}
func (s *CustomerServiceImpl) GetCustomerById(customerId string) (*entities.Customer, error) {
	var Customer *entities.Customer
	key := "customer:" + customerId
	data, err := s.customerRepository.RedisGetCustomerById(key)
	if err != nil {
		customer, err := s.customerRepository.GetCustomerById(customerId)
		if err != nil {
			return nil, &pkg.ErrNotFound{Message: "customer not found"}
		}
		Customer = customer
		data, _ := json.Marshal(&customer)
		value, _ := s.customerRepository.RedisSetCustomerById(key, data)
		log.Printf("set %s %s", key, value)
	} else {
		json.Unmarshal([]byte(data), &Customer)
	}
	return Customer, nil
}
