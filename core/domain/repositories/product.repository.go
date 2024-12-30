package repositories

import (
	"context"
	"encoding/json"
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts() ([]entities.Product, error)
	GetProductById(productId string) (*entities.Product, error)
	CreateProduct(product entities.Product) error
	UpdateProduct(productId string, body dtos.ProductUpdateRequestDto) error
	CountProducts() (int64, error)
	CountProductById(productId string) (int64, error)
	CountProductBySku(sku string) (int64, error)
	DeleteProduct(productId string) error
	AddFileImage(productId string, image string) error
	RedisGetProducts() ([]entities.Product, error)
	RedisGetProductsById(productId string) (*entities.Product, error)
	RedisSetProducts(data []byte) (string, error)
	RedisSetProductById(productId string, data []byte) (string, error)
	RedisRemoveProducts() (int64, error)
	RedisRemoveProductById(productId string) (int64, error)
}
type ProductRepositoriesImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewProductRepositories(db *gorm.DB, redis *redis.Client) ProductRepository {
	return &ProductRepositoriesImpl{db: db, redis: redis}
}
func (r *ProductRepositoriesImpl) GetProducts() ([]entities.Product, error) {
	var products []entities.Product
	err := r.db.Find(&products).Error
	return products, err
}
func (r *ProductRepositoriesImpl) GetProductById(productId string) (*entities.Product, error) {
	var product *entities.Product
	err := r.db.Where("id = ?", productId).Preload("ProductCategories.Categories").First(&product).Error
	return product, err
}
func (r *ProductRepositoriesImpl) CreateProduct(product entities.Product) error {
	return r.db.Create(&product).Error
}
func (r *ProductRepositoriesImpl) UpdateProduct(productId string, body dtos.ProductUpdateRequestDto) error {
	var product entities.Product
	return r.db.Model(&product).Where("id = ?", productId).Updates(&body).Error
}
func (r *ProductRepositoriesImpl) CountProducts() (int64, error) {
	var count int64
	var products entities.Product
	err := r.db.Model(&products).Count(&count).Error
	return count, err
}
func (r *ProductRepositoriesImpl) CountProductById(productId string) (int64, error) {
	var count int64
	var product entities.Product
	err := r.db.Model(&product).Where("id = ?", productId).Count(&count).Error
	return count, err
}
func (r *ProductRepositoriesImpl) CountProductBySku(sku string) (int64, error) {
	var count int64
	var product entities.Product
	err := r.db.Model(&product).Where("sku = ?", sku).Count(&count).Error
	return count, err
}
func (r *ProductRepositoriesImpl) DeleteProduct(productId string) error {
	var product entities.Product
	return r.db.Where("id = ?", productId).Delete(&product).Error
}
func (r *ProductRepositoriesImpl) AddFileImage(productId string, image string) error {
	var product entities.Product
	return r.db.Model(&product).Where("id = ?", productId).Update("image", image).Error
}
func (r *ProductRepositoriesImpl) RedisGetProducts() ([]entities.Product, error) {
	var products []entities.Product
	data, err := r.redis.Get(context.Background(), "products").Result()
	json.Unmarshal([]byte(data), &products)
	return products, err
}
func (r *ProductRepositoriesImpl) RedisGetProductsById(productId string) (*entities.Product, error) {
	var product *entities.Product
	data, err := r.redis.Get(context.Background(), productId).Result()
	json.Unmarshal([]byte(data), &product)
	return product, err
}
func (r *ProductRepositoriesImpl) RedisSetProducts(data []byte) (string, error) {
	exp := time.Minute * 30
	value, err := r.redis.Set(context.Background(), "products", data, exp).Result()
	return value, err
}
func (r *ProductRepositoriesImpl) RedisSetProductById(productId string, data []byte) (string, error) {
	exp := time.Minute * 30
	value, err := r.redis.Set(context.Background(), productId, data, exp).Result()
	return value, err
}
func (r *ProductRepositoriesImpl) RedisRemoveProducts() (int64, error) {
	value, err := r.redis.Del(context.Background(), "products").Result()
	return value, err
}
func (r *ProductRepositoriesImpl) RedisRemoveProductById(productId string) (int64, error) {
	value, err := r.redis.Del(context.Background(), productId).Result()
	return value, err
}
