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
