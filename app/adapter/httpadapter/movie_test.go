package httpadapter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Acova/movie-collection/app/domain"
	"github.com/Acova/movie-collection/app/port/mock"
	"github.com/gin-gonic/gin"
)

func TestHttpMovieReturnsDomainMovie(t *testing.T) {
	movie := &HttpMovie{
		Title:       "Inception",
		Director:    "Christopher Nolan",
		Synopsis:    "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.",
		ReleaseYear: 2010,
		Cast:        "Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page",
		Genre:       "Science Fiction",
		Rating:      8.8,
		Duration:    148,
		PosterURL:   "https://example.com/inception.jpg",
	}

	domainMovie := movie.ToDomain()

	if domainMovie.Title != movie.Title || domainMovie.Director != movie.Director ||
		domainMovie.Synopsis != movie.Synopsis || domainMovie.ReleaseYear != movie.ReleaseYear ||
		domainMovie.Cast != movie.Cast || domainMovie.Genre != movie.Genre ||
		domainMovie.Rating != movie.Rating || domainMovie.Duration != movie.Duration ||
		domainMovie.PosterURL != movie.PosterURL {
		t.Errorf("Expected domain movie to be %+v, but got %+v", movie.ToDomain(), domainMovie)
	}
}

func TestCreateMovie(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockMovieService := &mock.MockMovieService{
		Movies: []*domain.Movie{},
	}

	httpAdapter := NewHttpMovieAdapter(mockMovieService)

	movie := &HttpMovie{
		ID:          1,
		Title:       "Inception",
		Director:    "Christopher Nolan",
		Synopsis:    "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.",
		ReleaseYear: 2010,
		Cast:        "Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page",
		Genre:       "Science Fiction",
		Rating:      8.8,
		Duration:    148,
		PosterURL:   "https://example.com/inception.jpg",
	}

	body, err := json.Marshal(movie)
	if err != nil {
		t.Fatalf("Failed to marshal movie: %v", err)
	}

	request, _ := http.NewRequest("POST", "/movie", bytes.NewBuffer(body))
	mockResponseWriter := httptest.NewRecorder()
	mockContext, _ := gin.CreateTestContext(mockResponseWriter)
	mockContext.Request = request

	loginUser := &domain.User{
		ID:    1,
		Email: "test@test.com",
		Name:  "Test User",
	}
	mockContext.Set("id", loginUser)

	httpAdapter.CreateMovie(mockContext)

	if len(mockMovieService.Movies) != 1 {
		t.Errorf("Expected 1 movie in service, but got %d", len(mockMovieService.Movies))
	}
	if mockMovieService.Movies[0].Title != "Inception" {
		t.Errorf("Expected title to be 'Inception', but got '%s'", mockMovieService.Movies[0].Title)
	}
	if mockMovieService.Movies[0].Director != "Christopher Nolan" {
		t.Errorf("Expected director to be 'Christopher Nolan', but got '%s'", mockMovieService.Movies[0].Director)
	}
	if mockMovieService.Movies[0].Synopsis != "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO." {
		t.Errorf("Expected synopsis to be 'A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.', but got '%s'", mockMovieService.Movies[0].Synopsis)
	}
	if mockMovieService.Movies[0].ReleaseYear != movie.ReleaseYear {
		t.Errorf("Expected release year to be '%d', but got '%d'", movie.ReleaseYear, mockMovieService.Movies[0].ReleaseYear)
	}
	if mockMovieService.Movies[0].Cast != "Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page" {
		t.Errorf("Expected cast to be 'Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page', but got '%s'", mockMovieService.Movies[0].Cast)
	}
	if mockMovieService.Movies[0].Genre != "Science Fiction" {
		t.Errorf("Expected genre to be 'Science Fiction', but got '%s'", mockMovieService.Movies[0].Genre)
	}
	if mockMovieService.Movies[0].Rating != 8.8 {
		t.Errorf("Expected rating to be 8.8, but got %f", mockMovieService.Movies[0].Rating)
	}
	if mockMovieService.Movies[0].Duration != 148 {
		t.Errorf("Expected duration to be 148, but got %d", mockMovieService.Movies[0].Duration)
	}
	if mockMovieService.Movies[0].PosterURL != "https://example.com/inception.jpg" {
		t.Errorf("Expected poster URL to be 'https://example.com/inception.jpg', but got '%s'", mockMovieService.Movies[0].PosterURL)
	}
}

func TestListMovies(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockMovieService := &mock.MockMovieService{
		Movies: []*domain.Movie{
			{
				ID:          1,
				Title:       "Inception",
				Director:    "Christopher Nolan",
				Synopsis:    "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.",
				ReleaseYear: 2010,
				Cast:        "Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page",
				Genre:       "Science Fiction",
				Rating:      8.8,
				Duration:    148,
				PosterURL:   "https://example.com/inception.jpg",
			},
		},
	}

	httpAdapter := NewHttpMovieAdapter(mockMovieService)

	request, _ := http.NewRequest("GET", "/movie", nil)
	mockResponseWriter := httptest.NewRecorder()
	mockContext, _ := gin.CreateTestContext(mockResponseWriter)
	mockContext.Request = request

	loginUser := &domain.User{
		ID:    1,
		Email: "test@test.com",
		Name:  "Test User",
	}
	mockContext.Set("id", loginUser)

	httpAdapter.ListMovies(mockContext)
	movies := []*HttpMovie{}
	if err := json.Unmarshal(mockResponseWriter.Body.Bytes(), &movies); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}
	if len(movies) != 1 {
		t.Errorf("Expected 1 movie, but got %d", len(movies))
	}
	if movies[0].Title != "Inception" {
		t.Errorf("Expected title to be 'Inception', but got '%s'", movies[0].Title)
	}
	if movies[0].Director != "Christopher Nolan" {
		t.Errorf("Expected director to be 'Christopher Nolan', but got '%s'", movies[0].Director)
	}
	if movies[0].Synopsis != "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO." {
		t.Errorf("Expected synopsis to be 'A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.', but got '%s'", movies[0].Synopsis)
	}
	if movies[0].ReleaseYear != 2010 {
		t.Errorf("Expected release year to be '2010', but got '%d'", movies[0].ReleaseYear)
	}
	if movies[0].Cast != "Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page" {
		t.Errorf("Expected cast to be 'Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page', but got '%s'", movies[0].Cast)
	}
	if movies[0].Genre != "Science Fiction" {
		t.Errorf("Expected genre to be 'Science Fiction', but got '%s'", movies[0].Genre)
	}
	if movies[0].Rating != 8.8 {
		t.Errorf("Expected rating to be 8.8, but got %f", movies[0].Rating)
	}
	if movies[0].Duration != 148 {
		t.Errorf("Expected duration to be 148, but got %d", movies[0].Duration)
	}
	if movies[0].PosterURL != "https://example.com/inception.jpg" {
		t.Errorf("Expected poster URL to be 'https://example.com/inception.jpg', but got '%s'", movies[0].PosterURL)
	}
}

func TestGetMovie(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockMovieService := &mock.MockMovieService{
		Movies: []*domain.Movie{
			{
				ID:          1,
				Title:       "Inception",
				Director:    "Christopher Nolan",
				Synopsis:    "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.",
				ReleaseYear: 2010,
				Cast:        "Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page",
				Genre:       "Science Fiction",
				Rating:      8.8,
				Duration:    148,
				PosterURL:   "https://example.com/inception.jpg",
			},
		},
	}

	httpAdapter := NewHttpMovieAdapter(mockMovieService)

	request, _ := http.NewRequest("GET", "/movie/1", nil)
	mockResponseWriter := httptest.NewRecorder()
	mockContext, _ := gin.CreateTestContext(mockResponseWriter)
	mockContext.Request = request
	mockContext.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	loginUser := &domain.User{
		ID:    1,
		Email: "test@test.com",
		Name:  "Test User",
	}
	mockContext.Set("id", loginUser)

	httpAdapter.GetMovie(mockContext)

	movie := &HttpMovie{}
	if err := json.Unmarshal(mockResponseWriter.Body.Bytes(), movie); err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}
	if movie.Title != "Inception" {
		t.Errorf("Expected title to be 'Inception', but got '%s'", movie.Title)
	}
	if movie.Director != "Christopher Nolan" {
		t.Errorf("Expected director to be 'Christopher Nolan', but got '%s'", movie.Director)
	}
	if movie.Synopsis != "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO." {
		t.Errorf("Expected synopsis to be 'A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.', but got '%s'", movie.Synopsis)
	}
	if movie.ReleaseYear != 2010 {
		t.Errorf("Expected release year to be '2010', but got '%d'", movie.ReleaseYear)
	}
	if movie.Cast != "Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page" {
		t.Errorf("Expected cast to be 'Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page', but got '%s'", movie.Cast)
	}
	if movie.Genre != "Science Fiction" {
		t.Errorf("Expected genre to be 'Science Fiction', but got '%s'", movie.Genre)
	}
	if movie.Rating != 8.8 {
		t.Errorf("Expected rating to be 8.8, but got %f", movie.Rating)
	}
	if movie.Duration != 148 {
		t.Errorf("Expected duration to be 148, but got %d", movie.Duration)
	}
	if movie.PosterURL != "https://example.com/inception.jpg" {
		t.Errorf("Expected poster URL to be 'https://example.com/inception.jpg', but got '%s'", movie.PosterURL)
	}
}

func TestUpdateMovie(t *testing.T) {
	mockMovieService := &mock.MockMovieService{
		Movies: []*domain.Movie{
			{
				ID:          1,
				Title:       "Inception",
				Director:    "Christopher Nolan",
				Synopsis:    "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.",
				ReleaseYear: 2010,
				Cast:        "Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page",
				Genre:       "Science Fiction",
				Rating:      8.8,
				Duration:    148,
				PosterURL:   "https://example.com/inception.jpg",
				UserID:      1,
			},
		},
	}

	movie := &HttpMovie{
		ID:          1,
		Title:       "Inception Updated",
		Director:    "Christopher Nolan",
		Synopsis:    "An updated synopsis for Inception.",
		ReleaseYear: 2010,
		Cast:        "Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page",
		Genre:       "Science Fiction",
		Rating:      9.0,
		Duration:    150,
		PosterURL:   "https://example.com/inception_updated.jpg",
	}

	httpAdapter := NewHttpMovieAdapter(mockMovieService)

	body, err := json.Marshal(movie)
	if err != nil {
		t.Fatalf("Failed to marshal movie: %v", err)
	}
	request, _ := http.NewRequest("PUT", "/movie/1", bytes.NewBuffer(body))
	mockResponseWriter := httptest.NewRecorder()
	mockContext, _ := gin.CreateTestContext(mockResponseWriter)
	mockContext.Request = request
	mockContext.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	loginUser := &domain.User{
		ID:    1,
		Email: "test@test.com",
		Name:  "Test User",
	}
	mockContext.Set("id", loginUser)

	httpAdapter.UpdateMovie(mockContext)
	if len(mockMovieService.Movies) != 1 {
		t.Errorf("Expected 1 movie in service, but got %d", len(mockMovieService.Movies))
	}
	if mockMovieService.Movies[0].Title != "Inception Updated" {
		t.Errorf("Expected title to be 'Inception Updated', but got '%s'", mockMovieService.Movies[0].Title)
	}
	if mockMovieService.Movies[0].Rating != 9.0 {
		t.Errorf("Expected rating to be 9.0, but got %f", mockMovieService.Movies[0].Rating)
	}
	if mockMovieService.Movies[0].Duration != 150 {
		t.Errorf("Expected duration to be 150, but got %d", mockMovieService.Movies[0].Duration)
	}
	if mockMovieService.Movies[0].PosterURL != "https://example.com/inception_updated.jpg" {
		t.Errorf("Expected poster URL to be 'https://example.com/inception_updated.jpg', but got '%s'", mockMovieService.Movies[0].PosterURL)
	}
}

func TestDeleteMovie(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockMovieService := &mock.MockMovieService{
		Movies: []*domain.Movie{
			{
				ID:          1,
				Title:       "Inception",
				Director:    "Christopher Nolan",
				Synopsis:    "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a CEO.",
				ReleaseYear: 2010,
				Cast:        "Leonardo DiCaprio, Joseph Gordon-Levitt, Ellen Page",
				Genre:       "Science Fiction",
				Rating:      8.8,
				Duration:    148,
				PosterURL:   "https://example.com/inception.jpg",
				UserID:      1,
			},
		},
	}

	httpAdapter := NewHttpMovieAdapter(mockMovieService)

	request, _ := http.NewRequest("DELETE", "/movie/1", nil)
	mockResponseWriter := httptest.NewRecorder()
	mockContext, _ := gin.CreateTestContext(mockResponseWriter)
	mockContext.Request = request
	mockContext.Params = gin.Params{gin.Param{Key: "id", Value: "1"}}

	loginUser := &domain.User{
		ID:    1,
		Email: "test@test.com",
		Name:  "Test User",
	}
	mockContext.Set("id", loginUser)

	httpAdapter.DeleteMovie(mockContext)

	if len(mockMovieService.Movies) != 0 {
		t.Errorf("Expected no movies in service after deletion, but got %d", len(mockMovieService.Movies))
	}
}
