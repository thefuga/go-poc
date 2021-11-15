package http

import (
	. "go.uber.org/fx"
)

var (
	Module = Provide(
		NewUserFinderHandler,
	)

	Invokables = Invoke(
		InvokeUserFinderHandler,
	)
)
