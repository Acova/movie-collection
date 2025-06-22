package main

import (
	"github.com/Acova/movie-collection/app"
	"github.com/Acova/movie-collection/app/adapter/httpadapter"
	"github.com/Acova/movie-collection/app/adapter/postgresadapter"
	"github.com/Acova/movie-collection/app/port"
)

func main() {
	dbConnection := postgresadapter.NewPostgresDBConnection()
	postgresUserRepository := postgresadapter.NewPostgresAdapterUserRepository(dbConnection)
	userController := port.UserController{
		Repo: postgresUserRepository,
	}

	app := app.NewApp()
	app.RegisterPort(&userController)
	httpadapter.StartHttpServer(app)
}
