package main

import (
	"time"

	"github.com/Acova/movie-collection/app/adapter/postgresadapter"
	"github.com/Acova/movie-collection/app/util"
	"github.com/joho/godotenv"
)

// This script should ONLY be used for testing in an empty databse.
// It populates the database with some initial data for testing purposes.
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	postgresDbConnection, err := postgresadapter.NewPostgresDBConnection()
	if err != nil {
		panic("Error connecting to the database: " + err.Error())
	}

	testPassword, err := util.HashPassword("password1")
	if err != nil {
		panic("Error hashing password: " + err.Error())
	}

	users := []postgresadapter.PostgresUser{
		{
			ID:       1,
			Email:    "test1@test.es",
			Name:     "test1",
			Password: testPassword,
		},
		{
			ID:       2,
			Email:    "test2@test.es",
			Name:     "test2",
			Password: testPassword,
		},
	}

	for _, user := range users {
		result := postgresDbConnection.DB.Create(&user)
		if result.Error != nil {
			panic("Error creating user: " + result.Error.Error())
		}
	}

	movies := []postgresadapter.PostgresMovie{
		{
			Title:       "The Godfather",
			Director:    "Francis Ford Coppola",
			ReleaseDate: time.Date(1972, 3, 24, 0, 0, 0, 0, time.UTC),
			Cast:        "Marlon Brando, Al Pacino, James Caan",
			Genre:       "Crime, Drama",
			Synopsis:    "The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.",
			Rating:      9.2,
			Duration:    175,
			PosterURL:   "http://example.com/poster1.jpg",
			UserID:      1,
		},
		{
			Title:       "Matrix",
			Director:    "The Wachowskis",
			ReleaseDate: time.Date(1999, 3, 31, 0, 0, 0, 0, time.UTC),
			Cast:        "Keanu Reeves, Laurence Fishburne, Carrie-Anne Moss",
			Genre:       "Action, Sci-Fi",
			Synopsis:    "A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.",
			Rating:      8.7,
			Duration:    136,
			PosterURL:   "http://example.com/poster2.jpg",
			UserID:      2,
		},
		{
			Title:       "The Avengers",
			Director:    "Joss Whedon",
			ReleaseDate: time.Date(2012, 5, 4, 0, 0, 0, 0, time.UTC),
			Cast:        "Robert Downey Jr., Chris Evans, Scarlett Johansson",
			Genre:       "Action, Adventure, Sci-Fi",
			Synopsis:    "The Avengers assemble to save the world from Loki and his army of aliens.",
			Rating:      8.0,
			Duration:    143,
			PosterURL:   "http://example.com/poster3.jpg",
			UserID:      1,
		},
		{
			Title:       "Lord of the Rings: The Return of the King",
			Director:    "Peter Jackson",
			ReleaseDate: time.Date(2003, 12, 17, 0, 0, 0, 0, time.UTC),
			Cast:        "Elijah Wood, Viggo Mortensen, Ian McKellen",
			Genre:       "Action, Adventure, Drama",
			Synopsis:    "The final battle for Middle-earth begins.",
			Rating:      8.9,
			Duration:    201,
			PosterURL:   "http://example.com/poster4.jpg",
			UserID:      2,
		},
	}

	for _, movie := range movies {
		result := postgresDbConnection.DB.Create(&movie)
		if result.Error != nil {
			panic("Error creating movie: " + result.Error.Error())
		}
	}
}
