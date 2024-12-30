package services

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"
	"onshops/core/domain/repositories"
	"onshops/pkg"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductService interface {
	GetProducts() ([]entities.Product, error)
	GetProductById(productId string) (*entities.Product, error)
	CreateProduct(body dtos.ProductCreateRequestDto) error
	UpdateProduct(productId string, body dtos.ProductUpdateRequestDto) error
	DeleteProduct(productId string) error
	UploadFileImage(productId string, image *multipart.FileHeader) error
	UpdateFileImage(productId string, image *multipart.FileHeader) error
}
type ProductServiceImpl struct {
	productRepository repositories.ProductRepository
	validation        *validator.Validate
}

func NewProductService(productRepository *repositories.ProductRepository, validation *validator.Validate) ProductService {
	return &ProductServiceImpl{productRepository: *productRepository, validation: validation}
}
func (s *ProductServiceImpl) GetProducts() ([]entities.Product, error) {
	var Products []entities.Product
	data, err := s.productRepository.RedisGetProducts()
	if err != nil {
		products, err := s.productRepository.GetProducts()
		if err != nil {
			return nil, &pkg.ErrNotFound{Message: "products not found"}
		}
		Products = products
		data, _ := json.Marshal(products)
		value, _ := s.productRepository.RedisSetProducts(data)
		log.Printf("set products: %s", value)
	} else {
		Products = data
	}
	return Products, nil
}
func (s *ProductServiceImpl) GetProductById(productId string) (*entities.Product, error) {
	var Product *entities.Product
	id := "product:" + productId
	data, err := s.productRepository.RedisGetProductsById(id)
	if err != nil {
		product, err := s.productRepository.GetProductById(productId)
		if err != nil {
			return nil, &pkg.ErrNotFound{Message: "product not found"}
		}
		Product = product
		data, _ := json.Marshal(product)
		value, _ := s.productRepository.RedisSetProductById(id, data)
		log.Printf("set product:%s %s", productId, value)
	} else {
		Product = data
	}
	return Product, nil
}
func (s *ProductServiceImpl) CreateProduct(body dtos.ProductCreateRequestDto) error {
	if err := s.validation.Struct(body); err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	countProduct, _ := s.productRepository.CountProductBySku(body.Sku)
	if countProduct == 1 {
		return &pkg.ErrBadRequest{Message: "sku already exist"}
	}
	id := uuid.New().String()
	product := entities.Product{
		Id:          id,
		Sku:         body.Sku,
		Name:        body.Name,
		Price:       body.Price,
		Weight:      body.Weight,
		Description: body.Description,
		Stock:       body.Stock,
	}
	if err := s.productRepository.CreateProduct(product); err != nil {
		return err
	}
	value, _ := s.productRepository.RedisRemoveProducts()
	log.Printf("remove products: %v", value)
	return nil
}
func (s *ProductServiceImpl) UpdateProduct(productId string, body dtos.ProductUpdateRequestDto) error {
	if err := s.validation.Struct(body); err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	countProduct, _ := s.productRepository.CountProductById(productId)
	if countProduct == 0 {
		return &pkg.ErrBadRequest{Message: "product not found"}
	}
	if err := s.productRepository.UpdateProduct(productId, body); err != nil {
		return err
	}
	id := "product:" + productId
	value, _ := s.productRepository.RedisRemoveProductById(id)
	log.Printf("remove products: %v", value)
	return nil
}
func (s *ProductServiceImpl) DeleteProduct(productId string) error {
	product, err := s.productRepository.GetProductById(productId)
	if err != nil {
		return &pkg.ErrNotFound{Message: "product not found"}
	}

	if err := s.productRepository.DeleteProduct(productId); err != nil {
		return err
	}
	pkg.DeleteFile(product.Image)
	id := "product:" + productId
	value, _ := s.productRepository.RedisRemoveProductById(id)
	log.Printf("remove products: %v", value)
	return nil
}
func (s *ProductServiceImpl) UploadFileImage(productId string, image *multipart.FileHeader) error {
	countProduct, _ := s.productRepository.CountProductById(productId)
	if countProduct == 0 {
		return &pkg.ErrNotFound{Message: "product not found"}
	}
	filename, err := pkg.AddFileImage(image)
	if err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	if err := s.productRepository.AddFileImage(productId, filename); err != nil {
		pkg.DeleteFile(filename)
		return err
	}
	id := "product:" + productId
	value, _ := s.productRepository.RedisRemoveProductById(id)
	log.Printf("remove products: %v", value)
	return nil
}
func (s *ProductServiceImpl) UpdateFileImage(productId string, image *multipart.FileHeader) error {
	product, err := s.productRepository.GetProductById(productId)
	if err != nil {
		return &pkg.ErrNotFound{Message: "product not found"}
	}
	filename, err := pkg.AddFileImage(image)
	if err != nil {
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	if err := s.productRepository.AddFileImage(productId, filename); err != nil {
		pkg.DeleteFile(filename)
		return err
	}
	pkg.DeleteFile(product.Image)
	id := "product:" + productId
	value, _ := s.productRepository.RedisRemoveProductById(id)
	log.Printf("remove products: %v", value)
	return nil
}
