package repository

import "WEB1/internal/domain"


type MovieRepository interface {
	Create(movie domain.Movie) (domain.Movie, error)
	GetByID(id int) (domain.Movie, error)
	List() ([]domain.Movie, error)
}

type UserRepository interface {
	Create(user domain.User) (domain.User, error)
	GetByID(id uint) (domain.User, error)
	List() ([]domain.User, error)
} 

type ReviewRepository interface {
	Create(review domain.Review) (domain.Review, error)
	GetByID(id uint) (domain.Review, error)
	List() ([]domain.Review, error)
}