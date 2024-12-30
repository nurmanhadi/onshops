package services

import (
	"log"
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"
	"onshops/core/domain/repositories"
	"onshops/pkg"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type ProductCategoriesService interface {
	AddProductCategories(productId string, body dtos.ProductCategoriesRequestDto) error
	DeleteProductCategories(productId string, categoryId string) error
}
type ProductCategoriesServiceImpl struct {
	productCategoriesRepositories repositories.ProductCategoriesRepositories
	validation                    *validator.Validate
}

func NewProductCategoriesService(productCategoriesRepositories *repositories.ProductCategoriesRepositories, validation *validator.Validate) ProductCategoriesService {
	return &ProductCategoriesServiceImpl{productCategoriesRepositories: *productCategoriesRepositories, validation: validation}
}
func (s *ProductCategoriesServiceImpl) AddProductCategories(productId string, body dtos.ProductCategoriesRequestDto) error {
	if err := s.validation.Struct(body); err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	productCategories := entities.ProductCategories{
		CategoryId: body.CategoryId,
	}
	if err := s.productCategoriesRepositories.AddProductCategories(productCategories); err != nil {
		return err
	}
	keyProduct := "product:" + productId
	keyCategory := "category:" + string(rune(body.CategoryId))
	value, _ := s.productCategoriesRepositories.RedisRemoveProductAndCategoryById(keyProduct, keyCategory)
	log.Printf("remove %s and %s %v", keyProduct, keyCategory, value)
	return nil
}
func (s *ProductCategoriesServiceImpl) DeleteProductCategories(productId string, categoryId string) error {
	categoryID, _ := strconv.ParseUint(categoryId, 10, 64)
	if err := s.productCategoriesRepositories.DeleteProductCategories(productId, uint(categoryID)); err != nil {
		return err
	}
	keyProduct := "product:" + productId
	keyCategory := "category:" + categoryId
	value, _ := s.productCategoriesRepositories.RedisRemoveProductAndCategoryById(keyProduct, keyCategory)
	log.Printf("remove %s and %s %v", keyProduct, keyCategory, value)
	return nil
}
