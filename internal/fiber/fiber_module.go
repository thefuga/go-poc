package fiber

import (
	"context"

	. "go.uber.org/fx"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

var (
	Module     = Provide(NewFiber)
	Invokables = Invoke(InvokeFiber)
)

func NewFiber() *fiber.App {
	return fiber.New()
}

func InvokeFiber(
	app *fiber.App,
	lifecycle Lifecycle,
	configs *viper.Viper,
) {
	lifecycle.Append(Hook{
		OnStart: func(context.Context) error {
			go app.Listen(configs.GetString("app.fiber.address")) //nolint
			return nil
		},
		OnStop: func(context.Context) error { return app.Shutdown() },
	})
}
