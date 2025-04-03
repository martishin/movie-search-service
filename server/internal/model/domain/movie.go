package domain

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	RunTime     int       `json:"runtime"`
	MPAARating  string    `json:"mpaa_rating"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Video       string    `json:"video"`
	Genres      []*Genre  `json:"genres,omitempty"`
	UserRating  float64   `json:"user_rating"`
	// UserRating float64 `json:"-"`
}
