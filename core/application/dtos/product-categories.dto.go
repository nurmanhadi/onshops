package dtos

type ProductCategoriesRequestDto struct {
	CategoryId uint `json:"category_id" validate:"required"`
}
