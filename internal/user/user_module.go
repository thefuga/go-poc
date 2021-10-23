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
	)

	Invokables = fx.Options(
		repository.Invokables,
		http.Invokables,
	)
)
