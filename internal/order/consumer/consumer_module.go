package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewConfig,
	fx.Annotated{
		Target: NewConsumer,
		Name:   "creation",
	},
	fx.Annotated{
		Target: NewConsumer,
		Name:   "payment",
	},
	fx.Annotated{
		Target: NewConsumer,
		Name:   "cancellation",
	},
)

var Invokables = fx.Invoke(
	fx.Annotate(
		func(consumer *kafka.Consumer) error {
			return SubscribeTopics(consumer, []string{"creation"})
		},
		fx.ParamTags(`name:"creation"`),
	),
	fx.Annotate(
		func(consumer *kafka.Consumer) error {
			return SubscribeTopics(consumer, []string{"payment"})
		},
		fx.ParamTags(`name:"payment"`),
	),
	fx.Annotate(
		func(consumer *kafka.Consumer) error {
			return SubscribeTopics(consumer, []string{"cancellation"})
		},
		fx.ParamTags(`name:"cancellation"`),
	),

	fx.Annotate(
		RunConsumer,
		fx.ParamTags(`name:"creation"`),
	),
	fx.Annotate(
		RunConsumer,
		fx.ParamTags(`name:"payment"`),
	),
	fx.Annotate(
		RunConsumer,
		fx.ParamTags(`name:"cancellation"`),
	),
)
