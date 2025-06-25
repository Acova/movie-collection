package port

import "github.com/Acova/movie-collection/app/domain"

type MovieRepository interface {
	CreateMovie(movie *domain.Movie) error
	ListMovies() ([]*domain.Movie, error)
	GetMovieByTitle(id string) (*domain.Movie, error)
	UpdateMovie(movie *domain.Movie) error
	DeleteMovie(id string) error
}

type MovieService interface {
	CreateMovie(movie *domain.Movie) error
	ListMovies() ([]*domain.Movie, error)
	GetMovieByTitle(title string) (*domain.Movie, error)
	UpdateMovie(movie *domain.Movie) error
	DeleteMovie(id string) error
}
