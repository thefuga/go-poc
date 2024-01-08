package dto

import (
	"github.com/thefuga/go-poc/internal/order/event"
	"github.com/thefuga/go-poc/internal/order/model"

	proto "buf.build/gen/go/thefuga/go-poc/protocolbuffers/go/order/v1"
	"github.com/gofiber/fiber/v2"
)

type (
	CreateOrderRequest struct {
		UserID    int `json:"user_id"`
		ProductID int `json:"product_id"`
	}

	CreateOrderResponse struct {
		OrderID int `json:"order_id"`
	}
)

func (dto *CreateOrderRequest) FromProto(p *proto.CreateRequest) *CreateOrderRequest {
	dto.UserID = int(p.UserId)
	dto.ProductID = int(p.ProductId)

	return dto
}

func (dto CreateOrderRequest) ToEvent() event.Create {
	return event.Create{
		UserID:    int(dto.UserID),
		ProductID: int(dto.ProductID),
	}
}

func (dto *CreateOrderRequest) FromFiber(ctx *fiber.Ctx) error {
	return ctx.BodyParser(dto)
}

func (dto *CreateOrderResponse) FromOrder(m model.Order) *CreateOrderResponse {
	dto.OrderID = int(m.UserID)

	return dto
}

func (dto *CreateOrderResponse) ToProto() *proto.CreateResponse {
	return &proto.CreateResponse{OrderId: int64(dto.OrderID)}
}
