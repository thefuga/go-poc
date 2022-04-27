package repository

import (
	"github.com/thefuga/go-poc/internal/user/entity"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (UserRepository) FindByFirstName(firstName string) entity.User {
	return entity.User{
		FirstName: firstName,
	}
}
