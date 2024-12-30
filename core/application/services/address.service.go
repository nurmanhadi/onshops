package services

import (
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"
	"onshops/core/domain/repositories"
	"onshops/pkg"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type AddressService interface {
	GetAddresses(customerId string) ([]entities.Address, error)
	GetAddressById(addressId string) (*entities.Address, error)
	AddAddress(customerId string, body dtos.AddressAddRequestDto) error
	Updateddress(customerId string, addressId string, body dtos.AddressUpdateRequestDto) error
	DeleteAddress(customerId string, addressId string) error
}
type AddressServiceImpl struct {
	addressRepository  repositories.AddressRepositories
	customerRepository repositories.CustomerRepositories
	validation         *validator.Validate
}

func NewAddressService(addressRepository *repositories.AddressRepositories, customerRepository *repositories.CustomerRepositories, validation *validator.Validate) AddressService {
	return &AddressServiceImpl{addressRepository: *addressRepository, customerRepository: *customerRepository, validation: validation}
}
func (s *AddressServiceImpl) GetAddresses(customerId string) ([]entities.Address, error) {
	count, _ := s.customerRepository.CountCustomerById(customerId)
	if count == 0 {
		return nil, &pkg.ErrNotFound{Message: "customer not found"}
	}
	address, err := s.addressRepository.GetAddresses(customerId)
	if err != nil {
		return nil, err
	}
	return address, nil
}
func (s *AddressServiceImpl) GetAddressById(addressId string) (*entities.Address, error) {
	id, _ := strconv.ParseUint(addressId, 10, 64)
	address, err := s.addressRepository.GetAddressById(uint(id))
	if err != nil {
		return nil, &pkg.ErrNotFound{Message: "address not found"}
	}
	return address, nil
}
func (s *AddressServiceImpl) AddAddress(customerId string, body dtos.AddressAddRequestDto) error {
	if err := s.validation.Struct(&body); err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	count, _ := s.customerRepository.CountCustomerById(customerId)
	if count == 0 {
		return &pkg.ErrNotFound{Message: "customer not found"}
	}
	address := entities.Address{
		CustomerId:    customerId,
		RecipientName: body.RecipientName,
		Phone:         body.Phone,
		Street:        body.Street,
		City:          body.City,
		State:         body.State,
		PrtalCode:     body.PrtalCode,
		Country:       body.Country,
	}
	if err := s.addressRepository.AddAddress(address); err != nil {
		return err
	}
	key := "customer:" + customerId
	s.customerRepository.RedisRemoveCustomerById(key)
	return nil
}
func (s *AddressServiceImpl) Updateddress(customerId string, addressId string, body dtos.AddressUpdateRequestDto) error {
	if err := s.validation.Struct(&body); err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	count, _ := s.customerRepository.CountCustomerById(customerId)
	if count == 0 {
		return &pkg.ErrNotFound{Message: "customer not found"}
	}
	id, _ := strconv.ParseUint(addressId, 10, 64)
	if err := s.addressRepository.UpdateAddress(uint(id), body); err != nil {
		return err
	}
	key := "customer:" + customerId
	s.customerRepository.RedisRemoveCustomerById(key)
	return nil
}
func (s *AddressServiceImpl) DeleteAddress(customerId string, addressId string) error {
	count, _ := s.customerRepository.CountCustomerById(customerId)
	if count == 0 {
		return &pkg.ErrNotFound{Message: "customer not found"}
	}
	id, _ := strconv.ParseUint(addressId, 10, 64)
	if err := s.addressRepository.DeleteAddress(uint(id)); err != nil {
		return err
	}
	key := "customer:" + customerId
	s.customerRepository.RedisRemoveCustomerById(key)
	return nil
}
