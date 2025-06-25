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
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
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

	context.BindJSON(&user)
	a.port.CreateUser(user.ToDomain())
	context.IndentedJSON(http.StatusCreated, gin.H{"status": "User created"})
}
