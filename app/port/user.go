package port

import (
	"github.com/Acova/movie-collection/app/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	ListUsers() ([]*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}

type UserService interface {
	CreateUser(user *domain.User) error
	ListUsers() ([]*domain.User, error)
	GetLoginUser(email, password string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}
