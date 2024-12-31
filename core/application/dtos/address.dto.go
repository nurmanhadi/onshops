package dtos

type AddressAddRequestDto struct {
	RecipientName string `json:"recipient_name" validate:"required,min=1,max=100"`
	Phone         string `json:"phone" validate:"required,min=1,max=20"`
	Street        string `json:"street" validate:"required,min=1,max=255"`
	City          string `json:"city" validate:"required,min=1,max=50"`
	State         string `json:"state" validate:"required,min=1,max=50"`
	PortalCode    string `json:"portal_code" validate:"required,min=1,max=20"`
	Country       string `json:"country" validate:"required,min=1,max=50"`
}
type AddressUpdateRequestDto struct {
	RecipientName *string `json:"recipient_name" validate:"omitempty,min=1,max=100"`
	Phone         *string `json:"phone" validate:"omitempty,min=1,max=20"`
	Street        *string `json:"street" validate:"omitempty,min=1,max=255"`
	City          *string `json:"city" validate:"omitempty,min=1,max=50"`
	State         *string `json:"state" validate:"omitempty,min=1,max=50"`
	PortalCode    *string `json:"portal_code" validate:"omitempty,min=1,max=20"`
	Country       *string `json:"country" validate:"omitempty,min=1,max=50"`
}
