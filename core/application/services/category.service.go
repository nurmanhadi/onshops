package services

import (
	"encoding/json"
	"log"
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"
	"onshops/core/domain/repositories"
	"onshops/pkg"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type CategoryService interface {
	GetCategories() ([]entities.Categories, error)
	GetCategoriByID(categoryId string) (*entities.Categories, error)
	AddCategory(body dtos.CategoryRequestDto) error
	UpdateCategory(categoryId string, body dtos.CategoryRequestDto) error
	DeleteCategory(categoryId string) error
}
type CategoryServiceImpl struct {
	categoryRepository repositories.CategoryRepositories
	validation         *validator.Validate
}

func NewCategoryService(categoryRepository *repositories.CategoryRepositories, validation *validator.Validate) CategoryService {
	return &CategoryServiceImpl{categoryRepository: *categoryRepository, validation: validation}
}
func (s *CategoryServiceImpl) GetCategories() ([]entities.Categories, error) {
	categories, err := s.categoryRepository.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
func (s *CategoryServiceImpl) GetCategoriByID(categoryId string) (*entities.Categories, error) {
	id, _ := strconv.ParseUint(categoryId, 10, 64)
	key := "category:" + categoryId
	var Category *entities.Categories
	data, err := s.categoryRepository.RedisGetCategoryById(key)
	if err != nil {
		category, err := s.categoryRepository.GetCategoriByID(uint(id))
		if err != nil {
			return nil, &pkg.ErrNotFound{Message: "category not found"}
		}
		Category = category
		data, _ := json.Marshal(category)
		value, _ := s.categoryRepository.RedisSetCategoryById(key, data)
		log.Printf("set category:%s %v", categoryId, value)
	} else {
		json.Unmarshal([]byte(data), &Category)
	}
	return Category, nil
}
func (s *CategoryServiceImpl) AddCategory(body dtos.CategoryRequestDto) error {
	if err := s.validation.Struct(body); err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	category := entities.Categories{
		Name: *body.Name,
	}
	if err := s.categoryRepository.AddCategory(category); err != nil {
		return err
	}
	return nil
}
func (s *CategoryServiceImpl) UpdateCategory(categoryId string, body dtos.CategoryRequestDto) error {
	if err := s.validation.Struct(body); err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	id, _ := strconv.ParseUint(categoryId, 10, 64)
	countCategory, _ := s.categoryRepository.CountCategory(uint(id))
	if countCategory == 0 {
		return &pkg.ErrNotFound{Message: "category not found"}
	}
	if err := s.categoryRepository.UpdateCategory(uint(id), body); err != nil {
		return err
	}
	key := "category:" + categoryId
	value, _ := s.categoryRepository.RedisRemoveCategoryById(key)
	log.Printf("remove category:%s %v", categoryId, value)
	return nil
}
func (s *CategoryServiceImpl) DeleteCategory(categoryId string) error {
	id, _ := strconv.ParseUint(categoryId, 10, 64)
	countCategory, _ := s.categoryRepository.CountCategory(uint(id))
	if countCategory == 0 {
		return &pkg.ErrNotFound{Message: "category not found"}
	}
	if err := s.categoryRepository.DeleteCategory(uint(id)); err != nil {
		return err
	}
	key := "category:" + categoryId
	value, _ := s.categoryRepository.RedisRemoveCategoryById(key)
	log.Printf("remove category:%s %v", categoryId, value)
	return nil
}
