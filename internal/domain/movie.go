package domain

import "time"

type Movie struct {
	ID          uint
	Title       string
	Description string
	Genres      []string
	ReleaseDate time.Time
}
