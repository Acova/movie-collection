package main

import (
	"github.com/Acova/movie-collection/app"
	"github.com/Acova/movie-collection/app/adapter/http_adapter"
	"github.com/Acova/movie-collection/app/domain"
	"github.com/Acova/movie-collection/app/port"
)

type MockUserRepository struct {
	Users []domain.User
}

func (m *MockUserRepository) ListUsers() []domain.User {
	return m.Users
}

func (m *MockUserRepository) CreateUser(user domain.User) {
	m.Users = append(m.Users, user)
}

func main() {
	mockUserRepository := &MockUserRepository{
		Users: make([]domain.User, 0),
	}
	userController := &port.UserController{
		Repo: mockUserRepository,
	}

	app := app.NewApp()
	app.RegisterPort(userController)
	http_adapter.StartHttpServer(app)
}
