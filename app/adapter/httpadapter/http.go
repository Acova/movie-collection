package httpadapter

import (
	"fmt"
	"os"
	"time"

	"github.com/Acova/movie-collection/app/domain"
	"github.com/Acova/movie-collection/app/port"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func StartHttpServer(userService port.UserService) {
	httpUserAdapter := NewHttpUserAdapter(userService)

	// Create a new Gin engine
	engine := gin.Default()

	// Middleware to handle JWT
	jwtMiddleware, err := jwt.New(getJwtInitParams(userService))

	if err != nil {
		panic("JWT middleware initialization failed: " + err.Error())
	}

	engine.Use(handleMiddleware(jwtMiddleware))

	// Login route
	engine.POST("/login", jwtMiddleware.LoginHandler)

	// Refresh route
	engine.GET("/refresh_token", jwtMiddleware.MiddlewareFunc(), jwtMiddleware.RefreshHandler)

	// User routes
	usersRouterGroup := engine.Group("/user", jwtMiddleware.MiddlewareFunc())
	usersRouterGroup.GET("", httpUserAdapter.ListUsers)
	usersRouterGroup.POST("", httpUserAdapter.CreateUser)

	engine.Run("0.0.0.0:8081")
}

type LoginForm struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func getJwtInitParams(userService port.UserService) *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:       "movie-collection",
		Key:         []byte(os.Getenv("JWT_SECRET_KEY")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "id",
		Authenticator: func(c *gin.Context) (interface{}, error) {
			loginForm := LoginForm{}
			if err := c.ShouldBindJSON(&loginForm); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userEmail := loginForm.Email
			userPassword := loginForm.Password

			return userService.GetLoginUser(userEmail, userPassword)
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*domain.User); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Println("Calling PayloadFunc with data:", data)
			if user, ok := data.(*domain.User); ok {
				return jwt.MapClaims{
					"id":    user.Email,
					"name":  user.Name,
					"email": user.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			fmt.Println("Calling IdentityHandler")
			claims := jwt.ExtractClaims(c)
			return &domain.User{
				Email: claims["email"].(string),
				Name:  claims["name"].(string),
			}
		},
	}
}

func handleMiddleware(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := authMiddleware.MiddlewareInit()
		if err != nil {
			panic("JWT middleware initialization failed: " + err.Error())
		}
	}
}
