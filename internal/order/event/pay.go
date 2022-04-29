package event

type Pay struct {
	OrderID         int `json:"order_id"`
	PaymentMethodID int `json:"payment_method_id"`
}

func (event Pay) Validate() error {
	return nil
}
