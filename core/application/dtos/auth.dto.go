package dtos

type AuthRequestDto struct {
	Email    string `json:"email" validate:"required,email,min=1,max=100"`
	Password string `json:"password" validate:"required,min=1,max=100"`
}
