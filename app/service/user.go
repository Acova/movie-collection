package service

import (
	"github.com/Acova/movie-collection/app/domain"
	"github.com/Acova/movie-collection/app/port"
	"github.com/Acova/movie-collection/app/util"
)

type UserPort struct {
	Repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserPort {
	return &UserPort{
		Repo: repo,
	}
}

func (c *UserPort) ListUsers() []*domain.User {
	return c.Repo.ListUsers()
}

func (c *UserPort) CreateUser(user *domain.User) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		panic("Error hashing password: " + err.Error())
	}

	user.Password = hashedPassword
	c.Repo.CreateUser(user)
}

func (c *UserPort) GetLoginUser(email, password string) (*domain.User, error) {
	user, err := c.Repo.GetUserByEmail(email)
	if err != nil {
		return &domain.User{}, err
	}

	err = util.ComparePasswords(password, user.Password)
	if err != nil {
		return &domain.User{}, err
	}

	return user, nil
}

func (c *UserPort) GetUserByEmail(email string) (*domain.User, error) {
	user, err := c.Repo.GetUserByEmail(email)
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}
