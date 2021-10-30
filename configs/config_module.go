package configs

import (
	"github.com/spf13/viper"
	"github.com/thefuga/go-template/configs/app"

	"go.uber.org/fx"
)

var (
	Module = fx.Provide(
		// This enables you to inject Viper as a provider in case you wan't to
		// use your configs as a service.
		func() *viper.Viper {
			return viper.GetViper()
		},
	)

	Invokables = fx.Options(
		app.Invokables,
	)
)
