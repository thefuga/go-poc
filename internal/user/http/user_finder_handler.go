package http

import (
	"github.com/thefuga/go-template/internal/user/entity"

	"github.com/gofiber/fiber/v2"
)

type (
	UserFinderRepository interface {
		FindByFirstName(string) entity.User
	}

	UserFinderHandler struct {
		repository UserFinderRepository
	}
)

func InvokeUserFinderHandler(handler *UserFinderHandler, app *fiber.App) {
	app.Get("/users", handler.FindUser)
}

func NewUserFinderHandler(userFinder UserFinderRepository) *UserFinderHandler {
	return &UserFinderHandler{
		repository: userFinder,
	}
}

func (h UserFinderHandler) FindUser(ctx *fiber.Ctx) error {
	return ctx.JSON(h.repository.FindByFirstName(ctx.Query("first_name")))
}
