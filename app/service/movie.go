package service

import (
	"github.com/Acova/movie-collection/app/domain"
	"github.com/Acova/movie-collection/app/port"
)

type MovieService struct {
	Repo port.MovieRepository
}

func NewMovieService(repo port.MovieRepository) *MovieService {
	return &MovieService{
		Repo: repo,
	}
}

func (m *MovieService) CreateMovie(movie *domain.Movie) error {
	return m.Repo.CreateMovie(movie)
}

func (m *MovieService) ListMovies() ([]*domain.Movie, error) {
	movies, err := m.Repo.ListMovies()
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (m *MovieService) GetMovie(id uint) (*domain.Movie, error) {
	movie, err := m.Repo.GetMovie(id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (m *MovieService) UpdateMovie(movie *domain.Movie) error {
	return m.Repo.UpdateMovie(movie)
}

func (m *MovieService) DeleteMovie(movie *domain.Movie) error {
	return m.Repo.DeleteMovie(movie)
}
