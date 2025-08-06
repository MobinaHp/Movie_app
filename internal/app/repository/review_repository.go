package repository

import (
	"WEB1/internal/domain"
	"gorm.io/gorm"
)

type ReviewModel struct {
	gorm.Model
	MovieID  uint    `json:"movie_id"`
	UserID   uint    `json:"user_id"`
	Rating   float64 `json:"rating"`
	Comment  string  `json:"comment"`
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *reviewRepository {
	return &reviewRepository{db: db}
}

func (r *reviewRepository) Create(review domain.Review) (domain.Review, error) {
	model := ReviewModel{
		MovieID: review.MovieID,
		UserID:  review.UserID,
		Rating:  review.Rating,
		Comment: review.Comment,
	}
	if err := r.db.Create(&model).Error; err != nil {
		return domain.Review{}, err
	}
	return domain.Review{
		ID:      model.ID,
		MovieID: model.MovieID,
		UserID:  model.UserID,
		Rating:  model.Rating,
		Comment: model.Comment,
	}, nil
}

func (r *reviewRepository) GetByID(id uint) (domain.Review, error) {
	var model ReviewModel
	if err := r.db.First(&model, id).Error; err != nil {
		return domain.Review{}, err
	}
	return domain.Review{
		ID:      model.ID,
		MovieID: model.MovieID,
		UserID:  model.UserID,
		Rating:  model.Rating,
		Comment: model.Comment,
	}, nil
}

func (r *reviewRepository) List() ([]domain.Review, error) {
	var models []ReviewModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	var reviews []domain.Review
	for _, model := range models {
		reviews = append(reviews, domain.Review{
			ID:      model.ID,
			MovieID: model.MovieID,
			UserID:  model.UserID,
			Rating:  model.Rating,
			Comment: model.Comment,
		})
	}
	return reviews, nil
}
