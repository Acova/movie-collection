package httpadapter

import (
	"fmt"
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
	ID          uint      `json:"id"`
	Title       string    `json:"title" binding:"required,min=1,max=100"`
	Director    string    `json:"director" binding:"max=50"`
	Synopsis    string    `json:"synopsis" binding:"max=500"`
	ReleaseDate time.Time `json:"release_date"`
	Cast        string    `json:"cast" binding:"max=200"`
	Genre       string    `json:"genre" binding:"max=50"`
	Rating      float64   `json:"rating" binding:"min=0,max=10"`
	Duration    int       `json:"duration" binding:"min=0"`
	PosterURL   string    `json:"poster_url"`
}

func FromDomain(movie *domain.Movie) *HttpMovie {
	return &HttpMovie{
		ID:          movie.ID,
		Title:       movie.Title,
		Director:    movie.Director,
		Synopsis:    movie.Synopsis,
		ReleaseDate: movie.ReleaseDate,
		Cast:        movie.Cast,
		Genre:       movie.Genre,
		Rating:      movie.Rating,
		Duration:    movie.Duration,
		PosterURL:   movie.PosterURL,
	}
}

func (h *HttpMovie) ToDomain() *domain.Movie {
	return &domain.Movie{
		ID:          h.ID,
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

func (h *HttpMovieAdapter) CreateMovie(context *gin.Context) {
	movie := HttpMovie{}

	if err := context.BindJSON(&movie); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingMovie, err := h.movieService.ListMovies(map[string]string{"title": movie.Title})
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(existingMovie) > 0 {
		context.IndentedJSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("Movie with name `%s` already exists", movie.Title)})
		return
	}

	userValue, exists := context.Get("id")
	if !exists {
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user, ok := userValue.(*domain.User)
	if !ok {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data"})
		return
	}

	domainMovie := movie.ToDomain()
	domainMovie.UserID = user.ID
	err = h.movieService.CreateMovie(domainMovie)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	context.IndentedJSON(http.StatusCreated, FromDomain(domainMovie))
}

func (h *HttpMovieAdapter) ListMovies(context *gin.Context) {
	filter := make(map[string]string)
	if title := context.Query("title"); title != "" {
		filter["title"] = "%" + title + "%"
	}
	if director := context.Query("director"); director != "" {
		filter["director"] = "%" + director + "%"
	}
	if genre := context.Query("genre"); genre != "" {
		filter["genre"] = "%" + genre + "%"
	}
	if cast := context.Query("cast"); cast != "" {
		filter["cast"] = "%" + cast + "%"
	}

	domainMovies, err := h.movieService.ListMovies(filter)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	movies := make([]*HttpMovie, len(domainMovies))
	for i, movie := range domainMovies {
		movies[i] = FromDomain(movie)
	}

	context.IndentedJSON(http.StatusOK, movies)
}

func (h *HttpMovieAdapter) GetMovie(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	movie, err := h.movieService.GetMovie(uint(id))
	if err != nil {
		context.AbortWithError(http.StatusNotFound, err)
		return
	}

	context.IndentedJSON(http.StatusOK, FromDomain(movie))
}

func (h *HttpMovieAdapter) UpdateMovie(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	movieToUpdate, err := h.movieService.GetMovie(uint(id))
	if err != nil {
		context.AbortWithError(http.StatusNotFound, err)
		return
	}

	updatedMovie := HttpMovie{}
	if err := context.BindJSON(&updatedMovie); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedDomainMovie := updatedMovie.ToDomain()
	updatedDomainMovie.ID = uint(id)
	updatedDomainMovie.UserID = movieToUpdate.UserID // Preserve the user ID
	if err := h.movieService.UpdateMovie(updatedDomainMovie); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.IndentedJSON(http.StatusOK, FromDomain(updatedDomainMovie))
}

func (h *HttpMovieAdapter) DeleteMovie(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	movie, err := h.movieService.GetMovie(uint(id))
	if err != nil {
		context.AbortWithError(http.StatusNotFound, err)
		return
	}

	if err := h.movieService.DeleteMovie(movie); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"status": "Movie deleted"})
}
