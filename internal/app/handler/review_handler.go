package handler

import (
	"WEB1/internal/app/service"
	"net/http"
	"strconv"
	"WEB1/internal/dto"
	"github.com/gin-gonic/gin"
)

type reviewHandler struct {
	service service.ReviewService
}

func NewReviewHandler(service service.ReviewService) *reviewHandler {
	return &reviewHandler{service: service}
}

func (h *reviewHandler) CreateReview(c *gin.Context) {
	var req dto.CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.service.AddReview(req.MovieID, req.UserID, req.Rating, req.Comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}

func (h *reviewHandler) UpdateReview(c *gin.Context) {
	idStr := c.Param("id")  // اینجا
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}

	var req dto.UpdateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.service.UpdateReview(uint(id), req.Rating, req.Comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *reviewHandler) DeleteReview(c *gin.Context) {
	idStr := c.Param("id")  // اینجا
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}
	if err := h.service.DeleteReview(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *reviewHandler) GetReviewByID(c *gin.Context) {	
	idStr := c.Param("id") 
	id, err := strconv.Atoi(idStr)	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid review ID"})
		return
	}
	resp, err := h.service.GetReviewByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *reviewHandler) ListReviews(c *gin.Context) {
	reviews, err := h.service.ListReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}