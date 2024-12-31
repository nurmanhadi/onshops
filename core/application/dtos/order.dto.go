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
type OrderDetailRequestDto struct {
	OrderId     string `json:"order_id" validate:"required,min=1"`
	ProductId   string `json:"product_id" validate:"required,min=1"`
	Price       int64  `json:"price" validate:"required,min=1"`
	Sku         string `json:"sku" validate:"required,min=1"`
	Quantity    int    `json:"quantity" validate:"required,min=1"`
	GrossAmount int64  `json:"gross_amount" validate:"required,min=1"`
}
