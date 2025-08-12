package handler

import (
	"github.com/gin-gonic/gin"
)

type MovieHandler interface {
	CreateMovie(c *gin.Context)
	UpdateMovie(c *gin.Context)
	DeleteMovie(c *gin.Context)
	GetMovieByID(c *gin.Context)
	ListMovies(c *gin.Context)
}

type UserHandler interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	ListUsers(c *gin.Context)
}

type ReviewHandler interface {
	CreateReview(c *gin.Context)
	UpdateReview(c *gin.Context)
	DeleteReview(c *gin.Context)
	GetReviewByID(c *gin.Context)
	ListReviews(c *gin.Context)
}