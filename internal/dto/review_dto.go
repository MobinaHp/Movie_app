package dto


type CreateReviewRequest struct {
	MovieID uint    `json:"movie_id" binding:"required"`
	UserID  uint    `json:"user_id" binding:"required"`
	Rating  float64 `json:"rating" binding:"required"`
	Comment string  `json:"comment"`
}

// DTO for updating an existing review.
type UpdateReviewRequest struct {
	Rating  *float64 `json:"rating"`
	Comment *string  `json:"comment"`
}

// DTO for sending a review back to the client.
type ReviewResponse struct {
	ID        uint    `json:"id"`
	MovieID   uint    `json:"movie_id"`
	UserID    uint    `json:"user_id"`
	MovieName string  `json:"movie_name"`
	UserName  string  `json:"user_name"`
	Rating    float64 `json:"rating"`
	Comment   string  `json:"comment"`
}
