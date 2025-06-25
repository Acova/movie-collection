package main

import (
	"github.com/Acova/movie-collection/app/adapter/httpadapter"
	"github.com/Acova/movie-collection/app/adapter/postgresadapter"
	"github.com/Acova/movie-collection/app/service"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Initialize the application
	// app := app.NewApp() @TODO: Uncomment if you find any use for an application package

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

	// Initialize the controllers
	userService := service.NewUserService(postgresUserRepository)

	// Initialize the HTTP adapter
	httpadapter.StartHttpServer(userService)
}
