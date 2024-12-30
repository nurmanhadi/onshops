package repositories

import (
	entities "onshops/core/domain/entitites"

	"gorm.io/gorm"
)

type AuthRepositories interface {
	CreateCustomer(customer entities.Customer) error
	GetCustomerByEmail(email string) (*entities.Customer, error)
	CountCustomerByEmail(email string) (int64, error)
}
type AuthRepositoriesImpl struct {
	db *gorm.DB
}

func NewAuthRepositories(db *gorm.DB) AuthRepositories {
	return &AuthRepositoriesImpl{db: db}
}
func (r *AuthRepositoriesImpl) CreateCustomer(customer entities.Customer) error {
	return r.db.Create(&customer).Error
}
func (r *AuthRepositoriesImpl) GetCustomerByEmail(email string) (*entities.Customer, error) {
	var customer *entities.Customer
	err := r.db.Where("email = ?", email).First(&customer).Error
	return customer, err
}
func (r *AuthRepositoriesImpl) CountCustomerByEmail(email string) (int64, error) {
	var count int64
	err := r.db.Model(&entities.Customer{}).Where("email = ?", email).Count(&count).Error
	return count, err
}
