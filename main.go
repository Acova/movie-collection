package main

import (
	"github.com/Acova/movie-collection/app/adapter/httpadapter"
	"github.com/Acova/movie-collection/app/adapter/postgresadapter"
	"github.com/Acova/movie-collection/app/service"
	"github.com/joho/godotenv"
)

// @title Movie Collection API
// @version 1.0
// @description This is a simple API for managing a movie collection.

// @host localhost:8080
// @SecurityDefinitions.basic BasicAuth
// @tokenUrl /login
func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Initialize the PostgreSQL database adapter
	dbConnection, err := postgresadapter.NewPostgresDBConnection()
	if err != nil {
		panic("Error connecting to the database: " + err.Error())
	}

	// Initialize the repositories
	postgresUserRepository, err := postgresadapter.NewPostgresUserRepository(dbConnection)
	if err != nil {
		panic("Error creating user repository: " + err.Error())
	}

	postgresMovieRepository, err := postgresadapter.NewPostgresMovieRepository(dbConnection)
	if err != nil {
		panic("Error creating movie repository: " + err.Error())
	}

	// Initialize the controllers
	userService := service.NewUserService(postgresUserRepository)
	movieService := service.NewMovieService(postgresMovieRepository)

	// Initialize the HTTP adapter
	services := &httpadapter.HttpServices{
		UserService:  userService,
		MovieService: movieService,
	}
	httpadapter.StartHttpServer(services)
}
