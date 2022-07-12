/*
producer package TODO
*/
package producer

import (
	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/event"

	"go.uber.org/fx"
)

// Module defines the resolutions for processors of each event type.
var Module = fx.Provide(
	fx.Annotated{
		Target: NewProducer[event.Create],
		Name:   event.CreateAnnotation.String(),
	},
	fx.Annotated{
		Target: NewProducer[event.Pay],
		Name:   event.PayAnnotation.String(),
	},
	fx.Annotated{
		Target: NewProducer[event.Cancel],
		Name:   event.CancelAnnotation.String(),
	},
)

// Invokables defines mappings for channels of each event type.
var Invokables = fx.Invoke(
	fx.Annotate(
		func(
			p *Producer[event.Create],
			c channel.OrderEventChannel[event.Create],
			lifecycle fx.Lifecycle,
		) {
			p.RunProducer(c, event.CreateAnnotation.String(), lifecycle)
		},
		fx.ParamTags(event.CreateAnnotation.Tag(), channel.CreationProducer.Tag()),
	),
	fx.Annotate(
		func(
			p *Producer[event.Pay],
			c channel.OrderEventChannel[event.Pay],
			lifecycle fx.Lifecycle,
		) {
			p.RunProducer(c, event.PayAnnotation.String(), lifecycle)
		},
		fx.ParamTags(event.PayAnnotation.Tag(), channel.PaymentProducer.Tag()),
	),
	fx.Annotate(
		func(
			p *Producer[event.Cancel],
			c channel.OrderEventChannel[event.Cancel],
			lifecycle fx.Lifecycle,
		) {
			p.RunProducer(c, event.CancelAnnotation.String(), lifecycle)
		},
		fx.ParamTags(event.CancelAnnotation.Tag(), channel.CancellationProducer.Tag()),
	),
)
