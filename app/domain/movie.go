package domain

import "time"

type Movie struct {
	ID          uint
	Title       string
	Director    string
	ReleaseDate time.Time
	Cast        string
	Genre       string
	Synopsis    string
	Rating      float64
	Duration    int
	PosterURL   string
}
