package main

import (
	"WEB1/internal/app/handler"
	"WEB1/internal/app/repository"
	"WEB1/internal/app/service"
	"WEB1/internal/app"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Successfully connected to the database!")

	if err := db.AutoMigrate(
		&repository.MovieModel{},
		&repository.UserModel{},
		&repository.ReviewModel{},
	); err != nil {
		log.Fatalf("Failed to auto-migrate schema: %v", err)
	}

	//repositories
	movieRepo := repository.NewMovieRepository(db)
	userRepo := repository.NewUserRepository(db)
	reviewRepo := repository.NewReviewRepository(db)

	//services
	movieService := service.NewMovieService(movieRepo)
	userService := service.NewUserService(userRepo)
	reviewService := service.NewReviewService(reviewRepo, movieRepo, userRepo)

	//handlers
	movieHandler := handler.NewMovieHandler(movieService)
	userHandler := handler.NewUserHandler(userService)
	reviewHandler := handler.NewReviewHandler(reviewService)

	// Create router
	router := gin.Default()

	// Setup routes
	app.SetupRoutes(router, movieHandler, userHandler, reviewHandler)

	// Start server
	log.Println("Server starting on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
