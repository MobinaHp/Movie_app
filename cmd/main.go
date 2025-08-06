package main

import (
	"WEB1/internal/app/handler"
	"WEB1/internal/app/repository"
	"WEB1/internal/app/service"
	"fmt"
	"log"
	"net/http"
	"os"


	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Successfully connected to the database!")

	if err := db.AutoMigrate(&repository.MovieModel{}, &repository.UserModel{}, &repository.ReviewModel{}); err != nil {
		log.Fatalf("Failed to auto-migrate schema: %v", err)
	}

	//Movie
	movieRepo := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepo)
	movieHandler := handler.NewMovieHandler(movieService)

	//User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	//Review
	reviewRepo := repository.NewReviewRepository(db)
	reviewService := service.NewReviewService(reviewRepo)
	reviewHandler := handler.NewReviewHandler(reviewService)

	router := mux.NewRouter()

	//Movie
	router.HandleFunc("/movies", movieHandler.CreateMovie).Methods("POST")
	router.HandleFunc("/movies", movieHandler.ListMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", movieHandler.GetMovieByID).Methods("GET")

	//User
	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users", userHandler.ListUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")

	//Review
	router.HandleFunc("/reviews", reviewHandler.CreateReview).Methods("POST")
	router.HandleFunc("/reviews", reviewHandler.ListReviews).Methods("GET")
	router.HandleFunc("/reviews/{id}", reviewHandler.GetReviewByID).Methods("GET")
	

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}