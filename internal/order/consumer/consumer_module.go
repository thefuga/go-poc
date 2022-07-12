/*
consumer is the package for the default order events kafka consumer.
*/
package consumer

import (
	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/event"

	"go.uber.org/fx"
)

// Module registers a constructor for each consumer constrained by the event type.
var Module = fx.Provide(
	fx.Annotated{
		Target: NewConsumer[event.Create],
		Name:   event.CreateAnnotation.String(),
	},
	fx.Annotated{
		Target: NewConsumer[event.Pay],
		Name:   event.PayAnnotation.String(),
	},
	fx.Annotated{
		Target: NewConsumer[event.Cancel],
		Name:   event.CancelAnnotation.String(),
	},
)

// Invokables annotate each consumer with their respective event types to subscribe
// to their topics and run them with the appropriate channel. See RunConsumer for
// details on the channel injection.
var Invokables = fx.Invoke(
	fx.Annotate(
		func(c *Consumer[event.Create]) error {
			return SubscribeTopics(
				c.consumer, []string{event.CreateAnnotation.String()},
			)
		},
		fx.ParamTags(event.CreateAnnotation.Tag()),
	),
	fx.Annotate(
		func(c *Consumer[event.Pay]) error {
			return SubscribeTopics(
				c.consumer, []string{event.PayAnnotation.String()},
			)
		},
		fx.ParamTags(event.PayAnnotation.Tag()),
	),
	fx.Annotate(
		func(c *Consumer[event.Cancel]) error {
			return SubscribeTopics(
				c.consumer, []string{event.CancelAnnotation.String()},
			)
		},
		fx.ParamTags(event.CancelAnnotation.Tag()),
	),

	fx.Annotate(
		RunConsumer[event.Create],
		fx.ParamTags(
			event.CreateAnnotation.Tag(),
			channel.CreationConsumer.Tag(),
		),
	),
	fx.Annotate(
		RunConsumer[event.Pay],
		fx.ParamTags(
			event.PayAnnotation.Tag(),
			channel.PaymentConsumer.Tag(),
		),
	),
	fx.Annotate(
		RunConsumer[event.Cancel],
		fx.ParamTags(
			event.CancelAnnotation.Tag(),
			channel.CancellationConsumer.Tag(),
		),
	),
)
