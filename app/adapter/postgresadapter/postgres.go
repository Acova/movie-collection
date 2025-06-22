package postgresadapter

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDBConnection struct {
	DB *gorm.DB
}

func NewPostgresDBConnection() *PostgresDBConnection {
	dsn := "host=localhost user=user password=password dbname=movie_collection port=5432 sslmode=disable TimeZone=Europe/Lisbon"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database: " + err.Error())
	}

	db.AutoMigrate(&PostgresAdapterUser{})

	return &PostgresDBConnection{
		DB: db,
	}
}
