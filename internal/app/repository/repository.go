package repository

import "WEB1/internal/domain"


type MovieRepository interface {
    Create(m domain.Movie) (domain.Movie, error)
    Update(m domain.Movie) (domain.Movie, error)
	Delete(id uint) error
    GetByID(id uint) (domain.Movie, error)
    List() ([]domain.Movie, error)
}


type UserRepository interface {
	Create(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id uint) error
	GetByID(id uint) (domain.User, error)
	List() ([]domain.User, error)
}

type ReviewRepository interface {
	Create(review domain.Review) (domain.Review, error)
	Update(review domain.Review) (domain.Review, error)
	Delete(id uint) error
	GetByID(id uint) (domain.Review, error)
	List() ([]domain.Review, error)
}