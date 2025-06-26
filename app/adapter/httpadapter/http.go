package httpadapter

import (
	"os"
	"time"

	"github.com/Acova/movie-collection/app/domain"
	"github.com/Acova/movie-collection/app/port"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type HttpServices struct {
	UserService  port.UserService
	MovieService port.MovieService
}

func StartHttpServer(services *HttpServices) {
	httpUserAdapter := NewHttpUserAdapter(services.UserService)

	// Create a new Gin engine
	engine := gin.Default()

	// Middleware to handle JWT
	jwtMiddleware, err := jwt.New(getJwtInitParams(services.UserService))

	if err != nil {
		panic("JWT middleware initialization failed: " + err.Error())
	}

	engine.Use(handleMiddleware(jwtMiddleware))

	// Login route
	engine.POST("/login", jwtMiddleware.LoginHandler)

	// User registration route
	engine.POST("/user", httpUserAdapter.CreateUser)

	// Refresh route
	engine.GET("/refresh_token", jwtMiddleware.MiddlewareFunc(), jwtMiddleware.RefreshHandler)

	// User routes
	usersRouterGroup := engine.Group("/user", jwtMiddleware.MiddlewareFunc())
	usersRouterGroup.GET("", httpUserAdapter.ListUsers)

	// Movie routes
	httpMovieAdapter := NewHttpMovieAdapter(services.MovieService)
	moviesRouterGroup := engine.Group("/movie", jwtMiddleware.MiddlewareFunc())
	moviesRouterGroup.POST("", httpMovieAdapter.CreateMovie)
	moviesRouterGroup.GET("", httpMovieAdapter.ListMovies)
	moviesRouterGroup.GET("/:id", httpMovieAdapter.GetMovie)
	moviesRouterGroup.PUT("/:id", httpMovieAdapter.UpdateMovie)
	moviesRouterGroup.DELETE("/:id", httpMovieAdapter.DeleteMovie)

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
			if user, ok := data.(*domain.User); ok {
				return jwt.MapClaims{
					"id":    user.ID,
					"name":  user.Name,
					"email": user.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &domain.User{
				ID:    uint(claims["id"].(float64)),
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

func GetLoggedInUser(c *gin.Context) (*domain.User, bool) {
	userValue, exists := c.Get("id")
	if !exists {
		return nil, false
	}

	user, ok := userValue.(*domain.User)
	if !ok {
		return nil, false
	}

	return user, true
}
