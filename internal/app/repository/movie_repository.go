package repository

import (
	"strings"
	"time"
	"gorm.io/gorm"
	"WEB1/internal/domain"
)

type MovieModel struct {
	Title       string
	Description string
	Genres      string    // comma-separated
	ReleaseDate time.Time
	gorm.Model
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *movieRepository {
	return &movieRepository{db: db}
}

// Helper to convert domain.Movie to MovieModel
func toMovieModel(d domain.Movie) MovieModel {
    genresString := strings.Join(d.Genres, ",")
    return MovieModel{
        Title:       d.Title,
        Description: d.Description,
        Genres:      genresString,
        ReleaseDate: d.ReleaseDate,
    }
}

// Helper to convert MovieModel to domain.Movie
func toDomainMovie(m MovieModel) domain.Movie {
    genresSlice := strings.Split(m.Genres, ",")
    return domain.Movie{
        ID:          m.ID,
        Title:       m.Title,
        Description: m.Description,
        Genres:      genresSlice,
        ReleaseDate: m.ReleaseDate,
    }
}


func (r *movieRepository) Create(m domain.Movie) (domain.Movie, error) {
    model := toMovieModel(m)
    result := r.db.Create(&model)
    if result.Error != nil {
        return domain.Movie{}, result.Error
    }
    return toDomainMovie(model), nil
}

func (r *movieRepository) Update(m domain.Movie) (domain.Movie, error) {
    model := toMovieModel(m)
    result := r.db.Model(&model).Where("id = ?", model.ID).Updates(model)
    if result.Error != nil {
        return domain.Movie{}, result.Error
    }
    return toDomainMovie(model), nil
}

func (r *movieRepository) Delete(id uint) error {
    result := r.db.Delete(&MovieModel{}, id)
    return result.Error
}

func (r *movieRepository) GetByID(id uint) (domain.Movie, error) {
    var model MovieModel
    if err := r.db.First(&model, id).Error; err != nil {
        return domain.Movie{}, err
    }
    return toDomainMovie(model), nil
}

func (r *movieRepository) List() ([]domain.Movie, error) {
    var models []MovieModel
    if err := r.db.Find(&models).Error; err != nil {
        return nil, err
    }

    var movies []domain.Movie
    for _, model := range models {
        movies = append(movies, toDomainMovie(model))
    }
    return movies, nil
}