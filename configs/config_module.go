package configs

import (
	. "go.uber.org/fx"

	"github.com/spf13/viper"
	"github.com/thefuga/go-template/configs/app"
)

var (
	Module = Provide(
		// This enables you to inject Viper as a provider in case you want to
		// use your configs as a service.
		func() *viper.Viper {
			return viper.GetViper()
		},
	)

	Invokables = Options(
		app.Invokables,
	)
)
