package repositories

import (
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"

	"gorm.io/gorm"
)

type AddressRepositories interface {
	GetAddresses() ([]entities.Address, error)
	GetAddressById(addressId uint) (*entities.Address, error)
	AddAddress(address entities.Address) error
	UpdateAddress(addressId uint, body dtos.AddressUpdateRequestDto) error
	DeleteAddress(addressId uint) error
}
type AddressRepositoriesImpl struct {
	db *gorm.DB
}

func NewAddressRepositories(db *gorm.DB) AddressRepositories {
	return &AddressRepositoriesImpl{db: db}
}
func (r *AddressRepositoriesImpl) GetAddresses() ([]entities.Address, error) {
	var address []entities.Address
	err := r.db.Find(&address).Error
	return address, err
}
func (r *AddressRepositoriesImpl) GetAddressById(addressId uint) (*entities.Address, error) {
	var address *entities.Address
	err := r.db.Where("id = ?", addressId).First(&address).Error
	return address, err
}
func (r *AddressRepositoriesImpl) AddAddress(address entities.Address) error {
	return r.db.Create(&address).Error
}
func (r *AddressRepositoriesImpl) UpdateAddress(addressId uint, body dtos.AddressUpdateRequestDto) error {
	var address *entities.Address
	return r.db.Model(&address).Where("id = ?", addressId).Updates(&body).Error
}
func (r *AddressRepositoriesImpl) DeleteAddress(addressId uint) error {
	var address *entities.Address
	return r.db.Where("id = ?", addressId).Delete(&address).Error
}
