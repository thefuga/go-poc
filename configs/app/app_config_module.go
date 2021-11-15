package app

import (
	"context"
	"errors"

	. "go.uber.org/fx"

	"github.com/spf13/viper"
)

var Invokables = Invoke(InvokeAppConfigs)

func InvokeAppConfigs(lifecycle Lifecycle) {
	lifecycle.Append(Hook{
		OnStart: func(context.Context) error {
			if err := viper.BindEnv("app.env", "APP_ENV"); err != nil {
				return err
			}

			if !viper.IsSet("app.env") {
				return errors.New("The APP_ENV variable must be set!")
			}

			viper.SetConfigName(viper.GetString("app.env"))
			viper.SetConfigType("yaml")
			viper.AddConfigPath("./configs/app")

			return viper.MergeInConfig()
		},
	})
}
