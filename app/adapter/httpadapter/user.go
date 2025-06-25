package httpadapter

import (
	"net/http"

	"github.com/Acova/movie-collection/app/domain"
	"github.com/Acova/movie-collection/app/port"
	"github.com/gin-gonic/gin"
)

type HttpUserAdapter struct {
	port *port.UserPort
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

func (a *HttpUserAdapter) ListUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, a.port.ListUsers())
}

func (a *HttpUserAdapter) CreateUser(context *gin.Context) {
	var user HttpUser

	if err := context.BindJSON(&user); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	a.port.CreateUser(user.ToDomain())
	context.IndentedJSON(http.StatusCreated, gin.H{"status": "User created"})
}
