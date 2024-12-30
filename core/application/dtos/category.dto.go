package dtos

type CategoryRequestDto struct {
	Name *string `json:"name" validate:"omitempty,min=1,max=100"`
}
