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
	return make([]domain.User, 0) // @TODO: Implement this method
	// users, err := c.repo.ListUsers()
	// if err != nil {
	// 	return nil, err
	// }
	// return users, nil
}

func (c *UserPort) CreateUser(user domain.User) {
	c.Repo.CreateUser(user)
}
