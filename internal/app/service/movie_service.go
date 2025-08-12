package service

import (
	"WEB1/internal/domain"
	"WEB1/internal/app/repository"
	"fmt"
	"time"
)

type movieService struct {
	repo repository.MovieRepository
}

func NewMovieService(repo repository.MovieRepository) MovieService {
	return &movieService{repo: repo}
}

func (s *movieService) AddMovie(title, desc string, genres []string, release string) (domain.Movie, error) {
	releaseDate, err := time.Parse("2006", release)
	if err != nil {
		return domain.Movie{}, fmt.Errorf("invalid release date format: %w", err)
	}

	movie := domain.Movie{
		Title:       title,
		Description: desc,
		Genres:      genres,
		ReleaseDate: releaseDate,
	}

	movie, err = s.repo.Create(movie)
	if err != nil {
		return domain.Movie{}, err
	}
	return movie, nil
}

// UpdateMovie updates an existing movie.
func (s *movieService) UpdateMovie(id uint, title, desc string, genres []string, release string) (domain.Movie, error) {
	releaseDate, err := time.Parse("2006", release)
	if err != nil {
		return domain.Movie{}, fmt.Errorf("invalid release date format: %w", err)
	}

	movie := domain.Movie{
		ID:          id,
		Title:       title,
		Description: desc,
		Genres:      genres,
		ReleaseDate: releaseDate,
	}
	movie, err = s.repo.Update(movie)
	if err != nil {
		return domain.Movie{}, err
	}
	return movie, nil
}

// DeleteMovie deletes a movie by ID.
func (s *movieService) DeleteMovie(id uint) error {
	return s.repo.Delete(id)
}

// GetMovieByID retrieves a movie by ID.
func (s *movieService) GetMovieByID(id uint) (domain.Movie, error) {
	return s.repo.GetByID(id)
}

// ListMovies retrieves all movies.
func (s *movieService) ListMovies() ([]domain.Movie, error) {
	return s.repo.List()
}