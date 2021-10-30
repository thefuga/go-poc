package user

import (
	"github.com/thefuga/go-template/internal/user/http"
	"github.com/thefuga/go-template/internal/user/repository"

	"go.uber.org/fx"
)

var (
	Module = fx.Options(
		repository.Module,
		http.Module,
		// This maps unmaped interfaces to implementations without coupling
		// the actual implementation to the a client module.
		fx.Provide(
			func(r *repository.UserRepository) http.UserFinderRepository {
				return r
			},
		),
	)

	Invokables = fx.Options(
		repository.Invokables,
		http.Invokables,
	)
)
