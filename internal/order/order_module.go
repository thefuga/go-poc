package order

import (
	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/consumer"
	"github.com/thefuga/go-poc/internal/order/grpc"
	"github.com/thefuga/go-poc/internal/order/http"
	"github.com/thefuga/go-poc/internal/order/processor"
	"github.com/thefuga/go-poc/internal/order/producer"

	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		producer.Module,
		consumer.Module,
		processor.Module,
		channel.Module,
		grpc.Module,
		http.Module,
	)

	Invokables = fx.Options(
		producer.Invokables,
		consumer.Invokables,
		processor.Invokables,
		http.Invokables,
	)
)
