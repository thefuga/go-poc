package event

import "encoding/json"

type Pay struct {
	OrderID         int `json:"order_id"`
	PaymentMethodID int `json:"payment_method_id"`
}

func (event Pay) Validate() error {
	return nil
}

func (event Pay) Bytes() ([]byte, error) {
	return json.Marshal(event)
}
