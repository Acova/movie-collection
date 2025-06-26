package postgresadapter

import (
	"time"

	"github.com/Acova/movie-collection/app/domain"
	"gorm.io/gorm"
)

type PostgresMovie struct {
	gorm.Model
	ID          uint
	Title       string `gorm:"not null"`
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
		ID:          m.ID,
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

func FromDomain(movie *domain.Movie) (*PostgresMovie, error) {
	postgresMovie := &PostgresMovie{
		ID:          movie.ID,
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

	return postgresMovie, nil
}

type PostgresMovieRepository struct {
	postgres *PostgresDBConnection
}

func NewPostgresMovieRepository(postgres *PostgresDBConnection) (*PostgresMovieRepository, error) {
	return &PostgresMovieRepository{
		postgres: postgres,
	}, nil
}

func (repository *PostgresMovieRepository) ListMovies(filters map[string]string) ([]*domain.Movie, error) {
	var postgresMovies []PostgresMovie
	db := repository.postgres.DB
	for field, value := range filters {
		switch field {
		case "title":
			db = db.Where("title ILIKE ?", value)
		case "director":
			db = db.Where("director ILIKE ?", value)
		case "genre":
			db = db.Where("genre ILIKE ?", value)
		case "cast":
			db = db.Where("cast ILIKE ?", value)
		}
	}
	result := db.Find(&postgresMovies)
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
	postgresMovie, err := FromDomain(movie)
	if err != nil {
		return err
	}

	result := repository.postgres.DB.Create(&postgresMovie)
	return result.Error
}

func (repository *PostgresMovieRepository) GetMovie(id uint) (*domain.Movie, error) {
	postgresMovie := &PostgresMovie{}
	result := repository.postgres.DB.First(postgresMovie, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return postgresMovie.ToDomain(), nil
}

func (repository *PostgresMovieRepository) UpdateMovie(movie *domain.Movie) error {
	postgresMovie, err := FromDomain(movie)
	if err != nil {
		return err
	}

	result := repository.postgres.DB.Save(&postgresMovie)
	return result.Error
}

func (repository *PostgresMovieRepository) DeleteMovie(movie *domain.Movie) error {
	postgresMovie, err := FromDomain(movie)
	if err != nil {
		return err
	}

	result := repository.postgres.DB.Delete(&postgresMovie)
	return result.Error
}
