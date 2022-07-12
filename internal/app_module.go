package internal

import (
	"github.com/thefuga/go-poc/configs"
	"github.com/thefuga/go-poc/internal/fiber"
	"github.com/thefuga/go-poc/internal/grpc"
	"github.com/thefuga/go-poc/internal/kafka"
	"github.com/thefuga/go-poc/internal/order"
	"github.com/thefuga/go-poc/internal/user"

	"go.uber.org/fx"
)

var (
	ApplicationModule = fx.Options(
		configs.Module,
		fiber.Module,
		grpc.Module,
		kafka.Module,
		user.Module,
		order.Module,
	)

	ApplicationInvokables = fx.Options(
		configs.Invokables,
		fiber.Invokables,
		grpc.Invokables,
		user.Invokables,
		order.Invokables,
	)
)
