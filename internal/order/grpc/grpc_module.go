package grpc

import (
	"github.com/thefuga/go-poc/internal/order/channel"

	order "go.buf.build/grpc/go/thefuga/go-poc/order/v1"
	"go.uber.org/fx"
	"google.golang.org/grpc"
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
	func(server *grpc.Server, handler *OrderHandler) {
		order.RegisterOrderServiceServer(server, handler)
	},
)
