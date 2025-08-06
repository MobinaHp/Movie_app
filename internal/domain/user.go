package domain

import "time"

type User struct{
	ID uint
	Name string
	Email string
	HashedPassword string
	CreatedAt time.Time
}