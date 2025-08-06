package domain

import "time"

type Review struct {
	ID        uint      `json:"id"`
	MovieID  uint      `json:"movie_id"`
	UserID   uint      `json:"user_id"`
	Rating   float64   `json:"rating"`
	Comment  string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
