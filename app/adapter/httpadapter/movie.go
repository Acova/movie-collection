package httpadapter

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Acova/movie-collection/app/domain"
	"github.com/Acova/movie-collection/app/port"
	"github.com/gin-gonic/gin"
)

type HttpMovieAdapter struct {
	movieService port.MovieService
}

type HttpMovie struct {
	Title       string    `json:"title" binding:"required,min=1,max=100"`
	Director    string    `json:"director" binding:"min=1,max=50"`
	Synopsis    string    `json:"synopsis" binding:"min=1,max=500"`
	ReleaseDate time.Time `json:"release_date"`
	Cast        string    `json:"cast" binding:"min=1,max=200"`
	Genre       string    `json:"genre" binding:"min=1,max=50"`
	Rating      float64   `json:"rating" binding:"min=0,max=10"`
	Duration    int       `json:"duration" binding:"min=0"`
	PosterURL   string    `json:"poster_url"`
}

func (h *HttpMovie) ToDomain() *domain.Movie {
	return &domain.Movie{
		Title:       h.Title,
		Director:    h.Director,
		Synopsis:    h.Synopsis,
		ReleaseDate: h.ReleaseDate,
		Cast:        h.Cast,
		Genre:       h.Genre,
		Rating:      h.Rating,
		Duration:    h.Duration,
		PosterURL:   h.PosterURL,
	}
}

func NewHttpMovieAdapter(movieService port.MovieService) *HttpMovieAdapter {
	return &HttpMovieAdapter{
		movieService: movieService,
	}
}

func (h *HttpMovieAdapter) CreateMovie(context *gin.Context) error {
	movie := HttpMovie{}

	if err := context.BindJSON(&movie); err != nil {
		return context.AbortWithError(http.StatusBadRequest, err)
	}

	h.movieService.CreateMovie(movie.ToDomain())
	context.IndentedJSON(http.StatusCreated, gin.H{"status": "Movie created"})

	return nil
}

func (h *HttpMovieAdapter) ListMovies(context *gin.Context) error {
	movies, err := h.movieService.ListMovies()
	if err != nil {
		return context.AbortWithError(http.StatusInternalServerError, err)
	}

	context.IndentedJSON(http.StatusOK, movies)
	return nil
}

func (h *HttpMovieAdapter) GetMovie(context *gin.Context) error {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		return context.AbortWithError(http.StatusBadRequest, err)
	}

	movie, err := h.movieService.GetMovie(uint(id))
	if err != nil {
		return context.AbortWithError(http.StatusNotFound, err)
	}

	context.IndentedJSON(http.StatusOK, movie)
	return nil
}

func (h *HttpMovieAdapter) UpdateMovie(context *gin.Context) error {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		return context.AbortWithError(http.StatusBadRequest, err)
	}

	_, err = h.movieService.GetMovie(uint(id))
	if err != nil {
		return context.AbortWithError(http.StatusNotFound, err)
	}

	updatedMovie := HttpMovie{}
	if err := context.BindJSON(&updatedMovie); err != nil {
		return context.AbortWithError(http.StatusBadRequest, err)
	}

	if err := h.movieService.UpdateMovie(updatedMovie.ToDomain()); err != nil {
		return context.AbortWithError(http.StatusInternalServerError, err)
	}

	context.IndentedJSON(http.StatusOK, gin.H{"status": "Movie updated"})
	return nil
}

func (h *HttpMovieAdapter) DeleteMovie(context *gin.Context) error {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		return context.AbortWithError(http.StatusBadRequest, err)
	}

	movie, err := h.movieService.GetMovie(uint(id))
	if err != nil {
		return context.AbortWithError(http.StatusNotFound, err)
	}

	if err := h.movieService.DeleteMovie(movie); err != nil {
		return context.AbortWithError(http.StatusInternalServerError, err)
	}

	context.IndentedJSON(http.StatusOK, gin.H{"status": "Movie deleted"})
	return nil
}
