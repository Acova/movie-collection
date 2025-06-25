package mock

import (
	"errors"

	"github.com/Acova/movie-collection/app/domain"
)

type MockMovieRepository struct {
	Movies []*domain.Movie
}

func (m *MockMovieRepository) CreateMovie(movie *domain.Movie) error {
	m.Movies = append(m.Movies, movie)
	return nil
}

func (m *MockMovieRepository) ListMovies() ([]*domain.Movie, error) {
	return m.Movies, nil
}

func (m *MockMovieRepository) GetMovieByTitle(title string) (*domain.Movie, error) {
	for _, movie := range m.Movies {
		if movie.Title == title {
			return movie, nil
		}
	}
	return nil, errors.New("movie not found")
}

func (m *MockMovieRepository) UpdateMovie(movie *domain.Movie) error {
	for i, v := range m.Movies {
		if v.Title == movie.Title {
			m.Movies[i] = movie
			return nil
		}
	}
	return errors.New("movie not found")
}

func (m *MockMovieRepository) DeleteMovie(movie domain.Movie) error {
	for i, v := range m.Movies {
		if v.Title == movie.Title {
			m.Movies = append(m.Movies[:i], m.Movies[i+1:]...)
			return nil
		}
	}
	return errors.New("movie not found")
}

type MockMovieService struct {
	Movies []*domain.Movie
}

func (m *MockMovieService) CreateMovie(movie *domain.Movie) error {
	m.Movies = append(m.Movies, movie)
	return nil
}

func (m *MockMovieService) ListMovies() ([]*domain.Movie, error) {
	return m.Movies, nil
}

func (m *MockMovieService) GetMovieByTitle(title string) (*domain.Movie, error) {
	for _, movie := range m.Movies {
		if movie.Title == title {
			return movie, nil
		}
	}
	return nil, errors.New("movie not found")
}

func (m *MockMovieService) UpdateMovie(movie *domain.Movie) error {
	for i, v := range m.Movies {
		if v.Title == movie.Title {
			m.Movies[i] = movie
			return nil
		}
	}
	return errors.New("movie not found")
}

func (m *MockMovieService) DeleteMovie(movie domain.Movie) error {
	for i, v := range m.Movies {
		if v.Title == movie.Title {
			m.Movies = append(m.Movies[:i], m.Movies[i+1:]...)
			return nil
		}
	}
	return errors.New("movie not found")
}
