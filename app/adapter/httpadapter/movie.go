package httpadapter

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Acova/movie-collection/app/domain"
	"github.com/Acova/movie-collection/app/port"
	"github.com/gin-gonic/gin"
)

type HttpMovieAdapter struct {
	movieService port.MovieService
}

type HttpMovie struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title" binding:"required,min=1,max=100"`
	Director    string  `json:"director" binding:"max=50"`
	Synopsis    string  `json:"synopsis" binding:"max=500"`
	ReleaseYear int     `json:"release_year"`
	Cast        string  `json:"cast" binding:"max=200"`
	Genre       string  `json:"genre" binding:"max=50"`
	Rating      float64 `json:"rating" binding:"min=0,max=10"`
	Duration    int     `json:"duration" binding:"min=0"`
	PosterURL   string  `json:"poster_url"`
}

func FromDomain(movie *domain.Movie) *HttpMovie {
	return &HttpMovie{
		ID:          movie.ID,
		Title:       movie.Title,
		Director:    movie.Director,
		Synopsis:    movie.Synopsis,
		ReleaseYear: movie.ReleaseYear,
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
		ReleaseYear: h.ReleaseYear,
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

// @Summary Create a new movie
// @Description Create a new movie in the collection
// @Tags Movies
// @Accept json
// @Produce json
// @Param movie body HttpMovie true "Movie object"
// @Success 201 {object} HttpMovie
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /movies [post]
// @Security ApiKeyAuth
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

	user, loggedIn := GetLoggedInUser(context)
	if !loggedIn {
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
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

// @Summary List movies
// @Description List all movies with optional filters
// @Tags Movies
// @Accept json
// @Produce json
// @Param title query string false "Filter by movie title"
// @Param director query string false "Filter by movie director"
// @Param genre query string false "Filter by movie genre"
// @Param cast query string false "Filter by movie cast"
// @Success 200 {array} HttpMovie
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /movies [get]
// @Security ApiKeyAuth
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

// @Summary Get a movie by ID
// @Description Get details of a specific movie by its ID
// @Tags Movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} HttpMovie
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /movies/{id} [get]
// @Security ApiKeyAuth
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

// @Summary Update a movie
// @Description Update details of a specific movie by its ID
// @Tags Movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Param movie body HttpMovie true "Updated movie object"
// @Success 200 {object} HttpMovie
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /movies/{id} [put]
// @Security ApiKeyAuth
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

	user, loggedIn := GetLoggedInUser(context)
	if !loggedIn {
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if movieToUpdate.UserID != user.ID {
		context.IndentedJSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this movie"})
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

// @Summary Delete a movie
// @Description Delete a specific movie by its ID
// @Tags Movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /movies/{id} [delete]
// @Security ApiKeyAuth
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

	user, loggedIn := GetLoggedInUser(context)
	if !loggedIn {
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	if movie.UserID != user.ID {
		context.IndentedJSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this movie"})
		return
	}

	if err := h.movieService.DeleteMovie(movie); err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"status": "Movie deleted"})
}
