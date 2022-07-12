package event

import (
	"encoding/json"

	"github.com/thefuga/go-poc/internal/annotation"
)

const CancelAnnotation = annotation.Annotation("cancellation")

type Cancel struct {
	OrderID int `json:"order_id"`
}

func (event Cancel) Validate() error {
	return nil
}

func (event Cancel) Bytes() ([]byte, error) {
	return json.Marshal(event)
}
