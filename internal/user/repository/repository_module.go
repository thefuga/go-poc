package repository

import (
	"go.uber.org/fx"
)

var (
	Module = fx.Provide(
		NewUserRepository,
	)

	Invokables = fx.Invoke()
)
