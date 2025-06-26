package mock

import (
	"errors"

	"github.com/Acova/movie-collection/app/domain"
)

type MockUserRepository struct {
	Users []*domain.User
}

func (r *MockUserRepository) ListUsers() ([]*domain.User, error) {
	return r.Users, nil
}

func (r *MockUserRepository) CreateUser(user *domain.User) error {
	r.Users = append(r.Users, user)
	return nil
}

func (r *MockUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	for _, user := range r.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

type MockUserService struct {
	Users []*domain.User
}

func (m *MockUserService) CreateUser(user *domain.User) error {
	m.Users = append(m.Users, user)
	return nil
}

func (m *MockUserService) ListUsers() ([]*domain.User, error) {
	return m.Users, nil
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
