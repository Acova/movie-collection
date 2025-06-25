package httpadapter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Acova/movie-collection/app/domain"
	"github.com/Acova/movie-collection/app/port/mock"
	"github.com/gin-gonic/gin"
)

func TestHttpUserReturnsDomainUser(t *testing.T) {
	user := &HttpUser{
		Email:    "test@example.com",
		Name:     "Test User",
		Password: "password123",
	}

	domainUser := user.ToDomain()

	if domainUser.Email != user.Email || domainUser.Name != user.Name || domainUser.Password != user.Password {
		t.Errorf("Expected domain user to be %+v, but got %+v", user.ToDomain(), domainUser)
	}
}

func TestListUsers(t *testing.T) {
	// Mock the user service
	mockUserService := mock.MockUserService{
		Users: []*domain.User{
			{Email: "test@example.com", Name: "Test User", Password: "password123"},
		},
	}

	httpAdapter := NewHttpUserAdapter(&mockUserService)
	users := httpAdapter.userService.ListUsers()
	if len(users) != 1 {
		t.Errorf("Expected 1 user, but got %d", len(users))
	}
	if users[0].Email != "test@example.com" {
		t.Errorf("Expected email to be 'test@example.com', but got '%s'", users[0].Email)
	}
	if users[0].Name != "Test User" {
		t.Errorf("Expected name to be 'Test User', but got '%s'", users[0].Name)
	}
	if users[0].Password != "password123" {
		t.Errorf("Expected password to be 'password123', but got '%s'", users[0].Password)
	}
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUserService := &mock.MockUserService{
		Users: []*domain.User{},
	}

	httpAdapter := NewHttpUserAdapter(mockUserService)

	user := &HttpUser{
		Email:    "newuser@example.com",
		Name:     "New User",
		Password: "newpassword123",
	}

	body, _ := json.Marshal(user)
	request, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	mockResponseWriter := httptest.NewRecorder()
	mockContext, _ := gin.CreateTestContext(mockResponseWriter)
	mockContext.Request = request

	httpAdapter.CreateUser(mockContext)

	if len(mockUserService.Users) != 1 {
		t.Errorf("Expected 1 user, but got %d", len(mockUserService.Users))
	}

	if mockUserService.Users[0].Email != user.Email {
		t.Errorf("Expected email to be '%s', but got '%s'", user.Email, mockUserService.Users[0].Email)
	}

	if mockUserService.Users[0].Name != user.Name {
		t.Errorf("Expected name to be '%s', but got '%s'", user.Name, mockUserService.Users[0].Name)
	}

	if mockUserService.Users[0].Password != user.Password {
		t.Errorf("Expected password to be '%s', but got '%s'", user.Password, mockUserService.Users[0].Password)
	}
}
