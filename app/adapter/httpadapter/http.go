package httpadapter

import (
	"reflect"

	"github.com/Acova/movie-collection/app"
	"github.com/Acova/movie-collection/app/port"
	"github.com/gin-gonic/gin"
)

func StartHttpServer(app *app.App) {
	userHttpAdapter := &UserHttpAdapter{
		controller: app.GetPort(reflect.TypeOf(&port.UserPort{})).(*port.UserPort),
	}

	router := gin.Default()
	router.GET("/users", userHttpAdapter.ListUsers)
	router.Run("0.0.0.0:8081")
}
