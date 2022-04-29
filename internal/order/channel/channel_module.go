package channel

import (
	"github.com/thefuga/go-poc/internal/order/event"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotated{
		Target: NewOrderChannel[event.Create],
		Name:   "creation",
	},
	fx.Annotated{
		Target: NewOrderChannel[event.Pay],
		Name:   "payment",
	},
	fx.Annotated{
		Target: NewOrderChannel[event.Cancel],
		Name:   "cancellation",
	},
)
