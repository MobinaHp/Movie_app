package service

import (
	"WEB1/internal/dto"
	"WEB1/internal/domain"
)

type MovieService interface {
	AddMovie(title, desc string, genres []string, release string) (domain.Movie, error)
	UpdateMovie(id uint, title, desc string, genres []string, release string) (domain.Movie, error)
	DeleteMovie(id uint) error
	GetMovieByID(id uint) (domain.Movie, error)
	ListMovies() ([]domain.Movie, error)
}

type UserService interface {
	RegisterUser(name, email, password string) (domain.User, error)
	UpdateUser(id uint, name, email, password string) (domain.User, error)
	DeleteUser(id uint) error
	GetUserByID(id uint) (domain.User, error)
	ListUsers() ([]domain.User, error)
}

type ReviewService interface {
    AddReview(movieID, userID uint, rating float64, comment string) (dto.ReviewResponse, error)
    UpdateReview(id uint, rating *float64, comment *string) (dto.ReviewResponse, error)
    DeleteReview(id uint) error
    GetReviewByID(id uint) (dto.ReviewResponse, error)
    ListReviews() ([]dto.ReviewResponse, error)
}