package postgresadapter

import (
	"time"

	"github.com/Acova/movie-collection/app/domain"
	"gorm.io/gorm"
)

type PostgresMovie struct {
	gorm.Model
	ID          uint
	Title       string `gorm:"not null;uniqueIndex"`
	Director    string
	ReleaseDate time.Time
	Cast        string
	Genre       string
	Synopsis    string
	Rating      float64
	Duration    int
	PosterURL   string
}

func (PostgresMovie) TableName() string {
	return "movie"
}

func (m *PostgresMovie) ToDomain() *domain.Movie {
	return &domain.Movie{
		Title:       m.Title,
		Director:    m.Director,
		ReleaseDate: m.ReleaseDate,
		Cast:        m.Cast,
		Genre:       m.Genre,
		Synopsis:    m.Synopsis,
		Rating:      m.Rating,
		Duration:    m.Duration,
		PosterURL:   m.PosterURL,
	}
}

type PostgresMovieRepository struct {
	postgres *PostgresDBConnection
}

func NewPostgresMovieRepository(postgres *PostgresDBConnection) (*PostgresMovieRepository, error) {
	return &PostgresMovieRepository{
		postgres: postgres,
	}, nil
}

func (repository *PostgresMovieRepository) ListMovies() ([]*domain.Movie, error) {
	var postgresMovies []PostgresMovie
	result := repository.postgres.DB.Find(&postgresMovies)
	if result.Error != nil {
		return nil, result.Error
	}

	movies := make([]*domain.Movie, len(postgresMovies))
	for i, postgresMovie := range postgresMovies {
		movies[i] = postgresMovie.ToDomain()
	}

	return movies, nil
}

func (repository *PostgresMovieRepository) CreateMovie(movie *domain.Movie) error {
	postgresMovie := PostgresMovie{
		Title:       movie.Title,
		Director:    movie.Director,
		ReleaseDate: movie.ReleaseDate,
		Cast:        movie.Cast,
		Genre:       movie.Genre,
		Synopsis:    movie.Synopsis,
		Rating:      movie.Rating,
		Duration:    movie.Duration,
		PosterURL:   movie.PosterURL,
	}

	result := repository.postgres.DB.Create(&postgresMovie)
	return result.Error
}

func (repository *PostgresMovieRepository) GetMovieByTitle(title string) (*domain.Movie, error) {
	var postgresMovie PostgresMovie
	result := repository.postgres.DB.Where("title = ?", title).First(&postgresMovie)
	if result.Error != nil {
		return nil, result.Error
	}

	return postgresMovie.ToDomain(), nil
}
