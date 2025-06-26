package httpadapter

import (
	"fmt"
	"net/http"

	"github.com/Acova/movie-collection/app/domain"
	"github.com/Acova/movie-collection/app/port"
	"github.com/gin-gonic/gin"
)

type HttpUserAdapter struct {
	userService port.UserService
}

type HttpUser struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=8,max=40"`
}

func (u *HttpUser) ToDomain() *domain.User {
	return &domain.User{
		Email:    u.Email,
		Name:     u.Name,
		Password: u.Password,
	}
}

func NewHttpUserAdapter(userService port.UserService) *HttpUserAdapter {
	return &HttpUserAdapter{
		userService: userService,
	}
}

func (a *HttpUserAdapter) ListUsers(context *gin.Context) {
	users, err := a.userService.ListUsers()
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, users)
}

func (a *HttpUserAdapter) CreateUser(context *gin.Context) {
	var user HttpUser

	if err := context.BindJSON(&user); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := a.userService.CreateUser(user.ToDomain())
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create user: %v", err)})
		return
	}
	context.IndentedJSON(http.StatusCreated, gin.H{"status": "User created"})
}
