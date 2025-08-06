package domain

import "time"

type Movie struct {
	ID          int
	Title       string
	Description string
	Genres      []string
	ReleaseDate time.Time
}
