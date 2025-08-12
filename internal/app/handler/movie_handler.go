package handler

import (
	"net/http"
	"strconv"
	"WEB1/internal/app/service"
	"WEB1/internal/dto"
	"WEB1/internal/domain"
	"github.com/gin-gonic/gin"
)

type movieHandler struct {
	service service.MovieService
}

func NewMovieHandler(s service.MovieService) MovieHandler {
	return &movieHandler{service: s}
}

func toMovieResponse(d domain.Movie) dto.MovieResponse {
	return dto.MovieResponse{
		ID:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		Genres:      d.Genres,
		ReleaseDate: d.ReleaseDate.Format("2006-01-02"),
	}
}

func (h *movieHandler) CreateMovie(c *gin.Context) {
	var req dto.CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	movie, err := h.service.AddMovie(req.Title, req.Description, req.Genres, req.ReleaseDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, toMovieResponse(movie))
}

func (h *movieHandler) UpdateMovie(c *gin.Context) {
	idStr := c.Param("id") 
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	var req dto.CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie, err := h.service.UpdateMovie(uint(id), req.Title, req.Description, req.Genres, req.ReleaseDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, toMovieResponse(movie))
}

func (h *movieHandler) DeleteMovie(c *gin.Context) {
	idStr := c.Param("id")  // اینجا
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}
	if err := h.service.DeleteMovie(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil) 
}

func (h *movieHandler) GetMovieByID(c *gin.Context) {
	idStr := c.Param("id")  // اینجا
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}
	movie, err := h.service.GetMovieByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}
	c.JSON(http.StatusOK, toMovieResponse(movie))
}

func (h *movieHandler) ListMovies(c *gin.Context) {
	movies, err := h.service.ListMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var responses []dto.MovieResponse
	for _, m := range movies {
		responses = append(responses, toMovieResponse(m))
	}
	c.JSON(http.StatusOK, responses)
}