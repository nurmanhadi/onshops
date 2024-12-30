package dtos

type CustomerRequestDto struct {
	Name    *string `json:"name" validate:"omitempty,min=1,max=100"`
	Country *string `json:"country" validate:"omitempty,min=1,max=50"`
	Phone   *string `json:"phone" validate:"omitempty,min=1,max=20"`
}
