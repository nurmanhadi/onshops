package services

import (
	"fmt"
	"onshops/core/application/dtos"
	entities "onshops/core/domain/entitites"
	"onshops/core/domain/repositories"
	"onshops/pkg"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	AuthRegister(body dtos.AuthRequestDto) error
	AuthLogin(body dtos.AuthRequestDto) (string, error)
}
type AuthServiceImpl struct {
	authRepositories repositories.AuthRepositories
	validation       *validator.Validate
}

func NewAuthService(authRepositories *repositories.AuthRepositories, validation *validator.Validate) AuthService {
	return &AuthServiceImpl{authRepositories: *authRepositories, validation: validation}
}
func (s *AuthServiceImpl) AuthRegister(body dtos.AuthRequestDto) error {
	if err := s.validation.Struct(&body); err != nil {
		fmt.Println(err)
		return &pkg.ErrBadRequest{Message: err.Error()}
	}
	countCustomer, _ := s.authRepositories.CountCustomerByEmail(body.Email)
	if countCustomer == 1 {
		return &pkg.ErrBadRequest{Message: "email already exist"}
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	id := uuid.New().String()
	customer := entities.Customer{
		Id:       id,
		Email:    body.Email,
		Password: string(hashPassword),
	}
	if err := s.authRepositories.CreateCustomer(customer); err != nil {
		return err
	}
	return nil
}
func (s *AuthServiceImpl) AuthLogin(body dtos.AuthRequestDto) (string, error) {
	if err := s.validation.Struct(&body); err != nil {
		return "", &pkg.ErrBadRequest{Message: err.Error()}
	}
	customer, err := s.authRepositories.GetCustomerByEmail(body.Email)
	if err != nil {
		return "", &pkg.ErrBadRequest{Message: "email or password is wrong"}
	}
	comparePassword := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(body.Password))
	if comparePassword != nil {
		return "", &pkg.ErrBadRequest{Message: "email or password is wrong"}
	}
	accessToken, err := pkg.JwtGenerateAccessToken(customer.Id)
	if err != nil {
		return "", fmt.Errorf("cannot generate access token because: %v", err)
	}
	return accessToken, nil
}
