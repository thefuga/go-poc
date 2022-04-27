package user

import (
	"github.com/thefuga/go-poc/internal/user/http"
	"github.com/thefuga/go-poc/internal/user/repository"

	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		repository.Module,
		http.Module,
		fx.Provide(fx.Annotate(
			repository.NewUserRepository,
			fx.As(new(http.UserFinderRepository)),
		)),
	)

	Invokables = fx.Options(
		repository.Invokables,
		http.Invokables,
	)
)
