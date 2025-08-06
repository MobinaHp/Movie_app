package repository

import (
	"strings"
	"time"
	"gorm.io/gorm"
	"WEB1/internal/domain"
)

type MovieModel struct {
	ID          int       `gorm:"primaryKey"`
	Title       string
	Description string
	Genres      string    // comma-separated
	ReleaseDate time.Time
	gorm.Model
}



type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db: db}
}

func (r *movieRepository) Create(m domain.Movie) (domain.Movie, error) {
	model := MovieModel{
		Title:       m.Title,
		Description: m.Description,
		Genres:      strings.Join(m.Genres, ","),
		ReleaseDate: m.ReleaseDate,
	}
	if err := r.db.Create(&model).Error; err != nil {
		return domain.Movie{}, err
	}
	return domain.Movie{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
		Genres:      strings.Split(model.Genres, ","),
		ReleaseDate: model.ReleaseDate,
	}, nil
}

func (r *movieRepository) GetByID(id int)(domain.Movie, error) {
	var model MovieModel
	if err := r.db.First(&model, id).Error; err != nil {
		return domain.Movie{}, err
	}
	return domain.Movie{
		ID:          model.ID,
		Title:       model.Title,
		Description: model.Description,
		Genres:      strings.Split(model.Genres, ","),
		ReleaseDate: model.ReleaseDate,
	}, nil
}

func (r *movieRepository) List() ([]domain.Movie, error) {
	var models []MovieModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}
	var result []domain.Movie
	for _, m := range models {
		result = append(result, domain.Movie{
			ID:          m.ID,
			Title:       m.Title,
			Description: m.Description,
			Genres:      strings.Split(m.Genres, ","),
			ReleaseDate: m.ReleaseDate,
		})
	}
	return result, nil
}
