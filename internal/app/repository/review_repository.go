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

// Helper to convert domain.Review to ReviewModel
func toReviewModel(d domain.Review) ReviewModel {
	return ReviewModel{
		MovieID: d.MovieID,
		UserID:  d.UserID,
		Rating:  d.Rating,
		Comment: d.Comment,
	}
}

// Helper to convert ReviewModel to domain.Review
func toDomainReview(m ReviewModel) domain.Review {
	return domain.Review{
		ID:      m.ID,
		MovieID: m.MovieID,
		UserID:  m.UserID,
		Rating:  m.Rating,
		Comment: m.Comment,
	}
}

func (r *reviewRepository) Create(review domain.Review) (domain.Review, error) {
	model := toReviewModel(review)
	if err := r.db.Create(&model).Error; err != nil {
		return domain.Review{}, err
	}
	return toDomainReview(model), nil
}

func (r *reviewRepository) Update(review domain.Review) (domain.Review, error) {
	model := toReviewModel(review)
	if err := r.db.Save(&model).Error; err != nil {
		return domain.Review{}, err
	}
	return toDomainReview(model), nil
}

func (r *reviewRepository) Delete(id uint) error {
	if err := r.db.Delete(&ReviewModel{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *reviewRepository) GetByID(id uint) (domain.Review, error) {
	var model ReviewModel
	if err := r.db.First(&model, id).Error; err != nil {
		return domain.Review{}, err
	}
	return toDomainReview(model), nil
}

func (r *reviewRepository) List() ([]domain.Review, error) {
	var models []ReviewModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}

	var reviews []domain.Review
	for _, model := range models {
		reviews = append(reviews, toDomainReview(model))
	}
	return reviews, nil
}
