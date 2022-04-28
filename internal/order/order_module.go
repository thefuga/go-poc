package order

import (
	"github.com/thefuga/go-poc/internal/order/consumer"

	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		consumer.Module,
	)

	Invokables = fx.Options(
		consumer.Invokables,
	)
)
