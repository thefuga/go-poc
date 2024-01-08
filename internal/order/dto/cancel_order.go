package dto

import (
	proto "buf.build/gen/go/thefuga/go-poc/protocolbuffers/go/order/v1"
	"github.com/thefuga/go-poc/internal/order/event"
)

type (
	CancelOrderRequest struct {
		OrderID int `json:"order_id"`
	}

	CancelOrderResponse struct{}
)

func (dto *CancelOrderRequest) FromProto(p *proto.CancelRequest) *CancelOrderRequest {
	dto.OrderID = int(p.OrderId)

	return dto
}

func (dto CancelOrderRequest) ToEvent() event.Cancel {
	return event.Cancel{OrderID: int(dto.OrderID)}
}

func (dto *CancelOrderResponse) ToProto() *proto.CancelResponse {
	return &proto.CancelResponse{}
}
