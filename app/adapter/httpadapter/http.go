package httpadapter

import (
	"reflect"

	"github.com/Acova/movie-collection/app"
	"github.com/Acova/movie-collection/app/port"
	"github.com/gin-gonic/gin"
)

func StartHttpServer(app *app.App) {
	userHttpAdapter := &HttpUserAdapter{
		port: app.GetPort(reflect.TypeOf(&port.UserPort{})).(*port.UserPort),
	}

	router := gin.Default()

	// User routes
	usersRouterGroup := router.Group("/user")
	usersRouterGroup.GET("", userHttpAdapter.ListUsers)
	usersRouterGroup.POST("", userHttpAdapter.CreateUser)

	router.Run("0.0.0.0:8081")
}
