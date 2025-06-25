package main

import (
	"github.com/Acova/movie-collection/app/adapter/postgresadapter"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	postgresDbConnection := postgresadapter.NewPostgresDBConnection()
	postgresDbConnection.DB.AutoMigrate(&postgresadapter.PostgresUser{})
}
