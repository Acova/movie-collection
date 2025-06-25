package mock

import (
	"errors"

	"github.com/Acova/movie-collection/app/domain"
)

type MockUserService struct {
	Users []*domain.User
}

func (m *MockUserService) CreateUser(user *domain.User) {
	m.Users = append(m.Users, user)
}

func (m *MockUserService) ListUsers() []*domain.User {
	return m.Users
}

func (m *MockUserService) GetLoginUser(email, password string) (*domain.User, error) {
	for _, user := range m.Users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}
	return nil, errors.New("user not found or password incorrect")
}

func (m *MockUserService) GetUserByEmail(email string) (*domain.User, error) {
	for _, user := range m.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
