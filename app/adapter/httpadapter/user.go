package httpadapter

import (
	"net/http"

	"github.com/Acova/movie-collection/app/port"
	"github.com/gin-gonic/gin"
)

type UserHttpAdapter struct {
	controller *port.UserPort
}

func (a *UserHttpAdapter) ListUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, a.controller.ListUsers())
}
