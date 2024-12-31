package dtos

type PaymentRequestDto struct {
	OrderId     string      `json:"order_id" validate:"required,min=1,max=36"`
	GrossAmount int64       `json:"gross_amount" validate:"required"`
	ItemDetails ItemDetails `json:"item_details"`
}
type ItemDetails struct {
	ProductId string `json:"product_id" validate:"required,min=1,max=36"`
	Price     int64  `json:"price" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
	Name      string `json:"name" validate:"required,min=1,max=100"`
	Category  string `json:"category" validate:"required,min=1,max=50"`
}
type PaymentNotificationDto struct {
	TransactionTime   string `json:"transaction_time" validate:"required"`
	TransactionStatus string `json:"transaction_status" validate:"required"`
	TransactionID     string `json:"transaction_id" validate:"required"`
	StatusMessage     string `json:"status_message" validate:"required"`
	StatusCode        string `json:"status_code" validate:"required"`
	SignatureKey      string `json:"signature_key" validate:"required"`
	SettlementTime    string `json:"settlement_time" validate:"required"`
	PaymentType       string `json:"payment_type" validate:"required"`
	OrderID           string `json:"order_id" validate:"required"`
	MerchantID        string `json:"merchant_id" validate:"required"`
	GrossAmount       string `json:"gross_amount" validate:"required"`
	FraudStatus       string `json:"fraud_status" validate:"required"`
	Currency          string `json:"currency" validate:"required"`
}
