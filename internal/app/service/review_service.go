package service 

import (
	"WEB1/internal/domain"
	"WEB1/internal/app/repository"
)

type reviewService struct {
	repo repository.ReviewRepository
}

func NewReviewService(repo repository.ReviewRepository) *reviewService {
	return &reviewService{repo: repo}
}

func (s *reviewService) AddReview(movieID, userID uint, rating float64, comment string) (domain.Review, error) {
	review := domain.Review{
		MovieID: movieID,
		UserID:  userID,
		Rating:  rating,
		Comment: comment,
	}
	return s.repo.Create(review)
}

func (s *reviewService) GetReviewByID(id uint) (domain.Review, error) {
	return s.repo.GetByID(id)
}

func (s *reviewService) ListReviews() ([]domain.Review, error) {
	return s.repo.List()
}