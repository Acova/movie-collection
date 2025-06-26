package service

import (
	"testing"

	"github.com/Acova/movie-collection/app/domain"
	"github.com/Acova/movie-collection/app/port/mock"
)

func TestCreateMovie(t *testing.T) {
	mockRepository := &mock.MockMovieRepository{
		Movies: make([]*domain.Movie, 0),
	}

	movieService := NewMovieService(mockRepository)
	movie := &domain.Movie{
		ID:          1,
		Title:       "Inception",
		Synopsis:    "A mind-bending thriller",
		ReleaseYear: 2010,
		Director:    "Christopher Nolan",
		Rating:      8.8,
	}

	err := movieService.CreateMovie(movie)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(mockRepository.Movies) != 1 {
		t.Errorf("Expected 1 movie in repository, got %d", len(mockRepository.Movies))
	}
	if mockRepository.Movies[0].Title != movie.Title {
		t.Errorf("Expected movie title %s, got %s", movie.Title, mockRepository.Movies[0].Title)
	}
	if mockRepository.Movies[0].Synopsis != movie.Synopsis {
		t.Errorf("Expected movie synopsis %s, got %s", movie.Synopsis, mockRepository.Movies[0].Synopsis)
	}
	if mockRepository.Movies[0].ReleaseYear != movie.ReleaseYear {
		t.Errorf("Expected movie release year %v, got %v", movie.ReleaseYear, mockRepository.Movies[0].ReleaseYear)
	}
	if mockRepository.Movies[0].Director != movie.Director {
		t.Errorf("Expected movie director %s, got %s", movie.Director, mockRepository.Movies[0].Director)
	}
	if mockRepository.Movies[0].Rating != movie.Rating {
		t.Errorf("Expected movie rating %f, got %f", movie.Rating, mockRepository.Movies[0].Rating)
	}
}

func TestListMovies(t *testing.T) {
	mockRepository := &mock.MockMovieRepository{
		Movies: []*domain.Movie{
			{
				ID:          1,
				Title:       "Inception",
				Synopsis:    "A mind-bending thriller",
				ReleaseYear: 2010,
				Director:    "Christopher Nolan",
				Rating:      8.8,
			},
			{
				ID:          2,
				Title:       "The Matrix",
				Synopsis:    "A sci-fi classic",
				ReleaseYear: 1999,
				Director:    "The Wachowskis",
				Rating:      8.7,
			},
		},
	}

	movieService := NewMovieService(mockRepository)
	movies, err := movieService.ListMovies(make(map[string]string))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(movies) != 2 {
		t.Errorf("Expected 2 movies, got %d", len(movies))
	}
	for i, movie := range movies {
		if movie.Title != mockRepository.Movies[i].Title {
			t.Errorf("Expected movie title %s, got %s", mockRepository.Movies[i].Title, movie.Title)
		}
		if movie.Synopsis != mockRepository.Movies[i].Synopsis {
			t.Errorf("Expected movie synopsis %s, got %s", mockRepository.Movies[i].Synopsis, movie.Synopsis)
		}
		if movie.ReleaseYear != mockRepository.Movies[i].ReleaseYear {
			t.Errorf("Expected movie release year %v, got %v", mockRepository.Movies[i].ReleaseYear, movie.ReleaseYear)
		}
		if movie.Director != mockRepository.Movies[i].Director {
			t.Errorf("Expected movie director %s, got %s", mockRepository.Movies[i].Director, movie.Director)
		}
		if movie.Rating != mockRepository.Movies[i].Rating {
			t.Errorf("Expected movie rating %f, got %f", mockRepository.Movies[i].Rating, movie.Rating)
		}
	}
}

func TestGetMovie(t *testing.T) {
	mockRepository := &mock.MockMovieRepository{
		Movies: []*domain.Movie{
			{
				ID:          1,
				Title:       "Inception",
				Synopsis:    "A mind-bending thriller",
				ReleaseYear: 2010,
				Director:    "Christopher Nolan",
				Rating:      8.8,
			},
		},
	}

	movieService := NewMovieService(mockRepository)
	movie, err := movieService.GetMovie(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if movie.Title != "Inception" {
		t.Errorf("Expected movie title 'Inception', got %s", movie.Title)
	}
	if movie.Synopsis != "A mind-bending thriller" {
		t.Errorf("Expected movie synopsis 'A mind-bending thriller', got %s", movie.Synopsis)
	}
	if movie.Director != "Christopher Nolan" {
		t.Errorf("Expected movie director 'Christopher Nolan', got %s", movie.Director)
	}
	if movie.Rating != 8.8 {
		t.Errorf("Expected movie rating 8.8, got %f", movie.Rating)
	}
}

func TestUpdateMovie(t *testing.T) {
	mockRepository := &mock.MockMovieRepository{
		Movies: []*domain.Movie{
			{
				ID:          1,
				Title:       "Inception",
				Synopsis:    "A mind-bending thriller",
				ReleaseYear: 2010,
				Director:    "Christopher Nolan",
				Rating:      8.8,
			},
		},
	}

	movieService := NewMovieService(mockRepository)
	movie := &domain.Movie{
		ID:          1,
		Title:       "Inception",
		Synopsis:    "An updated synopsis",
		ReleaseYear: 2010,
		Director:    "Christopher Nolan",
		Rating:      9.0,
	}

	err := movieService.UpdateMovie(movie)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if mockRepository.Movies[0].Synopsis != movie.Synopsis {
		t.Errorf("Expected updated synopsis %s, got %s", movie.Synopsis, mockRepository.Movies[0].Synopsis)
	}
	if mockRepository.Movies[0].Rating != movie.Rating {
		t.Errorf("Expected updated rating %f, got %f", movie.Rating, mockRepository.Movies[0].Rating)
	}
}

func TestDeleteMovie(t *testing.T) {
	mockRepository := &mock.MockMovieRepository{
		Movies: []*domain.Movie{
			{
				ID:          1,
				Title:       "Inception",
				Synopsis:    "A mind-bending thriller",
				ReleaseYear: 2010,
				Director:    "Christopher Nolan",
				Rating:      8.8,
			},
		},
	}

	movieService := NewMovieService(mockRepository)
	movie := &domain.Movie{
		Title: "Inception",
	}

	err := movieService.DeleteMovie(movie)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(mockRepository.Movies) != 0 {
		t.Errorf("Expected 0 movies in repository, got %d", len(mockRepository.Movies))
	}
}
