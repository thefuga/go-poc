package configs

import (
	"github.com/thefuga/go-template/configs/app"

	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		app.Module,
	)

	Invokables = fx.Options(
		app.Invokables,
	)
)
