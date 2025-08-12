package app

import (
	"github.com/gin-gonic/gin"
	"WEB1/internal/app/handler"
)

func SetupRoutes(router *gin.Engine, movieHandler handler.MovieHandler, userHandler handler.UserHandler, reviewHandler handler.ReviewHandler) {
	api := router.Group("/api/v1") 

	// Movies Routes
	movies := api.Group("/movies")
	{
		movies.POST("/", movieHandler.CreateMovie)
		movies.PUT("/:id", movieHandler.UpdateMovie)
		movies.DELETE("/:id", movieHandler.DeleteMovie)
		movies.GET("/:id", movieHandler.GetMovieByID)
		movies.GET("/", movieHandler.ListMovies)
	}

	// Users Routes
	users := api.Group("/users")
	{
		users.POST("/", userHandler.CreateUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
		users.GET("/:id", userHandler.GetUserByID)
		users.GET("/", userHandler.ListUsers)
	}

	// Reviews Routes
	reviews := api.Group("/reviews")
	{
		reviews.POST("/", reviewHandler.CreateReview)
		reviews.PUT("/:id", reviewHandler.UpdateReview)
		reviews.DELETE("/:id", reviewHandler.DeleteReview)
		reviews.GET("/:id", reviewHandler.GetReviewByID)
		reviews.GET("/", reviewHandler.ListReviews)
	}
}
