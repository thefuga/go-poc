package http

import (
	"github.com/thefuga/go-poc/internal/order/channel"

	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewOrderHandler,
		fx.ParamTags(
			channel.CreationProducer.Tag(),
			channel.PaymentProducer.Tag(),
			channel.CancellationProducer.Tag(),
		),
	),
)

var Invokables = fx.Invoke(
	InvokeOrderHandler,
)
