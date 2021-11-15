package repository

import (
	. "go.uber.org/fx"
)

var (
	Module = Provide(
		NewUserRepository,
	)

	Invokables = Invoke()
)
