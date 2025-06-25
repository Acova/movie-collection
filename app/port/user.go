package port

import (
	"github.com/Acova/movie-collection/app/domain"
)

type UserRepository interface {
	ListUsers() []*domain.User
	CreateUser(user *domain.User)
	GetUserByEmail(email string) (*domain.User, error)
}

type UserService interface {
	CreateUser(user *domain.User)
	ListUsers() []*domain.User
	GetLoginUser(email, password string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
}
