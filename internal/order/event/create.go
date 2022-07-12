package event

import (
	"encoding/json"

	"github.com/thefuga/go-poc/internal/annotation"
)

const CreateAnnotation = annotation.Annotation("creation")

type Create struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
}

func (event Create) Validate() error {
	return nil
}

func (event Create) Bytes() ([]byte, error) {
	return json.Marshal(event)
}
