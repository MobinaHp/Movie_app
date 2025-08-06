package handler

import (
	"WEB1/internal/app/service"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

type reviewHandler struct {
	service service.ReviewService
}

func NewReviewHandler(service service.ReviewService) *reviewHandler {
	return &reviewHandler{service: service}
}

func (h *reviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	var req struct {
		MovieID uint    `json:"movie_id"`
		UserID  uint    `json:"user_id"`
		Rating  float64 `json:"rating"`
		Comment string  `json:"comment"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	review, err := h.service.AddReview(req.MovieID, req.UserID, req.Rating, req.Comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(review)
}

func (h *reviewHandler) GetReviewByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	review, err := h.service.GetReviewByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(review)
}

func (h *reviewHandler) ListReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := h.service.ListReviews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reviews)
}