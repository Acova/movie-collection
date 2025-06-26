package port

import "github.com/Acova/movie-collection/app/domain"

type MovieRepository interface {
	CreateMovie(movie *domain.Movie) error
	ListMovies() ([]*domain.Movie, error)
	GetMovie(id uint) (*domain.Movie, error)
	UpdateMovie(movie *domain.Movie) error
	DeleteMovie(movie *domain.Movie) error
}

type MovieService interface {
	CreateMovie(movie *domain.Movie) error
	ListMovies() ([]*domain.Movie, error)
	GetMovie(id uint) (*domain.Movie, error)
	UpdateMovie(movie *domain.Movie) error
	DeleteMovie(movie *domain.Movie) error
}
