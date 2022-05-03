package consumer

import (
	"github.com/thefuga/go-poc/internal/order/event"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewConfig,
	fx.Annotated{
		Target: NewConsumer[event.Create],
		Name:   "creation",
	},
	fx.Annotated{
		Target: NewConsumer[event.Pay],
		Name:   "payment",
	},
	fx.Annotated{
		Target: NewConsumer[event.Cancel],
		Name:   "cancellation",
	},
)

var Invokables = fx.Invoke(
	fx.Annotate(
		func(c *Consumer[event.Create]) error {
			return SubscribeTopics(c.consumer, []string{"creation"})
		},
		fx.ParamTags(`name:"creation"`),
	),
	fx.Annotate(
		func(c *Consumer[event.Pay]) error {
			return SubscribeTopics(c.consumer, []string{"payment"})
		},
		fx.ParamTags(`name:"payment"`),
	),
	fx.Annotate(
		func(c *Consumer[event.Cancel]) error {
			return SubscribeTopics(c.consumer, []string{"cancellation"})
		},
		fx.ParamTags(`name:"cancellation"`),
	),

	fx.Annotate(
		RunConsumer[event.Create],
		fx.ParamTags(`name:"creation"`, `name:"consumer-creation"`),
	),
	fx.Annotate(
		RunConsumer[event.Pay],
		fx.ParamTags(`name:"payment"`, `name:"consumer-payment"`),
	),
	fx.Annotate(
		RunConsumer[event.Cancel],
		fx.ParamTags(`name:"cancellation"`, `name:"consumer-cancellation"`),
	),
)
