package internal

import (
	"github.com/thefuga/go-poc/configs"
	"github.com/thefuga/go-poc/internal/fiber"
	"github.com/thefuga/go-poc/internal/order"
	"github.com/thefuga/go-poc/internal/user"

	"go.uber.org/fx"
)

var (
	ApplicationModule = fx.Options(
		configs.Module,
		fiber.Module,
		user.Module,
		order.Module,
	)

	ApplicationInvokables = fx.Options(
		configs.Invokables,
		fiber.Invokables,
		user.Invokables,
		order.Invokables,
	)
)
