package event

type Cancel struct {
	OrderID int `json:"order_id"`
}

func (event Cancel) Validate() error {
	return nil
}
