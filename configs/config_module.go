package configs

import (
	"go.uber.org/fx"

	"github.com/thefuga/go-poc/configs/app"
)

var (
	Module     = fx.Options(app.Module)
	Invokables = fx.Options()
)
