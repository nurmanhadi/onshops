package dtos

type OrderRequestDto struct {
	CustomerId      string `json:"customer_id" validate:"required"`
	ShippingAddress string `json:"shipping_address" validate:"required,min=1"`
	OrderAddress    string `json:"order_address" validate:"required,min=1"`
	OrderEmail      string `json:"order_email" validate:"required,email,min=1,max=100"`
}
type OrderUpdateRequestDto struct {
	Amount          *int64  `json:"amount" validate:"omitempty"`
	ShippingAddress *string `json:"shipping_address" validate:"omitempty,min=1"`
	OrderAddress    *string `json:"order_address" validate:"omitempty,min=1"`
	OrderEmail      *string `json:"order_email" validate:"omitempty,email,min=1,max=100"`
	OrderStatus     *string `json:"order_status" validate:"omitempty,min=1,max=12"`
}
