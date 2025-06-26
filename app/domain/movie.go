package domain

type Movie struct {
	ID          uint
	Title       string
	Director    string
	ReleaseYear int
	Cast        string
	Genre       string
	Synopsis    string
	Rating      float64
	Duration    int
	PosterURL   string
	UserID      uint
}
