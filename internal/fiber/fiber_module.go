package fiber

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewFiber)
	Invokables = fx.Invoke(InvokeFiber)
)

func NewFiber() *fiber.App {
	return fiber.New()
}

func InvokeFiber(app *fiber.App, lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go app.Listen(viper.GetString("app.fiber.address"))
			return nil
		},
		OnStop: func(context.Context) error { return app.Shutdown() },
	})
}
