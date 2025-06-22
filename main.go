package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello-world", helloWorld)
	router.Run("0.0.0.0:8080")
}

func helloWorld(c *gin.Context) {
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte("Hola, mundo!"))
}
