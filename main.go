package main

import (
	"github.com/Acova/movie-collection/app"
	"github.com/Acova/movie-collection/app/adapter/httpadapter"
	"github.com/Acova/movie-collection/app/adapter/postgresadapter"
	"github.com/Acova/movie-collection/app/port"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Initialize the application
	app := app.NewApp()

	// Initialize the PostgreSQL database adapter
	dbConnection := postgresadapter.NewPostgresDBConnection()
	postgresUserRepository := postgresadapter.NewPostgresUserRepository(dbConnection)

	// Initialize the controllers
	userPort := port.UserPort{
		Repo: postgresUserRepository,
	}

	// Register the ports
	app.RegisterPort(&userPort)

	// Initialize the HTTP adapter
	httpadapter.StartHttpServer(app)
}
