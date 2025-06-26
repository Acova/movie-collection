package postgresadapter

import (
	"testing"
	"time"

	"github.com/Acova/movie-collection/app/domain"
)

func TestPostgresMovieReturnsTableName(t *testing.T) {
	expectedTableName := "movie"
	actualTableName := PostgresMovie{}.TableName()

	if actualTableName != expectedTableName {
		t.Errorf("Expected table name '%s', but got '%s'", expectedTableName, actualTableName)
	}
}

func TestPostgresMovieToDomain(t *testing.T) {
	postgresMovie := PostgresMovie{
		Title:       "Inception",
		Director:    "Christopher Nolan",
		ReleaseDate: time.Now(),
		Cast:        "Leonardo DiCaprio, Joseph Gordon-Levitt",
		Genre:       "Sci-Fi",
		Synopsis:    "A mind-bending thriller",
		Rating:      8.8,
		Duration:    148,
		PosterURL:   "http://example.com/poster.jpg",
	}

	domainMovie := postgresMovie.ToDomain()

	if domainMovie.Title != postgresMovie.Title {
		t.Errorf("Expected title '%s', got '%s'", postgresMovie.Title, domainMovie.Title)
	}
	if domainMovie.Director != postgresMovie.Director {
		t.Errorf("Expected director '%s', got '%s'", postgresMovie.Director, domainMovie.Director)
	}
	if domainMovie.ReleaseDate != postgresMovie.ReleaseDate {
		t.Errorf("Expected release date '%v', got '%v'", postgresMovie.ReleaseDate, domainMovie.ReleaseDate)
	}
	if domainMovie.Cast != postgresMovie.Cast {
		t.Errorf("Expected cast '%s', got '%s'", postgresMovie.Cast, domainMovie.Cast)
	}
	if domainMovie.Genre != postgresMovie.Genre {
		t.Errorf("Expected genre '%s', got '%s'", postgresMovie.Genre, domainMovie.Genre)
	}
	if domainMovie.Synopsis != postgresMovie.Synopsis {
		t.Errorf("Expected synopsis '%s', got '%s'", postgresMovie.Synopsis, domainMovie.Synopsis)
	}
	if domainMovie.Rating != postgresMovie.Rating {
		t.Errorf("Expected rating %f, got %f", postgresMovie.Rating, domainMovie.Rating)
	}
	if domainMovie.Duration != postgresMovie.Duration {
		t.Errorf("Expected duration %d, got %d", postgresMovie.Duration, domainMovie.Duration)
	}
	if domainMovie.PosterURL != postgresMovie.PosterURL {
		t.Errorf("Expected poster URL '%s', got '%s'", postgresMovie.PosterURL, domainMovie.PosterURL)
	}
}

func TestPostgresMovieFromDomain(t *testing.T) {
	domainMovie := &domain.Movie{
		Title:       "Inception",
		Director:    "Christopher Nolan",
		ReleaseDate: time.Now(),
		Cast:        "Leonardo DiCaprio, Joseph Gordon-Levitt",
		Genre:       "Sci-Fi",
		Synopsis:    "A mind-bending thriller",
		Rating:      8.8,
		Duration:    148,
		PosterURL:   "http://example.com/poster.jpg",
	}

	postgresMovie, err := FromDomain(domainMovie)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if postgresMovie.Title != domainMovie.Title {
		t.Errorf("Expected title '%s', got '%s'", domainMovie.Title, postgresMovie.Title)
	}
	if postgresMovie.Director != domainMovie.Director {
		t.Errorf("Expected director '%s', got '%s'", domainMovie.Director, postgresMovie.Director)
	}
	if postgresMovie.ReleaseDate != domainMovie.ReleaseDate {
		t.Errorf("Expected release date '%v', got '%v'", domainMovie.ReleaseDate, postgresMovie.ReleaseDate)
	}
	if postgresMovie.Cast != domainMovie.Cast {
		t.Errorf("Expected cast '%s', got '%s'", domainMovie.Cast, postgresMovie.Cast)
	}
	if postgresMovie.Genre != domainMovie.Genre {
		t.Errorf("Expected genre '%s', got '%s'", domainMovie.Genre, postgresMovie.Genre)
	}
	if postgresMovie.Synopsis != domainMovie.Synopsis {
		t.Errorf("Expected synopsis '%s', got '%s'", domainMovie.Synopsis, postgresMovie.Synopsis)
	}
	if postgresMovie.Rating != domainMovie.Rating {
		t.Errorf("Expected rating %f, got %f", domainMovie.Rating, postgresMovie.Rating)
	}
	if postgresMovie.Duration != domainMovie.Duration {
		t.Errorf("Expected duration %d, got %d", domainMovie.Duration, postgresMovie.Duration)
	}
	if postgresMovie.PosterURL != domainMovie.PosterURL {
		t.Errorf("Expected poster URL '%s', got '%s'", domainMovie.PosterURL, postgresMovie.PosterURL)
	}
}
