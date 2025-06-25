package service

import (
	"testing"

	"github.com/Acova/movie-collection/app/domain"
	"github.com/Acova/movie-collection/app/port/mock"
	"github.com/Acova/movie-collection/app/util"
)

func TestListUsers(t *testing.T) {
	mockRepository := &mock.MockUserRepository{
		Users: []*domain.User{
			{Email: "user1@example.com", Password: "password1"},
			{Email: "user2@example.com", Password: "password2"},
		},
	}

	userService := NewUserService(mockRepository)
	users := userService.ListUsers()
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
	if users[0].Email != "user1@example.com" {
		t.Errorf("Expected user1@example.com, got %s", users[0].Email)
	}
	if users[1].Email != "user2@example.com" {
		t.Errorf("Expected user2@example.com, got %s", users[1].Email)
	}
	if users[0].Password != "password1" {
		t.Errorf("Expected password1, got %s", users[0].Password)
	}
	if users[1].Password != "password2" {
		t.Errorf("Expected password2, got %s", users[1].Password)
	}
}

func TestCreateUser(t *testing.T) {
	mockRepository := &mock.MockUserRepository{
		Users: []*domain.User{},
	}

	userService := NewUserService(mockRepository)
	newUser := &domain.User{Email: "user3@example.com", Password: "password3"}
	userService.CreateUser(newUser)

	if len(mockRepository.Users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(mockRepository.Users))
	}
	if mockRepository.Users[0].Email != "user3@example.com" {
		t.Errorf("Expected user3@example.com, got %s", mockRepository.Users[0].Email)
	}
	if ok := util.ComparePasswords("password3", mockRepository.Users[0].Password); nil != ok {
		t.Errorf("Expected password3, got %s", mockRepository.Users[0].Password)
	}
}

func TestGetLoginUser(t *testing.T) {
	password, ok := util.HashPassword("longpassword")
	if ok != nil {
		t.Fatalf("Error hashing password: %v", ok)
	}

	mockRepository := &mock.MockUserRepository{
		Users: []*domain.User{
			{Email: "user1@example.com", Password: password},
		},
	}

	userService := NewUserService(mockRepository)
	user, err := userService.GetLoginUser("user1@example.com", "longpassword")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user.Email != "user1@example.com" {
		t.Errorf("Expected user1@example.com, got %s", user.Email)
	}
	if ok := util.ComparePasswords("longpassword", user.Password); nil != ok {
		t.Errorf("Expected longpassword, got %s", user.Password)
	}
}

func TestGetUserByEmail(t *testing.T) {
	mockRepository := &mock.MockUserRepository{
		Users: []*domain.User{
			{Email: "user1@example.com", Password: "password"},
		},
	}

	userService := NewUserService(mockRepository)
	user, err := userService.GetUserByEmail("user1@example.com")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user.Email != "user1@example.com" {
		t.Errorf("Expected user1@example.com, got %s", user.Email)
	}
	if user.Password != "password" {
		t.Errorf("Expected password, got %s", user.Password)
	}
}
