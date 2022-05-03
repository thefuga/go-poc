package producer

import (
	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/event"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotated{
		Target: NewProducer[event.Create],
		Name:   "creation",
	},
	fx.Annotated{
		Target: NewProducer[event.Pay],
		Name:   "payment",
	},
	fx.Annotated{
		Target: NewProducer[event.Cancel],
		Name:   "cancellation",
	},
)

var Invokables = fx.Invoke(
	fx.Annotate(
		func(
			p *Producer[event.Create],
			c channel.OrderEventChannel[event.Create],
			lifecycle fx.Lifecycle,
		) {
			p.RunProducer(c, "creation", lifecycle)
		},
		fx.ParamTags(`name:"creation"`, `name:"producer-creation"`),
	),
	fx.Annotate(
		func(
			p *Producer[event.Pay],
			c channel.OrderEventChannel[event.Pay],
			lifecycle fx.Lifecycle,
		) {
			p.RunProducer(c, "payment", lifecycle)
		},
		fx.ParamTags(`name:"payment"`, `name:"producer-payment"`),
	),
	fx.Annotate(
		func(
			p *Producer[event.Cancel],
			c channel.OrderEventChannel[event.Cancel],
			lifecycle fx.Lifecycle,
		) {
			p.RunProducer(c, "cancellation", lifecycle)
		},
		fx.ParamTags(`name:"cancellation"`, `name:"producer-cancellation"`),
	),
)
