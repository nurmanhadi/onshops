package dtos

type ProductCreateRequestDto struct {
	Sku         string  `json:"sku" validate:"required,min=1,max=12"`
	Name        string  `json:"name" validate:"required,min=1,max=100"`
	Price       int64   `json:"price" validate:"required"`
	Weight      float64 `json:"weight" validate:"required"`
	Description string  `json:"description" validate:"required,min=1"`
	Stock       int     `json:"stock" validate:"required"`
}
type ProductUpdateRequestDto struct {
	Sku         *string  `json:"sku" validate:"omitempty,min=1,max=12"`
	Name        *string  `json:"name" validate:"omitempty,min=1,max=100"`
	Price       *int64   `json:"price" validate:"omitempty"`
	Weight      *float64 `json:"weight" validate:"omitempty"`
	Description *string  `json:"description" validate:"omitempty,min=1"`
	Stock       *int     `json:"stock" validate:"omitempty,min=1"`
}
