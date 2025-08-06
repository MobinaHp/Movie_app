package service

import "WEB1/internal/domain"

type MovieService interface {
	AddMovie(title, desc string, genres []string, release string) (domain.Movie, error)
	GetMovieByID(id int) (domain.Movie, error)
	ListMovies() ([]domain.Movie, error)
}

type UserService interface {
	RegisterUser(name, email, password string) (domain.User, error)
	GetUserByID(id uint) (domain.User, error)
	ListUsers() ([]domain.User, error)
}

type ReviewService interface {
	AddReview(movieID, userID uint, rating float64, comment string) (domain.Review, error)
	GetReviewByID(id uint) (domain.Review, error)
	ListReviews() ([]domain.Review, error)
}