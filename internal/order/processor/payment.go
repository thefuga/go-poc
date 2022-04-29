package processor

import (
	"context"
	"fmt"

	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/event"
	"go.uber.org/fx"
)

type Payment struct {
	orderChan channel.OrderEventChannel[event.Pay]
}

func NewPaymentProcessor(orderChan channel.OrderEventChannel[event.Pay]) *Payment {
	return &Payment{
		orderChan: orderChan,
	}
}

func ProcessPaymentEvents(p *Payment, lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{OnStart: func(context.Context) error {
		go func() { p.orderChan.Listen(p.Process) }()
		return nil
	}})
}

func (p Payment) Process(event event.Pay) {
	fmt.Printf("Paying order: %v\n", event)
}
