package event

import "encoding/json"

type Cancel struct {
	OrderID int `json:"order_id"`
}

func (event Cancel) Validate() error {
	return nil
}

func (event Cancel) Bytes() ([]byte, error) {
	return json.Marshal(event)
}
