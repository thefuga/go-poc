package user

import (
	"github.com/thefuga/go-template/internal/user/http"
	"github.com/thefuga/go-template/internal/user/repository"

	. "go.uber.org/fx"
)

var (
	Module = Options(
		repository.Module,
		http.Module,
		Provide(Annotate(
			repository.NewUserRepository,
			As(new(http.UserFinderRepository)),
		)),
	)

	Invokables = Options(
		repository.Invokables,
		http.Invokables,
	)
)
