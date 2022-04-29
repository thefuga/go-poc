package processor

import (
	"context"
	"fmt"

	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/event"
	"go.uber.org/fx"
)

type Creation struct {
	orderChan channel.OrderEventChannel[event.Create]
}

func NewCreationProcessor(orderChan channel.OrderEventChannel[event.Create]) *Creation {
	return &Creation{
		orderChan: orderChan,
	}
}

func ProcessCreationEvents(p *Creation, lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{OnStart: func(context.Context) error {
		go func() { p.orderChan.Listen(p.Process) }()
		return nil
	}})
}

func (p Creation) Process(event event.Create) {
	fmt.Printf("Creating order: %v\n", event)
}
