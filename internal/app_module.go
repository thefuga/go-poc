package internal

import (
	"github.com/thefuga/go-poc/configs"
	"github.com/thefuga/go-poc/internal/fiber"
	"github.com/thefuga/go-poc/internal/user"

	. "go.uber.org/fx"
)

var (
	ApplicationModule = Options(
		configs.Module,
		fiber.Module,
		user.Module,
	)

	ApplicationInvokables = Options(
		configs.Invokables,
		fiber.Invokables,
		user.Invokables,
	)
)
