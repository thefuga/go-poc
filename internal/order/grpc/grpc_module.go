package grpc

import (
	"github.com/thefuga/go-poc/internal/order/channel"

	order "buf.build/gen/go/thefuga/go-poc/grpc/go/order/v1/orderv1grpc"
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
