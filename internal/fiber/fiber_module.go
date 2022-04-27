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

func InvokeFiber(
	app *fiber.App,
	lifecycle fx.Lifecycle,
	configs *viper.Viper,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go app.Listen(configs.GetString("app.fiber.address")) //nolint
			return nil
		},
		OnStop: func(context.Context) error { return app.Shutdown() },
	})
}
