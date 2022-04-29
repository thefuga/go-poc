package event

type Create struct {
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
}

func (event Create) Validate() error {
	return nil
}
