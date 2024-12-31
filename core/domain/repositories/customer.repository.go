package repositories

import (
	"context"
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CustomerRepositories interface {
	UpdateCustomer(customerId string, body dtos.CustomerRequestDto) error
	GetCustomerById(customerId string) (*entities.Customer, error)
	CountCustomerById(customerId string) (int64, error)
	RedisGetCustomerById(customerId string) (string, error)
	RedisSetCustomerById(customerId string, data []byte) (string, error)
	RedisRemoveCustomerById(customerId string) (int64, error)
}
type CustomerRepositoriesImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewCustomerRepositories(db *gorm.DB, redis *redis.Client) CustomerRepositories {
	return &CustomerRepositoriesImpl{db: db, redis: redis}
}
func (r *CustomerRepositoriesImpl) UpdateCustomer(customerId string, body dtos.CustomerRequestDto) error {
	return r.db.Model(&entities.Customer{}).Where("id = ?", customerId).Updates(&body).Error
}
func (r *CustomerRepositoriesImpl) GetCustomerById(customerId string) (*entities.Customer, error) {
	var customer *entities.Customer
	err := r.db.Where("id = ?", customerId).Preload("Address").Preload("Orders").First(&customer).Error
	return customer, err
}
func (r *CustomerRepositoriesImpl) CountCustomerById(customerId string) (int64, error) {
	var count int64
	err := r.db.Model(&entities.Customer{}).Where("id = ?", customerId).Count(&count).Error
	return count, err
}
func (r *CustomerRepositoriesImpl) RedisGetCustomerById(customerId string) (string, error) {
	customer, err := r.redis.Get(context.Background(), customerId).Result()
	return customer, err
}
func (r *CustomerRepositoriesImpl) RedisSetCustomerById(customerId string, data []byte) (string, error) {
	exp := time.Minute * 30
	value, err := r.redis.Set(context.Background(), customerId, data, exp).Result()
	return value, err
}
func (r *CustomerRepositoriesImpl) RedisRemoveCustomerById(customerId string) (int64, error) {
	value, err := r.redis.Del(context.Background(), customerId).Result()
	return value, err
}
