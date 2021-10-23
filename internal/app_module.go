package internal

import (
	"github.com/thefuga/go-template/internal/fiber"
	"github.com/thefuga/go-template/internal/user"

	"go.uber.org/fx"
)

var (
	ApplicationModule = fx.Options(
		fiber.Module,
		user.Module,
	)

	ApplicationInvokables = fx.Options(
		fiber.Invokables,
		user.Invokables,
	)
)
