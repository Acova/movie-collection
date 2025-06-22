package port

import (
	"github.com/Acova/movie-collection/app/domain"
)

type UserRepository interface {
	ListUsers() []domain.User
	CreateUser(user domain.User)
}

type UserController struct {
	Repo UserRepository
}

func (c *UserController) GetPortName() string {
	return "user"
}

func (c *UserController) ListUsers() []domain.User {
	return make([]domain.User, 0) // @TODO: Implement this method
	// users, err := c.repo.ListUsers()
	// if err != nil {
	// 	return nil, err
	// }
	// return users, nil
}

func (c *UserController) CreateUser(user domain.User) {
	c.Repo.CreateUser(user)
}
