package repositories

import (
	"context"
	entities "onshops/core/domain/entitites"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ProductCategoriesRepositories interface {
	AddProductCategories(roductCategories entities.ProductCategories) error
	DeleteProductCategories(productId string, categoryId uint) error
	RedisRemoveProductAndCategoryById(productId string, categoryId string) (int64, error)
}
type ProductCategoriesRepositoriesImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewProductCategoriesRepositories(db *gorm.DB, redis *redis.Client) ProductCategoriesRepositories {
	return &ProductCategoriesRepositoriesImpl{db: db, redis: redis}
}
func (r *ProductCategoriesRepositoriesImpl) AddProductCategories(roductCategories entities.ProductCategories) error {
	return r.db.Create(&roductCategories).Error
}
func (r *ProductCategoriesRepositoriesImpl) DeleteProductCategories(productId string, categoryId uint) error {
	var productCategory entities.ProductCategories
	return r.db.Where("product_id = ?", productId).Where("category_id").Delete(&productCategory).Error
}
func (r *ProductCategoriesRepositoriesImpl) RedisRemoveProductAndCategoryById(productId string, categoryId string) (int64, error) {
	data, err := r.redis.Del(context.Background(), productId, categoryId).Result()
	return data, err
}
