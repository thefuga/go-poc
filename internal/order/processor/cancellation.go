package processor

import (
	"context"
	"fmt"

	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/event"
	"go.uber.org/fx"
)

type Cancellation struct {
	orderChan channel.OrderEventChannel[event.Cancel]
}

func NewCancellationProcessor(
	orderChan channel.OrderEventChannel[event.Cancel],
) *Cancellation {
	return &Cancellation{
		orderChan: orderChan,
	}
}

func ProcessCancellationEvents(p *Cancellation, lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{OnStart: func(context.Context) error {
		go func() { p.orderChan.Listen(p.Process) }()
		return nil
	}})
}

func (p Cancellation) Process(event event.Cancel) {
	fmt.Printf("Canceling order: %v\n", event)
}
