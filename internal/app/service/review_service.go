package service

import (
	"WEB1/internal/dto"
	"WEB1/internal/app/repository"
	"WEB1/internal/domain"
	"errors"
)

type reviewService struct {
	repo       repository.ReviewRepository
	movieRepo  repository.MovieRepository 
	userRepo   repository.UserRepository  
}

func NewReviewService(repo repository.ReviewRepository, movieRepo repository.MovieRepository, userRepo repository.UserRepository) *reviewService {
	return &reviewService{repo: repo, movieRepo: movieRepo, userRepo: userRepo}
}

// toReviewResponse یک متد کمکی برای ساخت DTO از دامین مدل‌ها
func (s *reviewService) toReviewResponse(review domain.Review) (dto.ReviewResponse, error) {
    movie, err := s.movieRepo.GetByID(review.MovieID)
    if err != nil {
        return dto.ReviewResponse{}, errors.New("movie not found for review")
    }
    user, err := s.userRepo.GetByID(review.UserID)
    if err != nil {
        return dto.ReviewResponse{}, errors.New("user not found for review")
    }

    return dto.ReviewResponse{
        ID:        review.ID,
        MovieID:   review.MovieID,
        UserID:    review.UserID,
        MovieName: movie.Title,
        UserName:  user.Name,
        Rating:    review.Rating,
        Comment:   review.Comment,
    }, nil
}

func (s *reviewService) AddReview(movieID, userID uint, rating float64, comment string) (dto.ReviewResponse, error) {
    review := domain.Review{
        MovieID: movieID,
        UserID:  userID,
        Rating:  rating,
        Comment: comment,
    }
    createdReview, err := s.repo.Create(review)
    if err != nil {
        return dto.ReviewResponse{}, err
    }
    return s.toReviewResponse(createdReview)
}

func (s *reviewService) UpdateReview(id uint, rating *float64, comment *string) (dto.ReviewResponse, error) {
    existingReview, err := s.repo.GetByID(id)
    if err != nil {
        return dto.ReviewResponse{}, errors.New("review not found")
    }

    if rating != nil {
        existingReview.Rating = *rating
    }
    if comment != nil {
        existingReview.Comment = *comment
    }

    updatedReview, err := s.repo.Update(existingReview)
    if err != nil {
        return dto.ReviewResponse{}, err
    }
    return s.toReviewResponse(updatedReview)
}

func (s *reviewService) DeleteReview(id uint) error {
	return s.repo.Delete(id)
}

func (s *reviewService) GetReviewByID(id uint) (dto.ReviewResponse, error) {
    review, err := s.repo.GetByID(id)
    if err != nil {
        return dto.ReviewResponse{}, err
    }
    return s.toReviewResponse(review)
}

func (s *reviewService) ListReviews() ([]dto.ReviewResponse, error) {
    reviews, err := s.repo.List()
    if err != nil {
        return nil, err
    }
    
    var responses []dto.ReviewResponse
    for _, review := range reviews {
        resp, err := s.toReviewResponse(review)
        if err != nil {
            return nil, err
        }
        responses = append(responses, resp)
    }
    return responses, nil
}