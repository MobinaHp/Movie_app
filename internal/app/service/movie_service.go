package service

import (
	"time"
	"WEB1/internal/domain"
	"WEB1/internal/app/repository"
)

type movieService struct {
	repo repository.MovieRepository
}

func NewMovieService(r repository.MovieRepository) MovieService {
	return &movieService{repo: r}
}

func (s *movieService) AddMovie(title, desc string, genres []string, release string) (domain.Movie, error) {
	date, err := time.Parse("2006-01-02", release)
	if err != nil {
		return domain.Movie{}, err
	}
	m := domain.Movie{
		Title:       title,
		Description: desc,
		Genres:      genres,
		ReleaseDate: date,
	}
	return s.repo.Create(m)
}

func (s *movieService) GetMovieByID(id int) (domain.Movie, error) {
	return s.repo.GetByID(id)
}

func (s *movieService) ListMovies() ([]domain.Movie, error) {
	return s.repo.List()
}
