package event

import "encoding/json"

type Create struct {
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
}

func (event Create) Validate() error {
	return nil
}

func (event Create) Bytes() ([]byte, error) {
	return json.Marshal(event)
}
