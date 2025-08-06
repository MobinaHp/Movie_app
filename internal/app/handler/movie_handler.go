package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"WEB1/internal/app/service"
)


type movieHandler struct {
	service service.MovieService
}

func NewMovieHandler(s service.MovieService) MovieHandler {
	return &movieHandler{service: s}
}

type createMovieRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Genres      []string `json:"genres"`
	ReleaseDate string   `json:"release_date"`
}

type movieResponse struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Genres      []string `json:"genres"`
	ReleaseDate string   `json:"release_date"`
}

func (h *movieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var req createMovieRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	movie, err := h.service.AddMovie(req.Title, req.Description, req.Genres, req.ReleaseDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := movieResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Genres:      movie.Genres,
		ReleaseDate: movie.ReleaseDate.Format("2006-01-02"),
	}
	json.NewEncoder(w).Encode(resp)
}

func (h *movieHandler) GetMovieByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	movie, err := h.service.GetMovieByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := movieResponse{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Genres:      movie.Genres,
		ReleaseDate: movie.ReleaseDate.Format("2006-01-02"),
	}
	json.NewEncoder(w).Encode(resp)
}

func (h *movieHandler) ListMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.service.ListMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var resp []movieResponse
	for _, m := range movies {
		resp = append(resp, movieResponse{
			ID:          m.ID,
			Title:       m.Title,
			Description: m.Description,
			Genres:      m.Genres,
			ReleaseDate: m.ReleaseDate.Format("2006-01-02"),
		})
	}
	json.NewEncoder(w).Encode(resp)
}
