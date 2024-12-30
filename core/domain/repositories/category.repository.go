package repositories

import (
	"context"
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CategoryRepositories interface {
	GetCategories() ([]entities.Categories, error)
	GetCategoriByID(categoryId uint) (*entities.Categories, error)
	AddCategory(category entities.Categories) error
	UpdateCategory(categoryId uint, body dtos.CategoryRequestDto) error
	DeleteCategory(categoryId uint) error
	CountCategory(categoryId uint) (int64, error)
	RedisGetCategoryById(categoryId string) (string, error)
	RedisSetCategoryById(categoryId string, data []byte) (string, error)
	RedisRemoveCategoryById(categoryId string) (int64, error)
}
type CategoryRepositoriesImpl struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewCategoryRepositories(db *gorm.DB, redis *redis.Client) CategoryRepositories {
	return &CategoryRepositoriesImpl{db: db, redis: redis}
}
func (r *CategoryRepositoriesImpl) GetCategories() ([]entities.Categories, error) {
	var categories []entities.Categories
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepositoriesImpl) GetCategoriByID(categoryId uint) (*entities.Categories, error) {
	var category *entities.Categories
	err := r.db.Where("id = ?", categoryId).Preload("ProductCategories.Product").First(&category).Error
	return category, err
}
func (r *CategoryRepositoriesImpl) AddCategory(category entities.Categories) error {
	return r.db.Create(&category).Error
}
func (r *CategoryRepositoriesImpl) UpdateCategory(categoryId uint, body dtos.CategoryRequestDto) error {
	var category *entities.Categories
	return r.db.Model(&category).Where("id = ?", categoryId).Updates(&body).Error
}
func (r *CategoryRepositoriesImpl) DeleteCategory(categoryId uint) error {
	var category *entities.Categories
	return r.db.Where("id = ?", categoryId).Delete(&category).Error
}
func (r *CategoryRepositoriesImpl) CountCategory(categoryId uint) (int64, error) {
	var count int64
	var category *entities.Categories
	err := r.db.Model(&category).Where("id = ?", categoryId).Count(&count).Error
	return count, err
}
func (r *CategoryRepositoriesImpl) RedisGetCategoryById(categoryId string) (string, error) {
	data, err := r.redis.Get(context.Background(), categoryId).Result()
	return data, err
}
func (r *CategoryRepositoriesImpl) RedisSetCategoryById(categoryId string, data []byte) (string, error) {
	exp := time.Hour * 30
	value, err := r.redis.Set(context.Background(), categoryId, data, exp).Result()
	return value, err
}
func (r *CategoryRepositoriesImpl) RedisRemoveCategoryById(categoryId string) (int64, error) {
	data, err := r.redis.Del(context.Background(), categoryId).Result()
	return data, err
}
