package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/dto"
	"github.com/thefuga/go-poc/internal/order/event"
)

type (
	OrderHandler struct {
		creationProducer     channel.OrderEventChannel[event.Create]
		paymentProducer      channel.OrderEventChannel[event.Pay]
		cancellationProducer channel.OrderEventChannel[event.Cancel]
	}
)

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

func InvokeOrderHandler(handler *OrderHandler, app *fiber.App) {
	app.Post("/orders", handler.Create)
}

func (h OrderHandler) Create(ctx *fiber.Ctx) error {
	var (
		request  dto.CreateOrderRequest
		response dto.CreateOrderResponse
	)

	if parseErr := request.FromFiber(ctx); parseErr != nil {
		return parseErr
	}

	h.creationProducer <- request.ToEvent()

	// for brevity this response will be empty
	return ctx.JSON(response)
}
