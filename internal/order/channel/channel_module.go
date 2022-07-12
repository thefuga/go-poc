package channel

import (
	"github.com/thefuga/go-poc/internal/annotation"
	"github.com/thefuga/go-poc/internal/order/event"

	"go.uber.org/fx"
)

const (
	CreationConsumer     = annotation.Annotation("consumer-creation")
	PaymentConsumer      = annotation.Annotation("consumer-payment")
	CancellationConsumer = annotation.Annotation("consumer-cancellation")

	CreationProducer     = annotation.Annotation("producer-creation")
	PaymentProducer      = annotation.Annotation("producer-payment")
	CancellationProducer = annotation.Annotation("producer-cancellation")
)

// Module for all order channels used throughout the application.
var Module = fx.Provide(
	fx.Annotated{
		Target: NewOrderChannel[event.Create],
		Name:   CreationConsumer.String(),
	},
	fx.Annotated{
		Target: NewOrderChannel[event.Pay],
		Name:   PaymentConsumer.String(),
	},
	fx.Annotated{
		Target: NewOrderChannel[event.Cancel],
		Name:   CancellationConsumer.String(),
	},

	fx.Annotated{
		Target: NewOrderChannel[event.Create],
		Name:   CreationProducer.String(),
	},
	fx.Annotated{
		Target: NewOrderChannel[event.Pay],
		Name:   PaymentProducer.String(),
	},
	fx.Annotated{
		Target: NewOrderChannel[event.Cancel],
		Name:   CancellationProducer.String(),
	},
)
