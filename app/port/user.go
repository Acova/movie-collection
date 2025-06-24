package port

import (
	"github.com/Acova/movie-collection/app/domain"
)

type UserRepository interface {
	ListUsers() []domain.User
	CreateUser(user domain.User)
	GetUserByEmail(email string) (domain.User, error)
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

func (c *UserPort) GetUserByEmail(email string) (domain.User, error) {
	user, err := c.Repo.GetUserByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
