package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"WEB1/internal/app/service"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) UserHandler {
	return &userHandler{service: s}
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.service.RegisterUser(req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.HashedPassword,
	}
	json.NewEncoder(w).Encode(resp)
}

func (h *userHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := UserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.HashedPassword,
	}
	json.NewEncoder(w).Encode(resp)
}

func (h *userHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.ListUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var resp []UserResponse
	for _, u := range users {
		resp = append(resp, UserResponse{
			ID:       u.ID,
			Name:     u.Name,
			Email:    u.Email,
			Password: u.HashedPassword,
		})
	}
	json.NewEncoder(w).Encode(resp)
}
