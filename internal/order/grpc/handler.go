package grpc

import (
	"context"

	order "buf.build/gen/go/thefuga/go-poc/protocolbuffers/go/order/v1"
	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/dto"
	"github.com/thefuga/go-poc/internal/order/event"
)

type OrderHandler struct {
	creationProducer     channel.OrderEventChannel[event.Create]
	paymentProducer      channel.OrderEventChannel[event.Pay]
	cancellationProducer channel.OrderEventChannel[event.Cancel]
}

func NewOrderHandler(
	creationProducer channel.OrderEventChannel[event.Create],
	paymentProducer channel.OrderEventChannel[event.Pay],
	cancellationProducer channel.OrderEventChannel[event.Cancel],
) *OrderHandler {
	return &OrderHandler{
		creationProducer:     creationProducer,
		paymentProducer:      paymentProducer,
		cancellationProducer: cancellationProducer,
	}
}
func (h OrderHandler) Create(
	ctx context.Context, in *order.CreateRequest,
) (*order.CreateResponse, error) {
	var (
		request  dto.CreateOrderRequest
		response dto.CreateOrderResponse
	)

	h.creationProducer <- request.FromProto(in).ToEvent()

	// for brevity this response will be empty
	return response.ToProto(), nil
}

func (h OrderHandler) Pay(
	ctx context.Context, in *order.PayRequest,
) (*order.PayResponse, error) {
	var (
		request  dto.PayOrderRequest
		response dto.PayOrderResponse
	)

	h.paymentProducer <- request.FromProto(in).ToEvent()

	// for brevity this response will be empty
	return response.ToProto(), nil
}

func (h OrderHandler) Cancel(
	ctx context.Context, in *order.CancelRequest,
) (*order.CancelResponse, error) {
	var (
		request  dto.CancelOrderRequest
		response dto.CancelOrderResponse
	)

	h.cancellationProducer <- request.FromProto(in).ToEvent()

	// for brevity this response will be empty
	return response.ToProto(), nil
}
