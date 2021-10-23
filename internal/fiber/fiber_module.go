package fiber

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"go.uber.org/fx"
)

func NewFiber(lifecycle fx.Lifecycle) *fiber.App {
	app := fiber.New()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go app.Listen("localhost:3001")
			return nil
		},
		OnStop: func(context.Context) error { return app.Shutdown() },
	})

	return app
}

var (
	Module     = fx.Provide(NewFiber)
	Invokables = fx.Invoke()
)
