package handler

import "net/http"

type MovieHandler interface {
	CreateMovie(w http.ResponseWriter, r *http.Request)
	GetMovieByID(w http.ResponseWriter, r *http.Request)
	ListMovies(w http.ResponseWriter, r *http.Request)
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUserByID(w http.ResponseWriter, r *http.Request)
	ListUsers(w http.ResponseWriter, r *http.Request)
}

type ReviewHandler interface {
	CreateReview(w http.ResponseWriter, r *http.Request)
	GetReviewByID(w http.ResponseWriter, r *http.Request)
	ListReviews(w http.ResponseWriter, r *http.Request)
}