package dto

import (
	"github.com/thefuga/go-poc/internal/order/event"
	proto "go.buf.build/grpc/go/thefuga/go-poc/order/v1"
)

type (
	PayOrderRequest struct {
		OrderID int `json:"order_id"`
	}

	PayOrderResponse struct{}
)

func (dto *PayOrderRequest) FromProto(p *proto.PayRequest) *PayOrderRequest {
	dto.OrderID = int(p.OrderId)

	return dto
}

func (dto PayOrderRequest) ToEvent() event.Pay {
	return event.Pay{OrderID: int(dto.OrderID)}
}

func (dto *PayOrderResponse) ToProto() *proto.PayResponse {
	return &proto.PayResponse{}
}
