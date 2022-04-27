package app

import (
	"errors"

	"go.uber.org/fx"

	"github.com/spf13/viper"
)

var Module = fx.Provide(
	// This enables you to inject Viper as a provider in case you want to
	// use your configs as a service.
	func() (*viper.Viper, error) {
		if err := viper.BindEnv("app.env", "APP_ENV"); err != nil {
			return nil, err
		}

		if !viper.IsSet("app.env") {
			return nil, errors.New("The APP_ENV variable must be set!")
		}

		viper.SetConfigName(viper.GetString("app.env"))
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./configs/app")

		if err := viper.MergeInConfig(); err != nil {
			return nil, err
		}

		return viper.GetViper(), nil
	},
)
