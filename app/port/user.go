package port

import (
	"github.com/Acova/movie-collection/app/domain"
)

type UserRepository interface {
	ListUsers() []domain.User
	CreateUser(user domain.User)
}

type UserPort struct {
	Repo UserRepository
}

func (c *UserPort) GetPortName() string {
	return "user"
}

func (c *UserPort) ListUsers() []domain.User {
	return c.Repo.ListUsers()
}

func (c *UserPort) CreateUser(user domain.User) {
	c.Repo.CreateUser(user)
}
