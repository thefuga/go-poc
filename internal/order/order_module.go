package order

import (
	"github.com/thefuga/go-poc/internal/order/channel"
	"github.com/thefuga/go-poc/internal/order/consumer"
	"github.com/thefuga/go-poc/internal/order/processor"

	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		consumer.Module,
		processor.Module,
		channel.Module,
	)

	Invokables = fx.Options(
		consumer.Invokables,
		processor.Invokables,
	)
)
